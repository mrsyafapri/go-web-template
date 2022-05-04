package controllers

import (
	"database/sql"
	"fmt"
	apps "go-web-template/server/apps/web"
	params "go-web-template/server/params/transaction"
	"go-web-template/server/services"
	"go-web-template/server/utils"
	"html/template"
	"log"
	"net/http"
	"path"
)

type TransactionController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
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
	transactions := services.NewTransactionService(t.DB).GetAllTransactions()

	web := apps.RenderWeb{
		Title: "Halaman Transaction",
		Data:  transactions,
	}
	err = tmpl.Execute(w, web)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *transactionController) Add(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("static", "pages/transactions/add.html"), utils.LayoutMaster)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		web := apps.RenderWeb{
			Title: "Tambah Transaksi",
		}
		err = tmpl.Execute(w, web)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var request params.TransactionCreate = params.TransactionCreate{
			EmployeeID: r.FormValue("employee"),
			MenuID:     r.FormValue("menu"),
		}
		transactionServices := services.NewTransactionService(t.DB)
		isSuccess := transactionServices.CreateNewTransaction(&request)
		msg := ""
		if isSuccess {
			msg = `
				<script>
					alert("Tambah data transaksi berhasil !")
					window.location.href="../transactions"
				</script>
			`
		} else {
			msg = `
				<script>
					alert("Tambah data transaksi gagal !")
					window.location.href="../transactions"
				</script>
			`
		}
		w.Write([]byte(msg))
	} else {
		msg := fmt.Sprintf("Method %s tidak diperbolehkan", method)
		w.Write([]byte(msg))
	}
}
