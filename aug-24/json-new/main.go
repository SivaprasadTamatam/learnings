package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`{"name":"John", "age":30}`)
	var raw json.RawMessage
	json.Unmarshal(data, &raw)
	fmt.Println(string(raw))

}
