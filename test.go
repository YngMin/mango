package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
	"ymgo/pkg/mango"
	"ymgo/pkg/options"
)

type PedometerInfo struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt       time.Time          `bson:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt"`
	Type            string             `bson:"type"`
	MaxPoint        int                `bson:"maxPoint"`
	MaxWalkingCount int                `bson:"maxWalkingCount"`
	ImageAdInfo     AdInfo             `bson:"imageAdInfo"`
	VideoAdInfo     AdInfo             `bson:"videoAdInfo"`
}

func (p *PedometerInfo) CollectionName() string {
	return "pedometerInfo"
}

type AdInfo struct {
	BenefitInfoList []BenefitInfo `bson:"benefitInfoList"`
}

type BenefitInfo struct {
	Point                  *int    `bson:"Point"`
	AdType                 *string `bson:"AdType"`
	WalkingCount           int     `bson:"WalkingCount"`
	AccumulateBenefitPoint int     `bson:"AccumulateBenefitPoint"`
}

func main() {

	clientOpts := options.Client().
		ApplyURI("mongodb://root:1234@localhost:27017")

	ctx := context.Background()
	defer ctx.Done()

	clientCTX := mango.NewContext(ctx)

	client, err := mango.NewClient(clientCTX, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("chlngersTest3")

	hex, _ := primitive.ObjectIDFromHex("65ba64579912a58c1b526fed")

	var dest PedometerInfo

	ctx = context.Background()
	defer ctx.Done()

	mangoCTX := mango.NewContext(ctx)
	mangoCTX.SetDatabase(db)
	mangoCTX.SetCollection(&PedometerInfo{})

	err = mango.Find(mangoCTX).
		Equals("_id", hex).
		One(&dest)

	if err != nil {
		log.Fatal(err)
	}

	PrintStruct(dest)

	result, err := mango.Update(mangoCTX).
		Equals("_id", hex).
		Set("maxPoint", 44).
		One()

	if err != nil {
		log.Fatal(err)
	}

	PrintStruct(result)
}

func PrintStruct(o any) {
	bytes, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		log.Print(err)
	} else {
		log.Print(string(bytes))
	}
}
