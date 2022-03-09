package middleware

import (
	"GuestLedgerBookApi/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

const connectionString = "mongodb+srv://admin:1234@cluster0.svtvd.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const databaseName = "GuestLedgerBook"
const collectionName = "Guests"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB'ye bağlandınız.")
	collection = client.Database(databaseName).Collection(collectionName)
	fmt.Println("Collection oluşturuldu.")
}

// Get all message from the Database
func getAllGuestsFromDB() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results
}
func GetAllGuests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllGuestsFromDB()
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		return
	}
}

// Insert message to Database
func insertOneGuest(guest models.Guest) {
	insertResult, err := collection.InsertOne(context.Background(), guest)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Send a message", insertResult.InsertedID)
}
func CreateGuest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var guest models.Guest
	_ = json.NewDecoder(r.Body).Decode(&guest)
	// fmt.Println(message, r.Body)
	insertOneGuest(guest)
	json.NewEncoder(w).Encode(guest)
}

// Delete message from database
func deleteOneGuest(guest string) {
	fmt.Println(guest)
	id, _ := primitive.ObjectIDFromHex(guest)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}
func DeleteGuest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneGuest(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
