// @title Go E-commerce API
// @version 1.0
// @description Authentication and E-commerce Backend built with Go
// @host localhost:8080
// @BasePath /
package main
import (
	"github.com/gorilla/mux"
	"net/http"
	"log" 
	"github.com/joho/godotenv"
	"github.com/MoyosoreCoder/go-ecommerce-api/middleware"
	"github.com/MoyosoreCoder/go-ecommerce-api/database"
	"github.com/MoyosoreCoder/go-ecommerce-api/handlers"
	_ "github.com/MoyosoreCoder/go-ecommerce-api/docs"
	// httpSwagger "github.com/swaggo/http-swagger" 
)

func main() {
        err := godotenv.Load()
        if err != nil {
                log.Println(".env not found")
        }
        database.ConnectDatabase()
        defer database.DB.Close()

        r := mux.NewRouter()

        r.HandleFunc("/register", handlers.RegisterUserHandler)
        r.HandleFunc("/login", handlers.LoginHandler)


        r.Handle(
        "/profile",
        middleware.AuthMiddleware(http.HandlerFunc(handlers.ProfileHandler)),
)

        log.Println("Server running on :8080")
        log.Fatal(http.ListenAndServe(":8080", r))
}
	