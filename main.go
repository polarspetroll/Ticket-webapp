package main

import (
	"admin"
	"fmt"
	"html/template"
	"list"
	"net/http"
	"os"
	"register"
	"tokenlist"
)

func main() {
	dbusr := os.Getenv("DBUSR")
	dbpasswd := os.Getenv("DBPWD")
	dbaddr := os.Getenv("DBADDR")
	dbname := os.Getenv("DBNAME")
	Database := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v", dbusr, dbpasswd, dbaddr, dbname)
	admin.DBcreds = Database
	register.DBcreds = Database
	list.D = Database
	tokenlist.D = Database

	http.HandleFunc("/", Redirect)
	http.HandleFunc("/register", register.Register)
	http.HandleFunc("/forget", register.ForgetTicket)
	http.HandleFunc("/500", InternalServerError)
	http.HandleFunc("/admin/", admin.AdminIndex)
	http.HandleFunc("/admin/check_in", admin.CheckIn)
	http.HandleFunc("/admin/ticket_verify/", admin.Admin)
	http.HandleFunc("/admin/token/", admin.TokenInsert)
	http.HandleFunc("/admin/delete", admin.Delete)
	http.HandleFunc("/admin/list", list.Table)
	http.HandleFunc("/admin/tokens", tokenlist.Table)
	http.ListenAndServe(":8080", nil)
	//// You can change the listener port here
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseGlob("templates/home.html"))
	template.Execute(w, "")

}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseGlob("templates/*.html"))
	template.ExecuteTemplate(w, "500.html", "")
}
