package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	env := []byte(`{"CodEmpresa": "0007"}`)
	resp, err := http.Post("http://172.28.189.39:3000/desligamentoProgramadoGetTotalItens", "application/json", bytes.NewBuffer(env))
	if err != nil {
		log.Fatal(err)
	}
	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\r\n", json)
}
