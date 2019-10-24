package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func startNetworkDamon() *sync.WaitGroup  {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connect: %v", err)
				continue
			}
			connectToService()
			fmt.Println(conn, "")
		}
	}()
	return &wg
}