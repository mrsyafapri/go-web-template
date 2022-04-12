package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"go-web-template/server/helper"
	params "go-web-template/server/params/menu"
	"go-web-template/server/services"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type MenuController interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
}

type menuController struct {
	DB *sql.DB
}

func NewMenuController(db *sql.DB) MenuController {
	return &menuController{db}
}

func (m *menuController) Add(w http.ResponseWriter, r *http.Request) {

	var request params.MenuCreate

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.HandleBadRequest(w, err)
		return
	}

	_, err = services.NewMenuService(m.DB).CreateNewMenu(&request)

	if err != nil {
		helper.HandleInternalServerError(w, errors.New("INTERNAL SERVER ERROR"))
		return
	} else {
		helper.HandleCreateSuccess(w, "Create new menu success !")
		return
	}
}

func (m *menuController) FindAll(w http.ResponseWriter, r *http.Request) {
	menus, err := services.NewMenuService(m.DB).GetAllMenu()

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			helper.HandleNotFound(w, errors.New("NO DATA"))

		} else {
			helper.HandleInternalServerError(w, err)
		}
		return
	}

	if len(*menus) == 0 {
		helper.HandleNotFound(w, errors.New("NO DATA"))
		return
	}

	helper.HandleSuccess(w, menus)
}

func (m *menuController) FindByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	// validasi, apakah id dari parameter merupakan UUID atau bukan
	_, err := uuid.Parse(id)
	if err != nil {
		helper.HandleBadRequest(w, err)
		return
	}

	menu, err := services.NewMenuService(m.DB).GetMenuByID(id)
	if err != nil {
		if err.Error() != sql.ErrNoRows.Error() {
			helper.HandleNotFound(w, errors.New("NO DATA"))

		} else {
			helper.HandleInternalServerError(w, err)
		}
		return
	}

	helper.HandleSuccess(w, menu)
}

func (m *menuController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	newID, err := uuid.Parse(id)
	if err != nil {
		helper.HandleBadRequest(w, err)
		return
	}

	var request params.MenuUpdate

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helper.HandleBadRequest(w, err)
		return
	}

	request.ID = newID

	_, err = services.NewMenuService(m.DB).UpdateMenuByID(&request)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			helper.HandleNotFound(w, errors.New("NO DATA"))
		} else {
			helper.HandleInternalServerError(w, err)
		}
		return
	}

	helper.HandleSuccess(w, "update success")
}

func (m *menuController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	_, err := uuid.Parse(id)

	if err != nil {
		helper.HandleBadRequest(w, err)
		return
	}

	_, err = services.NewMenuService(m.DB).DeleteMenuByID(id)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			helper.HandleNotFound(w, errors.New("NO DATA"))
		} else {
			helper.HandleInternalServerError(w, err)
		}
		return
	}

	helper.HandleSuccess(w, "delete success")
}
