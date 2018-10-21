package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/itstang/graphql/resolver"
	"github.com/itstang/graphql/schema"
)

// Server exposes HTTP endpoints
type Server struct {
	HTTP *http.Server
}

// NewServer allocates and creates a new Server
func NewServer() (*Server, error) {
	root, err := resolver.NewRoot()
	if err != nil {
		return nil, err
	}

	r := gin.Default()

	// GraphQL Explorer
	r.LoadHTMLGlob("internal/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "graphiql.tmpl", gin.H{
			"title": "GraphQL Explorer",
		})
	})

	// Main GraphQL Route
	r.POST("/gql", gin.WrapH(&relay.Handler{Schema: graphql.MustParseSchema(schema.String(), root)}))

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
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
