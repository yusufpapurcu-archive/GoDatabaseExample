package database

import (
	"context"
	"encoding/json"
	"time"

	"github.com/YTPSourceCode/DatabaseTest/models"

	"go.mongodb.org/mongo-driver/bson"
)

// GetAllUser function for Getting all Users
func GetAllUser() ([]string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second) // Context for Find function
	cur, err := users.Find(ctx, bson.D{})                               // Empty Find function
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)   // Close cursor
	var result models.User // Create Models.User for Decode result
	var list []string      // Decoded and Transformed Json's String
	for cur.Next(ctx) {    // Loop all cursors
		err := cur.Decode(&result) // Decode
		if err != nil {
			return nil, err
		}
		out, err := json.Marshal(result) // Transform JSON
		if err != nil {
			return nil, err
		}
		list = append(list, string(out)) // Storage the string array
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

// CreateUser function for Create User
func CreateUser(usr models.User) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // Context for Create function
	res, err := users.InsertOne(ctx, usr)                              // Create Function
	if err != nil {
		return nil, err
	}
	id := res.InsertedID
	return id, nil
}

// GetOneUser function for Serach User
func GetOneUser(usr models.User) (string, error) {
	var filtre bson.M               // Filtre variable created
	var result models.User          // For decode
	bytes, err := bson.Marshal(usr) // User convert the []byte
	if err != nil {
		return "", err
	}
	bson.Unmarshal(bytes, &filtre)                                     // Bytes convert filtre
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // Context for Serach
	err = users.FindOne(ctx, filtre).Decode(&result)                   // FindOne function and decode result
	if err != nil {
		return "", err
	}
	out, err := json.Marshal(&result) // Convert the Json
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// DeleteUser function for delete user
func DeleteUser(usr models.User) error {
	var filtre bson.M               // Filtre variable created
	var result models.User          // For decode
	bytes, err := bson.Marshal(usr) // User convert the []byte
	if err != nil {
		return err
	}
	bson.Unmarshal(bytes, &filtre)                                     // Bytes convert filtre
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // Context for Serach
	err = users.FindOne(ctx, filtre).Decode(&result)                   // FindOne function and decode result
	if err != nil {
		return err
	}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second) // Context for Delete
	_, err = users.DeleteOne(ctx, bson.D{{"_id", result.ID}})         // Delete User
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser function for update user
func UpdateUser(old, new models.User) error {
	var filtre bson.M               // Filtre variable created
	bytes, err := bson.Marshal(old) // OldUser convert the []byte
	if err != nil {
		return err
	}
	bson.Unmarshal(bytes, &filtre) // Bytes convert filtre
	var usr bson.M                 // User variable created
	bytes, err = bson.Marshal(new) // NewUser convert the []byte
	if err != nil {
		return err
	}
	bson.Unmarshal(bytes, &usr)
	update := bson.D{ // Updated Document
		{"$set", usr}, // $set paramater and usr variable setted
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // Context for Update function
	_, err = users.UpdateOne(ctx, filtre, update)                      // Update Document
	if err != nil {
		return err
	}
	return nil
}
