package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	jsonrpc string
	id      int
	result  string
}

func main() {

	jsonData := []byte(`{
		"jsonrpc":"2.0",
		"method":"eth_getBalance",
		"params":["0xea674fdde714fd979de3edf0f56aa9716b898ec8", "latest"],
		"id":1
	}`)

	resp, err := http.Post("http://10.1.1.20:8545", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	var m Message

	err = json.Unmarshal(body, &m)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m.result) // 为什么没放到m中去？寻求帮助 https://go.dev/blog/json
}
