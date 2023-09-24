package controllers

import (
	"net/http"

	"github.com/sriharshamur/lenslocked/views"
)

func StaticHandler(tpl views.Template, data func(r *http.Request) interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, data(r))
	}
}
