package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/distributors-admin/internal/config"
	"github.com/cifra-city/tokens"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func PlaceEmployeeDelete(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	_, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID) // TODO add auth for this user
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	placeId, err := uuid.Parse(chi.URLParam(r, "place_id"))
	if err != nil {
		log.Errorf("Failed to parse place id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Errorf("Failed to parse employee id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	employees, err := Server.MongoDB.PlacesEmployees.FilterByPlaceId(placeId).FilterByUserId(userId).Get(r.Context())
	if err != nil {
		log.Errorf("Failed to get place employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to get place employee"))
		return
	}

	if len(employees) != 1 {
		if len(employees) == 0 {
			httpkit.RenderErr(w, problems.NotFound())
		} else {
			log.Errorf("More than one employee found %s %s %s", placeId, userId, employees)
			httpkit.RenderErr(w, problems.InternalError())
		}
		return
	}

	_, err = Server.MongoDB.PlacesEmployees.FilterById(employees[0].ID).Delete(r.Context())
	if err != nil {
		log.Errorf("Failed to delete place employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to delete place employee"))
		return
	}

	httpkit.Render(w, http.StatusOK)
}
