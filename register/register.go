package register
import (
  "fmt"
  "net/http"
  "net/url"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "html/template"
  "crypto/rand"
  "unsafe"
)
type Ticket struct {
  Src string
  Qr string
}
var name, username, password, DBcreds, token, TicketType string
var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12890")
func Register(w http.ResponseWriter, r *http.Request) {
  Template := template.Must(template.ParseGlob("templates/register.html"))
  qr, _ := template.ParseFiles("templates/ticket.html")
  if r.Method == "POST" {
    tk := Ticket{Src: "", Qr: ""}
    r.ParseForm()
    name = r.PostForm.Get("name")
    username = r.PostForm.Get("username")
    password = r.PostForm.Get("password")
    token = r.PostForm.Get("token")
    TicketType = r.PostForm.Get("type")
    if len(name) > 30 {
      fmt.Fprintf(w, `<script>alert("name is too long")</script>`)
      return
    }else if len(username) > 30 {
      fmt.Fprintf(w, `<script>alert("username is too long")</script>`)
      return
    }else if len(password) > 50 {
      fmt.Fprintf(w, `<script>alert("password is too long")</script>`)
      return
    }
    if TicketType == "Single" {
      tk.Src = "https://fidebleb.sirv.com/tickets/Private_Single.jpg"
    }else if TicketType == "Couples" {
      tk.Src = "https://fidebleb.sirv.com/tickets/Private_Couples.jpg"
    }else {
      fmt.Fprintf(w, "Invalid ticket type")
      return
    }
    DB, err := sql.Open("mysql", DBcreds)
    if err != nil {
      fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
      fmt.Println(err)
      return
    }
    res, err := DB.Query("SELECT username FROM tickets WHERE username=?", username)
    if err != nil {
      fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
      fmt.Println(err)
      return
    }
    if res.Next() == true {
      fmt.Fprintf(w, `<script>alert("username exist");window.location.href = "/register";</script>`)
      return
    }
    res, _ = DB.Query("SELECT * FROM tokens WHERE token_id=?", token)
    if res.Next() == false {
      fmt.Fprintf(w, `<script>alert("Invalid Token");window.location.href = "/register";</script>`)
      return
    }
    res.Close()
    b := make([]byte, 40)
    rand.Read(b)
    for i := 0; i < 40; i++ {
      b[i] = alphabet[b[i]/5]
    }
    ticket := *(*string)(unsafe.Pointer(&b))
    res, err = DB.Query("INSERT INTO tickets(ticket, name, username, password, sold) VALUES(?, ?, ?, ?, true)", ticket, name, username, password)
    if err != nil {
      fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
      fmt.Println(err)
      return
    }
    res.Close()
    DB.Close()
    tk.Qr = url.QueryEscape(ticket)
    qr.Execute(w, tk)
    ticket = ""
  }else if r.Method == "GET" {
    Template.Execute(w, "")
  }
}


func ForgetTicket(w http.ResponseWriter, r *http.Request) {
  qr, _ := template.ParseFiles("templates/ticket.html")
  Template := template.Must(template.ParseGlob("templates/forget.html"))
  if r.Method == "GET" {
    Template.Execute(w, "")
  }else if r.Method == "POST" {
    r.ParseForm()
    username = r.PostForm.Get("username")
    password = r.PostForm.Get("password")
    db, err := sql.Open("mysql", DBcreds)
    if err != nil {
      fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
      panic(err)
      return
    }
    res, err := db.Query("SELECT ticket FROM tickets WHERE username=? AND password=? AND sold=1", username, password)
    if err != nil {
      fmt.Fprintf(w, `<script>window.location.href = "/500";</script>`)
      fmt.Println(err)
      return
    }
    if res.Next() == true {
      var ticket string
      res.Scan(&ticket)
      qr.Execute(w, url.QueryEscape(ticket))
    }else if res.Next() == false {
      fmt.Fprintf(w, `<script>alert("ther is no ticket with this information");window.location.href = "/forget"</script>`)
    }
    res.Close()
    db.Close()
  }
}
