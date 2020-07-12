package main

import (
	"flag"
	"strings"
	getdata "v2rss/getdata"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	var port string
	flag.StringVar(&port, "p", "3000", "端口号，默认为3000")
	flag.Parse()
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		var x bool = false
		var y bool = true
		var n string = c.DefaultQuery("n", "1")
		var w string = c.DefaultQuery("w", "0")
		var i string = c.DefaultQuery("i", "0")
		if w == "1" {
			x = true
		}
		if i == "1" {
			y = false
		}
		// fmt.Println(x, n, w)
		var data string = getdata.Start(n, x, y)
		c.String(200, data)
	})
	app.Run(strings.Join([]string{":", port}, ""))
}
