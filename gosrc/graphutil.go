package gosrc

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Node struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Value uint32 `json:"value"`
}

type Link struct {
	Source uint32 `json:"source"`
	Target uint32 `json:"target"`
	Weight uint32 `json:"weight"`
	Width  uint32 `json:"width"`
}

func GraphToNodeLinks(graph string) (nodes map[string]Node, links []Link) {
	nodes = map[string]Node{}
	links = []Link{}
	var key uint32 = 0
	for _, line := range strings.Split(graph, "\n") {
		if 1 > len(line) {
			continue
		}
		splitStr := strings.Split(line, " ")
		if 2 > len(splitStr) {
			panic(fmt.Sprintf("go mod graph output format error——%s", line))
		}
		var nKey uint32 = 0
		var nKey1 uint32 = 0
		node, ok := nodes[splitStr[0]]
		if ok {
			// 权重自增
			node.Value++
			nodes[splitStr[0]] = node
			nKey = node.Id
		} else {
			// 新增node
			var newNode = Node{
				Id:    key,
				Name:  splitStr[0],
				Value: 1,
			}
			nodes[splitStr[0]] = newNode
			nKey = key
			key++
		}

		node1, ok1 := nodes[splitStr[1]]
		if ok1 {
			node1.Value++
			nodes[splitStr[0]] = node
			nKey1 = node1.Id
		} else {
			// 新增node
			var newNode = Node{
				Id:    key,
				Name:  splitStr[1],
				Value: 1,
			}
			nodes[splitStr[1]] = newNode
			nKey1 = key
			key++
		}

		newLink := Link{
			Source: nKey,
			Target: nKey1,
			Weight: 1,
			Width: 1,
		}
		links = append(links, newLink)
	}
	return
}

func GetGoModGraph() string {
	var goModGraph string = ""

	go func() {
		var scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			goModGraph += scanner.Text() + "\n"
		}
	}()

	time.Sleep(1500 * time.Millisecond)
	if 1 > len(goModGraph) {
		cmd := exec.Command("go mod graph")
		if out, err := cmd.CombinedOutput(); nil != err {
			panic(fmt.Sprintf("go mod graph cmd run failed: %+v", err))
		} else {
			return string(out)
		}
	}
	return goModGraph
}
