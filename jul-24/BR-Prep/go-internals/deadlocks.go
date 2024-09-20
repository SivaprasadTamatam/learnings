package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	c := 0

	fmt.Scanf("%d", &c)
}
