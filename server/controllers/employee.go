package controllers

import (
	"database/sql"
	"fmt"
	apps "go-web-template/server/apps/web"
	params "go-web-template/server/params/employee"
	"go-web-template/server/services"
	"go-web-template/server/utils"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/google/uuid"
)

type EmployeeController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
}

type employeeController struct {
	DB *sql.DB
}

func NewEmployeeController(db *sql.DB) EmployeeController {
	return &employeeController{
		DB: db,
	}
}

func (e *employeeController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("static", "pages/employees/index.html"), utils.LayoutMaster)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	employees := services.NewEmployeeService(e.DB).GetAllEmployees()

	web := apps.RenderWeb{
		Title: "Halaman Employee",
		Data:  employees,
	}

	err = tmpl.Execute(w, web)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (e *employeeController) Add(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("static", "pages/employees/add.html"), utils.LayoutMaster)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		web := apps.RenderWeb{
			Title: "Tambah Pegawai",
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
		var request params.EmployeeCreate = params.EmployeeCreate{
			NIP:     r.Form.Get("nip"),
			Address: r.Form.Get("address"),
			Name:    r.Form.Get("name"),
		}
		employeeServices := services.NewEmployeeService(e.DB)
		isSuccess := employeeServices.CreateNewEmployee(&request)
		msg := ""
		if isSuccess {
			msg = `
				<script>
					alert("Tambah data pegawai berhasil !")
					window.location.href="../employees"
				</script>
			`
		} else {
			msg = `
				<script>
					alert("Tambah data pegawai gagal !")
					window.location.href="../employees"
				</script>
			`
		}
		w.Write([]byte(msg))
	} else {
		msg := fmt.Sprintf("Method %s tidak diperbolehkan", method)
		w.Write([]byte(msg))
	}
}

func (e *employeeController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query["id"][0]
	deleteData := services.NewEmployeeService(e.DB).DeleteEmbloyeeByID(id)
	msg := ""
	if deleteData {
		msg = `
			<script>
				alert("Hapus data pegawai berhasil !")
				window.location.href="../employees"
			</script>
		`
	} else {
		msg = `
			<script>
				alert("Hapus data pegawai gagal !")
				window.location.href="../employees"
			</script>
		`
	}
	w.Write([]byte(msg))
}

func (e *employeeController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	query := r.URL.Query()

	id := query.Get("id")

	if method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("static", "pages/employees/update.html"), utils.LayoutMaster)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employee := services.NewEmployeeService(e.DB).GetEmployeeByID(id)
		web := apps.RenderWeb{
			Title: "Halaman Detail Employee",
			Data:  employee,
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
		newID, err := uuid.Parse(id)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var request params.EmployeeUpdate = params.EmployeeUpdate{
			ID:      newID,
			NIP:     r.Form.Get("nip"),
			Address: r.Form.Get("address"),
			Name:    r.Form.Get("name"),
		}
		employeeServices := services.NewEmployeeService(e.DB)
		isSuccess := employeeServices.UpdateByID(&request)
		msg := ""
		if isSuccess {
			msg = `
				<script>
					alert("Ubah data pegawai berhasil !")
					window.location.href="../employees"
				</script>
			`
		} else {
			msg = `
				<script>
					alert("Ubah data pegawai gagal !")
					window.location.href="../employees"
				</script>
			`
		}
		w.Write([]byte(msg))
	} else {
		msg := fmt.Sprintf("Method %s tidak diperbolehkan", method)
		w.Write([]byte(msg))
	}
}
