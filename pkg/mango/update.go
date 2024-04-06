package mango

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateQuery[T any] struct {
	ctx        Context
	collection *Collection[T]
	filter     bson.M
	addToSet   bson.M
	push       bson.M
	set        bson.M
	unset      bson.M
}

func Update[T any](ctx Context, collection *Collection[T]) *UpdateQuery[T] {
	if collection == nil {
		return nil
	}
	return &UpdateQuery[T]{
		ctx:        ctx,
		collection: collection,
		filter:     make(bson.M),
		set:        make(bson.M),
		push:       make(bson.M),
		addToSet:   make(bson.M),
		unset:      make(bson.M),
	}
}

func (u *UpdateQuery[T]) Equals(key string, value any) *UpdateQuery[T] {
	u.filter[key] = value
	return u
}

func (u *UpdateQuery[T]) In(key string, values any) *UpdateQuery[T] {
	u.filter[key] = bson.M{
		"$in": values,
	}
	return u
}

func (u *UpdateQuery[T]) Like(key, value string, ignoreCase bool) *UpdateQuery[T] {
	regexValue := bson.M{
		"$regex": value,
	}
	if ignoreCase {
		regexValue["$options"] = "i"
	}
	u.filter[key] = regexValue
	return u
}

func (u *UpdateQuery[T]) GT(key string, value any) *UpdateQuery[T] {
	u.filter[key] = bson.M{
		"$gt": value,
	}
	return u
}

func (u *UpdateQuery[T]) GTE(key string, value any) *UpdateQuery[T] {
	u.filter[key] = bson.M{
		"$gte": value,
	}
	return u
}

func (u *UpdateQuery[T]) LT(key string, value any) *UpdateQuery[T] {
	u.filter[key] = bson.M{
		"$lt": value,
	}
	return u
}

func (u *UpdateQuery[T]) LTE(key string, value any) *UpdateQuery[T] {
	u.filter[key] = bson.M{
		"$lte": value,
	}
	return u
}

func (u *UpdateQuery[T]) AddToSet(key string, values any) *UpdateQuery[T] {
	u.addToSet[key] = bson.M{
		"$each": values,
	}
	return u
}

func (u *UpdateQuery[T]) Push(key string, values any) *UpdateQuery[T] {
	u.push[key] = bson.M{
		"$each": values,
	}
	return u
}

func (u *UpdateQuery[T]) Set(key string, value any) *UpdateQuery[T] {
	u.set[key] = value
	return u
}

func (u *UpdateQuery[T]) Unset(key string, value any) *UpdateQuery[T] {
	u.unset[key] = value
	return u
}

func (u *UpdateQuery[T]) One() (result *mongo.UpdateResult, err error) {
	result, err = u.collection.UpdateOne(u.ctx, u.filter, u.values())
	return
}

func (u *UpdateQuery[T]) Many() (result *mongo.UpdateResult, err error) {
	result, err = u.collection.UpdateMany(u.ctx, u.filter, u.values())
	return
}

func (u *UpdateQuery[T]) values() (m bson.M) {
	m = make(bson.M)
	if len(u.set) > 0 {
		m["$set"] = u.set
	}
	if len(u.unset) > 0 {
		m["$unset"] = u.unset
	}
	if len(u.addToSet) > 0 {
		m["$addToSet"] = u.addToSet
	}
	if len(u.push) > 0 {
		m["$push"] = u.push
	}
	return
}
