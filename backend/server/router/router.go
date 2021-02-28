package router

import (
	"forum/backend/server/database"
	"net/http"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./client/dist"))
	mux.Handle("/", fs)

	mux.HandleFunc("/api/posts", database.SendAllPosts)
	mux.HandleFunc("/api/create-post", database.InsertPost)

	return mux
}
