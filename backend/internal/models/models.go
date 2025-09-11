package models

import (
	"time"
)

// Owner model (merchant owner)
type Owner struct {
	OwnerID    string     `json:"owner_id" db:"owner_id"`
	Phone      string     `json:"phone" db:"phone"`
	Password   string     `json:"-" db:"password"` // Don't expose password in JSON
	IsVerified bool       `json:"is_verified" db:"is_verified"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	LastLogin  *time.Time `json:"last_login,omitempty" db:"last_login"`
}

// OTPVerification model
type OTPVerification struct {
	OTPID        string    `json:"otp_id" db:"otp_id"`
	OwnerID      *string   `json:"owner_id,omitempty" db:"owner_id"`
	Phone        string    `json:"phone" db:"phone"`
	OTPCode      string    `json:"otp_code" db:"otp_code"`
	PasswordHash string    `json:"-" db:"password_hash"`
	ExpiresAt    time.Time `json:"expires_at" db:"expires_at"`
	IsUsed       bool      `json:"is_used" db:"is_used"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// Merchant model
type Merchant struct {
	MerchantID   string    `json:"merchant_id" db:"merchant_id"`
	OwnerID      string    `json:"owner_id" db:"owner_id"`
	MerchantName string    `json:"merchant_name" db:"merchant_name"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// MenuItem model
type MenuItem struct {
	ItemID            string    `json:"item_id" db:"item_id"`
	MerchantID        string    `json:"merchant_id" db:"merchant_id"`
	ItemName          string    `json:"item_name" db:"item_name"`
	Cost              float64   `json:"cost" db:"cost"`                             // Purchase cost
	GeneralPriceSale  float64   `json:"general_price_sale" db:"general_price_sale"`   // General selling price
	DeliveryPriceSale float64   `json:"delivery_price_sale" db:"delivery_price_sale"` // Delivery selling price
	IsAvailable       bool      `json:"is_available" db:"is_available"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

// ExtraItem model
type ExtraItem struct {
	ExtraID           string    `json:"extra_id" db:"extra_id"`
	MerchantID        string    `json:"merchant_id" db:"merchant_id"`
	ItemName          string    `json:"item_name" db:"item_name"`
	GeneralPriceSale  float64   `json:"general_price_sale" db:"general_price_sale"`   // General selling price
	DeliveryPriceSale float64   `json:"delivery_price_sale" db:"delivery_price_sale"` // Delivery selling price
	IsAvailable       bool      `json:"is_available" db:"is_available"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

// Order model (simplified)
type Order struct {
	OrderID    string    `json:"order_id" db:"order_id"`
	MerchantID string    `json:"merchant_id" db:"merchant_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// Auth request/response models
type AuthRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Owner   *Owner `json:"owner,omitempty"`
	OTPCode string `json:"otp_code,omitempty"` // For testing only - remove in production
}

type OTPRequest struct {
	Phone string `json:"phone" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}

type OTPResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Phone   string `json:"phone,omitempty"`
	OTPCode string `json:"otp_code,omitempty"` // For testing only - remove in production
}

// Check phone request/response
type CheckPhoneRequest struct {
	Phone string `json:"phone" binding:"required"`
}

type CheckPhoneResponse struct {
	Exists bool `json:"exists"`
}

// Revenue and reporting models
type MerchantRevenue struct {
	MerchantID   string  `json:"merchant_id" db:"merchant_id"`
	MerchantName string  `json:"merchant_name" db:"merchant_name"`
	TotalIncome  float64 `json:"total_income" db:"total_income"`   // Total sales
	TotalCost    float64 `json:"total_cost" db:"total_cost"`       // Total costs
	NetRevenue   float64 `json:"net_revenue" db:"net_revenue"`     // Income - Cost
	OrderCount   int     `json:"order_count" db:"order_count"`
	Period       string  `json:"period" db:"period"`
}

// Dashboard summary for owner
type DashboardSummary struct {
	OwnerID        string            `json:"owner_id"`
	TotalMerchants int               `json:"total_merchants"`
	TotalRevenue   float64           `json:"total_revenue"`
	TotalIncome    float64           `json:"total_income"`
	TotalOutcome   float64           `json:"total_outcome"`
	Merchants      []MerchantSummary `json:"merchants"`
	FirstOrderDate *time.Time        `json:"first_order_date,omitempty"`
}

// Individual merchant summary
type MerchantSummary struct {
	MerchantID     string  `json:"merchant_id"`
	MerchantName   string  `json:"merchant_name"`
	Income         float64 `json:"income"`
	Outcome        float64 `json:"outcome"`
	Revenue        float64 `json:"revenue"`
	IncomeChange   float64 `json:"income_change,omitempty"`
	OutcomeChange  float64 `json:"outcome_change,omitempty"`
	RevenueChange  float64 `json:"revenue_change,omitempty"`
	Category       string  `json:"category,omitempty"`
}

// Daily Tracking model
type DailyTracking struct {
	TrackingID    string    `json:"tracking_id" db:"tracking_id"`
	MerchantID    string    `json:"merchant_id" db:"merchant_id"`
	TrackingDate  string    `json:"tracking_date" db:"tracking_date"`
	TotalIncome   float64   `json:"total_income" db:"total_income"`
	TotalOutcome  float64   `json:"total_outcome" db:"total_outcome"`
	NetRevenue    float64   `json:"net_revenue" db:"net_revenue"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// Income Transaction model
type IncomeTransaction struct {
	TransactionID   string    `json:"transaction_id" db:"transaction_id"`
	TrackingID      string    `json:"tracking_id" db:"tracking_id"`
	Description     string    `json:"description" db:"description"`
	Amount          float64   `json:"amount" db:"amount"`
	Category        string    `json:"category" db:"category"`
	TransactionDate string    `json:"transaction_date" db:"transaction_date"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Outcome Transaction model
type OutcomeTransaction struct {
	TransactionID   string    `json:"transaction_id" db:"transaction_id"`
	TrackingID      string    `json:"tracking_id" db:"tracking_id"`
	Description     string    `json:"description" db:"description"`
	Amount          float64   `json:"amount" db:"amount"`
	Category        string    `json:"category" db:"category"`
	TransactionDate string    `json:"transaction_date" db:"transaction_date"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}