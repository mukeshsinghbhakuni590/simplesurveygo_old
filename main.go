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
	surveyHandler := servicehandlers.SurveyService{}

	// Serves the API content
	http.Handle("/api/v1/ping/", pingHandler)
	http.Handle("/api/v1/authenticate/", authHandler)
	http.Handle("/api/v1/validate/", sessionHandler)
	http.Handle("/api/v1/survey/", surveyHandler)
    
	// Start Server
	http.ListenAndServe(":3000", nil)
}
