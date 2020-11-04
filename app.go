package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	getdata "v2rss/getdata"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	var (
		port   string
		config string
	)
	flag.StringVar(&port, "p", "3000", "端口号，默认为3000")
	flag.StringVar(&config, "c", "", "配置文件，默认空")
	flag.Parse()
	app := gin.New()
	app.GET("/", func(c *gin.Context) {
		var x bool = false
		var y int = 0
		var n string = c.DefaultQuery("n", "1")
		var w string = c.DefaultQuery("w", "0")
		var i string = c.DefaultQuery("i", "0")
		if w == "1" {
			x = true
		}
		y, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("err")
		}
		// fmt.Println(x, n, w)
		var data string = getdata.Start(n, x, y, config)
		c.String(200, data)
	})
	app.Run(strings.Join([]string{":", port}, ""))
}
