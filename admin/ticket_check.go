package admin
import (
  "net/http"
  "fmt"
  "html/template"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)
var username, password, ticket, token, DBcreds string
func Admin(w http.ResponseWriter, r *http.Request) {
  Template := template.Must(template.ParseGlob("templates/tickte_validate.html"))
  if r.Method == "GET" {
    Template.Execute(w, "")
  }else if r.Method == "POST" {
    r.ParseForm()
    username = r.PostForm.Get("username")
    password = r.PostForm.Get("password")
    ticket = r.PostForm.Get("ticket")
    db, err := sql.Open("mysql", DBcreds)
    if err != nil {
      fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
      return
    }
    defer db.Close()
    res, err := db.Query("SELECT ticket FROM tickets WHERE username=? AND password=? AND ticket=? AND sold=1 ", username, password, ticket)
    if err != nil {
      fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
      return
    }
    defer res.Close()
    if res.Next() == true {
      Template.Execute(w, "ticket exist!")
    }else if res.Next() == false {
      Template.Execute(w, "ticket doesn't exist")
    }
  }
}

func CheckIn(w http.ResponseWriter, r *http.Request) {
  template, _ := template.ParseFiles("templates/check_in.html")
  if r.Method == "GET" {
    template.Execute(w, "")
    return
  }else if r.Method == "POST" {
  r.ParseForm()
  ticket = r.PostForm.Get("ticketid")
  db, err := sql.Open("mysql", DBcreds)
  if err != nil {
    fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
    return
  }
  defer db.Close()
  q, _ := db.Query("SELECT ticket FROM tickets WHERE ticket=? AND checkin=false", ticket)
  if q.Next() == true {
    i, err := db.Prepare("UPDATE tickets SET checkin=true WHERE ticket=?")
    if err != nil {
      fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
      return
    }
    i.Exec(ticket)
    template.Execute(w, `Done!`)
    i.Close()
  }else if q.Next() == false {
    template.Execute(w, `there is no ticket with this id or it has already checked in`)
    return
  }
  q.Close()
 }
}

func TokenInsert(w http.ResponseWriter, r *http.Request) {
  tmp, _ := template.ParseFiles("templates/token.html")
  if r.Method == "GET" {
    tmp.Execute(w, "")
    return
  }
  r.ParseForm()
  token = r.PostForm.Get("token")
  if len(token) > 100 {
    tmp.Execute(w, "token id is too long")
    return
  }
  db, err := sql.Open("mysql", DBcreds)
  if err != nil {
    fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
    panic(err.Error())
    return
  }
  q, err := db.Query("SELECT * FROM tokens WHERE token_id=?", token)
  if err != nil {
    fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
    panic(err.Error())
    return
  }
  if q.Next() == true {
    tmp.Execute(w, "token exist!")
    return
  }
  q, err = db.Query("INSERT INTO tokens VALUES(?)", token)
  if err != nil {
    fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
    panic(err.Error())
    return
  }
  tmp.Execute(w, "Done!")
  db.Close()
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
  Template := template.Must(template.ParseGlob("templates/admin.html"))
  Template.Execute(w, "")
}
