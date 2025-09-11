package handlers

import (
	"encoding/json"
	"food-delivery-backend/internal/database"
	"food-delivery-backend/internal/models"
	"log"
	"net/http"
	"time"
)

type DashboardHandler struct {
	db *database.DB
}

func NewDashboardHandler(db *database.DB) *DashboardHandler {
	return &DashboardHandler{
		db: db,
	}
}

// GetDashboard returns dashboard data for an owner
func (h *DashboardHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Get owner_id from query parameter
	ownerID := r.URL.Query().Get("owner_id")
	if ownerID == "" {
		http.Error(w, "owner_id parameter required", http.StatusBadRequest)
		return
	}

	// Get merchant summary data with total sales from all orders
	// For simplicity, we calculate revenue as: (total menu items general_price_sale - cost) * number of orders
	merchantQuery := `
		SELECT 
			m.merchant_id,
			m.merchant_name,
			COUNT(DISTINCT o.order_id) as order_count,
			COALESCE(SUM(mi.general_price_sale), 0) as total_price_sale,
			COALESCE(SUM(mi.cost), 0) as total_cost
		FROM merchants m
		LEFT JOIN orders o ON m.merchant_id = o.merchant_id
		LEFT JOIN menu_items mi ON m.merchant_id = mi.merchant_id
		WHERE m.owner_id = ?
		GROUP BY m.merchant_id, m.merchant_name
	`

	rows, err := h.db.Query(merchantQuery, ownerID)
	if err != nil {
		log.Printf("Error fetching merchant data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var merchants []models.MerchantSummary
	var totalIncome, totalOutcome, totalRevenue float64

	for rows.Next() {
		var merchant models.MerchantSummary
		var orderCount int
		var totalPriceSale, totalCost float64
		
		err := rows.Scan(
			&merchant.MerchantID,
			&merchant.MerchantName,
			&orderCount,
			&totalPriceSale,
			&totalCost,
		)
		if err != nil {
			log.Printf("Error scanning merchant row: %v", err)
			continue
		}

		// Calculate actual revenue based on orders
		// Income = sum of all menu items general_price_sale * number of orders
		// Outcome = sum of all menu items cost * number of orders
		if orderCount > 0 {
			merchant.Income = totalPriceSale * float64(orderCount)
			merchant.Outcome = totalCost * float64(orderCount)
			merchant.Revenue = merchant.Income - merchant.Outcome
		} else {
			merchant.Income = 0
			merchant.Outcome = 0
			merchant.Revenue = 0
		}

		// Add mock change percentages for demo
		merchant.IncomeChange = 12.5
		merchant.OutcomeChange = 5.2
		merchant.RevenueChange = 18.7
		merchant.Category = "Restaurant"

		merchants = append(merchants, merchant)
		totalIncome += merchant.Income
		totalOutcome += merchant.Outcome
		totalRevenue += merchant.Revenue
	}

	// Get first order date for this owner
	var firstOrderDate *time.Time
	firstOrderQuery := `
		SELECT MIN(o.created_at) 
		FROM orders o 
		JOIN merchants m ON o.merchant_id = m.merchant_id 
		WHERE m.owner_id = ?
	`
	row := h.db.QueryRow(firstOrderQuery, ownerID)
	var tempDate time.Time
	if err := row.Scan(&tempDate); err == nil {
		firstOrderDate = &tempDate
	}

	// Create dashboard summary
	dashboard := models.DashboardSummary{
		OwnerID:        ownerID,
		TotalMerchants: len(merchants),
		TotalRevenue:   totalRevenue,
		TotalIncome:    totalIncome,
		TotalOutcome:   totalOutcome,
		Merchants:      merchants,
		FirstOrderDate: firstOrderDate,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dashboard)
}

// GetMerchantDetails returns detailed data for a specific merchant
func (h *DashboardHandler) GetMerchantDetails(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	merchantID := r.URL.Query().Get("merchant_id")
	if merchantID == "" {
		http.Error(w, "merchant_id parameter required", http.StatusBadRequest)
		return
	}

	// Get merchant with total revenue calculation
	query := `
		SELECT 
			m.merchant_id,
			m.merchant_name,
			COUNT(DISTINCT o.order_id) as order_count,
			COALESCE(SUM(mi.general_price_sale), 0) as total_price_sale,
			COALESCE(SUM(mi.cost), 0) as total_cost
		FROM merchants m
		LEFT JOIN orders o ON m.merchant_id = o.merchant_id
		LEFT JOIN menu_items mi ON m.merchant_id = mi.merchant_id
		WHERE m.merchant_id = ?
		GROUP BY m.merchant_id, m.merchant_name
	`

	var merchant models.MerchantSummary
	var orderCount int
	var totalPriceSale, totalCost float64
	
	row := h.db.QueryRow(query, merchantID)
	err := row.Scan(
		&merchant.MerchantID, 
		&merchant.MerchantName, 
		&orderCount,
		&totalPriceSale,
		&totalCost,
	)
	
	if err != nil {
		log.Printf("Error fetching merchant: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Calculate total revenue (not average)
	if orderCount > 0 {
		merchant.Income = totalPriceSale * float64(orderCount)
		merchant.Outcome = totalCost * float64(orderCount)
		merchant.Revenue = merchant.Income - merchant.Outcome
	} else {
		merchant.Income = 0
		merchant.Outcome = 0
		merchant.Revenue = 0
	}

	merchant.IncomeChange = 15.3
	merchant.OutcomeChange = 4.2
	merchant.RevenueChange = 22.5
	merchant.Category = "Restaurant"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(merchant)
}