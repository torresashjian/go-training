package crypto

import (
	"encoding/pem"
	"fmt"
	"crypto/x509"
)


// DecodeCertificate will return a decoded certificate from a pem file
func DecodePemCertificate(certStr string) error{
	block, _ := pem.Decode([]byte(certStr))
	if block == nil{
		return fmt.Errorf("Invalid certificate format")
	}

	fmt.Printf("BlockType of certificate: '%s'\n", block.Type)

	fmt.Printf("BlockHeaders of certificate: '%+v'\n", block.Headers)

	// Parse certificate

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}

	fmt.Printf("Subject of certificate: '%+v'\n", cert.Subject)
	fmt.Printf("Issuer of certificate: '%+v'\n", cert.Issuer)

	return nil
}