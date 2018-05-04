package servicehandlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"simplesurveygo/dao"
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
	return ResponseNotImplemented()
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
