package main

import (
	    "time"
        "context"
        "fmt"
        
    "go.mongodb.org/mongo-driver/bson" 
    "go.mongodb.org/mongo-driver/mongo" 
    "go.mongodb.org/mongo-driver/mongo/options"
)

type User struct{
  InvDt time.Time `bson:"invoice_date,omitempty" json:"invoice_date,omitempty"`
}

func main() {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
        if err != nil {
                panic(err)
        }
        
    testColl := client.Database("go-test").Collection("test")
    
    var user User
    dt,err := time.Parse(time.RFC3339,"2021-06-07T05:58:08.847Z")
    user.InvDt = dt
    inserg, err := testColl.InsertOne(context.TODO(), bson.D{{Key: "InvDt", Value: user.InvDt}})
    
    if err != nil {
            panic(err)
    }

	fmt.Println(inserg)
}
