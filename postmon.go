package postmongo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func BuscarCep(cep string) (map[string]interface{}, error) {

	res, err := http.Get("http://api.postmon.com.br/v1/cep/" + cep)

	if err != nil {
		return nil, err
	}

	var cep_result map[string]interface{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json_body := []byte(body)

	decoding := json.Unmarshal(json_body, &cep_result)

	return cep_result, decoding
}
