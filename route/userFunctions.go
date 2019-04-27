package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YTPSourceCode/DatabaseTest/database"
	"github.com/YTPSourceCode/DatabaseTest/models"
)

// GetUser function for Get AllUsers in database
func GetUser(w http.ResponseWriter, r *http.Request) {
	data, err := database.GetAllUser() // Getting all Users from database package's function
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "Some problems exist")
		return
	}
	fmt.Fprint(w, data)
}

// CreateUser function for Create user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var usr models.User                // Create User Struct for Decode Request.Body
	decoder := json.NewDecoder(r.Body) // Decode request.Body
	defer r.Body.Close()               // Close After Function
	err := decoder.Decode(&usr)        // Decode Request.Body to User struct
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Some problems exist")
		return
	}
	id, err := database.CreateUser(usr) // Create User from database package's function
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Some problems exist")
		return
	}
	fmt.Fprint(w, id)
}

// GetOneUser function for find one user in database
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	var usr models.User                // Create User Struct for Decode Request.Body
	decoder := json.NewDecoder(r.Body) // Decode request.Body
	defer r.Body.Close()               // Close After Function
	err := decoder.Decode(&usr)        // Decode Request.Body to User struct
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Some problems exist")
		return
	}
	result, err := database.GetOneUser(usr) // Get One user from database package's function
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Some problems exist")
		return
	}
	fmt.Fprint(w, result)
}

// DeleteUser function for delete one user in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var usr models.User                // Create User Struct for Decode Request.Body
	decoder := json.NewDecoder(r.Body) // Decode request.Body
	defer r.Body.Close()               // Close After Function
	err := decoder.Decode(&usr)        // Decode Request.Body to User struct
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Some problems exist")
		return
	}
	err = database.DeleteUser(usr) // Get One user from database package's function
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Some problems exist")
		return
	}
	fmt.Fprintf(w, "Kullanici Silindi")
}

// UpdateUser function for
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var tmp models.UpdateUser          // Create UpdateUser Struct for Decode Request.Body
	decoder := json.NewDecoder(r.Body) // Decode request.Body
	defer r.Body.Close()               // Close After Function
	err := decoder.Decode(&tmp)        // Decode Request.Body to User struct
	if err != nil {
		fmt.Print("Decode")
		fmt.Println(err)
		fmt.Fprintf(w, "Some problems exist")
		return
	}
	var new = models.User{ // Create Updated user
		Name:     tmp.NewName,
		Username: tmp.NewUsername,
		Password: tmp.NewPassword,
		Email:    tmp.NewEmail,
	}
	var old = models.User{ // Create old user struct for filter
		Name:     tmp.Name,
		Username: tmp.Username,
		Password: tmp.Password,
		Email:    tmp.Email,
	}
	err = database.UpdateUser(old, new) // Get One user from database package's function
	if err != nil {
		fmt.Print("Update User")
		fmt.Println(err)
		fmt.Fprintf(w, "Some problems exist")
		return
	}
	fmt.Fprint(w, "Basari ile guncelledik")
}
