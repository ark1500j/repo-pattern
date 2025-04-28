package main

import (
	"fmt"
	"log"

	// "net/http"
	"github.com/ark1500j/repo-pattern/initializers"
	h "github.com/ark1500j/repo-pattern/internal/handlers"
	rp "github.com/ark1500j/repo-pattern/internal/repositories"
	s "github.com/ark1500j/repo-pattern/internal/services"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	fmt.Println("hello world")
	repo := rp.NewRepository(initializers.DB)
	service := s.NewUserService(repo)

	handler := h.NewUserHandler(service)
	r := gin.Default()

	r.GET("/users", handler.GetUsers)
	r.GET("/users/:id", handler.GetUser)
	r.POST("/users", handler.CreateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	log.Fatal(r.Run())

}
