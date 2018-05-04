package servicehandlers

import (
	"encoding/json"
	"net/http"
	"simplesurveygo/dao"
	"simplesurveygo/validators"
)

type UserServiceHandler struct {
}

func (p UserServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

func (p UserServiceHandler) Get(r *http.Request) SrvcRes {
	token, ok := r.URL.Query()["token"]
	if (!ok) || len(token) > 1 || token[0]==""  {
		return SimpleBadRequest("parameters not passed properly")
	} else   {
	   if !validators.Validate_token(token[0]) {
          return UnauthorizedAccess("you have to login again")
	   } else {	
	   			user_info := dao.GetSessionDetails(token[0])
				response := Response200OK(user_info)
				response.Message = "User details"   
	  		    return response
	   }
	}
}	  

func (p UserServiceHandler) Put(r *http.Request) SrvcRes {
	ret := SrvcRes{}
	operation, ok := r.URL.Query()["operation"]
    if !ok || operation[0] == "" {
		ret = SimpleBadRequest("parameters not passed properly")
	} else if operation[0]=="answer" {
		decoder := json.NewDecoder(r.Body)
		var t dao.UserAnswer
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()	
		if !validators.Validate_token(t.Token) {
			ret = UnauthorizedAccess("you have to login again")
		} else {
			response := dao.Update_answer(t)
			ret = Response200OK(response)	
		}	
	} else if operation[0] == "question" {
		decoder := json.NewDecoder(r.Body)
		var t dao.UserQuestion
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()	
		if !validators.Validate_token(t.Token) {
			ret =  UnauthorizedAccess("you have to login again")
		} else {
			response := dao.Update_question(t)
			ret = Response200OK(response)		
		}
	}
	return ret
}		

func (p UserServiceHandler) Post(r *http.Request) SrvcRes {
	decoder := json.NewDecoder(r.Body)
	var t dao.UserCredentials
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
    if !validators.Validate_for_create_user(t) {
        return SimpleBadRequest("user already present")
	} else {
		user_info := dao.Create_user(t) 
		response := Response200OK(user_info)
		response.Message = "User created"
		return response
	}
}
