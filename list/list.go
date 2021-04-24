package list

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type output struct {
	Ticket   []string
	Name     []string
	Username []string
	Stat     []string
}

var t = []string{}
var u = []string{}
var n = []string{}
var s = []string{}
var cl1, cl2, cl3, cl4, D string

func Table(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/list.html")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", D)
	if err != nil {
		fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
		panic(err)
		return
	}
	q, _ := db.Query("SELECT ticket, username, name, checkin FROM tickets")
	for i := 0; q.Next(); i++ {
		q.Scan(&cl1, &cl2, &cl3, &cl4)
		t = append(t, cl1)
		u = append(u, cl2)
		n = append(n, cl3)
		s = append(s, cl4)

	}
	o := output{Ticket: t, Name: n, Username: u, Stat: s}
	tmp.ExecuteTemplate(w, "list.html", o)
	t, u, n, s = nil, nil, nil, nil

}
