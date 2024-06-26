package main

import (
	"RestGolang/handlers"
	"RestGolang/middleware"
	"RestGolang/server"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	// hub := websocket.NewHub()

	r.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)

	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)

	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/greet", handlers.InsertGreetHandler(s)).Methods(http.MethodPost)

	r.HandleFunc("/greet/{id}", handlers.GetGreetByIdHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/greet/{id}", handlers.UpdateGreetHandler(s)).Methods(http.MethodPut)

	r.HandleFunc("/greet/{id}", handlers.DeleteGreetHandler(s)).Methods(http.MethodDelete)

	r.HandleFunc("/greets", handlers.ListGreetHandler(s)).Methods(http.MethodGet)

	// r.HandleFunc("/ws", hub.HandleWebSocket)

}
