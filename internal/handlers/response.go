package handlers

import (
	"encoding/json"
	"go-api-tech-challenge/internal/models"
	"net/http"

	"github.com/go-chi/httplog/v2"
)

type outputCourse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// mapOutput maps a models.Course struct to an outputCourse struct.
func mapOutput(course models.Course) outputCourse {
	return outputCourse{
		ID:   course.ID,
		Name: course.Name,
	}
}

// mapMultipleOutput maps a slice of []models.Course to a slice of []outputCourse.
func mapMultipleOutput(course []models.Course) []outputCourse {
	coursesOut := make([]outputCourse, len(course))
	for i := 0; i < len(course); i++ {
		courseOut := mapOutput(course[i])
		coursesOut[i] = courseOut
	}

	return coursesOut
}

type responseCourse struct {
	Course outputCourse `json:"course"`
}

type responseCourses struct {
	Courses []outputCourse `json:"courses"`
}

type responseMsg struct {
	Message string `json:"message"`
}

type responseID struct {
	ObjectID int `json:"object_id"`
}

type responseErr struct {
	Error            string    `json:"error,omitempty"`
	ValidationErrors []problem `json:"validation_errors,omitempty"`
}

// encodeResponse encodes data as a JSON response.
func encodeResponse(w http.ResponseWriter, logger *httplog.Logger, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error("Error while marshaling data", "err", err, "data", data)
		http.Error(w, `{"Error": "Internal server error"}`, http.StatusInternalServerError)
	}
}
