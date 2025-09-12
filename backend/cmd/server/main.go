package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"food-delivery-backend/internal/database"
	"food-delivery-backend/internal/handlers"
	"food-delivery-backend/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.NewConnection(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	router := mux.NewRouter()
	
	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "database": "connected"}`))
	}).Methods("GET")

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db)
	dashboardHandler := handlers.NewDashboardHandler(db)
	merchantHandler := handlers.NewMerchantHandler(db)
	menuItemHandler := handlers.NewMenuItemHandler(db)
	extraItemHandler := handlers.NewExtraItemHandler(db)
	dailyTrackingHandler := handlers.NewDailyTrackingHandler(db)
	
	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()
	
	// CORS middleware for development
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// log.Printf("📡 Incoming request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
			
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			
			if r.Method == "OPTIONS" {
				log.Printf("✅ CORS preflight handled for %s", r.URL.Path)
				w.WriteHeader(http.StatusOK)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	})

	// Authentication routes with OPTIONS support for CORS
	api.HandleFunc("/auth/signin", authHandler.SignIn).Methods("POST", "OPTIONS")
	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST", "OPTIONS")
	api.HandleFunc("/auth/verify-otp", authHandler.VerifyOTP).Methods("POST", "OPTIONS")
	api.HandleFunc("/auth/check-phone", authHandler.CheckPhone).Methods("GET", "OPTIONS")

	// Dashboard routes
	api.HandleFunc("/dashboard", dashboardHandler.GetDashboard).Methods("GET", "OPTIONS")
	api.HandleFunc("/dashboard/merchant", dashboardHandler.GetMerchantDetails).Methods("GET", "OPTIONS")

	// Merchant routes
	api.HandleFunc("/merchants", merchantHandler.CreateMerchant).Methods("POST", "OPTIONS")
	api.HandleFunc("/merchants", merchantHandler.GetMerchantsByOwner).Methods("GET", "OPTIONS")
	api.HandleFunc("/merchants/{id}", merchantHandler.DeleteMerchant).Methods("DELETE", "OPTIONS")

	// Menu Items routes
	api.HandleFunc("/menu-items", menuItemHandler.CreateMenuItem).Methods("POST", "OPTIONS")
	api.HandleFunc("/menu-items", menuItemHandler.GetMenuItemsByMerchant).Methods("GET", "OPTIONS")
	api.HandleFunc("/menu-items/{id}", menuItemHandler.UpdateMenuItem).Methods("PUT", "OPTIONS")
	api.HandleFunc("/menu-items/{id}", menuItemHandler.DeleteMenuItem).Methods("DELETE", "OPTIONS")

	// Extra Items routes
	api.HandleFunc("/extra-items", extraItemHandler.CreateExtraItem).Methods("POST", "OPTIONS")
	api.HandleFunc("/extra-items", extraItemHandler.GetExtraItemsByMerchant).Methods("GET", "OPTIONS")
	api.HandleFunc("/extra-items/{id}", extraItemHandler.UpdateExtraItem).Methods("PUT", "OPTIONS")
	api.HandleFunc("/extra-items/{id}", extraItemHandler.DeleteExtraItem).Methods("DELETE", "OPTIONS")

	// Daily Tracking routes
	api.HandleFunc("/daily-tracking", dailyTrackingHandler.CreateDailyTracking).Methods("POST", "OPTIONS")
	api.HandleFunc("/daily-tracking", dailyTrackingHandler.GetDailyTrackings).Methods("GET", "OPTIONS")
	api.HandleFunc("/daily-tracking/{id}", dailyTrackingHandler.GetDailyTrackingDetails).Methods("GET", "OPTIONS")
	api.HandleFunc("/daily-tracking/{id}", dailyTrackingHandler.DeleteDailyTracking).Methods("DELETE", "OPTIONS")
	api.HandleFunc("/daily-tracking/{id}/income", dailyTrackingHandler.AddIncomeTransaction).Methods("POST", "OPTIONS")
	api.HandleFunc("/daily-tracking/{id}/outcome", dailyTrackingHandler.AddOutcomeTransaction).Methods("POST", "OPTIONS")
	api.HandleFunc("/daily-tracking/{id}/income/{transaction_id}", dailyTrackingHandler.DeleteIncomeTransaction).Methods("DELETE", "OPTIONS")
	api.HandleFunc("/daily-tracking/{id}/outcome/{transaction_id}", dailyTrackingHandler.DeleteOutcomeTransaction).Methods("DELETE", "OPTIONS")
	api.HandleFunc("/daily-tracking/weekly-summary", dailyTrackingHandler.GetWeeklyAggregatedData).Methods("GET", "OPTIONS")

	// Placeholder routes
	api.HandleFunc("/restaurants", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Restaurants endpoint - coming soon"}`))
	}).Methods("GET")

	api.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Orders endpoint - coming soon"}`))
	}).Methods("GET")

	api.HandleFunc("/revenue", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Revenue report endpoint - coming soon"}`))
	}).Methods("GET")

	log.Printf("Server starting on port %s", cfg.ServerPort)
	log.Printf("Health check available at: http://localhost:%s/health", cfg.ServerPort)
	log.Printf("API endpoints available at: http://localhost:%s/api/v1", cfg.ServerPort)
	
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}