package context

import (
	"context"
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-playground/universal-translator"
)

type Factory interface {
	CreateContext(replaceableCtx context.Context) (Context, error)
}

type factory struct {
	validator  *validator.Validate
	translator ut.Translator
}

func (f *factory) CreateContext(replaceableCtx context.Context) (Context, error) {
	return &ctx{
		replaceableCtx,
		f.validator,
		f.translator,
	}, nil
}

func NewFactory(v *validator.Validate, t ut.Translator) (Factory) {
	return &factory{v, t}
}
