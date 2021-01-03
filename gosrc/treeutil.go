package gosrc

import (
	"fmt"
	"strings"
)

type RouteNode struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Popularity uint32 `json:"popularity"`
	Route []uint32 `json:"-"`
}

type Tree struct {
	Id uint32 `json:"-"`
	Name string	`json:"name"`
	Popularity uint32 `json:"popularity"`
	Children []*Tree `json:"children"`
}

var routeNodes = map[string]RouteNode{}

func BuildTree(graph string) *Tree {
	var root = new(Tree)
	var roots = make(map[string]*Tree, 0)
	var key uint32 = 0
	for _, line := range strings.Split(graph, "\n") {
		if 1 > len(line) {
			continue
		}
		splitStr := strings.Split(line, " ")
		if 2 > len(splitStr) {
			panic(fmt.Sprintf("go mod graph output format error——%s", line))
		}

		var tree = new(Tree)
		var theRouteNode1 = GetRouteNode(splitStr[0])
		if nil == theRouteNode1 {
			// 寻找父节点
			tree.Name = splitStr[0]
			tree.Id = key
			tree.Popularity = 1

			if newRoot, ok := roots[splitStr[0]]; ok {
				newRoot.Popularity ++
			} else {
				newRoot = &Tree{
					Id: key,
					Name: splitStr[0],
					Popularity: 1,
				}
				roots[splitStr[0]] = newRoot
				InsertTreeRoute(newRoot, []uint32{})
			}
			if 1 > len(root.Name) {
				// 如果 root 为空
				root.Id = tree.Id
				root.Name = tree.Name
				root.Popularity = tree.Popularity
				InsertTreeRoute(tree, []uint32{})
				key ++
			}
		} else {
			// 重复正常
			//fmt.Println("重复的依赖节点")
		}

		var tree2 = new(Tree)
		// 依赖节点
		var theRouteNode2 = GetRouteNode(splitStr[1])
		if nil == theRouteNode2  {
			// 新节点
			tree2.Name = splitStr[1]
			tree2.Id = key
			tree2.Popularity = 1

			AppendTreeAfter(root, splitStr[0], tree2)
			key ++
		} else {
			// 按理说不可能
			//fmt.Println("重复的依赖节点", splitStr[1])
			//panic("重复的依赖节点")
		}
	}
	return root
}

func AppendTreeAfter(parentTreeNode *Tree, parent string, new *Tree) {
	routeKey := FindTreeRoute(parent)
	resultKeys := insertTree(parentTreeNode, routeKey, new)
	InsertTreeRoute(new, resultKeys)
}

func insertTree(parentTree *Tree, keys []uint32, new *Tree) (route []uint32) {
	if 1 > len(keys) {
		// 如果没有，直接挂根节点下
		if nil == parentTree.Children {
			parentTree.Children = make([]*Tree, 0)
		}
		parentTree.Children = append(parentTree.Children, new)
		parentTree.Popularity ++
		return []uint32{uint32(len(parentTree.Children)-1)}
	}
	theKey := keys[0]
	if nil != parentTree.Children && 0 < len(parentTree.Children) {
		if int(theKey) > len(parentTree.Children) {
			fmt.Println("out range of arr length")
		}
		var theChild = parentTree.Children[theKey]
		if 1 == len(keys) {
			if nil == theChild.Children {
				theChild.Children = make([]*Tree, 0)
			}
			theChild.Children = append(theChild.Children, new)
			parentTree.Popularity ++
			return []uint32{theKey, uint32(len(theChild.Children)-1)}
		} else {
			// 进入下级
			lastKeys := insertTree(theChild, keys[1:], new)
			return append([]uint32{theKey}, lastKeys...)
		}
	} else {

	}
	return
}

func GetRouteNode(key string) (*RouteNode) {
	if routeNode, ok := routeNodes[key]; ok {
		return &routeNode
	} else {
		// 没找到
		return nil
	}
	if "" == key {
		return nil
	} else {
		return nil
	}
}

func FindTreeRoute(key string) (route []uint32) {
	if routeNode, ok := routeNodes[key]; ok {
		return routeNode.Route
	}
	if "" == key {
		return []uint32{}
	} else {
		panic("有问题")
	}
}

func InsertTreeRoute(newTree *Tree, keys []uint32) {
	if the, ok := routeNodes[newTree.Name]; ok {
		the.Route = keys
	} else {
		routeNodes[newTree.Name] = RouteNode {
			Id: newTree.Id,
			Name: newTree.Name,
			Route: keys,
			Popularity: newTree.Popularity,
		}
	}
	return
}