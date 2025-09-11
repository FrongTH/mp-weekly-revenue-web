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

type MenuItemHandler struct {
	db *database.DB
}

func NewMenuItemHandler(db *database.DB) *MenuItemHandler {
	return &MenuItemHandler{db: db}
}

type MenuItem struct {
	ID                string  `json:"id"`
	MerchantID        string  `json:"merchant_id"`
	ItemName          string  `json:"item_name"`
	GeneralPriceSale  float64 `json:"general_price_sale"`
	DeliveryPriceSale float64 `json:"delivery_price_sale"`
	IsAvailable       bool    `json:"is_available"`
	CreatedAt         string  `json:"created_at,omitempty"`
}

type CreateMenuItemRequest struct {
	MerchantID        string  `json:"merchant_id"`
	ItemName          string  `json:"item_name"`
	GeneralPriceSale  float64 `json:"general_price_sale"`
	DeliveryPriceSale float64 `json:"delivery_price_sale"`
}

// CreateMenuItem handles POST /api/v1/menu-items
func (h *MenuItemHandler) CreateMenuItem(w http.ResponseWriter, r *http.Request) {
	log.Printf("📦 Creating menu item")

	var req CreateMenuItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("❌ Failed to parse request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.MerchantID == "" || req.ItemName == "" || req.GeneralPriceSale <= 0 || req.DeliveryPriceSale <= 0 {
		log.Printf("❌ Invalid menu item data")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Generate UUID for item
	itemID := uuid.New().String()

	// Insert menu item
	query := `
		INSERT INTO menu_items (item_id, merchant_id, item_name, cost, general_price_sale, delivery_price_sale, is_available) 
		VALUES (?, ?, ?, 0, ?, ?, TRUE)
	`
	_, err := h.db.Exec(query, itemID, req.MerchantID, req.ItemName, req.GeneralPriceSale, req.DeliveryPriceSale)
	if err != nil {
		log.Printf("❌ Failed to create menu item: %v", err)
		http.Error(w, "Failed to create menu item", http.StatusInternalServerError)
		return
	}

	response := MenuItem{
		ID:                itemID,
		MerchantID:        req.MerchantID,
		ItemName:          req.ItemName,
		GeneralPriceSale:  req.GeneralPriceSale,
		DeliveryPriceSale: req.DeliveryPriceSale,
		IsAvailable:       true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("✅ Menu item created: %s", itemID)
}

// GetMenuItemsByMerchant handles GET /api/v1/menu-items?merchant_id=xxx
func (h *MenuItemHandler) GetMenuItemsByMerchant(w http.ResponseWriter, r *http.Request) {
	merchantID := r.URL.Query().Get("merchant_id")
	if merchantID == "" {
		http.Error(w, "merchant_id is required", http.StatusBadRequest)
		return
	}

	log.Printf("📋 Getting menu items for merchant: %s", merchantID)

	query := `
		SELECT item_id, merchant_id, item_name, general_price_sale, delivery_price_sale, is_available, created_at
		FROM menu_items
		WHERE merchant_id = ?
		ORDER BY created_at DESC
	`

	rows, err := h.db.Query(query, merchantID)
	if err != nil {
		log.Printf("❌ Failed to get menu items: %v", err)
		http.Error(w, "Failed to retrieve menu items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []MenuItem
	for rows.Next() {
		var item MenuItem
		var createdAt sql.NullTime
		err := rows.Scan(&item.ID, &item.MerchantID, &item.ItemName, &item.GeneralPriceSale, &item.DeliveryPriceSale, &item.IsAvailable, &createdAt)
		if err != nil {
			log.Printf("❌ Failed to scan menu item: %v", err)
			continue
		}
		if createdAt.Valid {
			item.CreatedAt = createdAt.Time.Format("2006-01-02T15:04:05Z")
		}
		items = append(items, item)
	}

	// Return empty array if no items found
	if items == nil {
		items = []MenuItem{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// UpdateMenuItem handles PUT /api/v1/menu-items/{id}
func (h *MenuItemHandler) UpdateMenuItem(w http.ResponseWriter, r *http.Request) {
	// Extract item ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	itemID := pathParts[4]

	log.Printf("📝 Updating menu item: %s", itemID)

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

	params = append(params, itemID)
	query := "UPDATE menu_items SET " + strings.Join(updates, ", ") + " WHERE item_id = ?"

	result, err := h.db.Exec(query, params...)
	if err != nil {
		log.Printf("❌ Failed to update menu item: %v", err)
		http.Error(w, "Failed to update menu item", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Menu item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Menu item updated successfully"})
	log.Printf("✅ Menu item updated: %s", itemID)
}

// DeleteMenuItem handles DELETE /api/v1/menu-items/{id}
func (h *MenuItemHandler) DeleteMenuItem(w http.ResponseWriter, r *http.Request) {
	// Extract item ID from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	itemID := pathParts[4]

	log.Printf("🗑️ Deleting menu item: %s", itemID)

	query := "DELETE FROM menu_items WHERE item_id = ?"
	result, err := h.db.Exec(query, itemID)
	if err != nil {
		log.Printf("❌ Failed to delete menu item: %v", err)
		http.Error(w, "Failed to delete menu item", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Menu item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Menu item deleted successfully"})
	log.Printf("✅ Menu item deleted: %s", itemID)
}