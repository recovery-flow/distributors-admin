package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/distributors-admin/internal/config"
	"github.com/cifra-city/distributors-admin/internal/data/sql/repositories/sqlcore"
	"github.com/cifra-city/distributors-admin/internal/service/requests"
	"github.com/cifra-city/distributors-admin/resources"
	"github.com/cifra-city/tokens"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func EmployeeAdd(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	req, err := requests.NewEmployeeAdd(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	distributorId, err := uuid.Parse(req.Data.Id)
	if err != nil {
		log.Errorf("Failed to parse distributor id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	curEmployee, err := Server.SqlDB.DistributorsEmployees.GetByUser(r.Context(), distributorId, userID)
	if err != nil {
		log.Errorf("Failed to get distributor employee: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to get distributor employee"))
		return
	}
	if curEmployee.Role != sqlcore.RolesModerator && curEmployee.Role != sqlcore.RolesAdmin && curEmployee.Role != sqlcore.RolesOwner {
		httpkit.RenderErr(w, problems.Forbidden("User is have not permission"))
		return
	}

	if req.Data.Attributes.Role == string(sqlcore.RolesOwner) {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("You can't add owner user"))...)
		return
	}

	newEmployee, err := Server.SqlDB.DistributorsEmployees.Create(r.Context(), distributorId, userID, req.Data.Attributes.Role)
	if err != nil {
		if err.Error() == "role must be one of: admin, moderator, staff, member" {
			log.Errorf("Failed to create distributor staff: %v", err)
			httpkit.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		log.Errorf("Failed to create distributor staff: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to create distributor staff"))
		return
	}

	log.Infof("Staff added to distributor %v", distributorId)
	httpkit.Render(w, NewEmployeeResponse(newEmployee))
}

func NewEmployeeResponse(employees sqlcore.DistributorsEmployee) *resources.Employee {
	return &resources.Employee{
		Data: resources.EmployeeData{
			Id:   employees.ID.String(),
			Type: resources.DistributorEmployeeType,
			Attributes: resources.EmployeeDataAttributes{
				UserId:        employees.UserID.String(),
				Role:          string(employees.Role),
				DistributorId: employees.DistributorsID.String(),
				CreatedAt:     employees.CreatedAt.Time,
			},
		},
	}
}