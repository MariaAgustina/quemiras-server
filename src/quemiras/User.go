package main

import (
    "encoding/json"
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
)

//para que lo desencodee bien tengo que declarar las variables con mayuscula
type User struct {
    ID      string  `json:"id,omitempty"`
    Name    string  `json:"name"`
}

//TODO: in other file
type ProfileConfigLikes struct {

}

var usersArray []User

func GetUserEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range usersArray {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&User{})
}

func CreateUserEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var user User
    _ = json.NewDecoder(req.Body).Decode(&user)
    fmt.Printf("New user is %s", user.Name)
    user.ID = params["id"]
    usersArray = append(usersArray, user)
    json.NewEncoder(w).Encode(user)
}

func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range usersArray {
        if item.ID == params["id"] {
            usersArray = append(usersArray[:index], usersArray[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(usersArray)
}

func GetUsersArrayEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(usersArray)
}
