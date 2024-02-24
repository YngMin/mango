package ymgo

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Update struct {
	ctx      Context
	filter   bson.M
	addToSet bson.M
	push     bson.M
	set      bson.M
	unset    bson.M
}

func UpdateQuery(ctx Context) *Update {
	return &Update{
		ctx:      ctx,
		filter:   make(bson.M),
		set:      make(bson.M),
		push:     make(bson.M),
		addToSet: make(bson.M),
		unset:    make(bson.M),
	}
}

func (u *Update) Equals(key string, value any) *Update {
	u.filter[key] = value
	return u
}

func (u *Update) In(key string, values any) *Update {
	u.filter[key] = bson.M{
		"$in": values,
	}
	return u
}

func (u *Update) Like(key, value string, ignoreCase bool) *Update {
	regexValue := bson.M{
		"$regex": value,
	}
	if ignoreCase {
		regexValue["$options"] = "i"
	}
	u.filter[key] = regexValue
	return u
}

func (u *Update) GT(key string, value any) *Update {
	u.filter[key] = bson.M{
		"$gt": value,
	}
	return u
}

func (u *Update) GTE(key string, value any) *Update {
	u.filter[key] = bson.M{
		"$gte": value,
	}
	return u
}

func (u *Update) LT(key string, value any) *Update {
	u.filter[key] = bson.M{
		"$lt": value,
	}
	return u
}

func (u *Update) LTE(key string, value any) *Update {
	u.filter[key] = bson.M{
		"$lte": value,
	}
	return u
}

func (u *Update) AddToSet(key string, values any) *Update {
	u.addToSet[key] = bson.M{
		"$each": values,
	}
	return u
}

func (u *Update) Push(key string, values any) *Update {
	u.push[key] = bson.M{
		"$each": values,
	}
	return u
}

func (u *Update) Set(key string, value any) *Update {
	u.set[key] = value
	return u
}

func (u *Update) Unset(key string, value any) *Update {
	u.unset[key] = value
	return u
}

func (u *Update) UpdateOne(target any) (result UpdateResult, err error) {
	collection := u.ctx.db.Collection(getCollectionName(target))
	result, err = collection.UpdateOne(u.ctx, u.filter, u.updateValue())
	return
}

func (u *Update) UpdateMany(target any) (result UpdateResult, err error) {
	collection := u.ctx.db.Collection(getCollectionName(target))
	result, err = collection.UpdateMany(u.ctx, u.filter, u.updateValue())
	return
}

func (u *Update) updateValue() (m bson.M) {
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
