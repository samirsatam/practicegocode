package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// showSnippet handler
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// func createSnippet(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Create a new a snippet..."))
// }

func createSnippetv1(w http.ResponseWriter, r *http.Request) {
	//Make it so it responds to POST only.
	// MethodPost is a constant equal to the string "POST".
	if r.Method != http.MethodPost {
		// if not a POST request send a 405 [see below]
		// Tells the client which HTTP verbs are allowed.
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader can only be called once per response.
		// If not called explicitly, the firstcall to .Write will send a 200 ok
		w.WriteHeader(405)
		//Changing the response header map after a call to w.WriteHeader() or
		// w.Write() will have no effect on the headers that the user receives. You need to make
		// sure that your response header map contains all the headers you want before you call
		// these methods.
		w.Write([]byte("Method Not Allowed"))
	}
	w.Write([]byte("Create a new a snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	//Make it so it responds to POST only.
	// MethodPost is a constant equal to the string "POST".
	if r.Method != http.MethodPost {
		// if not a POST request send a 405 [see below]
		// Tells the client which HTTP verbs are allowed.
		w.Header().Set("Allow", http.MethodPost)
		// Use the http.Error() function to send a 405 status code and "Method Not
		// Allowed" string as the response body.
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new a snippet..."))
}
func howToSetOtherHeaderAttributes(h *http.Header) {
	// Set a new cache-control header. If an existing "Cache-Control" header exists
	// it will be overwritten.
	h.Set("Cache-Control", "public, max-age=31536000")
	// In contrast, the Add() method appends a new "Cache-Control" header and can
	// be called multiple times.
	h.Add("Cache-Control", "public")
	h.Add("Cache-Control", "max-age=31536000")
	// Delete all values for the "Cache-Control" header.
	h.Del("Cache-Control")
	// Retrieve the first value for the "Cache-Control" header.
	h.Get("Cache-Control")
	// set content-type specifically.
	h.Set("Content-Type", "application/json")
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	// The "/" is treated as a catch-all, and will be used for all unregistered routes.
	mux.HandleFunc("/", home)
	// Both examples are of type 'Fixed Paths'
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
