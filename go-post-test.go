package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

type configFile struct {
	URL       string `json:"url"`
	BodyType  string `json:"type"`
	BodyValue string `json:"value"`
}

func main() {
	var (
		fileName string
		cfg      []byte
		ret      []byte
		config   configFile
		resp     *http.Response
		err      error
	)
	// Flags de execução
	flag.StringVar(&fileName, "config", "", "nome do arquivo que contem o json com a configuração da chamada {url, type, value}")
	flag.StringVar(&config.URL, "url", "http://172.28.189.39:3000/desligamentoProgramadoGetTotalItens", "url que se deseja fazer um post test")
	flag.StringVar(&config.BodyType, "type", "application/json", "tipo de passagem de parametros do body")
	flag.StringVar(&config.BodyValue, "value", `{"CodEmpresa": "0007"}`, "valor a ser passado no body")
	flag.Parse() // sempre chamar o parse no final das definições

	if fileName != "" {
		// Tentando carregar o arquivo de configuração
		cfg, err = ioutil.ReadFile(fileName)
		if err != nil {
			log.Printf("Arquivo %s não encontrado.", fileName)
		} else {
			err = json.Unmarshal(cfg, &config)
			if err != nil {
				log.Printf("Não foi possível carregar o arquivo %s.\r\n", fileName)
				log.Println(err)
			} else {
				log.Printf("Arquivo de configurações %s carregado.\r\n", fileName)
			}
		}
	}
	log.Println("Chamando: ", config.URL)
	log.Println("Com o tipo: ", config.BodyType)
	log.Println("Valor: ", config.BodyValue)
	resp, err = http.Post(config.URL, config.BodyType, bytes.NewBuffer([]byte(config.BodyValue)))
	if err != nil {
		log.Fatalln(err)
	}
	ret, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Retorno: ", string(ret))
}
