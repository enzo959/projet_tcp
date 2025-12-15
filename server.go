package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func server() {
	ln, err := net.Listen("tcp", "10.36.0.35:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("en attente de connexion ")

	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	fmt.Println("Client connecté !")

	// Goroutine réception
	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, _ := reader.ReadString('\n')
			fmt.Print("Client: " + msg)
		}
	}()

	// Envoi
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text() + "\n"
		conn.Write([]byte(txt))
	}
}
