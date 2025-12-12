package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func client() {
	conn, err := net.Dial("tcp", ":8080") // <-- IP du serveur
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection établi !")

	// Réception
	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, _ := reader.ReadString('\n')
			fmt.Print("Serveur: " + msg)
		}
	}()

	// Envoi
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text() + "\n"
		conn.Write([]byte(txt))
	}
}
