package utils

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "invalid params",
	ERROR_EXIST_TAG:                "exist tag",
	ERROR_NOT_EXIST_TAG:            "not exist tag",
	ERROR_NOT_EXIST_ARTICLE:        "Article is not existing",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token auth fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token timeout",
	ERROR_AUTH_TOKEN:               "Token error",
	ERROR_AUTH:                     "User not exist",
	ERROR_EXIST_AUTH:               "User existing",
	ERROR_AUTH_PASSWORD:            "Password is wrong",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:   "Save image fail",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:  "Check image fail",
}

func GetMsg(code int) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}
	return MsgFlags[ERROR]
}
