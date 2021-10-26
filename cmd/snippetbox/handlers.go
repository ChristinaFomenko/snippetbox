package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//созд. функция обработчик "home", которая записывает байтовый слайс
//текст "..." как тело ответа
//обработчик главной страницы
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
	// функция http.NotFound() для возвращения клиенту ошибки 404.
	// Важно, чтобы мы завершили работу обработчика через return. Если мы забудем про "return", то обработчик
	// продолжит работу и выведет сообщение
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home_page_tmpl",
		"./ui/html/base_layout_tmpl",
		"./ui/html/footer_partial_tmpl",
	}

	ts, err := template.ParseFiles(files...) //Используем функцию template.ParseFiles() для чтения файла шаблона.
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}

	err = ts.Execute(w, nil)
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}
	//w.Write([]byte("Привет из Snippetbox!"))
}

//Обработчик для отображения содержимого
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Извлекаем значение параметра id из URL и попытаемся
	// конвертировать строку в integer используя функцию strconv.Atoi(). Если его нельзя
	// конвертировать в integer, или значение меньше 1, возвращаем ответ
	// 404 - страница не найдена!
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d", id)
}

//обраб для создания новой заметки
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Используем r.Method для проверки, использует ли запрос метод POST или нет. Обратите внимание,
	// что http.MethodPost является строкой и содержит текст "POST".
	if r.Method != http.MethodPost {
		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
		// карту HTTP-заголовков. Первый параметр - название заголовка, а
		// второй параметр - значение заголовка.
		w.Header().Set("Allow", http.MethodPost)
		// Используем функцию http.Error() для отправки кода состояния 405
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Форма создания новой заметки..."))
}
