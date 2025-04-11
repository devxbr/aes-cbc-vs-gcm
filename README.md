# AES CBC vs GCM Benchmark

Este projeto contém benchmarks para comparar o desempenho dos modos de operação CBC (Cipher Block Chaining) e GCM (Galois/Counter Mode) do algoritmo AES (Advanced Encryption Standard) em Go. O objetivo é medir o tempo necessário para criptografar e descriptografar dados de 1 MB usando AES-256.

## Estrutura do Código

### Constantes e Variáveis Globais

- **`dataSize`**: Define o tamanho dos dados a serem processados (1 MB).
- **`key`**: Chave de 256 bits (32 bytes) gerada aleatoriamente para AES-256.
- **`iv`**: Vetor de inicialização (IV) de 16 bytes para o modo CBC.

### Funções

#### AES CBC

- **`encryptCBC(plaintext []byte) []byte`**  
  Criptografa os dados no modo CBC.  
  - Aplica padding PKCS#7 para alinhar o tamanho do bloco.
  - Usa o vetor de inicialização `iv`.

- **`decryptCBC(ciphertext []byte) []byte`**  
  Descriptografa os dados no modo CBC.  
  - Remove o padding PKCS#7 após a descriptografia.

#### AES GCM

- **`encryptGCM(plaintext []byte) []byte`**  
  Criptografa os dados no modo GCM.  
  - Gera um nonce de 12 bytes aleatório para cada operação.

- **`decryptGCM(ciphertext []byte) []byte`**  
  Simula a descriptografia no modo GCM.  
  - Usa o mesmo nonce gerado na criptografia para fins de benchmark.

#### Padding para CBC

- **`pkcs7Pad(b []byte, blockSize int) []byte`**  
  Aplica padding PKCS#7 para alinhar o tamanho do bloco.

- **`pkcs7Unpad(b []byte) []byte`**  
  Remove o padding PKCS#7 após a descriptografia.

### Benchmarks

- **`BenchmarkEncryptCBC(b *testing.B)`**  
  Mede o tempo necessário para criptografar dados de 1 MB no modo CBC.

- **`BenchmarkDecryptCBC(b *testing.B)`**  
  Mede o tempo necessário para descriptografar dados de 1 MB no modo CBC.

- **`BenchmarkEncryptGCM(b *testing.B)`**  
  Mede o tempo necessário para criptografar dados de 1 MB no modo GCM.

- **`BenchmarkDecryptGCM(b *testing.B)`**  
  Mede o tempo necessário para descriptografar dados de 1 MB no modo GCM.

## Observações

- **DISCLAIMER**: Este código não é seguro para uso em produção. O nonce/IV deve ser gerado e armazenado corretamente para garantir a segurança. Aqui, o nonce é gerado aleatoriamente em ambos os lados (criptografia e descriptografia) apenas para fins de benchmark.
- O tamanho dos dados processados é fixado em 1 MB para garantir consistência nos testes.

## Como Executar os Benchmarks

1. Certifique-se de ter o Go instalado.
2. Execute o comando abaixo no terminal para rodar os benchmarks:

```sh
go test -bench=.