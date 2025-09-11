package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"food-delivery-backend/internal/database"
	"food-delivery-backend/internal/models"
	"food-delivery-backend/internal/services"
)

type AuthHandler struct {
	db         *database.DB
	otpService *services.OTPService
}

func NewAuthHandler(db *database.DB) *AuthHandler {
	return &AuthHandler{
		db:         db,
		otpService: services.NewOTPService(),
	}
}

// SignIn handles owner sign in
func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Check if owner exists
	var owner models.Owner
	var hashedPassword string
	
	query := `SELECT owner_id, phone, password, is_verified, created_at, updated_at, last_login 
	          FROM owners WHERE phone = ?`
	
	row := h.db.QueryRow(query, req.Phone)
	err := row.Scan(&owner.OwnerID, &owner.Phone, &hashedPassword, &owner.IsVerified, &owner.CreatedAt, &owner.UpdatedAt, &owner.LastLogin)
	
	if err == sql.ErrNoRows {
		response := models.AuthResponse{
			Success: false,
			Message: "This mobile phone is not member please register before.",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	if err != nil {
		log.Printf("Database error during sign in: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		response := models.AuthResponse{
			Success: false,
			Message: "Invalid password",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Update last login
	updateQuery := `UPDATE owners SET last_login = ? WHERE owner_id = ?`
	now := time.Now()
	_, err = h.db.Exec(updateQuery, now, owner.OwnerID)
	if err != nil {
		log.Printf("Failed to update last login: %v", err)
	}
	owner.LastLogin = &now

	// Successful sign in
	response := models.AuthResponse{
		Success: true,
		Message: "Sign in successful",
		Owner:   &owner,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Register handles owner registration by sending OTP
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Check if owner already exists
	var existingID string
	checkQuery := `SELECT owner_id FROM owners WHERE phone = ?`
	row := h.db.QueryRow(checkQuery, req.Phone)
	err := row.Scan(&existingID)
	
	if err == nil {
		response := models.OTPResponse{
			Success: false,
			Message: "This phone number is member please sign in or other number register",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	if err != sql.ErrNoRows {
		log.Printf("Database error during registration check: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Hash password for temporary storage
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Generate OTP
	otpCode := h.otpService.GenerateOTP()
	expiresAt := h.otpService.GetExpiryTime()
	otpID := uuid.New().String()

	// Delete any existing OTP for this phone number
	deleteQuery := `DELETE FROM otp_verifications WHERE phone = ?`
	_, err = h.db.Exec(deleteQuery, req.Phone)
	if err != nil {
		log.Printf("Failed to clean existing OTP: %v", err)
	}

	// Store OTP verification record
	insertQuery := `INSERT INTO otp_verifications (otp_id, phone, otp_code, password_hash, expires_at) VALUES (?, ?, ?, ?, ?)`
	_, err = h.db.Exec(insertQuery, otpID, req.Phone, otpCode, string(hashedPassword), expiresAt)
	if err != nil {
		log.Printf("Failed to store OTP: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Log OTP for testing (remove SMS sending)
	log.Printf("📱 Testing OTP for %s: %s (Valid for 5 minutes)", req.Phone, otpCode)

	// Successful OTP generated (return code for testing)
	response := models.OTPResponse{
		Success: true,
		Message: "OTP generated successfully. Use the code shown in the modal.",
		Phone:   req.Phone,
		OTPCode: otpCode, // For testing only - remove in production
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// VerifyOTP handles OTP verification and completes owner registration
func (h *AuthHandler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	var req models.OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Find OTP verification record
	var otp models.OTPVerification
	selectQuery := `SELECT otp_id, phone, otp_code, password_hash, expires_at, is_used 
	                FROM otp_verifications 
	                WHERE phone = ? AND otp_code = ? AND is_used = FALSE 
	                ORDER BY created_at DESC LIMIT 1`
	
	row := h.db.QueryRow(selectQuery, req.Phone, req.OTP)
	err := row.Scan(&otp.OTPID, &otp.Phone, &otp.OTPCode, &otp.PasswordHash, &otp.ExpiresAt, &otp.IsUsed)
	
	if err == sql.ErrNoRows {
		response := models.AuthResponse{
			Success: false,
			Message: "Invalid or expired OTP code",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	if err != nil {
		log.Printf("Database error during OTP verification: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Check if OTP is expired
	if time.Now().After(otp.ExpiresAt) {
		response := models.AuthResponse{
			Success: false,
			Message: "OTP code has expired",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Mark OTP as used
	updateQuery := `UPDATE otp_verifications SET is_used = TRUE WHERE otp_id = ?`
	_, err = h.db.Exec(updateQuery, otp.OTPID)
	if err != nil {
		log.Printf("Failed to mark OTP as used: %v", err)
	}

	// Create the owner account
	ownerID := uuid.New().String()
	insertOwnerQuery := `INSERT INTO owners (owner_id, phone, password, is_verified) VALUES (?, ?, ?, TRUE)`
	_, err = h.db.Exec(insertOwnerQuery, ownerID, otp.Phone, otp.PasswordHash)
	if err != nil {
		log.Printf("Failed to create owner: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Update OTP record with owner_id
	linkQuery := `UPDATE otp_verifications SET owner_id = ? WHERE otp_id = ?`
	_, err = h.db.Exec(linkQuery, ownerID, otp.OTPID)
	if err != nil {
		log.Printf("Failed to link OTP to owner: %v", err)
	}

	// Get the created owner
	var owner models.Owner
	selectOwnerQuery := `SELECT owner_id, phone, is_verified, created_at, updated_at, last_login 
	                    FROM owners WHERE owner_id = ?`
	
	row = h.db.QueryRow(selectOwnerQuery, ownerID)
	err = row.Scan(&owner.OwnerID, &owner.Phone, &owner.IsVerified, &owner.CreatedAt, &owner.UpdatedAt, &owner.LastLogin)
	if err != nil {
		log.Printf("Failed to get created owner: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Clean up used OTP
	cleanupQuery := `DELETE FROM otp_verifications WHERE phone = ? AND is_used = TRUE`
	_, err = h.db.Exec(cleanupQuery, otp.Phone)
	if err != nil {
		log.Printf("Failed to cleanup OTP records: %v", err)
	}

	// Successful registration
	response := models.AuthResponse{
		Success: true,
		Message: "Registration completed successfully",
		Owner:   &owner,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// CheckPhone checks if a phone number exists (for frontend validation)
func (h *AuthHandler) CheckPhone(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	phone := r.URL.Query().Get("phone")
	if phone == "" {
		http.Error(w, "Phone parameter required", http.StatusBadRequest)
		return
	}

	var existingID string
	checkQuery := `SELECT owner_id FROM owners WHERE phone = ?`
	row := h.db.QueryRow(checkQuery, phone)
	err := row.Scan(&existingID)
	
	exists := err == nil
	
	response := map[string]interface{}{
		"phone":  phone,
		"exists": exists,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}