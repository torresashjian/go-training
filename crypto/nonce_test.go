package crypto

import (
	"testing"
)

func TestGenerateNonce(t *testing.T) {
	nonce, err := GenerateNonce()
	if err != nil{
		t.Errorf("Error Generating Nonce, %s", err)
		t.Fail()
	}

	t.Logf("Nonce Equals: %v", string(nonce[:]))
}
