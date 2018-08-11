package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

// parserForm takes in a pointer to a request as well
// as a pointer to the destination
func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	if err := decoder.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
