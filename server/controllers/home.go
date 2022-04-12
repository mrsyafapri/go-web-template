package controllers

import (
	apps "go-web-template/server/apps/web"
	"go-web-template/server/utils"
	"html/template"
	"log"
	"net/http"
	"path"
)

type HomeController interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type homeController struct{}

func NewHomeController() HomeController {
	return &homeController{}
}

func (*homeController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("static", "pages/home/index.html"), utils.LayoutMaster)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	web := apps.RenderWeb{
		Title: "Halaman Home",
	}
	err = tmpl.Execute(w, web)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
