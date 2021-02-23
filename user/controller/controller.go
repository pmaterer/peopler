package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pmaterer/peopler/user"
)

type service interface {
	CreateUser(u user.User) error
	GetUser(id int64) (user.User, error)
	GetAllUsers() ([]user.User, error)
	UpdateUser(u user.User) error
	DeleteUser(id int64) error
}

type Controller struct {
	service service
}

func NewController(s service) *Controller {
	return &Controller{
		service: s,
	}
}

type Response struct {
	Status  int    `json:"-"`
	Message string `json:"message,omitempty"`
}

func writeErrorResponse(w http.ResponseWriter, code int, message string) {
	writeResponse(w, code, Response{Status: code, Message: message})
}

func writeResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (c *Controller) CreateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer r.Body.Close()

		var u user.User
		if err := json.Unmarshal(requestBody, &u); err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = c.service.CreateUser(u)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeResponse(w, http.StatusOK, Response{Message: "OK"})
	}
}

func (c *Controller) GetAllUsers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := c.service.GetAllUsers()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeResponse(w, http.StatusOK, users)
	}
}

func (c *Controller) GetUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		rawID := vars["id"]
		id, err := strconv.Atoi(rawID)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		user, err := c.service.GetUser(int64(id))
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeResponse(w, http.StatusOK, user)
	}
}

func (c *Controller) UpdateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		rawID := vars["id"]
		id, err := strconv.Atoi(rawID)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer r.Body.Close()

		var u user.User
		if err := json.Unmarshal(requestBody, &u); err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		u.ID = int64(id)

		err = c.service.UpdateUser(u)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeResponse(w, http.StatusOK, Response{Message: "OK"})
	}
}

func (c *Controller) DeleteUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		rawID := vars["id"]
		id, err := strconv.Atoi(rawID)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		err = c.service.DeleteUser(int64(id))
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeResponse(w, http.StatusOK, Response{Message: "OK"})
	}
}
