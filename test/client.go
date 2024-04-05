package main

import (
	"fmt"
	"net"
	"time"

	"github.com/SamuelAndreass/sampleproject/handle"
)

func main() {
	// untuk menkoneksikan jaringan ke server dengan local ip di port 1234
	// sama seperti net.Dial, net.DialTimeout menerima timeout artinya gagal terhubung
	// pada codingan ini menerima timeout ketika sudah 5 detik 
	// dan tidak terkoneksi jaringan manapun
	conn, err := net.DialTimeout("tcp", "localhost:1234", time.Second * 5)
	handle.HandleErr(err)
	defer conn.Close()
	fmt.Println("koneksi tertutup")

	// mengrim pesan ke server dalam bentuk array of byte
	conn.Write([]byte("Hello server"))

	// SetReadDeadline ditujuan ketika pesan gagal diterima dengan rentang waktu 5 detik
	// dari waktu sekarang / ketika pesan coba untuk dibaca / read akan menerima eror timeout
	err = conn.SetReadDeadline(time.Now().Add(time.Second* 5))
	if err != nil {
		if netErr, ok :=  err.(net.Error); netErr.Timeout() && ok {
			print("jaringan terputus")
		} 
		handle.HandleErr(err)
	}
	// menampung pesan dengan buffer dan print hasil pesan tersebut
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	handle.HandleErr(err)
	print("pesan balasan dari server: ", string(buffer[:n]))


}
