package main

import (
	b64 "encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		Help()
		return
	}
	option := os.Args[1]

	var data string
	if len(os.Args) > 2 {
		data = os.Args[2]
	}

	switch option {
	case "-e":
		Encode(data)
	case "-d":
		Decode(data)
	case "-h":
		Help()
	default:
		Help()
	}

}

func Encode(data string) {
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
}

func Decode(data string) {
	sDec, _ := b64.StdEncoding.DecodeString(data)
	fmt.Println(string(sDec))
}

func Help() {
	fmt.Println("b64 [option] [data]")
	fmt.Println("Options")
	fmt.Println("-e		encode")
	fmt.Println("-d		decode")
}
