package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

type PersonDeleter interface {
	DeletePerson(ctx context.Context, lastName string) error
}

// HandleUpdatePerson is a Handler that updates a given person
//
// @Summary		Deletes Person
// @Description	Deletes person
// @Tags		courses
// @Accept		json
// @Produce		json
// @Success		200		{object}	handlers.responseMsg
// @Failure		500		{object}	handlers.responseErr
// @Router		/api/course	[PUT]

func HandleDeletePerson(logger *httplog.Logger, service PersonDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// setup
		ctx := r.Context()
		lastName := chi.URLParam(r, "name")

		err := service.DeletePerson(ctx, lastName)
		if err != nil {
			logger.Error("error deleting person", "error", err)
			encodeResponse(w, logger, http.StatusInternalServerError, responseErr{
				Error: "Error deleting person",
			})
			return
		}

		// return response
		encodeResponse(w, logger, http.StatusOK, responseMsg{
			Message: "Person deleted successfully",
		})
	}
}
