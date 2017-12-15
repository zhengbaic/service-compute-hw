package service

import (
	"encoding/json"
	"fmt"
	"gorm-golang/entities"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type response struct {
	OK   bool
	Data interface{}
}

func jsonResponse(w http.ResponseWriter, resp *response) {
	encoder := json.NewEncoder(w)
	encoder.Encode(resp)
}

func failure(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	log.Printf("INFO: %s", message)
	jsonResponse(w, &response{
		OK:   false,
		Data: message,
	})
}

func illegalForm(w http.ResponseWriter, form url.Values, params ...string) bool {
	for _, p := range params {
		if form[p] == nil {
			failure(
				w, http.StatusBadRequest,
				fmt.Sprintf("illegal input: require field: %s", p))
			return true
		}
	}
	return false
}

func adduser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if illegalForm(w, req.Form, "username", "password") {
		return
	}

	u := entities.NewUser(req.Form["username"][0], req.Form["password"][0])
	dao := entities.GetDao()
	log.Printf(
		"INFO: adding user, username '%s', password '%s'",
		u.Username, u.Password)
	err := dao.SaveUser(u)
	if err != nil {
		failure(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		jsonResponse(w, &response{
			OK:   true,
			Data: "add user success",
		})
	}
}
func getuser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if illegalForm(w, req.Form, "userid") {
		return
	}
	uid, err := strconv.ParseUint(req.Form["userid"][0], 10, 64)
	if err != nil {
		failure(w, http.StatusBadRequest, err.Error())
		return
	}
	Mydao := entities.GetDao()
	u, err := Mydao.GetAUser(uid)
	if err != nil {
		failure(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		jsonResponse(w, &response{
			OK:   true,
			Data: u,
		})
	}
}
func getallusers(w http.ResponseWriter, req *http.Request) {
	Mydao := entities.GetDao()
	users, err := Mydao.GetAllUsers()
	if err != nil {
		failure(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		jsonResponse(w, &response{
			OK:   true,
			Data: users,
		})
	}
}
