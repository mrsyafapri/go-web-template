package controllers

import (
	apps "go-web-template/server/apps/web"
	"go-web-template/server/utils"
	"html/template"
	"log"
	"net/http"
	"path"
)

type TransactionController interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type transactionController struct{}

func NewTransactionController() TransactionController {
	return &transactionController{}
}

func (*transactionController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("static", "pages/transactions/index.html"), utils.LayoutMaster)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	web := apps.RenderWeb{
		Title: "Halaman Transaction",
	}
	err = tmpl.Execute(w, web)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
