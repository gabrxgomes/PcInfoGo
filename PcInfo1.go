package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"time"
)

func getLocalAddress() net.IP { //funcao pra pegar o ip local
	con, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		log.Fatal(error)
	}
	defer con.Close()

	localAddress := con.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}

func getossistem() {
	os := runtime.GOOS //para chamar o valor recebido por os temos que colocar runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Windows")
	case "darwin":
		fmt.Println("MAC operating system")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}

//---------------------------------------------------------------------------------------------
//função principal

func main() {
	currentTime := time.Now()

	ipString := getLocalAddress()
	//fmt.Println("Hello World!")
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Hostname: ", hostname)

	fmt.Println(" ")

	fmt.Println("Yor ip Address: ", ipString)

	fmt.Println(" ")

	fmt.Println("Yout Os sistem is: ", runtime.GOOS)

	fmt.Println(" ")

	fmt.Println("Current Time is: ")

	fmt.Printf("%d-%d-%d %d:%d:%d\n",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second())
}
