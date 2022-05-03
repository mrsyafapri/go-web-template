package controllers

import (
	"database/sql"
	apps "go-web-template/server/apps/web"
	"go-web-template/server/utils"
	"html/template"
	"log"
	"net/http"
	"path"
)

type TransactionController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
}

type transactionController struct {
	DB *sql.DB
}

func NewTransactionController(db *sql.DB) TransactionController {
	return &transactionController{
		DB: db,
	}
}

func (t *transactionController) Index(w http.ResponseWriter, r *http.Request) {
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

// berfungsi untuk menampilkan form dan melakukan proses tambah data ke database
func (t *transactionController) Add(w http.ResponseWriter, r *http.Request) {

}

// berfungsi untuk menghapus data employee berdasarkan ID nya
func (t *transactionController) DeleteByID(w http.ResponseWriter, r *http.Request) {

}

// berfungsi untuk mengubah data employee berdasarkan ID nya
func (t *transactionController) UpdateByID(w http.ResponseWriter, r *http.Request) {

}
