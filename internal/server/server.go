package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/gwuah/tinderclone/internal/handlers"
)

type Server struct {
	h          *handlers.Handler
	e          *gin.Engine
	srv        http.Server
	middleware []gin.HandlerFunc
}

func New(h *handlers.Handler, middleware ...gin.HandlerFunc) *Server {
	return &Server{
		h:          h,
		e:          gin.Default(),
		middleware : middleware,
	}
}

func (s *Server) setupMiddlewares(m []gin.HandlerFunc) {
	s.e.Use(m...)
}

func (s *Server) setupRoutes() {
	s.e.GET("/healthCheck", s.h.HealthCheck)
	s.e.POST("/createAccount", s.h.CreateAccount)
}

func (s *Server) Start() {
	s.setupMiddlewares(s.middleware)

	s.setupRoutes()

	s.srv = http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: s.e,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := s.srv.Close(); err != nil {
			log.Println("failed to shutdown server", err)
		}
	}()

	if err := s.srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("server closed after interruption")
		} else {
			log.Println("unexpected server shutdown. err:", err)
		}
	}
}

func (s *Server) Stop() error {
	return s.srv.Close()
}
