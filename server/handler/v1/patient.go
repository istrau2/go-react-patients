package handlerv1

import (
	"net/http"

	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/dto"
	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/store"
	"github.com/labstack/echo/v4"
)

type PatientHandlerGroup struct {
	Store *store.Store
}

func NewPatientHandlerGroup(st *store.Store) *PatientHandlerGroup {
	return &PatientHandlerGroup{
		Store: st,
	}
}

func (ag *PatientHandlerGroup) HandleRoutes(patient *echo.Group) {
	patient.GET("", ag.Match)
}

func (ag *PatientHandlerGroup) Match(c echo.Context) error {
	match := c.QueryParam("match")

	patients, count, err := ag.Store.Patient.ByParialName(match)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, dto.NewPatientListDTO_FM(patients, count))
}
