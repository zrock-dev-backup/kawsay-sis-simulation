package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// APIHandlers holds a reference to our in-memory data store.
type APIHandlers struct {
	Store *DataStore
}

// writeJSON is a helper to serialize data to JSON and write the HTTP response.
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// getOrgs handles requests for all organizations.
// @Summary Get all organizations
// @Description Retrieves a collection of all organizations, including schools and districts.
// @Tags Orgs
// @Produce json
// @Security ApiKeyAuth
// @Router /orgs [get]
func (h *APIHandlers) getOrgs(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string][]Org{"orgs": h.Store.Orgs})
}

// getOrg handles requests for a single organization by its SourcedId.
// @Summary Get a specific organization
// @Description Retrieves a single organization by its sourcedId.
// @Tags Orgs
// @Produce json
// @Param id path string true "SourcedId of the organization"
// @Success 200 {object} map[string]Org
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /orgs/{id} [get]
func (h *APIHandlers) getOrg(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, org := range h.Store.Orgs {
		if org.SourcedId == id {
			writeJSON(w, http.StatusOK, map[string]Org{"org": org})
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "Org not found"})
}

// getSchools handles requests for organizations of type 'school'.
// @Summary Get all schools
// @Description Retrieves a collection of all organizations with type 'school'.
// @Tags Schools
// @Produce json
// @Success 200 {object} map[string][]Org
// @Security ApiKeyAuth
// @Router /schools [get]
func (h *APIHandlers) getSchools(w http.ResponseWriter, r *http.Request) {
	var schools []Org
	for _, org := range h.Store.Orgs {
		if org.Type == "school" {
			schools = append(schools, org)
		}
	}
	writeJSON(w, http.StatusOK, map[string][]Org{"orgs": schools})
}

// getSchool handles requests for a single school by its SourcedId.
// @Summary Get a specific school
// @Description Retrieves a single school by its sourcedId.
// @Tags Schools
// @Produce json
// @Param id path string true "SourcedId of the school"
// @Success 200 {object} map[string]Org
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /schools/{id} [get]
func (h *APIHandlers) getSchool(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, org := range h.Store.Orgs {
		if org.SourcedId == id && org.Type == "school" {
			writeJSON(w, http.StatusOK, map[string]Org{"org": org})
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "School not found"})
}

// getUsers handles requests for all users.
// @Summary Get all users
// @Description Retrieves a collection of all users, including students and teachers.
// @Tags Users
// @Produce json
// @Success 200 {object} map[string][]User
// @Security ApiKeyAuth
// @Router /users [get]
func (h *APIHandlers) getUsers(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string][]User{"users": h.Store.Users})
}

// getUser handles requests for a single user by SourcedId.
// @Summary Get a specific user
// @Description Retrieves a single user by their sourcedId.
// @Tags Users
// @Produce json
// @Param id path string true "SourcedId of the user"
// @Success 200 {object} map[string]User
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /users/{id} [get]
func (h *APIHandlers) getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, user := range h.Store.Users {
		if user.SourcedId == id {
			writeJSON(w, http.StatusOK, map[string]User{"user": user})
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "User not found"})
}

// getTeachers handles requests for users with role 'teacher'.
// @Summary Get all teachers
// @Description Retrieves a collection of all users with the role 'teacher'.
// @Tags Teachers
// @Produce json
// @Success 200 {object} map[string][]User
// @Security ApiKeyAuth
// @Router /teachers [get]
func (h *APIHandlers) getTeachers(w http.ResponseWriter, r *http.Request) {
	var teachers []User
	for _, user := range h.Store.Users {
		if user.Role == "teacher" {
			teachers = append(teachers, user)
		}
	}
	writeJSON(w, http.StatusOK, map[string][]User{"users": teachers})
}

// getTeacher handles requests for a single teacher by SourcedId.
// @Summary Get a specific teacher
// @Description Retrieves a single teacher by their sourcedId.
// @Tags Teachers
// @Produce json
// @Param id path string true "SourcedId of the teacher"
// @Success 200 {object} map[string]User
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /teachers/{id} [get]
func (h *APIHandlers) getTeacher(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, user := range h.Store.Users {
		if user.SourcedId == id && user.Role == "teacher" {
			writeJSON(w, http.StatusOK, map[string]User{"user": user})
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "Teacher not found"})
}

// getStudents handles requests for users with role 'student'.
// @Summary Get all students
// @Description Retrieves a collection of all users with the role 'student'.
// @Tags Students
// @Produce json
// @Success 200 {object} map[string][]User
// @Security ApiKeyAuth
// @Router /students [get]
func (h *APIHandlers) getStudents(w http.ResponseWriter, r *http.Request) {
	var students []User
	for _, user := range h.Store.Users {
		if user.Role == "student" {
			students = append(students, user)
		}
	}
	writeJSON(w, http.StatusOK, map[string][]User{"users": students})
}

// getStudent handles requests for a single student by SourcedId.
// @Summary Get a specific student
// @Description Retrieves a single student by their sourcedId.
// @Tags Students
// @Produce json
// @Param id path string true "SourcedId of the student"
// @Success 200 {object} map[string]User
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /students/{id} [get]
func (h *APIHandlers) getStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, user := range h.Store.Users {
		if user.SourcedId == id && user.Role == "student" {
			writeJSON(w, http.StatusOK, map[string]User{"user": user})
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "Student not found"})
}

// getCourses handles requests for all courses.
// @Summary Get all courses
// @Description Retrieves a collection of all courses from the catalog.
// @Tags Courses
// @Produce json
// @Success 200 {object} map[string][]Course
// @Security ApiKeyAuth
// @Router /courses [get]
func (h *APIHandlers) getCourses(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string][]Course{"courses": h.Store.Courses})
}

// getCourse handles requests for a single course by SourcedId.
// @Summary Get a specific course
// @Description Retrieves a single course by its sourcedId.
// @Tags Courses
// @Produce json
// @Param id path string true "SourcedId of the course"
// @Success 200 {object} map[string]Course
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /courses/{id} [get]
func (h *APIHandlers) getCourse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, course := range h.Store.Courses {
		if course.SourcedId == id {
			writeJSON(w, http.StatusOK, map[string]Course{"course": course})
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "Course not found"})
}

// getClasses handles requests for all classes.
// @Summary Get all classes
// @Description Retrieves a collection of all scheduled classes.
// @Tags Classes
// @Produce json
// @Success 200 {object} map[string][]Class
// @Security ApiKeyAuth
// @Router /classes [get]
func (h *APIHandlers) getClasses(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string][]Class{"classes": h.Store.Classes})
}

// getClass handles requests for a single class by SourcedId.
// @Summary Get a specific class
// @Description Retrieves a single class by its sourcedId.
// @Tags Classes
// @Produce json
// @Param id path string true "SourcedId of the class"
// @Success 200 {object} map[string]Class
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /classes/{id} [get]
func (h *APIHandlers) getClass(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, class := range h.Store.Classes {
		if class.SourcedId == id {
			writeJSON(w, http.StatusOK, map[string]Class{"class": class})
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "Class not found"})
}

// getCategoriesForClass handles requests for categories for a given class.
// @Summary Get categories for a class
// @Description Retrieves a collection of grading categories for a given class.
// @Tags Classes
// @Produce json
// @Param id path string true "SourcedId of the class"
// @Success 200 {object} map[string][]Category
// @Security ApiKeyAuth
// @Router /classes/{id}/categories [get]
func (h *APIHandlers) getCategoriesForClass(w http.ResponseWriter, r *http.Request) {
	// In this mock, categories are global, not class-specific.
	// A real implementation would filter based on the class ID.
	writeJSON(w, http.StatusOK, map[string][]Category{"categories": h.Store.Categories})
}

// getEnrollments handles requests for all enrollments.
// @Summary Get all enrollments
// @Description Retrieves a collection of all user enrollments in classes.
// @Tags Enrollments
// @Produce json
// @Success 200 {object} map[string][]Enrollment
// @Security ApiKeyAuth
// @Router /enrollments [get]
func (h *APIHandlers) getEnrollments(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string][]Enrollment{"enrollments": h.Store.Enrollments})
}

// getEnrollment handles requests for a single enrollment by SourcedId.
// @Summary Get a specific enrollment
// @Description Retrieves a single enrollment by its sourcedId.
// @Tags Enrollments
// @Produce json
// @Param id path string true "SourcedId of the enrollment"
// @Success 200 {object} map[string]Enrollment
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /enrollments/{id} [get]
func (h *APIHandlers) getEnrollment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, enrollment := range h.Store.Enrollments {
		if enrollment.SourcedId == id {
			writeJSON(w, http.StatusOK, map[string]Enrollment{"enrollment": enrollment})
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "Enrollment not found"})
}

// getTerms handles requests for academic sessions of type 'term'.
// @Summary Get all terms
// @Description Retrieves a collection of all academic sessions with type 'term'.
// @Tags Academic Sessions
// @Produce json
// @Success 200 {object} map[string][]AcademicSession
// @Security ApiKeyAuth
// @Router /terms [get]
func (h *APIHandlers) getTerms(w http.ResponseWriter, r *http.Request) {
	var terms []AcademicSession
	for _, session := range h.Store.AcademicSessions {
		if session.Type == "term" {
			terms = append(terms, session)
		}
	}
	writeJSON(w, http.StatusOK, map[string][]AcademicSession{"academicSessions": terms})
}

// getTerm handles requests for a single term by SourcedId.
// @Summary Get a specific term
// @Description Retrieves a single term by its sourcedId.
// @Tags Academic Sessions
// @Produce json
// @Param id path string true "SourcedId of the term"
// @Success 200 {object} map[string]AcademicSession
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /terms/{id} [get]
func (h *APIHandlers) getTerm(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, session := range h.Store.AcademicSessions {
		if session.SourcedId == id && session.Type == "term" {
			writeJSON(w, http.StatusOK, map[string]AcademicSession{"academicSession": session})
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "Term not found"})
}

// getAcademicSessions handles requests for all academic sessions.
// @Summary Get all academic sessions
// @Description Retrieves a collection of all academic sessions of any type.
// @Tags Academic Sessions
// @Produce json
// @Success 200 {object} map[string][]AcademicSession
// @Security ApiKeyAuth
// @Router /academicSessions [get]
func (h *APIHandlers) getAcademicSessions(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string][]AcademicSession{"academicSessions": h.Store.AcademicSessions})
}

// getAcademicSession handles requests for a single academic session by SourcedId.
// @Summary Get a specific academic session
// @Description Retrieves a single academic session by its sourcedId.
// @Tags Academic Sessions
// @Produce json
// @Param id path string true "SourcedId of the academic session"
// @Success 200 {object} map[string]AcademicSession
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /academicSessions/{id} [get]
func (h *APIHandlers) getAcademicSession(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, session := range h.Store.AcademicSessions {
		if session.SourcedId == id {
			writeJSON(w, http.StatusOK, map[string]AcademicSession{"academicSession": session})
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "Academic Session not found"})
}

// getGradingPeriods handles requests for academic sessions of type 'gradingPeriod'.
// @Summary Get all grading periods
// @Description Retrieves a collection of all academic sessions with type 'gradingPeriod'.
// @Tags Academic Sessions
// @Produce json
// @Success 200 {object} map[string][]AcademicSession
// @Security ApiKeyAuth
// @Router /gradingPeriods [get]
func (h *APIHandlers) getGradingPeriods(w http.ResponseWriter, _ *http.Request) {
	var periods []AcademicSession
	for _, session := range h.Store.AcademicSessions {
		if session.Type == "gradingPeriod" {
			periods = append(periods, session)
		}
	}
	writeJSON(w, http.StatusOK, map[string][]AcademicSession{"academicSessions": periods})
}

// getGradingPeriod handles requests for a single grading period by SourcedId.
// @Summary Get a specific grading period
// @Description Retrieves a single grading period by its sourcedId.
// @Tags Academic Sessions
// @Produce json
// @Param id path string true "SourcedId of the grading period"
// @Success 200 {object} map[string]AcademicSession
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /gradingPeriods/{id} [get]
func (h *APIHandlers) getGradingPeriod(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for _, session := range h.Store.AcademicSessions {
		if session.SourcedId == id && session.Type == "gradingPeriod" {
			writeJSON(w, http.StatusOK, map[string]AcademicSession{"academicSession": session})
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "Grading Period not found"})
}
