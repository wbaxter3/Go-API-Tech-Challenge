package handlers

import (
	"context"
	"go-api-tech-challenge/internal/models"
	"net/http"

	"github.com/go-chi/httplog/v2"
)

type courseLister interface {
	ListCourses(ctx context.Context) ([]models.Course, error)
}

// HandleListCourses is a Handler that returns a list of all courses.
//
// @Summary		List all courses
// @Description	List all courses
// @Tags		courses
// @Accept		json
// @Produce		json
// @Success		200		{object}	handlers.responseCourses
// @Failure		500		{object}	handlers.responseErr
// @Router		/courses	[GET]
func HandleListCourses(logger *httplog.Logger, service courseLister) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// setup
		ctx := r.Context()

		// get values from database
		courses, err := service.ListCourses(ctx)
		if err != nil {
			logger.Error("error getting all locations", "error", err)
			encodeResponse(w, logger, http.StatusInternalServerError, responseErr{
				Error: "Error retrieving data",
			})
			return
		}

		// return response
		coursesOut := mapMultipleOutput(courses)
		encodeResponse(w, logger, http.StatusOK, responseCourses{
			Courses: coursesOut,
		})
	}
}
