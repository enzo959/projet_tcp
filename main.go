package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	mode := os.Args[1]

	switch mode {
	case "server":
		server()
	case "client":
		client()
	default:
		fmt.Println("erreur utilisez 'server' ou 'client'")
	}
}
