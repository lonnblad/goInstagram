package network

import (
	"io/ioutil"
	"log"
	"net/http"
)

type NetworkInterface interface {
	GetRequest(string) []byte
}

type NetworkStruct struct {
}

func (n *NetworkStruct) GetRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%+v\n", err)
		return nil
	}

	return body
}
