package main

import (
	"apimockgo/api"
	"log"
	"net/http"
)


func main(){
	service := api.Service {
		HTTPClient: http.DefaultClient,
	}

	//raiz do projeto e ação de sua chamada
	http.HandleFunc("/", service.GetRandomUser)
	log.Println("listando...")

	//tratamento de conexão aberta com panic
	panic(http.ListenAndServe("localhost: 7000", nil))
}