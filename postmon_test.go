// Copyright 2014 The Postmon API Team. All rights reserved.
// Use of this source code is governed by an Apache-style
// license tha can be found in the README.md file

package postmongo

import (
    "reflect"
    "testing"
)

// TestPostmon test if BuscarCep returns the expected result
func TestPostmon(t *testing.T) {
    info := map[string]interface{}{
        "complemento": "de 2161 ao fim - lado ímpar",
        "bairro":      "Cerqueira César",
        "cidade":      "São Paulo",
        "cep":         "01419101",
        "logradouro":  "Alameda Santos",
        "estado_info": map[string]interface{}{
            "area_km2":    "248.221,996",
            "codigo_ibge": "35",
            "nome":        "São Paulo",
        },
        "cidade_info": map[string]interface{}{
            "area_km2":    "1521,11",
            "codigo_ibge": "3550308",
        },
        "estado": "SP",
    }

    res, err := BuscarCep("01419101")
    if err != nil {
        t.Errorf("%v", err)
    }

    if equal := reflect.DeepEqual(res, info); equal == false {
        t.Errorf("CEP 01419101 = %v, want %v", res, info)
    }
}
