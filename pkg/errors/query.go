package errors

import (
	"fmt"
)

type BasicQueryInfo struct {
	collection string
	filter     any
	update     any
	doc        any
}

func (b *BasicQueryInfo) ErrorMSG() string {
	msg := "| {query info: "
	if b.filter != nil {
		msg += fmt.Sprintf(" filter: %+v", b.filter)
	}

	if b.update != nil {
		msg += fmt.Sprintf(", update: %+v", b.update)
	}

	if b.doc != nil {
		msg += fmt.Sprintf(", doc: %+v", b.doc)
	}
	msg += "}"
	return msg
}

type NotFound struct {
	BasicQueryInfo
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("%s not found. ", e.collection) + e.BasicQueryInfo.ErrorMSG()
}

type DuplicatedKey struct {
	BasicQueryInfo
	error
}
