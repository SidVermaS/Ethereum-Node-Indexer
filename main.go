package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/modules"
	"github.com/gofiber/fiber/v2"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"data": "hello "})
}

func main() {
	modules.InitializeAll()
	app := fiber.New()

	http.HandleFunc("/", GreetHandler)
	PORT := fmt.Sprintf("%s", os.Getenv(string(consts.API_PORT)))

	log.Printf("Server is running on PORT: %s...\n", PORT)
	err := app.Listen(":" + PORT)
	if err != nil {
		panic(err)
	}

}
