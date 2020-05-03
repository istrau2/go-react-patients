package handlerv1

import (
	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/store"
	"github.com/labstack/echo/v4"
)

type V1HandlerGroup struct {
	Store   *store.Store
	Patient *PatientHandlerGroup
}

func NewV1HandlerGroup(st *store.Store) *V1HandlerGroup {
	return &V1HandlerGroup{
		Store:   st,
		Patient: NewPatientHandlerGroup(st),
	}
}

func (ag *V1HandlerGroup) HandleRoutes(v1 *echo.Group) {
	patient := v1.Group("/patient")
	ag.Patient.HandleRoutes(patient)
}
