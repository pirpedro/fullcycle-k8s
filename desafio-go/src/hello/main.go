package main

 import (
 	"fmt"
 	"net/http"
 )

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	 	fmt.Fprintf(w, greetings("Code.education Rocks!"))
	})
	http.ListenAndServe(":8000", nil)
}

func greetings(s string) string{
	return "<b>" + s + "</b>"
}
