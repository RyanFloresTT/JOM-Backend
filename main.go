package main

import (
	"go-backend/controllers"
	"go-backend/initializers"
	"go-backend/middleware"
	"log"
	"net/http"
	"os"

	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/stripe/stripe-go/v79"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
	LoadProducts()
	UpdateProductPrices()
}

func main() {
	r := chi.NewRouter()

	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://localhost:3000", "http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	}).Handler(r)

	stripe.Key = os.Getenv("StripeAPIKey")

	r.Group(func(public chi.Router) {
		public.Get("/api/products", getProducts)
		public.Post("/create-checkout-session", createCheckoutSession)
		public.Post("/signup", controllers.Signup)
		public.Post("/login", controllers.Login)
		public.Post("/logout", controllers.Logout)
	})

	r.Group(func(protected chi.Router) {
		protected.Use(middleware.RequireAuth)
		protected.Get("/validate", controllers.Validate)
	})
	addr := "localhost:4242"

	log.Printf("Listening on http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
