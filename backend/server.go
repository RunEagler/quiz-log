package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"quiz-log/db"
	"quiz-log/graph"
	"quiz-log/graph/resolvers"
	"quiz-log/services"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Database connection
	dbConn, err := db.Connect(db.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     5432,
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "quizlog"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	// Initialize services
	quizService := services.NewQuizService(dbConn)
	questionService := services.NewQuestionService(dbConn)
	tagService := services.NewTagService(dbConn)
	attemptService := services.NewAttemptService(dbConn, questionService)
	statisticsService := services.NewStatisticsService(dbConn, attemptService)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &resolvers.Resolver{
			DB:                dbConn,
			QuizService:       quizService,
			QuestionService:   questionService,
			TagService:        tagService,
			AttemptService:    attemptService,
			StatisticsService: statisticsService,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
