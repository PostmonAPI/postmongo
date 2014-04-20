[![Build Status](https://travis-ci.org/PostmonAPI/postmon-golang.svg)](https://travis-ci.org/PostmonAPI/postmon-golang)

PostmonGo
===============

Go Wrapper for Postmon API

Uso
----
    
    package main
    
    import (
        "fmt"
        "github.com/PostmonAPI/postmongo"
    )

    func main() {
        res, err := postmongo.BuscarCep("01419101")
        fmt.Println(res)
        fmt.Println(err)
    }

Licença
--------

    Copyright 2014 The Postmon API Team

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
