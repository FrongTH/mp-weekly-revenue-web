package handlers

import (
	"encoding/json"
	"food-delivery-backend/internal/database"
	"log"
	"net/http"
	"strings"
	"github.com/google/uuid"
)

type MerchantHandler struct {
	db *database.DB
}

func NewMerchantHandler(db *database.DB) *MerchantHandler {
	return &MerchantHandler{
		db: db,
	}
}

type CreateMerchantRequest struct {
	OwnerID      string `json:"owner_id"`
	MerchantName string `json:"merchant_name"`
}

type MerchantResponse struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	MerchantID string `json:"merchant_id,omitempty"`
}

// CreateMerchant handles the creation of a new merchant
func (h *MerchantHandler) CreateMerchant(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Parse request body
	var req CreateMerchantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Validate input
	if req.OwnerID == "" || req.MerchantName == "" {
		response := MerchantResponse{
			Success: false,
			Message: "Owner ID and merchant name are required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if owner exists
	var ownerExists bool
	checkOwnerQuery := `SELECT EXISTS(SELECT 1 FROM owners WHERE owner_id = ?)`
	err := h.db.QueryRow(checkOwnerQuery, req.OwnerID).Scan(&ownerExists)
	if err != nil {
		log.Printf("Error checking owner: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if !ownerExists {
		response := MerchantResponse{
			Success: false,
			Message: "Owner not found",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generate merchant ID
	merchantID := uuid.New().String()

	// Insert merchant into database
	insertQuery := `
		INSERT INTO merchants (merchant_id, owner_id, merchant_name, created_at, updated_at)
		VALUES (?, ?, ?, NOW(), NOW())
	`
	
	_, err = h.db.Exec(insertQuery, merchantID, req.OwnerID, req.MerchantName)
	if err != nil {
		log.Printf("Error creating merchant: %v", err)
		response := MerchantResponse{
			Success: false,
			Message: "Failed to create merchant",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Printf("✅ Merchant created successfully: ID=%s, Name=%s, Owner=%s", merchantID, req.MerchantName, req.OwnerID)

	// Return success response
	response := MerchantResponse{
		Success:    true,
		Message:    "Merchant created successfully",
		MerchantID: merchantID,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// DeleteMerchant handles the deletion of a merchant
func (h *MerchantHandler) DeleteMerchant(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Extract merchant ID from URL path
	// Expected URL format: /api/v1/merchants/{merchantId}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Invalid merchant ID", http.StatusBadRequest)
		return
	}
	merchantID := pathParts[4]

	if merchantID == "" {
		response := MerchantResponse{
			Success: false,
			Message: "Merchant ID is required",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if merchant exists
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM merchants WHERE merchant_id = ?)`
	err := h.db.QueryRow(checkQuery, merchantID).Scan(&exists)
	if err != nil {
		log.Printf("Error checking merchant: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if !exists {
		response := MerchantResponse{
			Success: false,
			Message: "Merchant not found",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Delete merchant from database
	deleteQuery := `DELETE FROM merchants WHERE merchant_id = ?`
	_, err = h.db.Exec(deleteQuery, merchantID)
	if err != nil {
		log.Printf("Error deleting merchant: %v", err)
		response := MerchantResponse{
			Success: false,
			Message: "Failed to delete merchant",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Printf("✅ Merchant deleted successfully: ID=%s", merchantID)

	// Return success response
	response := MerchantResponse{
		Success: true,
		Message: "Merchant deleted successfully",
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetMerchantsByOwner returns all merchants for a specific owner
func (h *MerchantHandler) GetMerchantsByOwner(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	ownerID := r.URL.Query().Get("owner_id")
	if ownerID == "" {
		http.Error(w, "owner_id parameter required", http.StatusBadRequest)
		return
	}

	query := `
		SELECT merchant_id, merchant_name, created_at
		FROM merchants
		WHERE owner_id = ?
		ORDER BY created_at DESC
	`

	rows, err := h.db.Query(query, ownerID)
	if err != nil {
		log.Printf("Error fetching merchants: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Merchant struct {
		MerchantID   string `json:"merchant_id"`
		MerchantName string `json:"merchant_name"`
		CreatedAt    string `json:"created_at"`
	}

	var merchants []Merchant
	for rows.Next() {
		var m Merchant
		if err := rows.Scan(&m.MerchantID, &m.MerchantName, &m.CreatedAt); err != nil {
			log.Printf("Error scanning merchant: %v", err)
			continue
		}
		merchants = append(merchants, m)
	}

	// Return empty array if no merchants
	if merchants == nil {
		merchants = []Merchant{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(merchants)
}