package dao

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Alias    string `json:"alias"`
}

type Session struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Ans struct {
	Answer string
	Username string
} 

type Qstruct struct {
	Qid string
	Question string 
	Asker string
	Options []string 
	Answers []Ans 	
}


type SurveyPost struct {
	Title string  
	Status int
	Description string 
	Questions []Qstruct   
}



type SurveyOprPost struct {
	Username string
	Surveys []string 
}

type SurveyOprPut struct {
	Username string
	Surveys []string
	Token string
}

type UserQuestion struct {
	Token string
	Question string
	SurveyId string
}

type UserAnswer struct {
	SurveyId string
	Qid string
	Answer string 
	Token string
}



type UserSurveys struct {
	Username string 	`json:"username"`
	Surveys []string 	`json:"surveys"`
}

type Survey struct {
	SurveyId string 	`json:sid`
	Title string  		`json:"title"`
	Status int 			`json:"status"`
	Description string 	`json:"desc"`
	Questions []Qstruct `json:"questions"`  
}


func Update_question(t UserQuestion) Survey {
	user_info := GetSessionDetails(t.Token)
	qid := uuid.Must(uuid.NewV4()).String()
	q_var := Qstruct{
		Qid      : qid,  
		Question : t.Question,
		Asker    : user_info.Username,
		Options  : nil,
		Answers  : nil,
	}
	session := MgoSession.Clone()
	defer session.Close()
	s := session.DB("simplesurveys").C("survey")
	result := Get_survey(t.SurveyId)
	result.Questions  = append(result.Questions,q_var)
	s.Update(bson.M{"surveyid":t.SurveyId},bson.M{"$set":bson.M{"questions" : result.Questions}})	
	result = Get_survey(t.SurveyId)
	return result
}

func Update_answer(t UserAnswer) Survey {
    user_info := GetSessionDetails(t.Token)
	survey := Get_survey(t.SurveyId)
	ans := Ans{
		Answer : t.Answer,
		Username : user_info.Username,
	}
	qs := []Qstruct{}
	for _,v := range survey.Questions {
		if v.Qid == t.Qid {
		   v.Answers = append(v.Answers,ans) 
		   qs = append(qs,v)	
		} else {
			qs = append(qs,v)
		}
	} 
	session := MgoSession.Clone()
	defer session.Close()
	s := session.DB("simplesurveys").C("survey")
	fmt.Println(qs)
	s.Update(bson.M{"surveyid":t.SurveyId},bson.M{"$set":bson.M{"questions" : qs}})
	result := Get_survey(t.SurveyId)
	return result
}  

func Get_user_survey(username string) UserSurveys {
	session := MgoSession.Clone()
	defer session.Close()
    result := UserSurveys{}
	s := session.DB("simplesurveys").C("surveyOpr")
	s.Find(bson.M{"username": username}).One(&result)
	return result
}

func Create_link(t SurveyOprPost) UserSurveys {
	session := MgoSession.Clone()
	defer session.Close()

	s := session.DB("simplesurveys").C("surveyOpr")
	survey_opr_var := UserSurveys{
		Username : t.Username,
		Surveys  : t.Surveys, 
	}
	s.Insert(survey_opr_var)
	return survey_opr_var
}

func Update_user_surveys(t SurveyOprPut) UserSurveys {
	session := MgoSession.Clone()
	defer session.Close()
	result := UserSurveys{} 
	fmt.Println(t)
	s := session.DB("simplesurveys").C("surveyOpr")
	s.Find(bson.M{"username": t.Username}).One(&result)
	for _,v := range t.Surveys {
		result.Surveys = append(result.Surveys,v)
	} 
	fmt.Println(result)
	s.Update(bson.M{"username":t.Username},bson.M{"$set":bson.M{"surveys":result.Surveys}})
	return result
} 

func Create_survey(t SurveyPost) Survey {
	session := MgoSession.Clone()
	defer session.Close()

	s := session.DB("simplesurveys").C("survey")
	sid := uuid.Must(uuid.NewV4()).String()
	survey_post_var := Survey{
		 SurveyId : sid,
		 Title    : t.Title,
		 Status   : t.Status,
		 Description : t.Description,
		 Questions   : t.Questions, 
	}
	s.Insert(survey_post_var)
	return survey_post_var
} 



func AuthenticateUser(cred UserCredentials) string {
	session := MgoSession.Clone()
	defer session.Close()

	var response interface{}
	clctn := session.DB("simplesurveys").C("user")
	query := clctn.Find(bson.M{"username": cred.Username, "password": cred.Password})
	err := query.One(&response)
	uuidStr := uuid.Must(uuid.NewV4()).String()
	sessionStruct := Session{cred.Username, uuidStr}
	if err != nil {
		return ""
	}
	sessionClctn := session.DB("simplesurveys").C("session")
	sessionClctn.Insert(sessionStruct)
	return uuidStr
}



func GetSessionDetails(token string) UserCredentials {
	session := MgoSession.Clone()
	defer session.Close()

	var response Session
	sessionClctn := session.DB("simplesurveys").C("session")
	query := sessionClctn.Find(bson.M{"token": token})
	err := query.One(&response)
	if err != nil {
		return UserCredentials{}
	}

	var cred UserCredentials
	clctn := session.DB("simplesurveys").C("user")
	query = clctn.Find(bson.M{"username": response.Username})
	err = query.One(&cred)
	return cred
}


func Create_user(t UserCredentials) UserCredentials {
	session := MgoSession.Clone()
	defer session.Close()

	u := session.DB("simplesurveys").C("user")
	user_post_var := UserCredentials{
		Username : t.Username,
		Password : t.Password,
		Alias    : t.Alias,
	}
    u.Insert(user_post_var)
	return user_post_var
}  

func Get_user_by_username(uname string) UserCredentials {
	result := UserCredentials{} 
	session := MgoSession.Clone()
	defer session.Close()
    
	u := session.DB("simplesurveys").C("user")
	u.Find(bson.M{"username": uname}).One(&result)
    return result
}

func Remove_session(token string) {
	session := MgoSession.Clone()
	defer session.Close()
    
	u := session.DB("simplesurveys").C("session") 
	u.Remove(bson.M{"token":token})
}

func Get_all_survey() []Survey  { 
	session := MgoSession.Clone()
	defer session.Close()
    
	s := session.DB("simplesurveys").C("survey")
	result := []Survey{}
	s.Find(nil).All(&result)
   	return result
}

func Get_survey(sid string) Survey {
	session := MgoSession.Clone()
	defer session.Close()
    
	s := session.DB("simplesurveys").C("survey")
	result := Survey{}
	s.Find(bson.M{"surveyid" : sid}).One(&result)
	return result	
}  


