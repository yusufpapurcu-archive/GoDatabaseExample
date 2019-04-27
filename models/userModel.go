package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct for Create and Manage User data's
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`            // User ID
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`         // User Name
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`       // User Email
	Username string             `bson:"username,omitempty" json:"username,omitempty"` // User Username
	Password string             `bson:"password,omitempty" json:"password,omitempty"` // User Password
}

// UpdateUser struct for UpdateUser
type UpdateUser struct {
	Name        string `bson:"name,omitempty" json:"name,omitempty"`            // User Name
	Email       string `bson:"email,omitempty" json:"email,omitempty"`          // User Email
	Username    string `bson:"username,omitempty" json:"username,omitempty"`    // User Username
	Password    string `bson:"password,omitempty" json:"password,omitempty"`    // User Password
	NewName     string `bson:"name,omitempty" json:"newname,omitempty"`         // NewUser Name
	NewEmail    string `bson:"email,omitempty" json:"newemail,omitempty"`       // NewUser Email
	NewUsername string `bson:"username,omitempty" json:"newusername,omitempty"` // NewUser Username
	NewPassword string `bson:"password,omitempty" json:"newpassword,omitempty"` // NewUser Password
}
