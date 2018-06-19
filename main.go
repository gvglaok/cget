package main

import (
	//"io"
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	//"io/ioutil"
	"html/template"
	//_ "github.com/go-sql-driver/mysql"
	//"unicode/utf8" 哈哈
)

func main()  {
	serve()
	//view()
}

func dbdo(){

	//db,err := sql.Open("mysql" ,"root:root@/zjcms")
	db, err := sql.Open("sqlite3", "E:\\kwork\\goPro\\cget\\database\\cget.db?cache=shared&mode=memory")
	//db.SetMaxOpenConn(1)

	if err != nil {
		fmt.Printf("db connect error")
	}

	defer db.Close()

 
	type User struct {
		name,email  string 
		age int
	}
	var u User
	
	rows,err := db.Query("SELECT name,age FROM users")

	for rows.Next(){
		rows.Scan(&u.name,&u.age)
		 
		fmt.Printf("name is %s, age is %d \n", u.name ,u.age)
	}

} 

func serve() {
	route()
	http.ListenAndServe("localhost:8000",nil)
}

func route() {
	 
	http.HandleFunc("/",QR)
	 
}

//http server 处理
func index(w http.ResponseWriter, req *http.Request)  {
 
}

func QR(w http.ResponseWriter, req *http.Request) {
	var templ = template.Must(template.New("index").ParseFiles("./views./index.blade.html"))
    templ.Execute(w, req.FormValue("s"))
}

 


