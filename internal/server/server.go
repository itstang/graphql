package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server exposes HTTP endpoints
type Server struct {
	HTTP *http.Server
}

// NewServer allocates and creates a new Server
func NewServer() (*Server, error) {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return &Server{
		HTTP: srv,
	}, nil
}
