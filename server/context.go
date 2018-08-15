package server

import (
	"context"
	"io"
	"net/http"
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
)

type ContextFactoryInterface interface {
	CreateContext(replaceableCtx context.Context) (*Context, error)
}

type ContextInterface interface {
	context.Context
	ShouldBindJSON(r *io.Reader, i interface{}) error
	JSONResponse(w *http.ResponseWriter, i interface{}, statusCode int)
	BadJSONResponse(w *http.ResponseWriter, i interface{})
}

type Context struct {
	context.Context
	validator *validator.Validate
}

type ContextFactory struct {
	validator *validator.Validate
}

func (cf *ContextFactory) CreateContext(replaceableCtx context.Context) (*Context, error) {
	return &Context{
		replaceableCtx,
		validator.New(),
	}, nil
}

func (c *Context) ShouldBindJSON(r io.Reader, i interface{}) error {
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(i); nil != err {
		return err
	}

	return c.validator.Struct(i)
}

func BadJSONResponse(w *http.ResponseWriter, i interface{}) {

}

func (c *Context) JSONResponse(w *http.ResponseWriter, i interface{}) {

}
