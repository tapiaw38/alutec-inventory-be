package appcontext

import (
	"github.com/tapiaw38/alutec-inventory-be/internal/adapters/datasources/repositories"
)

type Context struct {
	Repositories *repositories.Repositories
}

type Factory func() *Context

func NewFactory(repos *repositories.Repositories) Factory {
	ctx := &Context{
		Repositories: repos,
	}
	return func() *Context {
		return ctx
	}
}
