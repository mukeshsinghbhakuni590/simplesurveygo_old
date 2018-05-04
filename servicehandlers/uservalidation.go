package servicehandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"simplesurveygo/dao"
	"simplesurveygo/validators"
)

type UserValidationHandler struct {
}

func (p UserValidationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p UserValidationHandler) Get(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p UserValidationHandler) Put(r *http.Request) SrvcRes {
	token, ok := r.URL.Query()["token"]
	if !ok || token[0] == "" {
		return SimpleBadRequest("parameters not passed properly")	
	} else {
		if !validators.Validate_session(token[0]) {
			return SimpleBadRequest("session not present")
		} else {
			dao.Remove_session(token[0])
			return Simple200OK("session deleted successfully")
		}
	}
}

func (p UserValidationHandler) Post(r *http.Request) SrvcRes {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var cred dao.UserCredentials
	err = json.Unmarshal(body, &cred)

	token := dao.AuthenticateUser(cred)

	if token == "" {
		return UnauthorizedAccess("Bad username or password")
	} else {
		return Response200OK(token)
	}

}
