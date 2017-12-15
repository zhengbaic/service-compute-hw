package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/painterdrown/service-computing/agenda/entity"
)

var user_list map[string]bool

type response struct {
	OK   bool
	Data interface{}
}

func LoginFirst(w http.ResponseWriter, username string) bool {
	if username == "" {
		FailResponse(w, http.StatusInternalServerError, "username is empty")
		return false
	}
	if user_list[username] == false {
		FailResponse(w, http.StatusInternalServerError, "log in plz")
		return false
	}
	return true
}
func JsonResponse(w http.ResponseWriter, res *response) {
	encoder := json.NewEncoder(w)
	encoder.Encode(res)
}
func SuccessResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	JsonResponse(w, &response{
		OK:   true,
		Data: data,
	})
}
func FailResponse(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	log.Printf("INFO: %s", message)
	JsonResponse(w, &response{
		OK:   false,
		Data: message,
	})
}
func CheckRequire(w http.ResponseWriter, form url.Values, require ...string) bool {
	for _, item := range require {
		if form[item] == nil {
			FailResponse(
				w, http.StatusBadRequest,
				fmt.Sprintf("[ERROR] missing: %s", item))
			return true
		}
	}
	return false
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if CheckRequire(w, r.Form, "username", "password") {
		return
	}
	err := entity.Register(r.Form["username"][0], r.Form["password"][0])
	log.Printf(
		"INFO: user register: username '%s', password '%s'",
		r.Form["username"][0], r.Form["password"][0])
	if err != nil {
		FailResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		SuccessResponse(w, "user register success")
	}
}
func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := entity.GetAllUsers()
	if err != nil {
		FailResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		SuccessResponse(w, users)
	}

}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	if LoginFirst(w, username) == false {
		return
	}
	err := entity.DeleteUser(username)
	if err != nil {
		FailResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		log.Printf(
			"INFO: user delete: username '%s'",
			username)
		SuccessResponse(w, "user delete success")
	}
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if CheckRequire(w, r.Form, "username", "password") {
		return
	}
	if user_list[r.Form["username"][0]] == true {
		FailResponse(w, http.StatusInternalServerError, "u have log in !")
		return
	}
	users, err := entity.GetAllUsers()
	if err != nil {
		FailResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	for _, item := range users {
		if r.Form["username"][0] == item.Username && r.Form["password"][0] == item.Password {
			user_list[item.Username] = true
			log.Printf(
				"INFO: user log in: username '%s'",
				r.Form["username"][0])
			SuccessResponse(w, "login success")
			return
		}
	}
	FailResponse(w, http.StatusInternalServerError, "wrong username/password")
}
func CreateMeetingHandler(w http.ResponseWriter, r *http.Request) {
	title := r.Form["title"][0]
	username := mux.Vars(r)["username"]
	participators := r.Form["participator"][0]
	stime := r.Form["stime"][0]
	etime := r.Form["etime"][0]
	/*ps := strings.FieldsFunc(participators, func(arg1 rune) bool {
		return arg1 == ','
	})*/
	if LoginFirst(w, username) == false {
		return
	}
	if title == "" || participators == "" || etime == "" || stime == "" {
		FailResponse(w, http.StatusInternalServerError, "title/participator/stime/etime is empty")
		return
	}
	participators += ","
	err := entity.CreateMeeting(username, title, participators, stime, etime)
	if err != nil {
		FailResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		log.Printf(
			"INFO: user create meeting: username '%s', title '%s', ps '%s', time '%s'-'%s'",
			username, title, participators, stime, etime)
		SuccessResponse(w, "add meetings success")
	}
}
func DeleteMeetingHandler(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	if LoginFirst(w, username) == false {
		return
	}
	err := entity.DeleteMeetings(username)
	if err != nil {
		FailResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		log.Printf(
			"INFO: user delete all meeting: username '%s'",
			username)
		SuccessResponse(w, "delete meetings success")
	}
}
func QueryMeetingHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := mux.Vars(r)["username"]
	if LoginFirst(w, username) == false {
		return
	}
	if CheckRequire(w, r.Form, "stime", "etime") {
		return
	}
	meetings, err := entity.QueryMeetings(username, r.Form["stime"][0], r.Form["etime"][0])
	if err != nil {
		FailResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		SuccessResponse(w, meetings)
	}
}
func AddParticipatorsHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	username := mux.Vars(r)["username"]
	participator := mux.Vars(r)["participator"]
	if LoginFirst(w, username) == false {
		return
	}
	if title == "" || username == "" || participator == "" {
		FailResponse(w, http.StatusInternalServerError, "title/username/participator is empty")
		return
	}
	err := entity.AddParticipator(username, title, participator)
	if err != nil {
		FailResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		SuccessResponse(w, "Add Participators Success")
	}
}
func DeleteParticipatorsHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	username := mux.Vars(r)["username"]
	participator := mux.Vars(r)["participator"]
	if LoginFirst(w, username) == false {
		return
	}
	if title == "" || username == "" || participator == "" {
		FailResponse(w, http.StatusInternalServerError, "title/username/participator is empty")
		return
	}
	err := entity.DeleteParticipator(username, title, participator)
	if err != nil {
		FailResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		SuccessResponse(w, "Delete Participators Success")
	}
}
func DropMeetingHandler(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	username := mux.Vars(r)["username"]
	if LoginFirst(w, username) == false {
		return
	}
	if title == "" {
		FailResponse(w, http.StatusInternalServerError, "title is empty")
		return
	}
	err := entity.DeleteMeeting(username, title)
	if err != nil {
		FailResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		SuccessResponse(w, "Delete Success")
	}
}

func router() http.Handler {
	user_list = make(map[string]bool)
	user_list["init"] = false
	router := mux.NewRouter()
	router.HandleFunc("/users", RegisterHandler).Methods("POST")                //注册
	router.HandleFunc("/users", GetAllUserHandler).Methods("GET")               //获取所有用户
	router.HandleFunc("/users/{username}", DeleteUserHandler).Methods("DELETE") //删除用户
	router.HandleFunc("/users/signin", LoginHandler).Methods("POST")            //登陆

	router.HandleFunc("/meetings/{username}", CreateMeetingHandler).Methods("POST")                                              //创建会议
	router.HandleFunc("/meetings/{username}", DeleteMeetingHandler).Methods("DELETE")                                            //清除用户的所有会议
	router.HandleFunc("/meetings/{username}", QueryMeetingHandler).Methods("GET")                                                //查询会议
	router.HandleFunc("/meetings/{username}/{title}/participators/{participator}", AddParticipatorsHandler).Methods("POST")      //添加参与者
	router.HandleFunc("/meetings/{username}/{title}/participators/{participator}", DeleteParticipatorsHandler).Methods("DELETE") //删除参与者
	router.HandleFunc("/meetings/{username}/{title}", DropMeetingHandler).Methods("DELETE")                                      //删除，退出会议

	return router
}
