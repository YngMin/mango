package mango

import (
	"go.mongodb.org/mongo-driver/bson"
)

type UpdateQuery struct {
	ctx      Context
	filter   bson.M
	addToSet bson.M
	push     bson.M
	set      bson.M
	unset    bson.M
}

func Update(ctx Context) *UpdateQuery {
	return &UpdateQuery{
		ctx:      ctx,
		filter:   make(bson.M),
		set:      make(bson.M),
		push:     make(bson.M),
		addToSet: make(bson.M),
		unset:    make(bson.M),
	}
}

func (u *UpdateQuery) Equals(key string, value any) *UpdateQuery {
	u.filter[key] = value
	return u
}

func (u *UpdateQuery) In(key string, values any) *UpdateQuery {
	u.filter[key] = bson.M{
		"$in": values,
	}
	return u
}

func (u *UpdateQuery) Like(key, value string, ignoreCase bool) *UpdateQuery {
	regexValue := bson.M{
		"$regex": value,
	}
	if ignoreCase {
		regexValue["$options"] = "i"
	}
	u.filter[key] = regexValue
	return u
}

func (u *UpdateQuery) GT(key string, value any) *UpdateQuery {
	u.filter[key] = bson.M{
		"$gt": value,
	}
	return u
}

func (u *UpdateQuery) GTE(key string, value any) *UpdateQuery {
	u.filter[key] = bson.M{
		"$gte": value,
	}
	return u
}

func (u *UpdateQuery) LT(key string, value any) *UpdateQuery {
	u.filter[key] = bson.M{
		"$lt": value,
	}
	return u
}

func (u *UpdateQuery) LTE(key string, value any) *UpdateQuery {
	u.filter[key] = bson.M{
		"$lte": value,
	}
	return u
}

func (u *UpdateQuery) AddToSet(key string, values any) *UpdateQuery {
	u.addToSet[key] = bson.M{
		"$each": values,
	}
	return u
}

func (u *UpdateQuery) Push(key string, values any) *UpdateQuery {
	u.push[key] = bson.M{
		"$each": values,
	}
	return u
}

func (u *UpdateQuery) Set(key string, value any) *UpdateQuery {
	u.set[key] = value
	return u
}

func (u *UpdateQuery) Unset(key string, value any) *UpdateQuery {
	u.unset[key] = value
	return u
}

func (u *UpdateQuery) One() (result UpdateResult, err error) {
	collection := u.ctx.db.Collection(u.ctx.collection)
	result, err = collection.UpdateOne(u.ctx, u.filter, u.values())
	return
}

func (u *UpdateQuery) Many() (result UpdateResult, err error) {
	collection := u.ctx.db.Collection(u.ctx.collection)
	result, err = collection.UpdateMany(u.ctx, u.filter, u.values())
	return
}

func (u *UpdateQuery) values() (m bson.M) {
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
