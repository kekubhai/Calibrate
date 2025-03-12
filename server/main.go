package main
import (
	"net/http"
)
func handler(w httpResponseWriter, r *httpRequest) {

}
func main (){
	http.HandleFunc("w", handler){
		http.ListenServer(":8080", nil )
	}
}