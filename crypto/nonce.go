package crypto

import (
	"io"
	"crypto/rand"
)

const(
	NoneSize = 24
)

// GenerateNonce generates a random nonce
func GenerateNonce() (*[NoneSize]byte, error) {

	nonce := new([NoneSize]byte)
	_, err := io.ReadFull(rand.Reader, nonce[:])
	if err != nil{
		return nil, err
	}

	return nonce, nil
}
