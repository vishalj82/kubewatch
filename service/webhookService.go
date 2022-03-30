package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	v1 "k8s.io/api/admission/v1"
	"io/ioutil"
)

func server(w http.ResponseWriter, req *http.Request) {
    fmt.Println("Serving request", req.URL)
	data, err := ioutil.ReadAll(req.Body)
	reqJson := string(data)
    fmt.Println("Request byte", reqJson)
	deserializer := codecs.UniversalDeserializer()
	obj, gvk, err := deserializer.Decode(data, nil, nil)
	fmt.Println("Request object", obj)
	if err != nil {
	    msg := fmt.Sprintf("Request could not be decoded: %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	if arv1, ok := obj.(*v1.AdmissionReview); ok {
        responseAdmissionReview := &v1.AdmissionReview{}
        responseAdmissionReview.SetGroupVersionKind(*gvk)
		responseAdmissionReview.Response = allowAll(*arv1)
		responseAdmissionReview.Response.UID = arv1.Request.UID
		respBytes, _ := json.Marshal(responseAdmissionReview)
		s := string(respBytes)
		fmt.Println("Response byte", s)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respBytes)
		return
	}
	w.Write([]byte("<h1>Welcome to my web server Again!</h1>"))
}

func main() {
	http.HandleFunc("/", server)
	http.HandleFunc("/pods", server)
	err := http.ListenAndServeTLS(":443", "/opt/kubewatch/certs/tls.crt", "/opt/kubewatch/certs/tls.key", nil)
    if err != nil {
        fmt.Println("ListenAndServe: ", err)
    }
}
