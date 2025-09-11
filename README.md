# Food Delivery Revenue Tracking

A full-stack application for tracking food delivery revenue with Vue.js frontend and Go backend.

## 📋 Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Installation & Setup](#installation--setup)
- [Running the Application](#running-the-application)
- [Development](#development)
- [API Endpoints](#api-endpoints)
- [Database Schema](#database-schema)
- [Environment Configuration](#environment-configuration)

## 🛠 Prerequisites

Before running this project locally, ensure you have the following installed:

### Required Software
- **Node.js**: Version ^20.19.0 or >=22.12.0
- **pnpm**: Package manager for frontend dependencies
- **Go**: Version 1.21.0 or higher
- **MySQL**: Version 5.7 or higher

### Installation Commands
```bash
# Install Node.js (visit https://nodejs.org)
# Install pnpm
npm install -g pnpm

# Install Go (visit https://golang.org/dl)
# Install MySQL (visit https://dev.mysql.com/downloads/mysql)
```

## 📁 Project Structure

```
food-delivery-revenue-tracking/
├── src/                          # Frontend (Vue.js)
│   ├── main.js                   # Application entry point
│   ├── App.vue                   # Root component
│   ├── router/                   # Vue Router configuration
│   └── __tests__/               # Unit tests
├── backend/                      # Backend (Go)
│   ├── cmd/server/              # Application entry point
│   ├── internal/                # Internal packages
│   │   ├── database/           # Database connection
│   │   └── models/             # Data models
│   ├── pkg/config/             # Configuration management
│   ├── database/               # SQL schema and migrations
│   ├── .env.example           # Environment variables template
│   └── go.mod                 # Go dependencies
├── package.json                # Frontend dependencies
├── vite.config.js             # Vite configuration
└── README.md                  # This file
```

## ⚙️ Installation & Setup

### 1. Clone the Repository
```bash
git clone <repository-url>
cd food-delivery-revenue-tracking
```

### 2. Frontend Setup
```bash
# Install frontend dependencies
pnpm install
```

### 3. Backend Setup
```bash
# Navigate to backend directory
cd backend

# Install Go dependencies
go mod tidy

# Return to root directory
cd ..
```

### 4. Database Setup

#### Start MySQL Service
```bash
# On Windows (if using MySQL service)
net start mysql

# On macOS (if using Homebrew)
brew services start mysql

# On Linux (Ubuntu/Debian)
sudo systemctl start mysql
```

#### Create Database and Tables
```bash
# Connect to MySQL
mysql -u root -p

# Create database and tables using the schema file
source backend/database/schema.sql

# Exit MySQL
exit
```

### 5. Environment Configuration
```bash
# Copy environment template
cp backend/.env.example backend/.env

# Edit the .env file with your database credentials
# Default values should work if you're using standard MySQL setup
```

## 🚀 Running the Application

### Development Mode (Recommended)

#### Terminal 1 - Start Backend Server
```bash
cd backend
go run cmd/server/main.go
```
The backend server will start on `http://localhost:8080`

#### Terminal 2 - Start Frontend Development Server
```bash
# From root directory
pnpm dev
```
The frontend will start on `http://localhost:5173` with hot reload

### Production Mode

#### Build Frontend
```bash
pnpm build
```

#### Build Backend Binary
```bash
cd backend
go build -o bin/server cmd/server/main.go
./bin/server
```

## 🔧 Development

### Frontend Commands
```bash
# Development server with hot reload
pnpm dev

# Build for production
pnpm build

# Preview production build
pnpm preview

# Run unit tests
pnpm test:unit

# Lint code (with auto-fix)
pnpm lint

# Format code with Prettier
pnpm format
```

### Backend Commands
```bash
cd backend

# Run development server
go run cmd/server/main.go

# Build binary
go build -o bin/server cmd/server/main.go

# Run tests (if available)
go test ./...

# Update dependencies
go mod tidy
```

### Database Commands
```bash
# Connect to MySQL
mysql -u root -p

# Use the database
use food_delivery_revenue_database;

# Show tables
show tables;

# Reset database (reload schema)
source backend/database/schema.sql
```

## 🌐 API Endpoints

The backend server provides the following REST API endpoints:

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check endpoint |
| GET | `/api/v1/restaurants` | Get all restaurants |
| GET | `/api/v1/orders` | Get all orders |
| GET | `/api/v1/revenue` | Get revenue reports |

### Example API Calls
```bash
# Health check
curl http://localhost:8080/health

# Get restaurants
curl http://localhost:8080/api/v1/restaurants

# Get orders
curl http://localhost:8080/api/v1/orders

# Get revenue reports
curl http://localhost:8080/api/v1/revenue
```

## 🗄️ Database Schema

The application uses the following main tables:

- **restaurants**: Store restaurant information
- **orders**: Store order details and revenue data
- **revenue_reports**: Store aggregated revenue reports

See `backend/database/schema.sql` for complete schema definition and sample data.

## 🔐 Environment Configuration

### Backend Environment Variables (`.env`)
```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_mysql_password
DB_NAME=food_delivery_revenue_database

# Server Configuration
SERVER_PORT=8080
```

## 🧪 Testing

### Frontend Tests
```bash
# Run unit tests
pnpm test:unit

# Run tests in watch mode
pnpm test:unit --watch
```

### Backend Tests
```bash
cd backend
go test ./...
```

## 🐛 Troubleshooting

### Common Issues

#### Port Already in Use
```bash
# Check what's using port 8080
netstat -ano | findstr :8080  # Windows
lsof -i :8080                 # macOS/Linux

# Kill the process if needed
```

#### Database Connection Issues
1. Ensure MySQL is running
2. Check database credentials in `.env` file
3. Verify database exists: `show databases;`
4. Check if tables exist: `use food_delivery_revenue_database; show tables;`

#### Frontend Build Issues
```bash
# Clear node_modules and reinstall
rm -rf node_modules pnpm-lock.yaml
pnpm install
```

#### Backend Build Issues
```bash
# Clean Go module cache
go clean -modcache
go mod tidy
```

## 📚 Additional Resources

- [Vue.js Documentation](https://vuejs.org/)
- [Vite Documentation](https://vite.dev/)
- [Go Documentation](https://golang.org/doc/)
- [MySQL Documentation](https://dev.mysql.com/doc/)
- [Gorilla Mux Router](https://github.com/gorilla/mux)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request

## 📝 License

This project is private and proprietary.