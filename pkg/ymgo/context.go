package ymgo

import "context"

type Context struct {
	context.Context
	db *Database
}

func NewContext(ctx context.Context) Context {
	return Context{Context: ctx}
}

func (ctx *Context) SetDatabase(db *Database) {
	ctx.db = db
}
