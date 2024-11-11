package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("views/index.html", "views/episodes.html")
	router.Static("/views", "./views")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/episodes", func(c *gin.Context) {
		c.HTML(http.StatusOK, "episodes.html", nil)
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
