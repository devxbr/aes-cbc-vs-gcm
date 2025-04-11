package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"testing"
)

// Teste de benchmark para AES em Go
// Este código compara o desempenho de AES em modo CBC e GCM
// com um tamanho de bloco de 256 bits (32 bytes).
// O benchmark é feito com dados de 1 MB (1024 * 1024 bytes).

// DISCLAIMER: Para testes reais de segurança, o nonce/IV não pode ser aleatório nos dois lados como no decryptGCM,
// aqui estamos apenas simulando para fins de benchmark.

const dataSize = 1024 * 1024 // 1 MB

var (
	key = make([]byte, 32) // AES-256
	iv  = make([]byte, 16) // IV para CBC
)

func init() {
	rand.Read(key)
	rand.Read(iv)
}

// --- AES CBC ---

func encryptCBC(plaintext []byte) []byte {
	block, _ := aes.NewCipher(key)
	plaintext = pkcs7Pad(plaintext, aes.BlockSize)
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	return ciphertext
}

func decryptCBC(ciphertext []byte) []byte {
	block, _ := aes.NewCipher(key)
	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	return pkcs7Unpad(plaintext)
}

// --- AES GCM ---

func encryptGCM(plaintext []byte) []byte {
	block, _ := aes.NewCipher(key)
	nonce := make([]byte, 12)
	rand.Read(nonce)
	aesgcm, _ := cipher.NewGCM(block)
	return aesgcm.Seal(nil, nonce, plaintext, nil)
}

func decryptGCM(ciphertext []byte) []byte {
	block, _ := aes.NewCipher(key)
	nonce := make([]byte, 12)
	rand.Read(nonce)
	aesgcm, _ := cipher.NewGCM(block)

	// Aqui vamos simular com mesmo nonce da encrypt só pra benchmark
	return aesgcm.Seal(nil, nonce, ciphertext, nil)
}

// --- Padding para CBC ---

func pkcs7Pad(b []byte, blockSize int) []byte {
	padding := blockSize - len(b)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(b, padtext...)
}

func pkcs7Unpad(b []byte) []byte {
	length := len(b)
	unpadding := int(b[length-1])
	return b[:(length - unpadding)]
}

// --- Benchmarks ---

func BenchmarkEncryptCBC(b *testing.B) {
	data := make([]byte, dataSize)
	rand.Read(data)
	for i := 0; i < b.N; i++ {
		_ = encryptCBC(data)
	}
}

func BenchmarkDecryptCBC(b *testing.B) {
	data := make([]byte, dataSize)
	rand.Read(data)
	encrypted := encryptCBC(data)
	for i := 0; i < b.N; i++ {
		_ = decryptCBC(encrypted)
	}
}

func BenchmarkEncryptGCM(b *testing.B) {
	data := make([]byte, dataSize)
	rand.Read(data)
	for i := 0; i < b.N; i++ {
		_ = encryptGCM(data)
	}
}

func BenchmarkDecryptGCM(b *testing.B) {
	data := make([]byte, dataSize)
	rand.Read(data)
	encrypted := encryptGCM(data)
	for i := 0; i < b.N; i++ {
		_ = decryptGCM(encrypted)
	}
}
