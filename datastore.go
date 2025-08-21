package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// BaseModel provides fields common to most OneRoster objects.
// @Description Common fields for most OneRoster objects.
type BaseModel struct {
	SourcedId        string    `json:"sourcedId"`
	Status           string    `json:"status"`
	DateLastModified time.Time `json:"dateLastModified"`
	Metadata         any       `json:"metadata"`
}

// GUIDRef is a reference to another object in the system.
// @Description A reference to another OneRoster object.
type GUIDRef struct {
	Href      string `json:"href"`
	SourcedId string `json:"sourcedId"`
	Type      string `json:"type"`
}

// Org represents an organization, like a school or district.
// @Description Represents an organization, such as a school or district.
type Org struct {
	BaseModel
	Name       string    `json:"name"`
	Type       string    `json:"type"` // e.g., 'school', 'district'
	Identifier string    `json:"identifier"`
	Parent     *GUIDRef  `json:"parent,omitempty"`
	Children   []GUIDRef `json:"children,omitempty"`
}

// User represents a person, like a student or teacher.
// @Description Represents a person within the system, such as a student or a teacher.
type User struct {
	BaseModel
	Username    string    `json:"username"`
	UserIds     []any     `json:"userIds"`
	EnabledUser bool      `json:"enabledUser"`
	GivenName   string    `json:"givenName"`
	FamilyName  string    `json:"familyName"`
	Role        string    `json:"role"` // 'student', 'teacher'
	Identifier  string    `json:"identifier"`
	Email       string    `json:"email"`
	Orgs        []GUIDRef `json:"orgs"`
}

// Course represents a course catalog entry.
// @Description Represents a course in the course catalog.
type Course struct {
	BaseModel
	Title        string    `json:"title"`
	SchoolYear   *GUIDRef  `json:"schoolYear,omitempty"`
	CourseCode   string    `json:"courseCode"`
	Grades       []string  `json:"grades,omitempty"`
	Subjects     []string  `json:"subjects,omitempty"`
	SubjectCodes []string  `json:"subjectCodes,omitempty"`
	Resources    []GUIDRef `json:"resources,omitempty"`
}

// Class represents a specific instance of a course.
// @Description Represents a specific instance of a course for a particular term and school.
type Class struct {
	BaseModel
	Title        string    `json:"title"`
	ClassCode    string    `json:"classCode"`
	ClassType    string    `json:"classType"` // 'homeroom', 'scheduled'
	Location     string    `json:"location"`
	Grades       []string  `json:"grades"`
	Subjects     []string  `json:"subjects"`
	Course       GUIDRef   `json:"course"`
	School       GUIDRef   `json:"school"`
	Terms        []GUIDRef `json:"terms"`
	SubjectCodes []string  `json:"subjectCodes,omitempty"`
	Periods      []string  `json:"periods,omitempty"`
	Resources    []GUIDRef `json:"resources,omitempty"`
}

// Enrollment links a user to a class in a specific role.
// @Description Represents the link between a user and a class for a specific role.
type Enrollment struct {
	BaseModel
	User      GUIDRef  `json:"user"`
	Class     GUIDRef  `json:"class"`
	School    GUIDRef  `json:"school"`
	Role      string   `json:"role"`
	Primary   bool     `json:"primary"`
	BeginDate string   `json:"beginDate"`
	EndDate   string   `json:"endDate"`
}

// AcademicSession represents a time period like a term or semester.
// @Description Represents a time period in the academic calendar, such as a term, semester, or grading period.
type AcademicSession struct {
	BaseModel
	Title     string    `json:"title"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	Type      string    `json:"type"` // 'gradingPeriod', 'semester', 'schoolYear', 'term'
	Parent    *GUIDRef  `json:"parent,omitempty"`
	Children  []GUIDRef `json:"children,omitempty"`
	SchoolYear string   `json:"schoolYear"`
}

// Category represents a grading category for a class.
// @Description Represents a grading category within a class.
type Category struct {
	BaseModel
	Title  string `json:"title"`
	Weight int    `json:"weight"`
}

// DataStore holds all our in-memory mock data.
type DataStore struct {
	Orgs             []Org
	Users            []User
	Courses          []Course
	Classes          []Class
	Enrollments      []Enrollment
	AcademicSessions []AcademicSession
	Categories       []Category
}

// NewDataStore creates and populates a DataStore with a large volume of mock data.
func NewDataStore() *DataStore {
	ds := &DataStore{}

	// --- Generate Orgs (Schools) ---
	for i := 1; i <= 10; i++ {
		schoolId := uuid.New().String()
		ds.Orgs = append(ds.Orgs, Org{
			BaseModel:  BaseModel{SourcedId: schoolId, Status: "active", DateLastModified: time.Now()},
			Name:       fmt.Sprintf("School #%d", i),
			Type:       "school",
			Identifier: fmt.Sprintf("SCH%03d", i),
		})
	}

	// --- Generate Users (Students & Teachers) ---
	// 1000 Students
	for i := 1; i <= 1000; i++ {
		userId := uuid.New().String()
		school := ds.Orgs[i%len(ds.Orgs)] // Assign student to a school
		ds.Users = append(ds.Users, User{
			BaseModel:   BaseModel{SourcedId: userId, Status: "active", DateLastModified: time.Now()},
			Username:    fmt.Sprintf("student%d", i),
			EnabledUser: true,
			GivenName:   "Student",
			FamilyName:  fmt.Sprintf("User%d", i),
			Role:        "student",
			Identifier:  fmt.Sprintf("STU%04d", i),
			Email:       fmt.Sprintf("student%d@example.com", i),
			Orgs:        []GUIDRef{{Href: "/orgs/" + school.SourcedId, SourcedId: school.SourcedId, Type: "org"}},
		})
	}
	// 250 Teachers
	for i := 1; i <= 250; i++ {
		userId := uuid.New().String()
		school := ds.Orgs[i%len(ds.Orgs)] // Assign teacher to a school
		ds.Users = append(ds.Users, User{
			BaseModel:   BaseModel{SourcedId: userId, Status: "active", DateLastModified: time.Now()},
			Username:    fmt.Sprintf("teacher%d", i),
			EnabledUser: true,
			GivenName:   "Teacher",
			FamilyName:  fmt.Sprintf("User%d", i),
			Role:        "teacher",
			Identifier:  fmt.Sprintf("TCH%04d", i),
			Email:       fmt.Sprintf("teacher%d@example.com", i),
			Orgs:        []GUIDRef{{Href: "/orgs/" + school.SourcedId, SourcedId: school.SourcedId, Type: "org"}},
		})
	}

	// --- Generate Academic Sessions (Terms) ---
	for i := 1; i <= 4; i++ {
		termId := uuid.New().String()
		ds.AcademicSessions = append(ds.AcademicSessions, AcademicSession{
			BaseModel: BaseModel{SourcedId: termId, Status: "active", DateLastModified: time.Now()},
			Title:     fmt.Sprintf("Fall Semester 202%d", i+4),
			Type:      "term",
			StartDate: fmt.Sprintf("202%d-09-01", i+4),
			EndDate:   fmt.Sprintf("202%d-12-20", i+4),
			SchoolYear: fmt.Sprintf("202%d", i+4),
		})
	}

	// --- Generate Courses ---
	for i := 1; i <= 50; i++ {
		courseId := uuid.New().String()
		ds.Courses = append(ds.Courses, Course{
			BaseModel:  BaseModel{SourcedId: courseId, Status: "active", DateLastModified: time.Now()},
			Title:      fmt.Sprintf("Course %d", i),
			CourseCode: fmt.Sprintf("CRS%03d", i),
			Subjects:   []string{"General"},
		})
	}

	// --- Generate Classes ---
	for i := 1; i <= 500; i++ {
		classId := uuid.New().String()
		course := ds.Courses[i%len(ds.Courses)]
		school := ds.Orgs[i%len(ds.Orgs)]
		term := ds.AcademicSessions[i%len(ds.AcademicSessions)]
		ds.Classes = append(ds.Classes, Class{
			BaseModel: BaseModel{SourcedId: classId, Status: "active", DateLastModified: time.Now()},
			Title:     course.Title,
			ClassCode: fmt.Sprintf("%s-S%d", course.CourseCode, i),
			ClassType: "scheduled",
			Course:    GUIDRef{Href: "/courses/" + course.SourcedId, SourcedId: course.SourcedId, Type: "course"},
			School:    GUIDRef{Href: "/schools/" + school.SourcedId, SourcedId: school.SourcedId, Type: "school"},
			Terms:     []GUIDRef{{Href: "/terms/" + term.SourcedId, SourcedId: term.SourcedId, Type: "term"}},
			Grades:    []string{"10"},
			Subjects:  []string{"General"},
		})
	}

	// --- Generate Categories ---
	ds.Categories = append(ds.Categories,
		Category{BaseModel: BaseModel{SourcedId: uuid.New().String()}, Title: "Homework", Weight: 20},
		Category{BaseModel: BaseModel{SourcedId: uuid.New().String()}, Title: "Exams", Weight: 50},
		Category{BaseModel: BaseModel{SourcedId: uuid.New().String()}, Title: "Participation", Weight: 30},
	)

	return ds
}
