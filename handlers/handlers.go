package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-land/user-service/dao"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUserInfo(resp http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	userName := params["name"]

	addJsonContentType(resp)

	if _, ok := dao.UserInfo[userName]; !ok {
		notFound(resp, userName)
		return
	}

	userData := dao.UserInfo[userName]

	content, err := json.Marshal(userData)

	if err != nil {
		internalError(resp, userName)
		return
	}

	ok(resp, string(content))
}

type ErrorMessage struct {
	ErrorCode   string `json:"errorCode"`
	Description string `json:"description"`
}

func addJsonContentType(resp http.ResponseWriter) {
	resp.Header().Add("Content-Type", "application/json")
}

func createErrorMessage(errorCode string, description string) string {
	errorBody, _ := json.Marshal(ErrorMessage{
		ErrorCode:   errorCode,
		Description: description,
	})

	return string(errorBody)
}

func ok(resp http.ResponseWriter, body string) {
	resp.WriteHeader(http.StatusOK)
	fmt.Fprint(resp, string(body))
}

func notFound(resp http.ResponseWriter, userName string) {
	resp.WriteHeader(http.StatusNotFound)
	fmt.Fprint(resp, createErrorMessage("USER_NOT_FOUND", "Can't find user with name: "+userName))
}

func internalError(resp http.ResponseWriter, userName string) {
	resp.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(resp, createErrorMessage("MARSHALLING_ERROR", "Can't marshall user with name: "+userName))
}
