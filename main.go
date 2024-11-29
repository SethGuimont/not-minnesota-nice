package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/views", "./views")
	router.StaticFS("/media", http.Dir("media"))
	router.LoadHTMLGlob("views/*.tmpl")

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowMethods = []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"}
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Home",
		})
	})

	router.GET("/episodes", func(c *gin.Context) {
		c.HTML(http.StatusOK, "episodes.tmpl", gin.H{
			"title": "Episodes",
		})
	})

	router.GET("/latest", func(c *gin.Context) {
		c.HTML(http.StatusOK, "latest.tmpl", gin.H{
			"title": "Latest Video",
		})
	})

	router.GET("/stream/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		file, err := os.Open("videos/" + filename)
		if err != nil {
			c.String(http.StatusNotFound, "Video not found.")
			return
		}
		defer file.Close()

		c.Header("Context-Type", "video/mp4")
		buffer := make([]byte, 64*1024) // 64kb buffer size
		io.CopyBuffer(c.Writer, file, buffer)
	})

	router.Run(":8080")
}
