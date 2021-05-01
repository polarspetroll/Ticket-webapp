package tokenlist

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type output struct {
	Token []string
	Used  []string
}

var t = []string{}
var u = []string{}
var cl1, cl2, D string

func Table(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/tokenlist.html")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", D)
	if err != nil {
		fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
		panic(err)
		return
	}
	q, _ := db.Query("SELECT token_id, used FROM tokens")
	for i := 0; q.Next(); i++ {
		q.Scan(&cl1, &cl2)
		t = append(t, cl1)
		u = append(u, cl2)

	}
	o := output{Token: t, Used: u}
	tmp.ExecuteTemplate(w, "tokenlist.html", o)
	t, u = nil, nil
}
