package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongo.org/mongo-driver/mongo"
	"go.mongo.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb+srv://sagarts:"+os.Getenv("DB_PASS")+"@cluster0.ihtfff2.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// IMPORTANT
var collection *mongo.Collection 

// connect with mongoDB
func init(){
	// client option
	clientOption := options.Client().ApplyURI(connectionString)
 
	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready.")

}