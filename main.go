package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//serve()
	dbdo()
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

	var users = map[int]User{}

	for rows.Next() {
		index := 0
		rows.Scan(&u.name, &u.age)

		rows.Scan(users[index].name, users[index].age)
		index++
		//fmt.Printf("name is %s, age is %d \n", u.name, u.age)
	}

	for index := 0; index < len(users); index++ {
		fmt.Print(users[index].name)
	}

}

func serve() {
	route()
	http.ListenAndServe("localhost:8000", nil)
}

func route() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/users", users)

}

//dealHtmlFile 处理页面内 资源文件
func dealHtmlFile(r *http.Request) (data []byte, contentType string) {

	path := r.URL.Path[1:]
	log.Println(path)
	data, err1 := ioutil.ReadFile(string(path))

	if err1 != nil {
		fmt.Println("file read error")
	}

	//var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".svg") {
		contentType = "image/svg+xml"
	} else {
		contentType = "text/plain"
	}

	//file = data

	return
}

//Index http server 方法
func Index(w http.ResponseWriter, r *http.Request) {

	data, contentType := dealHtmlFile(r)

	w.Header().Add("Content Type", contentType)
	w.Write(data)

	var tpl, err = template.ParseFiles("./views/index.html")
	if err != nil || tpl == nil {
		fmt.Print("解析错误")
	}

	tpl.Execute(w, "index")

}

func users(w http.ResponseWriter, r *http.Request) {

	/* data, ct := dealHtmlFile(r)

	w.Header().Add("Content Type", ct)
	w.Write(data) */

	var tpl, err = template.ParseFiles("./views/users.html")
	if err != nil || tpl == nil {
		fmt.Print("解析错误")
	}

	type users struct {
		Name string
		Age  int
	}

	u := users{"admin", 21}
	fmt.Print(u.Name)

	tpl.Execute(w, u)

	//直接输出字符串
	//w.Write([]byte("User list!"))

}
