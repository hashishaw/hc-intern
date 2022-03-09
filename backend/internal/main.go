package main

import (
  "github.com/hashishaw/hc-intern/backend/pkg/swagger/server/restapi"
  "github.com/hashishaw/hc-intern/backend/pkg/swagger/server/restapi/operations"
  "log"

  "github.com/go-openapi/loads"
  "github.com/go-openapi/runtime/middleware"
)

func main() {

  // Initialize Swagger
  swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
  if err != nil {
    log.Fatalln(err)
  }

  api := operations.NewInterviewAPIAPI(swaggerSpec)
  server := restapi.NewServer(api)

  defer func() {
    if err := server.Shutdown(); err != nil {
      // error handle
      log.Fatalln(err)
    }
  }()

  server.Port = 8080

  api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)

  api.GetAPIInterviewIDHandler = operations.GetAPIInterviewIDHandlerFunc(GetInterview)

  api.PostAPIScheduleCandidateIDHandler = operations.PostAPIScheduleCandidateIDHandlerFunc(GetInterviewIdByCandidateId)

  // Start server which listening
  if err := server.Serve(); err != nil {
    log.Fatalln(err)
  }
}

//Health route returns OK
func Health(operations.CheckHealthParams) middleware.Responder {
  return operations.NewCheckHealthOK().WithPayload("OK")
}

func GetInterview(interview operations.GetAPIInterviewIDParams) middleware.Responder {
  if interview.ID == "foobar" {
    return operations.NewGetAPIInterviewIDNoContent()
  }
  return operations.NewGetAPIInterviewIDNotFound()
}

func GetInterviewIdByCandidateId(p operations.PostAPIScheduleCandidateIDParams) middleware.Responder {

  var interviewId string
  if p.CandidateID != "" {
    // TODO: get candidate info from greenhouse
    interviewId = "foobar"
  }

  return operations.NewPostAPIScheduleCandidateIDOK().WithPayload(interviewId)
}
