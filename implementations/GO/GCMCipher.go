// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	key := make([]byte, 16)
	block, err := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	fmt.Println(key)
	fmt.Println(err)
	data := []byte("this is just an example anas")
	ciphertext := gcm.Seal(nil, nonce, data, nil)
	fmt.Printf("ciphertext: %x\n", ciphertext)
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	fmt.Printf("plaintext: %s\n", plaintext)
	fmt.Println(gcm.NonceSize())
}
