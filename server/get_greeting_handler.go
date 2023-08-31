package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/towa1204/go-swagger-tutorial/server/gen/restapi/factory"
)

func GetGreeting(p factory.GetGreetingParams) middleware.Responder {
	payload := *p.Name
	return factory.NewGetGreetingOK().WithPayload(payload)
}
