package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/hello", helloFunc)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Listen create failure：",err.Error())
	}
}

func helloFunc(w http.ResponseWriter, r *http.Request)  {
	out := ""

	fmt.Println("Print Header parameter list：")
	out = out + fmt.Sprintf("Print Header parameter list：\n")
	if len(r.Header) > 0 {
		for k,v := range r.Header {
			fmt.Printf("%s=%s\n", k, v[0])
			out = out + fmt.Sprintf("%s=%s\n", k, v[0])
		}
	}
	fmt.Println("Print Form parameter list：")
	out = out + fmt.Sprintf("Print Form parameter list：\n")
	r.ParseForm()
	if len(r.Form) > 0 {
		for k,v := range r.Form {
			fmt.Printf("%s=%s\n", k, v[0])
			out = out + fmt.Sprintf("%s=%s\n", k, v[0])
		}
	}
	//verify username & password，if reply session in the response header
	//If request failure, relay StatusUnauthorized code
	w.WriteHeader(http.StatusOK)
	if (r.Form.Get("user") == "admin") && (r.Form.Get("pass") == "888") {
		w.Write([]byte("hello, verify success!\n" + out))
	} else {
		w.Write([]byte("hello, verify failure！"))
	}
}