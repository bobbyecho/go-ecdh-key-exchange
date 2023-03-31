package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
)

func AesGcmDecrypt(key []byte, nonce []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	sealer, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := sealer.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func AesGcmEncrypt(key []byte, payload []byte) ([]byte, error) {
	nonce := make([]byte, 12)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	sealer, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	//encodedNonce := base64.StdEncoding.EncodeToString(nonce)
	//encodedCipher := base64.StdEncoding.EncodeToString(sealer.Seal(nil, nonce, payload, nil))

	//return []byte(encodedNonce + "." + encodedCipher), nil

	hexNonce := hex.EncodeToString(nonce)
	hexCipher := hex.EncodeToString(sealer.Seal(nil, nonce, payload, nil))

	return []byte(hexNonce + "." + hexCipher), nil
}
