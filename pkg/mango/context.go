package mango

import "context"

type Context struct {
	context.Context
	db         *Database
	collection ICollection
}

func NewContext(ctx context.Context) Context {
	return Context{Context: ctx}
}

func (ctx *Context) SetDatabase(db *Database) {
	ctx.db = db
}

func (ctx *Context) SetCollection(o ICollection) {
	ctx.collection = o
}
