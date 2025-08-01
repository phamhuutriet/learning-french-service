# 📘 Apprendre.ai - French Learning Service

A smart, AI-assisted French learning web app with flashcards, pronunciation tools, spaced repetition, and learning analytics.

## 🚀 Project Structure

```
learning-french-service/
├── cmd/server/          # Application entrypoints
├── internal/
│   ├── models/          # Database models (GORM)
│   ├── handlers/        # HTTP handlers (controllers)
│   ├── middleware/      # HTTP middleware
│   ├── services/        # Business logic
│   └── database/        # Database connection & config
├── pkg/
│   ├── utils/           # Utility functions
│   └── types/           # Shared types
├── configs/             # Configuration files
├── migrations/          # Database migrations
├── main.go              # Application entry point
└── api_design.md        # API documentation
```

## 🗄️ Database Models

The project includes complete database models for:

-   **Users** - User accounts and preferences
-   **Decks** - Vocabulary collections
-   **Words** - French vocabulary with AI enrichment
-   **Questions** - Auto-generated learning questions (3 per word)
-   **QuestionReviews** - SM-2 spaced repetition tracking
-   **QuestionAttempts** - Individual learning sessions with AI grading
-   **PronunciationAttempts** - Voice practice with AI analysis
-   **UserStats** - Learning analytics and insights
-   **LearningStreaks** - Motivation tracking

## 🛠️ Setup Instructions

### Prerequisites

-   Go 1.21+
-   PostgreSQL 14+

### Installation

1. **Clone the repository** (if not already done)

    ```bash
    git clone <repository-url>
    cd learning-french-service
    ```

2. **Install dependencies**

    ```bash
    go mod download
    ```

3. **Set up PostgreSQL**

    ```bash
    # Create database
    createdb apprendre_db

    # Or using psql
    psql -c "CREATE DATABASE apprendre_db;"
    ```

4. **Configure environment**

    ```bash
    cp configs/database.example.env .env
    # Edit .env with your database credentials
    ```

5. **Run the application**
    ```bash
    go run main.go
    ```

The application will:

-   Connect to PostgreSQL
-   Run automatic migrations
-   Create database indexes
-   Start the Gin server on port 8080

### Test the Setup

```bash
# Health check
curl http://localhost:8080/health

# API info
curl http://localhost:8080/api/v1/
```

## 📡 API Endpoints

See [api_design.md](./api_design.md) for complete API documentation including:

-   Authentication & user management
-   Deck & word management
-   AI enrichment workflow
-   Learning & question attempts
-   Statistics & analytics

## 🔧 Environment Variables

Key configuration options:

```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=apprendre_db

# Server
PORT=8080
GIN_MODE=debug
```

## 📋 Next Steps

1. **Implement handlers** - Create REST API endpoints
2. **Add authentication** - JWT middleware and user auth
3. **AI integration** - Connect OpenAI for word enrichment
4. **Pronunciation API** - Integrate pronunciation service
5. **Frontend** - Build React/Vue frontend
6. **Testing** - Add unit and integration tests

## 🎯 Features

-   ✅ Database models with relationships
-   ✅ GORM integration with PostgreSQL
-   ✅ Automatic migrations and indexing
-   ✅ Clean project structure
-   ⏳ REST API endpoints (next step)
-   ⏳ AI word enrichment
-   ⏳ Spaced repetition algorithm
-   ⏳ Pronunciation analysis
