package dao

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type Qstruct struct {
	Question []string
	Options []string 
}

type SurveyCredentials struct {
	SurveyId bson.ObjectId          `bson:"_id" json:"_id" ,omitempty`
	Users []string                  `json:"usernames"`
	Title string 					`json : "title"`
	Status int 						`json : "status"`
	Description string				`json : "desc"`
	Questions []Qstruct             `json : "questions"`
}

type SurveyPostStruct struct {         
	Users []string                 
	Title string 					
	Status int 						
	Description string				
	Questions []Qstruct
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Alias    string `json:"alias"`
}

type Session struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}


func Create_survey(t SurveyPostStruct) SurveyCredentials {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("surveyData").C("survey")
	sid := bson.NewObjectId()
	post_var := SurveyCredentials{
		SurveyId : sid,
		Title : t.Title,
		Status : t.Status,
		Description : t.Description, 
		Questions : t.Questions, 
	}
	err = c.Insert(post_var)
	if err != nil {
		panic(err)
	}
	return post_var
} 

func Get_survey_data_by_id(sid string) SurveyCredentials{
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("surveyData").C("survey")
	result :=  SurveyCredentials{}
	c.Find(bson.M{"_id": bson.ObjectIdHex(sid)}).One(&result)
  	survey_info_var :=  SurveyCredentials{
		SurveyId : result.SurveyId,
		Title : result.Title,
		Status : result.Status,
		Description : result.Description,
        Questions : result.Questions,
	}
	return survey_info_var
}

func Get_session_by_token(token string) Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("surveyData").C("session")
	result := Session{}
	c.Find(bson.M{"token": token}).One(&result)
	session_info := Session{
		Username : result.Username,
		Token : result.Token,
	}
	return session_info
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
