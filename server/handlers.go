package server

import (
	"html/template"
	"net/http"
	"sync"
)

func (s *server) handleTemplate(data interface{}, files ...string) http.HandlerFunc {
	var (
		init sync.Once
		tpl  *template.Template
		err  error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tpl, err = s.parseTemplates(w, files...)
			if err != nil {
				s.logErr("error parsing template files", err)
			}
		})
		w.Header().Set("Content-Type", "text/html")
		err = tpl.ExecuteTemplate(w, "app", s.parseData(nil))
		if err != nil {
			s.logErr("error executing template", err)
		}
	}
}
