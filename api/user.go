package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HTTPClientInterface interface {
	Do(req *http.Request) (*http.Response, error)

}

type Service struct{
	HTTPClient HTTPClientInterface
}

func (s *Service) GetRandomUser (w http.ResponseWriter, r *http.Request) {
	apiURL := "httos://"
}