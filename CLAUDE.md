# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Frontend (Vue.js)
- **Install dependencies**: `pnpm install`
- **Development server**: `pnpm dev` - Starts Vite dev server with hot reload
- **Build for production**: `pnpm build` - Compiles and minifies for production
- **Preview production build**: `pnpm preview` - Preview production build locally
- **Run unit tests**: `pnpm test:unit` - Runs Vitest unit tests
- **Lint code**: `pnpm lint` - Runs ESLint with auto-fix
- **Format code**: `pnpm format` - Formats code with Prettier

### Backend (Go)
- **Install dependencies**: `cd backend && go mod tidy`
- **Run server**: `cd backend && go run cmd/server/main.go` - Starts Go server on port 8080
- **Build binary**: `cd backend && go build -o bin/server cmd/server/main.go`
- **Connect MySQL Database**: `mysql -u root -p {password}`
- **Database setup**: `source backend/database/schema.sql` - Creates database and tables

## Project Architecture

This is a full-stack food delivery revenue tracking application:

### Frontend (Vue 3 + Vite)
- **Framework**: Vue 3 with Composition API (`<script setup>` syntax)
- **Build Tool**: Vite with Vue plugin and devtools
- **Routing**: Vue Router with history mode (currently no routes defined)
- **Testing**: Vitest with jsdom environment for unit tests
- **Linting**: ESLint with Vue, Vitest, and Prettier configurations
- **Package Manager**: pnpm (workspace configuration present)

### Backend (Go + MySQL)
- **Language**: Go 1.21+
- **Database**: MySQL with `food_delivery_revenue_database`
- **Router**: Gorilla Mux for HTTP routing
- **Database Driver**: go-sql-driver/mysql
- **Configuration**: Environment variables with .env support (godotenv)
- **Architecture**: Clean architecture with separated concerns:
  - `cmd/server/` - Application entry point
  - `internal/database/` - Database connection logic
  - `internal/models/` - Data structures (Restaurant, Order, RevenueReport)
  - `pkg/config/` - Configuration management
  - `database/` - SQL schema and migrations

## Code Structure

### Frontend
- `src/main.js` - Application entry point, mounts Vue app with router
- `src/App.vue` - Root Vue component (currently minimal)
- `src/router/index.js` - Vue Router configuration (no routes defined yet)
- `src/__tests__/` - Unit test files location

### Backend
- `backend/cmd/server/main.go` - HTTP server with basic REST API endpoints
- `backend/internal/database/connection.go` - MySQL connection management
- `backend/internal/models/models.go` - Data models for restaurants, orders, and revenue reports
- `backend/pkg/config/config.go` - Configuration loading with environment variables
- `backend/database/schema.sql` - Database schema with sample data
- `backend/.env.example` - Environment variable template

## Database Setup

1. Ensure MySQL is running on localhost:3306
2. Create `.env` file in backend/ based on `.env.example`
3. Run: `mysql -u root -p < backend/database/schema.sql`
4. Default database name: `food_delivery_revenue_database`

## API Endpoints

- `GET /health` - Health check endpoint
- `GET /api/v1/restaurants` - Restaurants endpoint (placeholder)
- `GET /api/v1/orders` - Orders endpoint (placeholder)
- `GET /api/v1/revenue` - Revenue reports endpoint (placeholder)

The project requires Node.js ^20.19.0 || >=22.12.0 for frontend and Go 1.21+ for backend.