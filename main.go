package main

import (
	"fmt"

	"github.com/go-university/cryptit/decrypt"
	"github.com/go-university/cryptit/encrypt"
)

func main() {
	myMessage := "Hey this is a secret message!"
	encrypted := encrypt.Nimbus(myMessage)
	decrypted := decrypt.Nimbus(encrypted)
	fmt.Printf("Actual message: %v\nEncrypted: %v\nDecrypted: %v", myMessage, encrypted, decrypted)
}
