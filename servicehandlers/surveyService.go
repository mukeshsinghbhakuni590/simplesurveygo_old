package servicehandlers

import (
	"encoding/json"
	"net/http"
	"simplesurveygo/dao"
	"simplesurveygo/validators"
)


type SurveyService struct {
}

func (p SurveyService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p SurveyService) Get(r *http.Request) SrvcRes {
	option ,ok1 := r.URL.Query()["option"]
	token  ,ok2 := r.URL.Query()["token"]

	if (!ok1 || !ok2) || (option[0] == "" || token[0] == "") {
		return SimpleBadRequest("parameters not passed properly")
	} else {
        if !validators.Validate_token(token[0]) {
			return  UnauthorizedAccess("you have to login again")
		} else {
		   if option[0] == "all" {
			   all_survey := dao.Get_all_survey()
			   return Response200OK(all_survey)
		   } else {
			   survey_info := dao.Get_survey(option[0])
			   return Response200OK(survey_info)
		   }		
		}
	}
}	  

func (p SurveyService) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
	
}

func (p SurveyService) Post(r *http.Request) SrvcRes {
	decoder := json.NewDecoder(r.Body)
	var t dao.SurveyPost
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	s_info := dao.Create_survey(t)
	service := Response200OK(s_info)
	service.Message = "survey created"
	return service 	
}
