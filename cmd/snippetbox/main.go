package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Запуск сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

//import (
//	"log"
//	"net/http"
//)
//
//func main() {
//	// Используется функция http.NewServeMux() для инициализации нового роутера, затем
//	// функцию "home" регистрируется как обработчик для URL-шаблона "/".
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", home)
//	//регистрируем 2 новых обработчика и соотв url-шаблоны в маршрутизаторе servemux
//	mux.HandleFunc("/snippet", showSnippet)
//	mux.HandleFunc("/snippet/create", createSnippet)
//
//	// Инициализируем FileServer, он будет обрабатывать
//	// HTTP-запросы к статическим файлам из папки "./ui/static".
//	// Обратите внимание, что переданный в функцию http.Dir путь
//	// является относительным корневой папке проекта
//	fileServer := http.FileServer(http.Dir("./ui/static/"))
//	// Используем функцию mux.Handle() для регистрации обработчика для
//	// всех запросов, которые начинаются с "/static/". Мы убираем
//	// префикс "/static" перед тем как запрос достигнет http.FileServer
//	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
//	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
//	err := http.ListenAndServe(":4000", mux)
//	log.Fatal(err)
//}
