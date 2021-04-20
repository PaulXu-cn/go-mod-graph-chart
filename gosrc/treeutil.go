package gosrc

import (
	"fmt"
	"strings"
)

type RouteNode struct {
	Id    uint32   `json:"id"`
	Name  string   `json:"name"`
	Value uint32   `json:"value"`
	Route []uint32 `json:"-"`
}

type Tree struct {
	Id       uint32  `json:"-"`
	Name     string  `json:"name"`
	Value    uint32  `json:"value"`
	Children []*Tree `json:"children"`
}

var routeNodes = map[string]RouteNode{}

func BuildTree(graph string) (root *Tree, depth uint32, width uint32, repeatDependNodes map[uint32]*Tree) {
	root = new(Tree)
	repeatDependNodes = map[uint32]*Tree{}
	var roots = make(map[string]*Tree, 0)
	var key uint32 = 0
	for _, line := range strings.Split(graph, "\n") {
		if 1 > len(line) {
			continue
		}
		splitStr := strings.Split(line, " ")
		if 2 > len(splitStr) 	{
			panic(fmt.Sprintf("go mod graph output format error——%s", line))
		}

		var tree = new(Tree)
		var theRouteNode1 = GetRouteNode(splitStr[0])
		if nil == theRouteNode1 {
			// 寻找父节点
			tree.Name = splitStr[0]
			tree.Id = key
			tree.Value = 1

			if newRoot, ok := roots[splitStr[0]]; ok {
				newRoot.Value++
			} else {
				newRoot = &Tree{
					Id:    key,
					Name:  splitStr[0],
					Value: 1,
				}
				roots[splitStr[0]] = newRoot
				InsertTreeRoute(newRoot, []uint32{})
			}
			if 1 > len(root.Name) {
				// 如果 root 为空
				root.Id = tree.Id
				root.Name = tree.Name
				root.Value = tree.Value
				InsertTreeRoute(tree, []uint32{})
				key ++
			}
		} else {
			// 重复正常
			//fmt.Println("重复的依赖节点")
		}

		var tree2 = new(Tree)
		tree2.Name = splitStr[1]
		tree2.Id = key
		tree2.Value = 1

		// 依赖节点
		var theRouteNode2 = GetRouteNode(splitStr[1])
		if nil == theRouteNode2  {
			// 新节点
			AppendTreeAfter(root, splitStr[0], tree2)
			key ++
		} else {
			// 这是规则外的
			fmt.Println("重复的依赖节点", splitStr[1])
			if nil == repeatDependNodes || 1 > len(repeatDependNodes) {
				repeatDependNodes = make(map[uint32]*Tree, 0)
			}
			var newRepeatDependNode = Tree{}
			if value, ok := repeatDependNodes[theRouteNode1.Id]; !ok {
				newRepeatDependNode = Tree{
					Name: tree2.Name,
					Value: tree2.Value,
					Id: tree2.Id,
					Children: []*Tree{
						{
							Name: theRouteNode1.Name,
							Value: theRouteNode1.Value,
							Id: theRouteNode1.Id,
						},
					},
				}
				repeatDependNodes[tree2.Id] = &newRepeatDependNode
			} else {
				value.Children = append(value.Children, &Tree{
					Name: theRouteNode1.Name,
					Value: theRouteNode1.Value,
					Id: theRouteNode1.Id,
				})
			}
		}
	}
	depth, width = CalculateDepthHeight(root)
	return
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
		parentTree.Value++
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
			parentTree.Value++
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
			Id:    newTree.Id,
			Name:  newTree.Name,
			Route: keys,
			Value: newTree.Value,
		}
	}
	return
}

func CalculateDepthHeight(root *Tree) (depth uint32, width uint32) {
	depth = calcDepth(root.Children)
	var treeWidths = &[]uint32{1}
	calcWidth(root.Children, treeWidths, 1)
	width = (*treeWidths)[0]
	for _, item := range *treeWidths {
		if width < item {
			width = item
		}
	}
	return
}

func calcDepth(node []*Tree) (depth uint32) {
	if nil != node && 0 < len(node) {
		// 有
		var maxDepth uint32 = 0
		var theDepth uint32 = 0
		for _, item := range node {
			theDepth = calcDepth(item.Children)
			if theDepth > maxDepth {
				// 如果更深
				maxDepth = theDepth
			}
		}
		return maxDepth + 1
	}
	// 没有
	return 1
}

func calcWidth(node []*Tree, treeWidths *[]uint32, level int) {
	if nil != node && 0 < len(node) {
		// 有
		var theWidth uint32 = uint32(len(node))
		if level < len(*treeWidths) {
			var originWidth = (*treeWidths)[level]
			(*treeWidths)[level] = originWidth + theWidth
		} else {
			*treeWidths = append(*treeWidths, theWidth)
		}
		for _, item := range node {
			calcWidth(item.Children, treeWidths, level + 1)
		}
	}
}
