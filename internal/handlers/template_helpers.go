package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var functions template.FuncMap

func renderPage(w http.ResponseWriter, status int, page string, data any) error {
	ts, err := template.New(page).Funcs(functions).ParseFiles("./ui/html/base.gohtml")
	if err != nil {
		return err
	}

	ts, err = ts.ParseGlob("./ui/html/partials/*.gohtml")
	if err != nil {
		return err
	}

	ts, err = ts.ParseFiles(fmt.Sprintf("./ui/html/pages/%s", page))
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	err = ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	_, err = buf.WriteTo(w)
	return err
}
