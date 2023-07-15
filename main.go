package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, r.URL, " ", r.Method, " ", r.Host, "\n")
	fmt.Fprint(w, "<h1>That's root main page</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch , email me at <a href=\"mailto:serjant.saygon@gmail.com\">serjant.saygon@gmail.com</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	fmt.Fprint(w, "<h1>FAQ Page</h1><ul>"+
		"<li>Q: <b>Is there a free version ?</b></li>"+
		"<li>A: Yes, we offer a free trial for 30 days ?</li>"+
		"<li>Q: <b>Is there a free version ?</b></li>"+
		"<li>A: yes ?</li>"+
		"<li>Q: <b>Is there a free version ?</b></li>"+
		"<li>A: Is there a free version ?</li>"+"</ul>")
}

//	func pathHandler(w http.ResponseWriter, r *http.Request) {
//		switch r.URL.Path {
//		case "/":
//			homeHandler(w, r)
//		case "/contacts":
//			contactHandler(w, r)
//		default:
//			http.Error(w, "Page not found", http.StatusNotFound)
//		}
//	}
type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contacts":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting server on :3001...")
	http.ListenAndServe(":3001", router)
}
