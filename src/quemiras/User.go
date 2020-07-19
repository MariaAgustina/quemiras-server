package main

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

//para que lo desencodee bien tengo que declarar las variables con mayuscula
type User struct {
    ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Name string             `json:"name,omitempty" bson:"name,omitempty"`
}

//TODO: in other file
type ProfileConfigLikes struct {

}

var usersArray []User

func GetUserEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
    params := mux.Vars(request)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    var user User
    collection := client.Database("quemiras").Collection("users")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    err := collection.FindOne(ctx, User{ID: id}).Decode(&user)
    if err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }
    json.NewEncoder(response).Encode(user)
}

func CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
    var user User
    _ = json.NewDecoder(request.Body).Decode(&user)
    collection := client.Database("quemiras").Collection("users")
    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
    result, _ := collection.InsertOne(ctx, user)
    fmt.Printf("New user inserted %s \n", user.Name)
    json.NewEncoder(response).Encode(result)
}

//TODO
// func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
//     params := mux.Vars(req)
//     for index, item := range usersArray {
//         if item.ID == params["id"] {
//             usersArray = append(usersArray[:index], usersArray[index+1:]...)
//             break
//         }
//     }
//     json.NewEncoder(w).Encode(usersArray)
// }

func GetUsersArrayEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("content-type", "application/json")
    var usersArray []User
    collection := client.Database("quemiras").Collection("users")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var user User
        cursor.Decode(&user)
        usersArray = append(usersArray, user)
    }
    if err := cursor.Err(); err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
        return
    }
    json.NewEncoder(response).Encode(usersArray)
}
