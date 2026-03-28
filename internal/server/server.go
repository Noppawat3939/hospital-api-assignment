package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

func New(db *gorm.DB) *Server {
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	r.RedirectTrailingSlash = true
	r.Use(gin.Logger(), gin.Recovery(), CORS())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	s := &Server{db: db, router: r}
	s.registerRoutes()

	return s
}

func (s *Server) Start(port string) error {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: s.router,
	}

	fmt.Printf("Starting server on port: %v\n", port)
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("start server error: %w", err)
	}

	return nil
}
