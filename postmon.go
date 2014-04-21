// Copyright 2014 The Postmon API Team. All rights reserved.
// Use of this source code is governed by an Apache-style
// license that can be found in the README.md file

package postmongo

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

// BuscarCep returns a map with all results about the CEP
// An error will be launched if any problems occur
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
