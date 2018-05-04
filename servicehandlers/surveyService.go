package servicehandlers

import (
	"encoding/json"
	"net/http"
	"simplesurveygo/dao"
	"simplesurveygo/validators"
)

type SurveyService struct {
}

func (p  SurveyService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p  SurveyService) Get(r *http.Request) SrvcRes {
	token, ok1 := r.URL.Query()["token"]
	session , ok2 := r.URL.Query()["session"]
	surveyid , ok3 := r.URL.Query()["surveyid"]
    if (!ok1 || !ok2 || !ok3) {
		return SimpleBadRequest("parameters not passed") 
	} else {
        if validators.Validate_user_session(token[0],session[0]) {
		   survey_info := dao.Get_survey_data_by_id(surveyid[0])
		   s_info, _ := json.Marshal(survey_info)
		   return Simple200OK(string(s_info))
		} else {
			return UnauthorizedAccess("you have to login again")
		} 
	}
}

func (p  SurveyService) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

func (p  SurveyService) Post(r *http.Request) SrvcRes {
	decoder := json.NewDecoder(r.Body)
	var t dao.SurveyPostStruct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	s_info, _ := json.Marshal(dao.Create_survey(t))
	return Simple200OK(string(s_info)) 
}

