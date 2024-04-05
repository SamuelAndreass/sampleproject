package main

import (
	"net"
	"time"

	"github.com/SamuelAndreass/sampleproject/handle"
)

func main() {
	// membuat koneksi baru paada localhost / ip kita dan pada port 1234
	net, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		print(err)
		return
	}
	handle.HandleErr(err)
	defer net.Close()

	// menghandle pesan dari client 
	// for loop infinit bertujuan agar bisa menerima beberapa pesan dari client
	for {
		conn, err := net.Accept()
		if err != nil {
			print(err)
			return
		}
	// go routine untuk membaca pesan
		go handleServer(conn)
	}

}

func handleServer(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		print(err)
		return
	}

	// setwritedeadline ditujuan ketika pesan gagal terkirim dengan rentang waktu 5 detik
	// dari waktu sekarang / ketika pesan coba untuk dikirim / write
	err = conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	if err != nil {
		if netErr, ok := err.(net.Error); netErr.Timeout() && ok {
			print("jaringan terputus")
		}
		handle.HandleErr(err)

	}
	// setelah pesan diterima dari client akan dikirim kembali ke client
	print("pesan dari client: ", string(buffer[:n]))
	_, err = conn.Write(buffer[:n])
	handle.HandleErr(err)

}
