package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id 		int			`json:"id"`
	Title string	`json:"title"`
	Text 	string	`json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{{Id: 1, Title: "Title 1", Text: "Someone"}}
}

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var result, err = json.Marshal(posts)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	var post Post
	var decodeError = json.NewDecoder(request.Body).Decode(&post)
	response.Header().Set("Content-Type", "application/json")

	if decodeError != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling the request."}`))
		return
	}

	post.Id = len(posts) + 1
	posts = append(posts, post)

	var result, err = json.Marshal(post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling the post"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
}