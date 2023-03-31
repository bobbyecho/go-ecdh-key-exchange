package utils

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"key-exchange/crypto"
	"strings"
)

func EncryptPayload(sessionID string, payload any) []byte {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil
	}

	secretKey, err := crypto.EcdhGetSecretKey(sessionID)
	if err != nil {
		return nil
	}

	cipherText, err := crypto.AesGcmEncrypt(secretKey, payloadBytes)
	if err != nil {
		return nil
	}

	fmt.Printf("(encrypt) encrypted payload: %s\n", cipherText)

	return cipherText
}

func DecryptPayload(sessionID string, payload string) []byte {
	fmt.Printf("(decrypt - before) encrypted payload: %s\n", payload)

	splitResponse := strings.Split(payload, ".")

	ivSplit := splitResponse[0]
	cipherSplit := splitResponse[1]

	ivPart, _ := hex.DecodeString(ivSplit)
	cipherTextPart, _ := hex.DecodeString(cipherSplit)

	secretKey, err := crypto.EcdhGetSecretKey(sessionID)
	if err != nil {
		return nil
	}

	plainText, err := crypto.AesGcmDecrypt(secretKey, ivPart, cipherTextPart)
	if err != nil {
		return nil
	}

	fmt.Printf("(decrypt - after) decrypted: %s\n", plainText)

	return plainText
}
