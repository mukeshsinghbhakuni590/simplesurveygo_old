package main

import (
	"net/http"
	"simplesurveygo/servicehandlers"
)

func main() {

	// Serves the html pages
	http.Handle("/", http.FileServer(http.Dir("./static")))

	pingHandler := servicehandlers.PingHandler{}
	authHandler := servicehandlers.UserValidationHandler{}
	sessionHandler := servicehandlers.SessionHandler{}
	userServiceHandler := servicehandlers.UserServiceHandler{}
	surveyService := servicehandlers.SurveyService{}
	surveyOperation := servicehandlers.SurveyOperations{}
	
	// Serves the API content
	http.Handle("/api/v1/ping/", pingHandler)
	http.Handle("/api/v1/authenticate/", authHandler)
	http.Handle("/api/v1/validate/", sessionHandler)
	http.Handle("/api/v1/user/", userServiceHandler)
	http.Handle("/api/v1/survey/", surveyService)
	http.Handle("/api/v1/surveyOpr/", surveyOperation)
	// Start Server
	http.ListenAndServe(":3000", nil)
}
