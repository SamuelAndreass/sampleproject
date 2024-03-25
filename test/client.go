package main

import (
	"fmt"
	"net"

	"github.com/SamuelAndreass/sampleproject/handle"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	handle.HandleErr(err)
	defer conn.Close()
	fmt.Println("koneksi tertutup")

	conn.Write([]byte("Hello server"))

}
