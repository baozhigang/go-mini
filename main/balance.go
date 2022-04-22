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
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
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

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body)) // {"jsonrpc":"2.0","id":1,"result":"0x3d69b16b62d47b5ea0"}

	var m Message

	err = json.Unmarshal(body, &m)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m.Result) // 0x3d69b16b62d47b5ea0  16进制转10进制即可
}
