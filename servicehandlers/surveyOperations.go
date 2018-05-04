package servicehandlers

import (
	"encoding/json"
	"net/http"
	"simplesurveygo/dao"
	"simplesurveygo/validators"
)


type SurveyOperations struct {
}

func (p  SurveyOperations) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p  SurveyOperations) Get(r *http.Request) SrvcRes {
	token ,ok1 := r.URL.Query()["token"]
	if !ok1 || token[0] == "" {
       return SimpleBadRequest("parameters not passed properly") 
	} else {
       if !validators.Validate_session(token[0]) {
		 return UnauthorizedAccess("you have to login again")  
	   } else {
         session_info := dao.GetSessionDetails(token[0])
		 user_survey := dao.Get_user_survey(session_info.Username)
		 return Response200OK(user_survey)
	   } 
	}
}	  

func (p  SurveyOperations) Put(r *http.Request) SrvcRes {
	decoder := json.NewDecoder(r.Body)
	var t dao.SurveyOprPut
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	if !validators.Validate_session(t.Token) {
		return  UnauthorizedAccess("you have to login again")
	} else {
	   user_surveys := dao.Update_user_surveys(t)
	   return Response200OK(user_surveys)
	}
}

func (p  SurveyOperations) Post(r *http.Request) SrvcRes {
	decoder := json.NewDecoder(r.Body)
	var t dao.SurveyOprPost
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	response := dao.Create_link(t)
	return Response200OK(response)  	
}
