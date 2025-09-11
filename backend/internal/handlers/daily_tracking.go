package handlers

import (
	"encoding/json"
	"food-delivery-backend/internal/database"
	"food-delivery-backend/internal/models"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type DailyTrackingHandler struct {
	db *database.DB
}

func NewDailyTrackingHandler(db *database.DB) *DailyTrackingHandler {
	return &DailyTrackingHandler{db: db}
}

type CreateDailyTrackingRequest struct {
	MerchantID   string `json:"merchant_id"`
	TrackingDate string `json:"tracking_date"`
}

// CreateDailyTracking handles POST /api/v1/daily-tracking
func (h *DailyTrackingHandler) CreateDailyTracking(w http.ResponseWriter, r *http.Request) {
	log.Printf("📅 Creating daily tracking")

	var req CreateDailyTrackingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("❌ Failed to parse request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.MerchantID == "" || req.TrackingDate == "" {
		log.Printf("❌ Missing required fields")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Check if tracking already exists for this date
	checkQuery := `
		SELECT tracking_id FROM daily_trackings 
		WHERE merchant_id = ? AND tracking_date = ?
	`
	var existingID string
	err := h.db.QueryRow(checkQuery, req.MerchantID, req.TrackingDate).Scan(&existingID)
	if err == nil {
		// Tracking already exists
		log.Printf("⚠️ Daily tracking already exists: %s", existingID)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"tracking_id": existingID,
			"message":     "Daily tracking already exists for this date",
			"exists":      true,
		})
		return
	}

	// Generate UUID for tracking
	trackingID := uuid.New().String()

	// Insert daily tracking
	query := `
		INSERT INTO daily_trackings (
			tracking_id, merchant_id, tracking_date, 
			total_income, total_outcome, net_revenue
		) VALUES (?, ?, ?, 0, 0, 0)
	`
	_, err = h.db.Exec(query, trackingID, req.MerchantID, req.TrackingDate)
	if err != nil {
		log.Printf("❌ Failed to create daily tracking: %v", err)
		http.Error(w, "Failed to create daily tracking", http.StatusInternalServerError)
		return
	}

	response := models.DailyTracking{
		TrackingID:   trackingID,
		MerchantID:   req.MerchantID,
		TrackingDate: req.TrackingDate,
		TotalIncome:  0,
		TotalOutcome: 0,
		NetRevenue:   0,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("✅ Daily tracking created: %s", trackingID)
}

// GetDailyTrackings handles GET /api/v1/daily-tracking?merchant_id=xxx&start_date=xxx&end_date=xxx
func (h *DailyTrackingHandler) GetDailyTrackings(w http.ResponseWriter, r *http.Request) {
	merchantID := r.URL.Query().Get("merchant_id")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	
	if merchantID == "" {
		http.Error(w, "merchant_id is required", http.StatusBadRequest)
		return
	}

	log.Printf("📋 Getting daily trackings for merchant: %s", merchantID)

	var query string
	var args []interface{}

	if startDate != "" && endDate != "" {
		query = `
			SELECT tracking_id, merchant_id, tracking_date, 
				   total_income, total_outcome, net_revenue, created_at, updated_at
			FROM daily_trackings
			WHERE merchant_id = ? AND tracking_date BETWEEN ? AND ?
			ORDER BY tracking_date DESC
		`
		args = []interface{}{merchantID, startDate, endDate}
	} else {
		query = `
			SELECT tracking_id, merchant_id, tracking_date, 
				   total_income, total_outcome, net_revenue, created_at, updated_at
			FROM daily_trackings
			WHERE merchant_id = ?
			ORDER BY tracking_date DESC
		`
		args = []interface{}{merchantID}
	}

	rows, err := h.db.Query(query, args...)
	if err != nil {
		log.Printf("❌ Failed to get daily trackings: %v", err)
		http.Error(w, "Failed to retrieve daily trackings", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var trackings []models.DailyTracking
	for rows.Next() {
		var tracking models.DailyTracking
		err := rows.Scan(
			&tracking.TrackingID, &tracking.MerchantID, &tracking.TrackingDate,
			&tracking.TotalIncome, &tracking.TotalOutcome, &tracking.NetRevenue,
			&tracking.CreatedAt, &tracking.UpdatedAt,
		)
		if err != nil {
			log.Printf("❌ Failed to scan tracking: %v", err)
			continue
		}
		trackings = append(trackings, tracking)
	}

	// Return empty array if no trackings found
	if trackings == nil {
		trackings = []models.DailyTracking{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trackings)
}

// GetDailyTrackingDetails handles GET /api/v1/daily-tracking/{id}
func (h *DailyTrackingHandler) GetDailyTrackingDetails(w http.ResponseWriter, r *http.Request) {
	// Extract tracking ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	trackingID := pathParts[4]

	log.Printf("📊 Getting daily tracking details: %s", trackingID)

	// Get tracking info
	var tracking models.DailyTracking
	query := `
		SELECT tracking_id, merchant_id, tracking_date, 
			   total_income, total_outcome, net_revenue, created_at, updated_at
		FROM daily_trackings
		WHERE tracking_id = ?
	`
	err := h.db.QueryRow(query, trackingID).Scan(
		&tracking.TrackingID, &tracking.MerchantID, &tracking.TrackingDate,
		&tracking.TotalIncome, &tracking.TotalOutcome, &tracking.NetRevenue,
		&tracking.CreatedAt, &tracking.UpdatedAt,
	)
	if err != nil {
		log.Printf("❌ Daily tracking not found: %v", err)
		http.Error(w, "Daily tracking not found", http.StatusNotFound)
		return
	}

	// Get income transactions
	incomeQuery := `
		SELECT transaction_id, description, amount, category, transaction_date, created_at, updated_at
		FROM income_transactions
		WHERE tracking_id = ?
		ORDER BY transaction_date DESC
	`
	incomeRows, err := h.db.Query(incomeQuery, trackingID)
	if err != nil {
		log.Printf("⚠️ Failed to get income transactions: %v", err)
	}
	defer incomeRows.Close()

	var incomeTransactions []models.IncomeTransaction
	for incomeRows.Next() {
		var transaction models.IncomeTransaction
		err := incomeRows.Scan(
			&transaction.TransactionID, &transaction.Description,
			&transaction.Amount, &transaction.Category, &transaction.TransactionDate,
			&transaction.CreatedAt, &transaction.UpdatedAt,
		)
		if err != nil {
			continue
		}
		transaction.TrackingID = trackingID
		incomeTransactions = append(incomeTransactions, transaction)
	}

	// Get outcome transactions
	outcomeQuery := `
		SELECT transaction_id, description, amount, category, transaction_date, created_at, updated_at
		FROM outcome_transactions
		WHERE tracking_id = ?
		ORDER BY transaction_date DESC
	`
	outcomeRows, err := h.db.Query(outcomeQuery, trackingID)
	if err != nil {
		log.Printf("⚠️ Failed to get outcome transactions: %v", err)
	}
	defer outcomeRows.Close()

	var outcomeTransactions []models.OutcomeTransaction
	for outcomeRows.Next() {
		var transaction models.OutcomeTransaction
		err := outcomeRows.Scan(
			&transaction.TransactionID, &transaction.Description,
			&transaction.Amount, &transaction.Category, &transaction.TransactionDate,
			&transaction.CreatedAt, &transaction.UpdatedAt,
		)
		if err != nil {
			continue
		}
		transaction.TrackingID = trackingID
		outcomeTransactions = append(outcomeTransactions, transaction)
	}

	// Build response
	response := map[string]interface{}{
		"tracking":      tracking,
		"income_items":  incomeTransactions,
		"outcome_items": outcomeTransactions,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// AddIncomeTransaction handles POST /api/v1/daily-tracking/{id}/income
func (h *DailyTrackingHandler) AddIncomeTransaction(w http.ResponseWriter, r *http.Request) {
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

	// Update total income in daily_trackings
	updateQuery := `
		UPDATE daily_trackings 
		SET total_income = total_income + ?,
			net_revenue = total_income + ? - total_outcome
		WHERE tracking_id = ?
	`
	_, err = h.db.Exec(updateQuery, req.Amount, req.Amount, trackingID)
	if err != nil {
		log.Printf("⚠️ Failed to update totals: %v", err)
	}

	response := models.IncomeTransaction{
		TransactionID:   transactionID,
		TrackingID:      trackingID,
		Description:     req.Description,
		Amount:          req.Amount,
		Category:        req.Category,
		TransactionDate: req.TransactionDate,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("✅ Income transaction added: %s", transactionID)
}

// AddOutcomeTransaction handles POST /api/v1/daily-tracking/{id}/outcome
func (h *DailyTrackingHandler) AddOutcomeTransaction(w http.ResponseWriter, r *http.Request) {
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

	// Update total outcome in daily_trackings
	updateQuery := `
		UPDATE daily_trackings 
		SET total_outcome = total_outcome + ?,
			net_revenue = total_income - (total_outcome + ?)
		WHERE tracking_id = ?
	`
	_, err = h.db.Exec(updateQuery, req.Amount, req.Amount, trackingID)
	if err != nil {
		log.Printf("⚠️ Failed to update totals: %v", err)
	}

	response := models.OutcomeTransaction{
		TransactionID:   transactionID,
		TrackingID:      trackingID,
		Description:     req.Description,
		Amount:          req.Amount,
		Category:        req.Category,
		TransactionDate: req.TransactionDate,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("✅ Outcome transaction added: %s", transactionID)
}

// DeleteDailyTracking handles DELETE /api/v1/daily-tracking/{id}
func (h *DailyTrackingHandler) DeleteDailyTracking(w http.ResponseWriter, r *http.Request) {
	// Extract tracking ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	trackingID := pathParts[4]

	log.Printf("🗑️ Deleting daily tracking: %s", trackingID)

	// Start a transaction
	tx, err := h.db.Begin()
	if err != nil {
		log.Printf("❌ Failed to start transaction: %v", err)
		http.Error(w, "Failed to delete daily tracking", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	// Delete income transactions first (due to foreign key constraints)
	_, err = tx.Exec("DELETE FROM income_transactions WHERE tracking_id = ?", trackingID)
	if err != nil {
		log.Printf("❌ Failed to delete income transactions: %v", err)
		http.Error(w, "Failed to delete daily tracking", http.StatusInternalServerError)
		return
	}

	// Delete outcome transactions
	_, err = tx.Exec("DELETE FROM outcome_transactions WHERE tracking_id = ?", trackingID)
	if err != nil {
		log.Printf("❌ Failed to delete outcome transactions: %v", err)
		http.Error(w, "Failed to delete daily tracking", http.StatusInternalServerError)
		return
	}

	// Delete the daily tracking itself
	result, err := tx.Exec("DELETE FROM daily_trackings WHERE tracking_id = ?", trackingID)
	if err != nil {
		log.Printf("❌ Failed to delete daily tracking: %v", err)
		http.Error(w, "Failed to delete daily tracking", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Daily tracking not found", http.StatusNotFound)
		return
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		log.Printf("❌ Failed to commit transaction: %v", err)
		http.Error(w, "Failed to delete daily tracking", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Daily tracking deleted successfully"})
	log.Printf("✅ Daily tracking deleted: %s", trackingID)
}

// DeleteIncomeTransaction handles DELETE /api/v1/daily-tracking/{tracking_id}/income/{transaction_id}
func (h *DailyTrackingHandler) DeleteIncomeTransaction(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 6 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	
	trackingID := pathParts[4]
	transactionID := pathParts[6]
	
	log.Printf("🗑️ Deleting income transaction: %s from tracking: %s", transactionID, trackingID)
	
	// Start transaction
	tx, err := h.db.Begin()
	if err != nil {
		log.Printf("❌ Failed to begin transaction: %v", err)
		http.Error(w, "Failed to delete income transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()
	
	// Get the transaction amount before deleting
	var amount float64
	err = tx.QueryRow("SELECT amount FROM income_transactions WHERE transaction_id = ? AND tracking_id = ?", 
		transactionID, trackingID).Scan(&amount)
	if err != nil {
		log.Printf("❌ Failed to find income transaction: %v", err)
		http.Error(w, "Income transaction not found", http.StatusNotFound)
		return
	}
	
	// Delete the income transaction
	result, err := tx.Exec("DELETE FROM income_transactions WHERE transaction_id = ? AND tracking_id = ?", 
		transactionID, trackingID)
	if err != nil {
		log.Printf("❌ Failed to delete income transaction: %v", err)
		http.Error(w, "Failed to delete income transaction", http.StatusInternalServerError)
		return
	}
	
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Income transaction not found", http.StatusNotFound)
		return
	}
	
	// Update daily tracking totals
	_, err = tx.Exec(`
		UPDATE daily_trackings 
		SET total_income = total_income - ?, 
			net_revenue = net_revenue - ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE tracking_id = ?
	`, amount, amount, trackingID)
	if err != nil {
		log.Printf("❌ Failed to update daily tracking totals: %v", err)
		http.Error(w, "Failed to update tracking totals", http.StatusInternalServerError)
		return
	}
	
	if err = tx.Commit(); err != nil {
		log.Printf("❌ Failed to commit transaction: %v", err)
		http.Error(w, "Failed to delete income transaction", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Income transaction deleted successfully"})
	log.Printf("✅ Income transaction deleted: %s", transactionID)
}

// DeleteOutcomeTransaction handles DELETE /api/v1/daily-tracking/{tracking_id}/outcome/{transaction_id}
func (h *DailyTrackingHandler) DeleteOutcomeTransaction(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 6 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	
	trackingID := pathParts[4]
	transactionID := pathParts[6]
	
	log.Printf("🗑️ Deleting outcome transaction: %s from tracking: %s", transactionID, trackingID)
	
	// Start transaction
	tx, err := h.db.Begin()
	if err != nil {
		log.Printf("❌ Failed to begin transaction: %v", err)
		http.Error(w, "Failed to delete outcome transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()
	
	// Get the transaction amount before deleting
	var amount float64
	err = tx.QueryRow("SELECT amount FROM outcome_transactions WHERE transaction_id = ? AND tracking_id = ?", 
		transactionID, trackingID).Scan(&amount)
	if err != nil {
		log.Printf("❌ Failed to find outcome transaction: %v", err)
		http.Error(w, "Outcome transaction not found", http.StatusNotFound)
		return
	}
	
	// Delete the outcome transaction
	result, err := tx.Exec("DELETE FROM outcome_transactions WHERE transaction_id = ? AND tracking_id = ?", 
		transactionID, trackingID)
	if err != nil {
		log.Printf("❌ Failed to delete outcome transaction: %v", err)
		http.Error(w, "Failed to delete outcome transaction", http.StatusInternalServerError)
		return
	}
	
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Outcome transaction not found", http.StatusNotFound)
		return
	}
	
	// Update daily tracking totals
	_, err = tx.Exec(`
		UPDATE daily_trackings 
		SET total_outcome = total_outcome - ?, 
			net_revenue = net_revenue + ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE tracking_id = ?
	`, amount, amount, trackingID)
	if err != nil {
		log.Printf("❌ Failed to update daily tracking totals: %v", err)
		http.Error(w, "Failed to update tracking totals", http.StatusInternalServerError)
		return
	}
	
	if err = tx.Commit(); err != nil {
		log.Printf("❌ Failed to commit transaction: %v", err)
		http.Error(w, "Failed to delete outcome transaction", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Outcome transaction deleted successfully"})
	log.Printf("✅ Outcome transaction deleted: %s", transactionID)
}

// GetWeeklyAggregatedData handles GET /api/v1/daily-tracking/weekly-summary?merchant_id=xxx&start_date=xxx&end_date=xxx
func (h *DailyTrackingHandler) GetWeeklyAggregatedData(w http.ResponseWriter, r *http.Request) {
	merchantID := r.URL.Query().Get("merchant_id")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	
	if merchantID == "" || startDate == "" || endDate == "" {
		http.Error(w, "merchant_id, start_date, and end_date are required", http.StatusBadRequest)
		return
	}

	log.Printf("📊 Getting weekly summary for merchant: %s from %s to %s", merchantID, startDate, endDate)

	query := `
		SELECT 
			SUM(total_income) as total_income,
			SUM(total_outcome) as total_outcome,
			SUM(net_revenue) as net_revenue,
			COUNT(*) as tracking_days
		FROM daily_trackings
		WHERE merchant_id = ? AND tracking_date BETWEEN ? AND ?
	`

	var totalIncome, totalOutcome, netRevenue float64
	var trackingDays int

	err := h.db.QueryRow(query, merchantID, startDate, endDate).Scan(&totalIncome, &totalOutcome, &netRevenue, &trackingDays)
	if err != nil {
		log.Printf("❌ Failed to get weekly summary: %v", err)
		http.Error(w, "Failed to retrieve weekly summary", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"merchant_id":    merchantID,
		"start_date":     startDate,
		"end_date":       endDate,
		"total_income":   totalIncome,
		"total_outcome":  totalOutcome,
		"net_revenue":    netRevenue,
		"tracking_days":  trackingDays,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}