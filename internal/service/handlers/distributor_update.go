package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/distributors-admin/internal/config"
	"github.com/cifra-city/distributors-admin/internal/data/sql/repositories/sqlcore"
	"github.com/cifra-city/distributors-admin/internal/service/requests"
	"github.com/cifra-city/tokens"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func DistributorUpdate(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	req, err := requests.NewDistributorUpdate(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	distributorId, err := uuid.Parse(chi.URLParam(r, "distributor_id"))
	if err != nil {
		log.Errorf("Failed to parse distributor id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	employee, err := Server.SqlDB.DistributorsEmployees.GetByUser(r.Context(), distributorId, userId)
	if err != nil {
		log.Errorf("Failed to get distributor staff: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to get distributor staff"))
		return
	}

	if employee.Role != sqlcore.RolesAdmin && employee.Role != sqlcore.RolesOwner {
		log.Errorf("User is not allowed to update distributor")
		httpkit.RenderErr(w, problems.Forbidden("User is not allowed to update distributor"))
		return
	}

	distributor, err := Server.SqlDB.Distributors.UpdateName(r.Context(), distributorId, req.Data.Attributes.Name)
	if err != nil {
		log.Errorf("Failed to update distributor: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to update distributor"))
		return
	}

	httpkit.Render(w, NewDistributorResponse(distributor))
}
