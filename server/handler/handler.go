package handler

import (
	handlerv1 "bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/handler/v1"
	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/store"
	"github.com/labstack/echo/v4"
)

type HandlerGroup struct {
	Store *store.Store
	V1    *handlerv1.V1HandlerGroup
}

func NewHandlerGroup(st *store.Store) *HandlerGroup {
	return &HandlerGroup{
		Store: st,
		V1:    handlerv1.NewV1HandlerGroup(st),
	}
}

func (ag *HandlerGroup) HandleRoutes(all *echo.Group) {
	v1 := all.Group("/v1")
	ag.V1.HandleRoutes(v1)
}
