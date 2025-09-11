-- Create database
CREATE DATABASE IF NOT EXISTS food_delivery_revenue_database;
USE food_delivery_revenue_database;


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