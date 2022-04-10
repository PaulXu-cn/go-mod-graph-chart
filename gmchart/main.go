package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	dist "github.com/PaulXu-cn/go-mod-graph-chart/godist"
	src "github.com/PaulXu-cn/go-mod-graph-chart/gosrc"
)

var (
	debug int    = 0
	keep  int    = 0
	port  int    = 0
	mode  string = "tree"
)

func init() {
	flag.IntVar(&debug, "debug", 0, "whether debug model")
	flag.IntVar(&keep, "keep", 0, "set no zero http server never exit")
	flag.IntVar(&port, "port", 0, "set http server port or random port number")
	flag.StringVar(&mode, "mode", "tree", "work mode, [tree/graph]")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var header = w.Header()
	header.Add("Content-type", "text/html; charset=utf-8")
	var indexStr = dist.GetFile("index.html")
	w.Write([]byte(indexStr))
	w.WriteHeader(http.StatusOK)
}

func MainJsHandler(w http.ResponseWriter, r *http.Request) {
	var header = w.Header()
	header.Add("Content-type", "text/javascript; charset=utf-8")
	var indexStr = dist.GetFile("main.js")
	w.Write([]byte(indexStr))
	w.WriteHeader(http.StatusOK)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	var header = w.Header()
	header.Add("content-type", "text/json; charset=utf-8")
	var jsonStr = "{\"message\": \"pong\"}"
	w.Write([]byte(jsonStr))
	w.WriteHeader(http.StatusOK)
}

type GraphData struct {
	Nodes []src.Node `json:"nodes"`
	Links []src.Link `json:"links"`
	Num   uint32     `json:"num"`
}

type GraphJson struct {
	Message string    `json:"message"`
	Data    GraphData `json:"data"`
}

type TreeData struct {
	Tree  src.Tree `json:"tree"`
	Depth uint32   `json:"depth"`
	Width uint32   `json:"width"`
}

type TreeJson struct {
	Message string   `json:"message"`
	Data    TreeData `json:"data"`
}

type AnTreeData struct {
	Tree  map[string]*src.Tree `json:"tree"`
}

type AnTreeJson struct {
	Message string   `json:"message"`
	Data    AnTreeData `json:"data"`
}

func main() {
	flag.Parse()
	fmt.Println("go mod graph version v0.5.3")
	var goModGraph string = src.GetGoModGraph()

	// nodes and links
	nodes, links := src.GraphToNodeLinks(goModGraph)
	nodeArr := map[uint32]src.Node{}
	for _, item := range nodes {
		nodeArr[item.Id] = item
	}
	nodeSortArr := []src.Node{}
	for key := 0; key < len(nodes); key++ {
		nodeSortArr = append(nodeSortArr, nodeArr[uint32(key)])
	}

	// tree
	tree, depth, width, anotherTree := src.BuildTree(goModGraph)

	if 0 < debug {
		// 如果是 debug 模式
		re, _ := json.Marshal(&nodeSortArr)
		fmt.Printf("nodes : %s\n\n", string(re))
		re, _ = json.Marshal(&links)
		fmt.Printf("links : %s\n\n", string(re))
		re, _ = json.Marshal(&tree)
		fmt.Printf("tree : %s\n", string(re))
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/main.js", MainJsHandler)
	mux.HandleFunc("/ping", PingHandler)

	mux.HandleFunc("/graph.json", func(w http.ResponseWriter, r *http.Request) {
		var header = w.Header()
		header.Add("Content-type", "text/javascript; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		var graph = GraphJson{
			Message: "success",
			Data: GraphData{
				Nodes: nodeSortArr,
				Links: links,
				Num:   uint32(len(nodeSortArr)),
			},
		}
		var graphStr, _ = json.Marshal(graph)
		w.Write(graphStr)
	})

	mux.HandleFunc("/tree.json", func(w http.ResponseWriter, r *http.Request) {
		var header = w.Header()
		header.Add("Content-type", "text/javascript; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		var tree = TreeJson{
			Message: "success",
			Data: TreeData{
				Tree:  *tree,
				Depth: depth,
				Width: width,
			},
		}
		var treeStr, _ = json.Marshal(tree)
		w.Write(treeStr)
	})

	var host = "0.0.0.0"
	// 监听并在 0.0.0.0:xx 上启动服务
	li, err := net.Listen("tcp", net.JoinHostPort(host, strconv.Itoa(port)))
	if nil != err {
		fmt.Printf("gmchart server listen err(%v)\n", err)
	}

	mux.HandleFunc("/an-tree.json", func (w http.ResponseWriter, r *http.Request) {
		var header = w.Header()
		header.Add("Content-type", "text/javascript; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		var tree = AnTreeJson{
			Message: "success",
			Data: AnTreeData{
				Tree: anotherTree,
			},
		}
		var treeStr, _ = json.Marshal(tree)
		w.Write(treeStr)
	})

	server := &http.Server{
		Handler: mux,
	}
	addr := li.Addr().String()
	var addrs = strings.Split(addr, ":")
	var printAddr = "http://127.0.0.1:" + addrs[len(addrs)-1] // 这里127.0.0.1 写死，兼容win
	if strings.ToLower(mode) == "graph" {
		printAddr += "/#graph"
	}

	go func() error {
		// open it by default browser
		return src.OpenBrowser(printAddr)
	}()

	if 1 > keep {
		go func() error {
			fmt.Printf("the go mod graph will top in 60s\nvisit it by %s\n", printAddr)
			time.Sleep(60 * time.Second)
			os.Exit(0)
			return nil
		}()
	}

	err = server.Serve(li)
	if err != nil {
		fmt.Printf("gmchart server start err(%v)\n", err)
	}
}
