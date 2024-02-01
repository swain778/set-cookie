package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{}
	cookie.Name = "accessToken"
	cookie.Value = "ro8BS6Hiivgzy8Xuu09JDjlNLnSLldY5"
	cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"
	http.SetCookie(w, &cookie)

	cookie2 := http.Cookie{}
	cookie2.Name = "page"
	cookie2.Value = "GoLinuxCloud"
	cookie2.Expires = time.Now().Add(365 * 24 * time.Hour)
	cookie2.Secure = false
	cookie2.HttpOnly = true
	cookie2.Path = "/"
	http.SetCookie(w, &cookie2)

	fmt.Fprintf(w, "this is cookies!\n")
}

func printCookie(w http.ResponseWriter, req *http.Request) {
	var returnStr string
	for _, cookie := range req.Cookies() {
		returnStr = returnStr + cookie.Name + ":" + cookie.Value + "\n"
	}
	fmt.Fprintf(w, returnStr)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/print", printCookie)
	http.ListenAndServe(":8000", nil)
}
