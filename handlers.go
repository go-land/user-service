package main

import (
	"algorithms/dao"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUserInfo(resp http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	userName := params["name"]

	content := "<html><body><h1>User info: " + dao.UserInfo[userName] + "</h1></body></html> "

	resp.WriteHeader(http.StatusOK)

	fmt.Fprint(resp, content);
}
