package handlers

import (
	"database/sql"
	"encoding/json"
	"food-delivery-backend/internal/database"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type ExtraItemHandler struct {
	db *database.DB
}

func NewExtraItemHandler(db *database.DB) *ExtraItemHandler {
	return &ExtraItemHandler{db: db}
}

type ExtraItem struct {
	ID                string  `json:"id"`
	MerchantID        string  `json:"merchant_id"`
	ItemName          string  `json:"item_name"`
	GeneralPriceSale  float64 `json:"general_price_sale"`
	DeliveryPriceSale float64 `json:"delivery_price_sale"`
	IsAvailable       bool    `json:"is_available"`
	CreatedAt         string  `json:"created_at,omitempty"`
}

type CreateExtraItemRequest struct {
	MerchantID        string  `json:"merchant_id"`
	ItemName          string  `json:"item_name"`
	GeneralPriceSale  float64 `json:"general_price_sale"`
	DeliveryPriceSale float64 `json:"delivery_price_sale"`
}

// CreateExtraItem handles POST /api/v1/extra-items
func (h *ExtraItemHandler) CreateExtraItem(w http.ResponseWriter, r *http.Request) {
	log.Printf("📦 Creating extra item")

	var req CreateExtraItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("❌ Failed to parse request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.MerchantID == "" || req.ItemName == "" || req.GeneralPriceSale <= 0 || req.DeliveryPriceSale <= 0 {
		log.Printf("❌ Invalid extra item data")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Generate UUID for item
	extraID := uuid.New().String()

	// Insert extra item
	query := `
		INSERT INTO extra_items (extra_id, merchant_id, item_name, general_price_sale, delivery_price_sale, is_available) 
		VALUES (?, ?, ?, ?, ?, TRUE)
	`
	_, err := h.db.Exec(query, extraID, req.MerchantID, req.ItemName, req.GeneralPriceSale, req.DeliveryPriceSale)
	if err != nil {
		log.Printf("❌ Failed to create extra item: %v", err)
		http.Error(w, "Failed to create extra item", http.StatusInternalServerError)
		return
	}

	response := ExtraItem{
		ID:                extraID,
		MerchantID:        req.MerchantID,
		ItemName:          req.ItemName,
		GeneralPriceSale:  req.GeneralPriceSale,
		DeliveryPriceSale: req.DeliveryPriceSale,
		IsAvailable:       true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("✅ Extra item created: %s", extraID)
}

// GetExtraItemsByMerchant handles GET /api/v1/extra-items?merchant_id=xxx
func (h *ExtraItemHandler) GetExtraItemsByMerchant(w http.ResponseWriter, r *http.Request) {
	merchantID := r.URL.Query().Get("merchant_id")
	if merchantID == "" {
		http.Error(w, "merchant_id is required", http.StatusBadRequest)
		return
	}

	log.Printf("📋 Getting extra items for merchant: %s", merchantID)

	query := `
		SELECT extra_id, merchant_id, item_name, general_price_sale, delivery_price_sale, is_available, created_at
		FROM extra_items
		WHERE merchant_id = ?
		ORDER BY created_at DESC
	`

	rows, err := h.db.Query(query, merchantID)
	if err != nil {
		log.Printf("❌ Failed to get extra items: %v", err)
		http.Error(w, "Failed to retrieve extra items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []ExtraItem
	for rows.Next() {
		var item ExtraItem
		var createdAt sql.NullTime
		err := rows.Scan(&item.ID, &item.MerchantID, &item.ItemName, &item.GeneralPriceSale, &item.DeliveryPriceSale, &item.IsAvailable, &createdAt)
		if err != nil {
			log.Printf("❌ Failed to scan extra item: %v", err)
			continue
		}
		if createdAt.Valid {
			item.CreatedAt = createdAt.Time.Format("2006-01-02T15:04:05Z")
		}
		items = append(items, item)
	}

	// Return empty array if no items found
	if items == nil {
		items = []ExtraItem{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// UpdateExtraItem handles PUT /api/v1/extra-items/{id}
func (h *ExtraItemHandler) UpdateExtraItem(w http.ResponseWriter, r *http.Request) {
	// Extract item ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	extraID := pathParts[4]

	log.Printf("📝 Updating extra item: %s", extraID)

	var req struct {
		ItemName          string  `json:"item_name"`
		GeneralPriceSale  float64 `json:"general_price_sale"`
		DeliveryPriceSale float64 `json:"delivery_price_sale"`
		IsAvailable       *bool   `json:"is_available"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Build update query dynamically
	updates := []string{}
	params := []interface{}{}

	if req.ItemName != "" {
		updates = append(updates, "item_name = ?")
		params = append(params, req.ItemName)
	}
	if req.GeneralPriceSale > 0 {
		updates = append(updates, "general_price_sale = ?")
		params = append(params, req.GeneralPriceSale)
	}
	if req.DeliveryPriceSale > 0 {
		updates = append(updates, "delivery_price_sale = ?")
		params = append(params, req.DeliveryPriceSale)
	}
	if req.IsAvailable != nil {
		updates = append(updates, "is_available = ?")
		params = append(params, *req.IsAvailable)
	}

	if len(updates) == 0 {
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	params = append(params, extraID)
	query := "UPDATE extra_items SET " + strings.Join(updates, ", ") + " WHERE extra_id = ?"

	result, err := h.db.Exec(query, params...)
	if err != nil {
		log.Printf("❌ Failed to update extra item: %v", err)
		http.Error(w, "Failed to update extra item", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Extra item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Extra item updated successfully"})
	log.Printf("✅ Extra item updated: %s", extraID)
}

// DeleteExtraItem handles DELETE /api/v1/extra-items/{id}
func (h *ExtraItemHandler) DeleteExtraItem(w http.ResponseWriter, r *http.Request) {
	// Extract item ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	extraID := pathParts[4]

	log.Printf("🗑️ Deleting extra item: %s", extraID)

	query := "DELETE FROM extra_items WHERE extra_id = ?"
	result, err := h.db.Exec(query, extraID)
	if err != nil {
		log.Printf("❌ Failed to delete extra item: %v", err)
		http.Error(w, "Failed to delete extra item", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Extra item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Extra item deleted successfully"})
	log.Printf("✅ Extra item deleted: %s", extraID)
}