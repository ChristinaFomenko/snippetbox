package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	//flag.Parse() для извлечения флага из командной строки.
	flag.Parse()
	//созд логера для записи инфо

	//можно осущ.логирование в папку
	//f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	//infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Запуск сервера на http://127.0.0.1%s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
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
