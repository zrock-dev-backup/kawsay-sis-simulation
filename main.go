package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "go-oneroster-mock/docs" // Import generated docs
)

// @title OneRoster Mock API
// @version 1.0
// @description This is a mock server for the OneRoster v1p1 API specification.
// @description It serves a large, procedurally generated, in-memory dataset for testing purposes.

// @contact.name Kawsay Developer Agent
// @contact.url http://www.example.com
// @contact.email dev.agent@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5100
// @BasePath /ims/oneroster/v1p1

// --- AÑADE ESTAS LÍNEAS PARA LA AUTENTICACIÓN ---
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// --------------------------------------------------

func main() {
	log.Println("Generating mock data store...")
	store := NewDataStore()
	log.Printf("Data generation complete. %d users, %d orgs, %d classes loaded.", len(store.Users), len(store.Orgs), len(store.Classes))

	handlers := &APIHandlers{Store: store}

	r := chi.NewRouter()

	// --- Middleware ---
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// CORS for frontend development
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173", "http://localhost:5100"}, // Add your C# dev server port if needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// --- Mock Authentication Middleware ---
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Swagger UI assets don't need auth
			if strings.HasPrefix(r.URL.Path, "/swagger/") {
				next.ServeHTTP(w, r)
				return
			}
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Unauthorized: Missing Authorization header", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	// --- API Routes ---
	r.Route("/ims/oneroster/v1p1", func(r chi.Router) {
		// Orgs & Schools
		r.Get("/orgs", handlers.getOrgs)
		r.Get("/orgs/{id}", handlers.getOrg)
		r.Get("/schools", handlers.getSchools)
		r.Get("/schools/{id}", handlers.getSchool)

		// Users, Teachers, Students
		r.Get("/users", handlers.getUsers)
		r.Get("/users/{id}", handlers.getUser)
		r.Get("/teachers", handlers.getTeachers)
		r.Get("/teachers/{id}", handlers.getTeacher)
		r.Get("/students", handlers.getStudents)
		r.Get("/students/{id}", handlers.getStudent)

		// Courses & Classes
		r.Get("/courses", handlers.getCourses)
		r.Get("/courses/{id}", handlers.getCourse)
		r.Get("/classes", handlers.getClasses)
		r.Get("/classes/{id}", handlers.getClass)
		r.Get("/classes/{id}/categories", handlers.getCategoriesForClass)

		// Enrollments
		r.Get("/enrollments", handlers.getEnrollments)
		r.Get("/enrollments/{id}", handlers.getEnrollment)

		// Academic Sessions, Terms, Grading Periods
		r.Get("/terms", handlers.getTerms)
		r.Get("/terms/{id}", handlers.getTerm)
		r.Get("/academicSessions", handlers.getAcademicSessions)
		r.Get("/academicSessions/{id}", handlers.getAcademicSession)
		r.Get("/gradingPeriods", handlers.getGradingPeriods)
		r.Get("/gradingPeriods/{id}", handlers.getGradingPeriod)
	})

	// --- Swagger UI Route ---
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Println("Starting server on :5100...")
	if err := http.ListenAndServe(":5100", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
