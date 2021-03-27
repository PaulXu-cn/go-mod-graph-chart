package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	dist "github.com/PaulXu-cn/go-mod-graph-chart/godist"
	src "github.com/PaulXu-cn/go-mod-graph-chart/gosrc"
)

var (
	debug int = 0
	keep int = 0
)

func init() {
	flag.IntVar(&debug, "debug", 0, "is debug model")
	flag.IntVar(&keep, "keep", 0, "start http server not exit")
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
	header.Add("content-type","text/json; charset=utf-8")
	var jsonStr = "{\"message\": \"pong\"}"
	w.Write([]byte(jsonStr))
	w.WriteHeader(http.StatusOK)
}

type GraphData struct {
	Nodes []src.Node `json:"nodes""`
	Links []src.Link `json:"links""`
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

func main() {
	flag.Parse()
	var goModGraph string = src.GetGoModGraph()
	var commands = map[string]string{
		"windows": "start",
		"darwin":  "open",
		"linux":   "xdg-open",
	}

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
	tree, depth, width := src.BuildTree(goModGraph)

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

	mux.HandleFunc("/graph.json", func (w http.ResponseWriter, r *http.Request) {
		var header = w.Header()
		header.Add("Content-type", "text/javascript; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		var graph = GraphJson{
			Message: "success",
			Data: GraphData{
				Nodes: nodeSortArr,
				Links: links,
				Num: uint32(len(nodeSortArr)),
			},
		}
		var graphStr, _ = json.Marshal(graph)
		w.Write(graphStr)
	})

	mux.HandleFunc("/tree.json", func (w http.ResponseWriter, r *http.Request) {
		var header = w.Header()
		header.Add("Content-type", "text/javascript; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		var tree = TreeJson{
			Message: "success",
			Data: TreeData{
				Tree: *tree,
				Depth: depth,
				Width: width,
			},
		}
		var treeStr, _ = json.Marshal(tree)
		w.Write(treeStr)
	})

	var host = "0.0.0.0"
	var port = "60306"
	// 监听并在 0.0.0.0:8080 上启动服务
	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: mux,
	}

	go func() error {
		run, ok := commands[runtime.GOOS]
		if !ok {
			return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
		}
		// open it by default browser
		cmd := exec.Command(run, "http://127.0.0.1:"+port)
		return cmd.Start()
	}()

	if 1 > keep {
		go func() error {
			fmt.Printf("the go mod graph will top in 60s\nvisit http://127.0.0.1:%s\n", port)
			time.Sleep(60 * time.Second)
			os.Exit(0)
			return nil
		}()
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("gmchart server start err(%v)\n",  err)
	}
}
