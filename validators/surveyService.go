package validators

import (
	"simplesurveygo/dao"
	"reflect"
)


func Validate_user_session(token string,username string) bool {
	session_info := dao.Get_session_by_token(token)
	if reflect.DeepEqual(session_info, (dao.Session{})) {
        return false
	} else if session_info.Username == username {
        	return true
	} else {
	       return false		
	}
} 