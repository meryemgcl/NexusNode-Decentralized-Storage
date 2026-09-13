package crypto_test

import (
	"bytes"
	"testing"

	"github.com/meryemgcl/NexusNode-Decentralized-Storage/crypto"
)

func TestEncryptDecryptRoundtrip(t *testing.T) {
	key := []byte("12345678901234567890123456789012") // 32 bytes
	plaintext := []byte("NexusNode AES-256 GCM test data!")

	ciphertext, err := crypto.Encrypt(key, plaintext)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	if bytes.Equal(ciphertext, plaintext) {
		t.Fatal("ciphertext should differ from plaintext")
	}

	recovered, err := crypto.Decrypt(key, ciphertext)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if !bytes.Equal(recovered, plaintext) {
		t.Fatalf("round-trip mismatch: got %q, want %q", recovered, plaintext)
	}
}

func TestEncryptRequires32ByteKey(t *testing.T) {
	_, err := crypto.Encrypt([]byte("short"), []byte("data"))
	if err == nil {
		t.Fatal("expected error for short key")
	}
}

func TestDecryptWithWrongKeyFails(t *testing.T) {
	key1 := []byte("12345678901234567890123456789012")
	key2 := []byte("99999999999999999999999999999999")

	ciphertext, err := crypto.Encrypt(key1, []byte("secret"))
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	_, err = crypto.Decrypt(key2, ciphertext)
	if err == nil {
		t.Fatal("expected error when decrypting with wrong key")
	}
}

func TestNonceIsRandomized(t *testing.T) {
	key := []byte("12345678901234567890123456789012")
	data := []byte("same data")

	ct1, _ := crypto.Encrypt(key, data)
	ct2, _ := crypto.Encrypt(key, data)

	if bytes.Equal(ct1, ct2) {
		t.Fatal("expected different ciphertexts due to random nonce")
	}
}
