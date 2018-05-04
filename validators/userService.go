package validators

import (
   "simplesurveygo/dao"
   "reflect"
)

func Validate_token(token string) bool {
	user_info := dao.GetSessionDetails(token)
    if reflect.DeepEqual(user_info, (dao.UserCredentials{})) {
		return false
	} else {
		return true
	}
}

func Validate_for_create_user(t dao.UserCredentials) bool {
   user_info := dao.Get_user_by_username(t.Username)
   if reflect.DeepEqual(user_info, (dao.UserCredentials{})) {
	return true
   } else {
	return false
   }
}

func Validate_session(token string) bool {
	user_session_info := dao.GetSessionDetails(token)
	if reflect.DeepEqual(user_session_info, (dao.UserCredentials{})) {
		return false
	} else {
		return true		
	}
}