package handlers

import (
	"encoding/json"
	"food-delivery-backend/internal/database"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type WeeklyTrackingHandler struct {
	db *database.DB
}

func NewWeeklyTrackingHandler(db *database.DB) *WeeklyTrackingHandler {
	return &WeeklyTrackingHandler{db: db}
}

type WeeklyTracking struct {
	TrackingID    string    `json:"tracking_id"`
	MerchantID    string    `json:"merchant_id"`
	WeekLabel     string    `json:"week_label"`
	DateRange     string    `json:"date_range"`
	StartDate     string    `json:"start_date"`
	EndDate       string    `json:"end_date"`
	TotalIncome   float64   `json:"total_income"`
	TotalOutcome  float64   `json:"total_outcome"`
	NetRevenue    float64   `json:"net_revenue"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

type CreateWeeklyTrackingRequest struct {
	MerchantID string `json:"merchant_id"`
	WeekLabel  string `json:"week_label"`
	DateRange  string `json:"date_range"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

type IncomeTransaction struct {
	TransactionID   string    `json:"transaction_id"`
	TrackingID      string    `json:"tracking_id"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	Category        string    `json:"category"`
	TransactionDate string    `json:"transaction_date"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

type OutcomeTransaction struct {
	TransactionID   string    `json:"transaction_id"`
	TrackingID      string    `json:"tracking_id"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	Category        string    `json:"category"`
	TransactionDate string    `json:"transaction_date"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

// CreateWeeklyTracking handles POST /api/v1/weekly-tracking
func (h *WeeklyTrackingHandler) CreateWeeklyTracking(w http.ResponseWriter, r *http.Request) {
	log.Printf("📅 Creating weekly tracking")

	var req CreateWeeklyTrackingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("❌ Failed to parse request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.MerchantID == "" || req.StartDate == "" || req.EndDate == "" {
		log.Printf("❌ Missing required fields")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Check if tracking already exists for this period
	checkQuery := `
		SELECT tracking_id FROM weekly_trackings 
		WHERE merchant_id = ? AND start_date = ? AND end_date = ?
	`
	var existingID string
	err := h.db.QueryRow(checkQuery, req.MerchantID, req.StartDate, req.EndDate).Scan(&existingID)
	if err == nil {
		// Tracking already exists
		log.Printf("⚠️ Weekly tracking already exists: %s", existingID)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"tracking_id": existingID,
			"message":     "Weekly tracking already exists for this period",
			"exists":      true,
		})
		return
	}

	// Generate UUID for tracking
	trackingID := uuid.New().String()

	// Insert weekly tracking
	query := `
		INSERT INTO weekly_trackings (
			tracking_id, merchant_id, week_label, date_range, 
			start_date, end_date, total_income, total_outcome, net_revenue
		) VALUES (?, ?, ?, ?, ?, ?, 0, 0, 0)
	`
	_, err = h.db.Exec(query, 
		trackingID, req.MerchantID, req.WeekLabel, req.DateRange,
		req.StartDate, req.EndDate,
	)
	if err != nil {
		log.Printf("❌ Failed to create weekly tracking: %v", err)
		http.Error(w, "Failed to create weekly tracking", http.StatusInternalServerError)
		return
	}

	response := WeeklyTracking{
		TrackingID:   trackingID,
		MerchantID:   req.MerchantID,
		WeekLabel:    req.WeekLabel,
		DateRange:    req.DateRange,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		TotalIncome:  0,
		TotalOutcome: 0,
		NetRevenue:   0,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("✅ Weekly tracking created: %s", trackingID)
}

// GetWeeklyTrackings handles GET /api/v1/weekly-tracking?merchant_id=xxx
func (h *WeeklyTrackingHandler) GetWeeklyTrackings(w http.ResponseWriter, r *http.Request) {
	merchantID := r.URL.Query().Get("merchant_id")
	if merchantID == "" {
		http.Error(w, "merchant_id is required", http.StatusBadRequest)
		return
	}

	log.Printf("📋 Getting weekly trackings for merchant: %s", merchantID)

	query := `
		SELECT tracking_id, merchant_id, week_label, date_range, 
			   start_date, end_date, total_income, total_outcome, net_revenue, created_at
		FROM weekly_trackings
		WHERE merchant_id = ?
		ORDER BY start_date DESC
	`

	rows, err := h.db.Query(query, merchantID)
	if err != nil {
		log.Printf("❌ Failed to get weekly trackings: %v", err)
		http.Error(w, "Failed to retrieve weekly trackings", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var trackings []WeeklyTracking
	for rows.Next() {
		var tracking WeeklyTracking
		err := rows.Scan(
			&tracking.TrackingID, &tracking.MerchantID, &tracking.WeekLabel,
			&tracking.DateRange, &tracking.StartDate, &tracking.EndDate,
			&tracking.TotalIncome, &tracking.TotalOutcome, &tracking.NetRevenue,
			&tracking.CreatedAt,
		)
		if err != nil {
			log.Printf("❌ Failed to scan tracking: %v", err)
			continue
		}
		trackings = append(trackings, tracking)
	}

	// Return empty array if no trackings found
	if trackings == nil {
		trackings = []WeeklyTracking{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trackings)
}

// GetWeeklyTrackingDetails handles GET /api/v1/weekly-tracking/{id}
func (h *WeeklyTrackingHandler) GetWeeklyTrackingDetails(w http.ResponseWriter, r *http.Request) {
	// Extract tracking ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	trackingID := pathParts[4]

	log.Printf("📊 Getting weekly tracking details: %s", trackingID)

	// Get tracking info
	var tracking WeeklyTracking
	query := `
		SELECT tracking_id, merchant_id, week_label, date_range, 
			   start_date, end_date, total_income, total_outcome, net_revenue, created_at
		FROM weekly_trackings
		WHERE tracking_id = ?
	`
	err := h.db.QueryRow(query, trackingID).Scan(
		&tracking.TrackingID, &tracking.MerchantID, &tracking.WeekLabel,
		&tracking.DateRange, &tracking.StartDate, &tracking.EndDate,
		&tracking.TotalIncome, &tracking.TotalOutcome, &tracking.NetRevenue,
		&tracking.CreatedAt,
	)
	if err != nil {
		log.Printf("❌ Weekly tracking not found: %v", err)
		http.Error(w, "Weekly tracking not found", http.StatusNotFound)
		return
	}

	// Get income transactions
	incomeQuery := `
		SELECT transaction_id, description, amount, category, transaction_date
		FROM income_transactions
		WHERE tracking_id = ?
		ORDER BY transaction_date DESC
	`
	incomeRows, err := h.db.Query(incomeQuery, trackingID)
	if err != nil {
		log.Printf("⚠️ Failed to get income transactions: %v", err)
	}
	defer incomeRows.Close()

	var incomeTransactions []IncomeTransaction
	for incomeRows.Next() {
		var transaction IncomeTransaction
		err := incomeRows.Scan(
			&transaction.TransactionID, &transaction.Description,
			&transaction.Amount, &transaction.Category, &transaction.TransactionDate,
		)
		if err != nil {
			continue
		}
		transaction.TrackingID = trackingID
		incomeTransactions = append(incomeTransactions, transaction)
	}

	// Get outcome transactions
	outcomeQuery := `
		SELECT transaction_id, description, amount, category, transaction_date
		FROM outcome_transactions
		WHERE tracking_id = ?
		ORDER BY transaction_date DESC
	`
	outcomeRows, err := h.db.Query(outcomeQuery, trackingID)
	if err != nil {
		log.Printf("⚠️ Failed to get outcome transactions: %v", err)
	}
	defer outcomeRows.Close()

	var outcomeTransactions []OutcomeTransaction
	for outcomeRows.Next() {
		var transaction OutcomeTransaction
		err := outcomeRows.Scan(
			&transaction.TransactionID, &transaction.Description,
			&transaction.Amount, &transaction.Category, &transaction.TransactionDate,
		)
		if err != nil {
			continue
		}
		transaction.TrackingID = trackingID
		outcomeTransactions = append(outcomeTransactions, transaction)
	}

	// Build response
	response := map[string]interface{}{
		"tracking":     tracking,
		"income_items": incomeTransactions,
		"outcome_items": outcomeTransactions,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// AddIncomeTransaction handles POST /api/v1/weekly-tracking/{id}/income
func (h *WeeklyTrackingHandler) AddIncomeTransaction(w http.ResponseWriter, r *http.Request) {
	// Extract tracking ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 6 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	trackingID := pathParts[4]

	log.Printf("💰 Adding income transaction to tracking: %s", trackingID)

	var req struct {
		Description     string  `json:"description"`
		Amount          float64 `json:"amount"`
		Category        string  `json:"category"`
		TransactionDate string  `json:"transaction_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate UUID for transaction
	transactionID := uuid.New().String()

	// Insert transaction
	query := `
		INSERT INTO income_transactions (
			transaction_id, tracking_id, description, amount, category, transaction_date
		) VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := h.db.Exec(query, transactionID, trackingID, req.Description, req.Amount, req.Category, req.TransactionDate)
	if err != nil {
		log.Printf("❌ Failed to add income transaction: %v", err)
		http.Error(w, "Failed to add income transaction", http.StatusInternalServerError)
		return
	}

	// Update total income in weekly_trackings
	updateQuery := `
		UPDATE weekly_trackings 
		SET total_income = total_income + ?,
			net_revenue = total_income + ? - total_outcome
		WHERE tracking_id = ?
	`
	_, err = h.db.Exec(updateQuery, req.Amount, req.Amount, trackingID)
	if err != nil {
		log.Printf("⚠️ Failed to update totals: %v", err)
	}

	response := IncomeTransaction{
		TransactionID:   transactionID,
		TrackingID:      trackingID,
		Description:     req.Description,
		Amount:          req.Amount,
		Category:        req.Category,
		TransactionDate: req.TransactionDate,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("✅ Income transaction added: %s", transactionID)
}

// AddOutcomeTransaction handles POST /api/v1/weekly-tracking/{id}/outcome
func (h *WeeklyTrackingHandler) AddOutcomeTransaction(w http.ResponseWriter, r *http.Request) {
	// Extract tracking ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 6 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	trackingID := pathParts[4]

	log.Printf("💸 Adding outcome transaction to tracking: %s", trackingID)

	var req struct {
		Description     string  `json:"description"`
		Amount          float64 `json:"amount"`
		Category        string  `json:"category"`
		TransactionDate string  `json:"transaction_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate UUID for transaction
	transactionID := uuid.New().String()

	// Insert transaction
	query := `
		INSERT INTO outcome_transactions (
			transaction_id, tracking_id, description, amount, category, transaction_date
		) VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := h.db.Exec(query, transactionID, trackingID, req.Description, req.Amount, req.Category, req.TransactionDate)
	if err != nil {
		log.Printf("❌ Failed to add outcome transaction: %v", err)
		http.Error(w, "Failed to add outcome transaction", http.StatusInternalServerError)
		return
	}

	// Update total outcome in weekly_trackings
	updateQuery := `
		UPDATE weekly_trackings 
		SET total_outcome = total_outcome + ?,
			net_revenue = total_income - (total_outcome + ?)
		WHERE tracking_id = ?
	`
	_, err = h.db.Exec(updateQuery, req.Amount, req.Amount, trackingID)
	if err != nil {
		log.Printf("⚠️ Failed to update totals: %v", err)
	}

	response := OutcomeTransaction{
		TransactionID:   transactionID,
		TrackingID:      trackingID,
		Description:     req.Description,
		Amount:          req.Amount,
		Category:        req.Category,
		TransactionDate: req.TransactionDate,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("✅ Outcome transaction added: %s", transactionID)
}

// DeleteWeeklyTracking handles DELETE /api/v1/weekly-tracking/{id}
func (h *WeeklyTrackingHandler) DeleteWeeklyTracking(w http.ResponseWriter, r *http.Request) {
	// Extract tracking ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	trackingID := pathParts[4]

	log.Printf("🗑️ Deleting weekly tracking: %s", trackingID)

	// Start a transaction
	tx, err := h.db.Begin()
	if err != nil {
		log.Printf("❌ Failed to start transaction: %v", err)
		http.Error(w, "Failed to delete weekly tracking", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Delete income transactions first (due to foreign key constraints)
	_, err = tx.Exec("DELETE FROM income_transactions WHERE tracking_id = ?", trackingID)
	if err != nil {
		log.Printf("❌ Failed to delete income transactions: %v", err)
		http.Error(w, "Failed to delete weekly tracking", http.StatusInternalServerError)
		return
	}

	// Delete outcome transactions
	_, err = tx.Exec("DELETE FROM outcome_transactions WHERE tracking_id = ?", trackingID)
	if err != nil {
		log.Printf("❌ Failed to delete outcome transactions: %v", err)
		http.Error(w, "Failed to delete weekly tracking", http.StatusInternalServerError)
		return
	}

	// Delete the weekly tracking itself
	result, err := tx.Exec("DELETE FROM weekly_trackings WHERE tracking_id = ?", trackingID)
	if err != nil {
		log.Printf("❌ Failed to delete weekly tracking: %v", err)
		http.Error(w, "Failed to delete weekly tracking", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Weekly tracking not found", http.StatusNotFound)
		return
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		log.Printf("❌ Failed to commit transaction: %v", err)
		http.Error(w, "Failed to delete weekly tracking", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Weekly tracking deleted successfully"})
	log.Printf("✅ Weekly tracking deleted: %s", trackingID)
}