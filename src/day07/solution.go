package day07

import (
	"aoc2017/src/util"
	"fmt"
	"strconv"
	"strings"
)

type Set map[string]bool
type TreeMap map[string]*TreeNode

type TreeNode struct {
	children    Set
	name        string
	weight      int
	totalWeight int
}

type Weight map[string]int

func Solve() {
	input := util.GetLines("day07", "\n")
	tm := make(TreeMap, 0)
	for _, line := range input {
		tm.add(line)
	}
	root := tm.findRoot()
	fmt.Println("Part 1 = ", root)
	name, weights := tm.findUnstable(root)
	fmt.Println("Part 2 = ", name, tm[name].weight-unstableDiff(name, weights))
}

func (tm TreeMap) add(line string) {
	parts := strings.Split(line, "->")
	new := &TreeNode{
		children: make(Set, 0)}
	nw := strings.Split(parts[0], " ")
	new.name = strings.Trim(nw[0], " ")
	new.weight, _ = strconv.Atoi(strings.Trim(nw[1], " ()"))
	if len(parts) == 2 {
		new.children.addAll(parts[1])
	}
	tm[new.name] = new
}

func (ref Set) addAll(in string) {
	each := strings.Split(in, ",")
	for _, part := range each {
		ref[strings.Trim(part, " ")] = true
	}
}

func (ref TreeMap) findRoot() string {
	for key, node := range ref {
		if len(node.children) == 0 {
			continue
		}
		exists := false
		for _, node := range ref {
			if node.children[key] {
				exists = true
				break
			}
		}
		if !exists {
			return key
		}
	}
	return "None"
}

// Find a node where a single child is of different weight,
// and ensure that the balance above that child is correct
func (tm TreeMap) findUnstable(name string) (string, Weight) {
	parent := tm[name]
	unstableNode, weights := parent.unbalancedChild(tm)
	if unstableNode != "None" {
		unstableAbove, _ := tm[unstableNode].unbalancedChild(tm)
		if unstableAbove == "None" {
			return unstableNode, weights
		}
		return tm.findUnstable(unstableNode)
	}
	if unstableNode == "None" {
		return parent.name, weights
	}
	return "None", make(Weight, 0)
}

func (node TreeNode) unbalancedChild(tm TreeMap) (string, Weight) {
	// Collect total weight for all children
	elems := make(Weight, 0)
	for name := range node.children {
		elems[name] = tm[name].calcWeight(tm)
	}
	// Check if weights are equal
	previous := -1
	for _, weight := range elems {
		if previous != 0 && previous != weight {
			return unique(elems), elems
		}
		previous = weight
	}
	return "None", elems
}

func (node *TreeNode) calcWeight(tm TreeMap) int {
	if len(node.children) == 0 {
		return node.weight
	}
	if node.totalWeight != 0 {
		return node.totalWeight
	}
	sum := 0
	for name, childNode := range tm {
		if node.children[name] {
			sum += childNode.calcWeight(tm)
		}
	}
	node.totalWeight = sum + node.weight
	return node.totalWeight
}

func unique(children Weight) string {
	seen := make(map[int]string)
	unique := make(map[int]string)
	for key, value := range children {
		if _, ok := seen[value]; ok {
			delete(unique, value)
		} else {
			seen[value] = key
			unique[value] = key
		}
	}
	for _, first := range unique {
		return first
	}
	return "None"
}

// Subtract other weight from the unstable to get the unstableDiff
func unstableDiff(name string, weights Weight) int {
	for otherName, weight := range weights {
		if otherName != name {
			return weights[name] - weight
		}
	}
	return -1
}
