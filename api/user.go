package api

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type HTTPClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

type Service struct {
	HTTPClient HTTPClientInterface
}

func (s *Service) GetRandomUser(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://randomuser.me/api"

	//tratando erro de requests
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		log.Println(err)
	}

	//tratando erro de response
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	//fechando a conexão após execuções
	defer res.Body.Close()

	//tratamento de erro ao retornar os dados do arquivo
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	//convertendo para texto os dados do arquivo
	fmt.Println(string(data))

	//retorno de sucesso da chamada da API + retorno arquivo
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
