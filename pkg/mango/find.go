package mango

import (
	"go.mongodb.org/mongo-driver/bson"
)

type FindQuery[T any] struct {
	ctx        Context
	collection *Collection[T]
	filter     bson.M
	sort       []string
}

func Find[T any](ctx Context, collection *Collection[T]) *FindQuery[T] {
	if collection == nil {
		return nil
	}
	return &FindQuery[T]{
		ctx:        ctx,
		collection: collection,
		filter:     make(bson.M),
		sort:       make([]string, 0),
	}
}

func (f *FindQuery[T]) Equals(key string, value any) *FindQuery[T] {
	f.filter[key] = value
	return f
}

func (f *FindQuery[T]) In(key string, values any) *FindQuery[T] {
	f.filter[key] = bson.M{
		"$in": values,
	}
	return f
}

func (f *FindQuery[T]) Like(key, value string, ignoreCase bool) *FindQuery[T] {
	regexValue := bson.M{
		"$regex": value,
	}
	if ignoreCase {
		regexValue["$options"] = "i"
	}
	f.filter[key] = regexValue
	return f
}

func (f *FindQuery[T]) GT(key string, value any) *FindQuery[T] {
	f.filter[key] = bson.M{
		"$gt": value,
	}
	return f
}

func (f *FindQuery[T]) GTE(key string, value any) *FindQuery[T] {
	f.filter[key] = bson.M{
		"$gte": value,
	}
	return f
}

func (f *FindQuery[T]) LT(key string, value any) *FindQuery[T] {
	f.filter[key] = bson.M{
		"$lt": value,
	}
	return f
}

func (f *FindQuery[T]) LTE(key string, value any) *FindQuery[T] {
	f.filter[key] = bson.M{
		"$lte": value,
	}
	return f
}

func (f *FindQuery[T]) Sort(key string, desc bool) *FindQuery[T] {
	if desc {
		key = "-" + key
	}
	f.sort = append(f.sort, key)
	return f
}

func (f *FindQuery[T]) One() (document T, err error) {
	document, err = f.collection.FindOne(f.ctx, f.filter)
	if err != nil {
		return
	}
	return
}

func (f *FindQuery[T]) All() (documents []T, err error) {
	documents, err = f.collection.Find(f.ctx, f.filter)
	if err != nil {
		return
	}
	return
}
