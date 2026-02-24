//package main
package main
import (
	"net/http"
	"log"
	"github.com/MoyosoreCoder/go-ecommerce-api/internal/handler"
)
func main(){
	// Register route
	http.HandleFunc("/regiter", handler.RegisterUserHandler)

	// Start server
	log.Println("Server running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}