package context

import (
	"context"
	"encoding/json"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type Context interface {
	context.Context
	ShouldBindJSON(r *http.Request, i interface{}) error
	ShouldBindQuery(r *http.Request, i interface{}) error
	JSONResponse(w http.ResponseWriter, i interface{}, statusCode int)
	BadJSONResponse(w http.ResponseWriter, err error)
	InternalErrorResponse(w http.ResponseWriter)
}

type ctx struct {
	context.Context
	validator  *validator.Validate
	translator ut.Translator
}

func (c *ctx) ShouldBindJSON(r *http.Request, i interface{}) error {
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(i); nil != err {
		return err
	}

	return c.validator.Struct(i)
}

func (c *ctx) ShouldBindQuery(r *http.Request, i interface{}) error {
	if err := mapForm(i, r.URL.Query()); nil != err {
		return err
	}

	return c.validator.Struct(i)
}

func (c *ctx) BadJSONResponse(w http.ResponseWriter, err error) {
	errors := make([]string, 0, 3)
	statusCode := 400

	switch err.(type) {
	case *json.SyntaxError:
		statusCode = 400
		errors = append(errors, "Invalid JSON request")
	case validator.ValidationErrors:
		statusCode = 422
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			errors = append(errors, e.Translate(c.translator))
		}
	default:
		statusCode = 422
		errors = append(errors, err.Error())
	}

	m := map[string][]string{
		"errors": errors,
	}

	c.JSONResponse(w, &m, statusCode)
}

func (c *ctx) InternalErrorResponse(w http.ResponseWriter) {
	errors := make([]string, 1, 1)
	errors[0] = "Internal Error. Look at the server logs";
	m := map[string][]string{
		"errors": errors,
	}

	c.JSONResponse(w, &m, 500)
}

func (c *ctx) JSONResponse(w http.ResponseWriter, i interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if nil == i {
		w.Write([]byte("{}"))
		return
	}

	r, _ := json.Marshal(&i)
	w.Write(r)
}
