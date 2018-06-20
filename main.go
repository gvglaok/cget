package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	//"io"
	//"io/ioutil"
	//_ "github.com/go-sql-driver/mysql"
	//"unicode/utf8" 哈哈
)

func main() {
	serve()
	//view()
	//QR()
}

func dbdo() {

	//db,err := sql.Open("mysql" ,"root:root@/zjcms")
	db, err := sql.Open("sqlite3", "E:\\kwork\\goPro\\cget\\database\\cget.db?cache=shared&mode=memory")
	//db.SetMaxOpenConn(1)

	if err != nil {
		fmt.Printf("db connect error")
	}

	defer db.Close()

	type User struct {
		name, email string
		age         int
	}
	var u User

	rows, err := db.Query("SELECT name,age FROM users")

	for rows.Next() {
		rows.Scan(&u.name, &u.age)

		fmt.Printf("name is %s, age is %d \n", u.name, u.age)
	}

}

func serve() {
	route()
	http.ListenAndServe("localhost:8000", nil)
}

func route() {

	http.HandleFunc("/", QR)

}

//QR http server 方法
func QR(w http.ResponseWriter, req *http.Request) {
	//var tpfile := ioutil.ReadFile("E:\\kwork\\goPro\\cget\\views\\index.html")
	var templ, err = template.ParseFiles("./views/index.html")
	if err != nil || templ == nil {
		fmt.Print("解析错误")
	}

	//var templ = template.Must(template.New("index").ParseFiles("E:\\kwork\\goPro\\cget\\views\\index.html"))
	//var templ = template.Must(template.New("index").Parse("hello html template"))
	templ.Execute(w, "index")
	fmt.Print("解析 ok")

}
