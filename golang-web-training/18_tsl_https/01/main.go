package main

import "net/http"

func main() {
	http.HandleFunc("/", foo)
	//generate certs
	//go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server .\n"))
}
