package main1
import (
	"fmt"
	"log"
	"net/http"
)
func formHndler(w http.ResponseWriter,r *http.Request){
	if err := r.ParseForm(); err !=nil {
		fmt.Fprintf(w,"ParseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w,"POST request successeful\n")
	name :=r.FormValue("name")
	address :=r.FormValue("address")

	fmt.Fprintf(w,"Name= %s\n",name)
	fmt.Fprintf(w,"Address= %s\n",address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
if r.URL.Path!="/hello"{
	http.Error(w,"404 not found",http.StatusNotFound)
return
}
if r.Method != "GET"{
	http.Error(w,"Method is not supported",http.StatusNotFound)
	return
}
fmt.Fprintf(w,"hello!")
 }


func main(){

fileserver:= http.FileServer(http.Dir("./ststic"))
http.Handle("/",fileserver)
http.HandleFunc("/form",formHndler)
http.HandleFunc("/hello",helloHandler)
 fmt.Printf("start server port 8000\n")
 if err:=http.ListenAndServe(":8000",nil); err !=nil{
log.Fatal(err)
 }

}