package main

import (
	"log"
	"net/http"
)

func main() {
	// Используется функция http.NewServeMux() для инициализации нового роутера, затем
	// функцию "home" регистрируется как обработчик для URL-шаблона "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	//регистрируем 2 новых обработчика и соотв url-шаблоны в маршрутизаторе servemux
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
