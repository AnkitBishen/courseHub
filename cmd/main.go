package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/AnkitBishen/courseHub/internal/cros"
	"github.com/AnkitBishen/courseHub/internal/handler/auth"
	mysqldb "github.com/AnkitBishen/courseHub/internal/storage/mySqlDb"
	"github.com/joho/godotenv"
)

func main() {

	// configuration stuffs
	envLoadErr := godotenv.Load(".env")
	if envLoadErr != nil {
		panic(envLoadErr.Error())
	}

	// connect to database
	_, dbErr := mysqldb.New()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	// Start the web server
	router := http.NewServeMux()

	// manage routes
	router.HandleFunc("POST /api/auth/register", auth.Register())

	// handle cros error
	crosRouter := cros.HandleCros(router)

	slog.Info("Server started", "version", os.Getenv("Version"), "port", os.Getenv("HttpServerPort"))
	err := http.ListenAndServe(os.Getenv("HttpServerPort"), crosRouter)
	if err != nil {
		panic(err.Error())
	}
}
