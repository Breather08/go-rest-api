package router

import (
	"forum/backend/server/database"
	"net/http"
)

func commonHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./client/dist/index.html");
}

func Router() http.Handler {
	mux := http.NewServeMux()

	fs := http.StripPrefix("/", http.FileServer(http.Dir("./client/dist")))
	mux.Handle("/", fs)
	mux.HandleFunc("/profile", commonHandler)
	mux.HandleFunc("/about", commonHandler)
	mux.HandleFunc("/api/posts", database.SendAllPosts)
	mux.HandleFunc("/api/create-post", database.InsertPost)

	return mux
}
