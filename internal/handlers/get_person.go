package handlers

import (
	"context"
	"go-api-tech-challenge/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

type PersonGetter interface {
	GetPersonByName(ctx context.Context, name string) (models.Person, error)
}

// HandleListPersons is a Handler that returns a list of all persons.
//
// @Summary		List all persons
// @Description	List all persons
// @Tags		courses
// @Accept		json
// @Produce		json
// @Success		200		{object}	handlers.responsePersons
// @Failure		500		{object}	handlers.responseErr
// @Router		/api/course	[GET]
func HandleGetPersonByName(logger *httplog.Logger, service PersonGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// setup
		ctx := r.Context()
		name := chi.URLParam(r, "name")

		// get values from database
		persons, err := service.GetPersonByName(ctx, name)
		if err != nil {
			logger.Error("error getting all courses", "error", err)
			encodeResponse(w, logger, http.StatusInternalServerError, responseErr{
				Error: "Error retrieving data",
			})
			return
		}

		// return response
		personOut := mapOutputPerson(persons)
		encodeResponse(w, logger, http.StatusOK, responsePerson{
			Person: personOut,
		})
	}
}
