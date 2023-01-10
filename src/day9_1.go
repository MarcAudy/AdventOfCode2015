package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func day9_1() {

	type nodeInfo struct {
		nodeName string
		costs    map[int]int
	}
	var nodeMap = map[string]int{}
	var nodeInfoMap = map[int]nodeInfo{}

	getNodeID := func(nodeName string) int {
		if nid, ok := nodeMap[nodeName]; ok {
			return nid
		}
		nid := len(nodeMap)
		nodeMap[nodeName] = nid
		return nid
	}

	exp := regexp.MustCompile(`(.*) to (.*) = (.*)`)

	for _, line := range getInput("day9_input.txt") {
		matches := exp.FindStringSubmatch(line)

		nid1 := getNodeID(matches[1])
		nid2 := getNodeID(matches[2])
		cost, _ := strconv.Atoi(matches[3])

		if _, ok := nodeInfoMap[nid1]; !ok {
			nodeInfoMap[nid1] = nodeInfo{matches[1], make(map[int]int)}
		}
		if _, ok := nodeInfoMap[nid2]; !ok {
			nodeInfoMap[nid2] = nodeInfo{matches[2], make(map[int]int)}
		}

		nodeInfoMap[nid1].costs[nid2] = cost
		nodeInfoMap[nid2].costs[nid1] = cost
	}

	type route struct {
		curNode int
		cost    int
		visited set[int]
	}

	var routesToConsider []route

	nodeCount := len(nodeMap)
	for i := 0; i < nodeCount; i++ {
		for j := 0; j < nodeCount; j++ {
			if i != j {
				routesToConsider = append(routesToConsider, route{j, nodeInfoMap[i].costs[j], NewSet(i, j)})
			}
		}
	}

	for true {
		sort.Slice(routesToConsider, func(i, j int) bool { return routesToConsider[i].cost > routesToConsider[j].cost })

		r := Pop(&routesToConsider)

		if len(r.visited) == nodeCount {
			fmt.Println(r.cost)
			break
		}

		for i := 0; i < nodeCount; i++ {
			if !r.visited.Contains(i) {
				newVisited := CopyMap(r.visited)
				newVisited.Add(i)
				routesToConsider = append(routesToConsider, route{i, r.cost + nodeInfoMap[r.curNode].costs[i], newVisited})
			}
		}
	}

}
