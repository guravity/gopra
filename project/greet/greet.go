package main

import (
	"fmt"
	// "bytes"
	"os"
	"io"
	"net/http"
)

func Greet (writer io.Writer, name string){
	// *bytes.Bufferをio.Writerに変更→bytes.Bufferの挙動と、os.Stdoutの挙動の両方を実現.便利ね.
	fmt.Fprintf(writer ,"Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request){
	Greet(w, "world")
}

func main(){ // go run greet.goで動く
	if os.Args[1] == "os"{
		Greet(os.Stdout, "Elodie")
	} else if os.Args[1] == "net" {
		http.ListenAndServe(":8000", http.HandlerFunc(MyGreeterHandler))
	}
}