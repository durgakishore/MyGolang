package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type data struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"Name"`
	Age         int32              `bson:"Age"`
	EmpID       int32              `bson:"EmpID"`
	Designation string             `bson:"Designation"`
	Comments    Review             `bson:"Comments"`
}

//Review struct -- holds comments data
type Review struct { //should have use capital letter otherwise it will not export
	Comment string `bson:"comment"`
	By      string `bson:"by"`
}

func main() {

	dboperations := database{}
	dboperations.GetConnection("27017", "mydb", "empDetails")

	doc := data{
		Name:        "Kishore",
		Age:         25,
		EmpID:       111,
		Designation: "SSE",
		Comments:    Review{Comment: "good ", By: "Ravi"},
	}
	dboperations.InsertOne(doc)

	document := []data{
		data{
			Name:        "Kishore1",
			Age:         33,
			EmpID:       222,
			Designation: "SSE",
			Comments:    Review{Comment: "good ", By: "Ravi"},
		},
		data{
			Name:        "Kishore2",
			Age:         32,
			EmpID:       333,
			Designation: "SSE",
			Comments:    Review{Comment: "nice", By: "Sunil"},
		},

		data{
			Name:        "Kishore3",
			Age:         33,
			EmpID:       444,
			Designation: "SSE",
			Comments:    Review{Comment: "still needs improvement ", By: "Ravi"},
		},

		data{
			Name:        "Kishore4",
			Age:         34,
			EmpID:       555,
			Designation: "SSE",
			Comments:    Review{Comment: "improved alot", By: "Ravi"},
		},
	}

	//https://stackoverflow.com/questions/44319906/why-golang-struct-array-cannot-be-assigned-to-an-interface-array

	d := make([]interface{}, len(document))
	for i, v := range document {
		d[i] = v
	}

	res, err := dboperations.InsertMany(d)

	if err != nil {
		fmt.Println("error while inserting multiple documents")
	} else {
		for _, uid := range res {
			fmt.Println(uid.(primitive.ObjectID).Hex())
		}
	}

	//docID := "5f6edeaf607aca0ef1bb4fd5"
	//objID, _ := primitive.ObjectIDFromHex(docID)

	filter := bson.M{"Age": bson.M{"$eq": "Kishore3"}}
	update := bson.M{"$set": bson.M{"Age": 30}}

	err = dboperations.UpdateOne(filter, update)
	if err != nil {
		fmt.Println("Error while updatOne execution")
	}

	filter = bson.M{"Age": bson.M{"$lte": 32}}
	update = bson.M{"$set": bson.M{"Designation": "SE", "Comments.commet": "Focus more on design"}}

	err = dboperations.UpdateMany(filter, update)
	if err != nil {
		fmt.Println("Error while updatOne execution")
	}

	filter = bson.M{"EmpID": bson.M{"$eq": 333}}
	dboperations.DeleteOne(filter)
	if err != nil {
		fmt.Println("Error while DeleteOne execution")
	}

	filter = bson.M{"Age": bson.M{"$lte": 33}}
	dboperations.DeleteMany(filter)
	if err != nil {
		fmt.Println("Error while DeleteOne execution")
	}

	filter = bson.M{} //Delete all the data from a collection  like RemoveAll()
	dboperations.DeleteMany(filter)
	if err != nil {
		fmt.Println("Error while Deletemany execution")
	}
}
