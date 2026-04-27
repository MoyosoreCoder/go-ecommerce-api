// @title Go E-commerce API
// @version 1.0
// @description Authentication and E-commerce Backend built with Go
// @host localhost:8080
// @BasePath /
package main
import (
	"net/http"
	"log" 
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/database"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/handler"
	_ "github.com/MoyosoreCoder/go-ecommerce-api/docs"
	httpSwagger "github.com/swaggo/http-swagger" 
)

func main(){
	database.Connect()
	// Register route
	http.HandleFunc("/register", handler.RegisterUserHandler)
	//login route
	http.HandleFunc("/login", handler.LoginUserHandler)
	// Swagger route
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Start server
	log.Println("Server running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
 		log.Fatal(err)
	}
}