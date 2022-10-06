package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MtaData struct {
	IP       string `json:"ip,omitempty" bson:"ip"`
	Hostname string `json:"hostname,omitempty" bson:"hostname"`
	Active   bool   `json:"active,omitempty" bson:"active"`
}

const connectionString = "mongodb+srv://shivendutyagi:Shivendutyagi@cluster0.xcpoa1r.mongodb.net/test"

const dbName = "mta-optimizer"
const colName = "mta"
const X = 1

func (m *MtaData) getMongoConnection() *mongo.Collection {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	collection := client.Database(dbName).Collection(colName)
	return collection

}

func (m *MtaData) GetAllmtas() []string {
	mtaCollection := m.getMongoConnection()
	var resmaptrue = map[string]int{}
	var resmapfalse = map[string]int{}
	// val := os.Getenv("X")
	// intvar, err := strconv.Atoi(val)
	// fmt.Println(val)
	cur, err := mtaCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	num := 1
	for cur.Next(context.Background()) {
		var mta MtaData
		err := cur.Decode(&mta)
		if err != nil {
			log.Fatal(err)
		}
		if mta.Active {
			if count, ok := resmaptrue[mta.Hostname]; ok {
				count += num
				resmaptrue[mta.Hostname] = count
			} else {
				resmaptrue[mta.Hostname] = num
			}
		} else {
			if count, ok := resmapfalse[mta.Hostname]; ok {
				count += num
				resmapfalse[mta.Hostname] = count
			} else {
				resmapfalse[mta.Hostname] = num
			}
		}
	}

	var resultarray []string
	for k, v := range resmaptrue {
		if v <= X {
			resultarray = append(resultarray, k)
		}
	}
	for k, v := range resmapfalse {
		if _, ok := resmaptrue[k]; ok {
			continue
		}
		if v <= X {
			resultarray = append(resultarray, k)
		}
	}

	defer cur.Close(context.Background())

	return resultarray
}

func GetMtaobj() MtaData {
	return MtaData{}
}
