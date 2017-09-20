package main

import (
	_ "beegoAdmin/routers"
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
)

func main() {

	beego.ErrorHandler("404", page_not_found)
	beego.ErrorHandler("403", page_note_permission)
	beego.ErrorHandler("500", page_internal_error)
	beego.Run()
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.tpl").ParseFiles("views/404.tpl")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}

func page_note_permission(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("403.tpl").ParseFiles("views/403.tpl")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}

func page_internal_error(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("500.tpl").ParseFiles("views/500.tpl")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}
