package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		log.Fatal("PORT not found in local")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))
	  v1Router := chi.NewRouter()
	  v1Router.Get("/healthz", HandlerReadiness)
	  v1Router.Get("/error", HandlerErr)
	  router.Mount("/v1", v1Router);
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portNumber,
	}
	log.Printf("[+] server runing...... (PORT number:%s)", portNumber)
	if err := srv.ListenAndServe();err != nil{
		log.Fatal(err)
	}
	fmt.Println("PORT:", portNumber)
}
