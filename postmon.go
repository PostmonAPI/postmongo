package postmon

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

func BuscarCep(cep string) (map[string]interface{}, error) {

    res, _ := http.Get("http://api.postmon.com.br/v1/cep/" + cep)

    var cep_result map[string]interface{}

    body, _ := ioutil.ReadAll(res.Body)

    json_body := []byte(body)

    decoding := json.Unmarshal(json_body, &cep_result)

    return cep_result, decoding
}
