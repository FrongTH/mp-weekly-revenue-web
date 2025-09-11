-- Create database
CREATE DATABASE IF NOT EXISTS food_delivery_revenue_database;
USE food_delivery_revenue_database;

-- ==========================
-- OWNERS (merchant owners)
-- ==========================
CREATE TABLE owners (
    owner_id CHAR(36) PRIMARY KEY, -- UUID or string ID
    phone VARCHAR(15) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    is_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    last_login TIMESTAMP NULL,
    INDEX idx_phone (phone)
);

-- OTP verification table
CREATE TABLE otp_verifications (
    otp_id CHAR(36) PRIMARY KEY,
    owner_id CHAR(36) NULL,
    phone VARCHAR(15) NOT NULL,
    otp_code VARCHAR(6) NOT NULL,
    password_hash VARCHAR(255) NOT NULL, -- Store password temporarily until verification
    expires_at TIMESTAMP NOT NULL,
    is_used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_phone (phone),
    INDEX idx_otp_code (otp_code),
    INDEX idx_expires_at (expires_at),
    FOREIGN KEY (owner_id) REFERENCES owners(owner_id) ON DELETE CASCADE
);

-- ==========================
-- MERCHANT STRUCTURE
-- ==========================

-- Merchants (owned by owners, no address)
CREATE TABLE merchants (
    merchant_id CHAR(36) PRIMARY KEY,
    owner_id CHAR(36) NOT NULL,
    merchant_name VARCHAR(150) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (owner_id) REFERENCES owners(owner_id) ON DELETE CASCADE
);

-- Menu Items (with cost & sale price)
CREATE TABLE menu_items (
    item_id CHAR(36) PRIMARY KEY,
    merchant_id CHAR(36) NOT NULL,
    item_name VARCHAR(150) NOT NULL,
    cost DECIMAL(10,2) NOT NULL,              -- purchase cost
    general_price_sale DECIMAL(10,2) NOT NULL,  -- general selling price
    delivery_price_sale DECIMAL(10,2) NOT NULL, -- delivery selling price
    is_available BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id) ON DELETE CASCADE
);

-- Orders (minimal: no order_status, only timestamp + merchant)
CREATE TABLE orders (
    order_id CHAR(36) PRIMARY KEY,
    merchant_id CHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id) ON DELETE CASCADE
);

-- Extra Items (additional items for menu)
CREATE TABLE extra_items (
    extra_id CHAR(36) PRIMARY KEY,
    merchant_id CHAR(36) NOT NULL,
    item_name VARCHAR(150) NOT NULL,
    general_price_sale DECIMAL(10,2) NOT NULL,  -- general selling price
    delivery_price_sale DECIMAL(10,2) NOT NULL, -- delivery selling price
    is_available BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id) ON DELETE CASCADE
);

-- ==========================
-- WEEKLY TRACKING STRUCTURE
-- ==========================

-- Weekly Tracking Periods
CREATE TABLE weekly_trackings (
    tracking_id CHAR(36) PRIMARY KEY,
    merchant_id CHAR(36) NOT NULL,
    week_label VARCHAR(50) NOT NULL,
    date_range VARCHAR(50) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    total_income DECIMAL(10,2) DEFAULT 0,
    total_outcome DECIMAL(10,2) DEFAULT 0,
    net_revenue DECIMAL(10,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id) ON DELETE CASCADE,
    INDEX idx_merchant_dates (merchant_id, start_date, end_date)
);

-- Income Transactions
CREATE TABLE income_transactions (
    transaction_id CHAR(36) PRIMARY KEY,
    tracking_id CHAR(36) NOT NULL,
    description VARCHAR(255) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    category VARCHAR(50) NOT NULL,
    transaction_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (tracking_id) REFERENCES weekly_trackings(tracking_id) ON DELETE CASCADE,
    INDEX idx_tracking_date (tracking_id, transaction_date)
);

-- Outcome Transactions (Expenses)
CREATE TABLE outcome_transactions (
    transaction_id CHAR(36) PRIMARY KEY,
    tracking_id CHAR(36) NOT NULL,
    description VARCHAR(255) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    category VARCHAR(50) NOT NULL,
    transaction_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (tracking_id) REFERENCES weekly_trackings(tracking_id) ON DELETE CASCADE,
    INDEX idx_tracking_date (tracking_id, transaction_date)
);