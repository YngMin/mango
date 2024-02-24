package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
	"ymgo/pkg/options"
	"ymgo/pkg/ymgo"
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

	clientOpts := options.Client()
	clientOpts.ApplyURI("mongodb://root:1234@localhost:27017")

	ctx := context.Background()
	defer ctx.Done()

	mongoCTX := ymgo.NewContext(ctx)

	client, err := ymgo.NewClient(mongoCTX, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("chlngersTest3")

	hex, _ := primitive.ObjectIDFromHex("65ba64579912a58c1b526fed")

	var res PedometerInfo

	mongoCTX.SetDatabase(db)

	err = ymgo.FindQuery(mongoCTX).
		Equals("_id", hex).
		FindOne(&res)

	if err != nil {
		log.Fatal(err)
	}

	PrintStruct(res)
}

func PrintStruct(o any) {
	bytes, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		log.Print(err)
	} else {
		log.Print(string(bytes))
	}
}
