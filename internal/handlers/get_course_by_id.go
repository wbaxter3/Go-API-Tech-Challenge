package handlers

import (
	"context"
	"go-api-tech-challenge/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/httplog/v2"
)

type courseGetter interface {
	GetCourseByID(ctx context.Context, ID int) (models.Course, error)
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
func HandleGetCourseByID(logger *httplog.Logger, service courseGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		// setup
		idString := chi.URLParam(r, "ID")
		ID, err := strconv.Atoi(idString)
		if err != nil {
			logger.Error("error getting ID", "error", err)
			encodeResponse(w, logger, http.StatusBadRequest, responseErr{
				Error: "Not a valid ID",
			})
			return
		}
		// get values from database
		course, err := service.GetCourseByID(ctx, ID)
		if err != nil {
			log.Println(err)
			logger.Error("error getting all courses", "error", err)
			encodeResponse(w, logger, http.StatusInternalServerError, responseErr{
				Error: "Error retrieving data",
			})
			return
		}

		// return response
		coursesOut := mapOutput(course)
		encodeResponse(w, logger, http.StatusOK, responseCourse{
			Course: coursesOut,
		})
	}
}
