package errors

import "errors"

var (
	ErrNeedDatabase       = errors.New("need database")
	ErrNeedCollectionName = errors.New("need collection name")
)
