package server 

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/aabdullahgungor/mybookcase/router"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port: "8000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := router.ConfigRoutes(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run("localhost:" + s.port))
}