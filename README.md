# Quiz Log

A quiz creation app for reviewing what you've learned

## Features

### Core Features
- Create, edit, and delete quizzes
- Create questions (multiple choice, short answer, true/false)
- Take quizzes and get scored results

### Question Management
- Tag/category classification
- Difficulty settings (Easy/Medium/Hard)
- Import/export questions (JSON format)

### Learning Management
- Record learning history
- Track accuracy rates
- Review incorrect questions
- Category-based statistics

## Tech Stack

### Backend
- Go 1.25
- gqlgen (GraphQL)
- PostgreSQL
- sql-migrate (migration tool)

### Frontend
- React 18
- TypeScript
- Relay (GraphQL Client)
- React Router
- Vite

## Setup

### Prerequisites
- Go 1.21 or higher
- Node.js 18 or higher
- PostgreSQL 14 or higher

### Database Setup

```bash
# Create PostgreSQL database
createdb quizlog

# Run migrations
cd backend
make migrate-up
```

### Backend Setup

```bash
cd backend

# Install dependencies
make install

# Configure environment variables
cp .env.example .env
# Edit .env file to set database connection information

# Generate GraphQL code
make generate

# Start server
make run
```

Server runs at http://localhost:8080
GraphQL Playground: http://localhost:8080/

### Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Run Relay compiler
npm run relay

# Start development server
npm run dev
```

Frontend runs at http://localhost:5173

## Development

### Modifying GraphQL Schema

1. Edit `backend/graph/schema/schema.graphqls`
2. Regenerate code with `cd backend && make generate`
3. Regenerate Relay types with `cd frontend && npm run relay`

### Database Migrations

Migrations are managed using sql-migrate.

```bash
cd backend

# Check migration status
make migrate-status

# Run migrations
make migrate-up

# Rollback migrations
make migrate-down

# Create new migration
make migrate-new
# Or directly: sql-migrate new migration_name
```

New migration files are created in `backend/db/migrations/`.
Follow the sql-migrate format (`-- +migrate Up` and `-- +migrate Down`).

## Project Structure

```
quiz-log/
├── backend/
│   ├── db/              # Database connection and migrations
│   ├── graph/           # GraphQL schema and resolvers
│   ├── server.go        # Main server
│   └── Makefile
└── frontend/
    ├── src/
    │   ├── components/  # React components
    │   ├── App.tsx
    │   └── main.tsx
    └── package.json
```

## Next Steps

1. Implement resolvers (backend/graph/*.resolvers.go)
2. Implement frontend components
3. Question import/export functionality
4. Learning statistics visualization
