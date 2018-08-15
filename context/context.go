package context

import (
	"context"
	"encoding/json"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"net/http"
)

type ContextFactoryInterface interface {
	CreateContext(replaceableCtx context.Context) (*Context, error)
}

type ContextInterface interface {
	context.Context
	ShouldBindJSON(r io.Reader, i interface{}) error
	JSONResponse(w http.ResponseWriter, i interface{}, statusCode int)
	BadJSONResponse(w http.ResponseWriter, err error)
}

type Context struct {
	context.Context
	validator  *validator.Validate
	translator ut.Translator
}

type ContextFactory struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func (cf *ContextFactory) CreateContext(replaceableCtx context.Context) (*Context, error) {
	return &Context{
		replaceableCtx,
		cf.Validator,
		cf.Translator,
	}, nil
}

func (c *Context) ShouldBindJSON(r io.Reader, i interface{}) error {
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(i); nil != err {
		return err
	}

	return c.validator.Struct(i)
}

func (c *Context) BadJSONResponse(w http.ResponseWriter, err error) {
	errors := make([]string, 0, 3)
	statusCode := 400

	switch err.(type) {
	case *json.SyntaxError:
		statusCode = 400
		errors = append(errors, "Invalid JSON request")
	case *validator.ValidationErrors:
		statusCode = 422
		for _, e := range err.(validator.ValidationErrors) {
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

func (c *Context) JSONResponse(w http.ResponseWriter, i interface{}, statusCode int) {
	w.WriteHeader(statusCode)

	if nil == i {
		w.Write([]byte("{}"))
	}

	r, _ := json.Marshal(&i)
	w.Write(r)
}
