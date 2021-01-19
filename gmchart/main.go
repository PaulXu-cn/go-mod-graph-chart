package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"

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

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-type", "text/html; charset=utf-8")
		c.String(200, dist.GetFile("index.html"))
	})
	r.GET("/main.js", func(c *gin.Context) {
		c.Header("Content-type", "text/javascript; charset=utf-8")
		c.String(200, dist.GetFile("main.js"))
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/graph.json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
			"data": struct {
				Nodes []src.Node `json:"nodes""`
				Links []src.Link `json:"links""`
				Num uint32 `json:"num"`
			}{
				Nodes: nodeSortArr,
				Links: links,
				Num: uint32(len(nodeSortArr)),
			},
		})
	})
	r.GET("/tree.json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
			"data": struct {
				Tree src.Tree `json:"tree"`
				Depth uint32 `json:"depth"`
				Width uint32 `json:"width"`
			}{
				Tree: *tree,
				Depth: depth,
				Width: width,
			},
		})
	})

	var host = "0.0.0.0"
	var port = "60306"

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

	r.Run(host + ":" + port) // 监听并在 0.0.0.0:8080 上启动服务
}
