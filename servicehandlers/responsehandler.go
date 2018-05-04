package servicehandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type SrvcRes struct {
	Code     int
	Response interface{}
	Message  string
	Headers  map[string]string
}

func marshalResponse(r interface{}) ([]byte, error) {
	return json.MarshalIndent(r, "", "")
}

func (s SrvcRes) RenderResponse(w http.ResponseWriter) {
	if s.Headers == nil {
		s.Headers = map[string]string{"Content-Type": "application/json"}
	}
	for h, val := range s.Headers {
		w.Header().Set(h, val)
	}

	formatted := bson.M{
		"responseData": s.Response,
		"message":      s.Message,
		"status":       true}

	data, _ := marshalResponse(formatted)
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))

	/* HTTP STATUS CODE*/
	w.WriteHeader(s.Code)

	/* HTTP Response Body*/
	fmt.Fprint(w, string(data))
}

func Simple200OK(message string) SrvcRes {
	return SrvcRes{http.StatusOK, "{}", message, nil}
}

func SimpleBadRequest(message string) SrvcRes {
	return SrvcRes{http.StatusBadRequest, "{}", message, nil}
}

func InternalServerError(message string) SrvcRes {
	return SrvcRes{http.StatusInternalServerError, "{}", message, nil}
}

func Response200OK(response interface{}) SrvcRes {
	return SrvcRes{http.StatusOK, response, "OK", nil}
}

func ResponseNotImplemented() SrvcRes {
	return SrvcRes{http.StatusNotImplemented, "{}", "Method not implementd", nil}
}

func UnauthorizedAccess(message string) SrvcRes {
	return SrvcRes{http.StatusUnauthorized, "{}", message, nil}
}
