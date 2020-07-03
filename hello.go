package main

import (
        "encoding/json"
        "net/http"
	"fmt"
)

type Post struct {
    Title  string `json:"Title"`
    Author string `json:"Author"`
    Text   string `json:"Text"`
}


func PostsHandler(w http.ResponseWriter, r *http.Request) {
 
    posts := []Post{
        Post{"Post one", "John", "This is first post."},
        Post{"Post two", "Jane", "This is second post."},
        Post{"Post three", "John", "This is another post."},
    }
 
    json.NewEncoder(w).Encode(posts)
}

func main() {
	fmt.Println("started server in localhost:5051")
	http.HandleFunc("/posts", PostsHandler)
	http.ListenAndServe(":5051", nil)
}
