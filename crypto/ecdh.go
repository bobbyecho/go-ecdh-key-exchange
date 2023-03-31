package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/google/uuid"
	"sync"
)

var ecdhCurve ecdh.Curve

type EcdhKeyMap struct {
	sync.Mutex
	value map[string][]byte
}

var ecdhKeyMap EcdhKeyMap

func (c *EcdhKeyMap) Make() {
	c.value = make(map[string][]byte)
}

func (c *EcdhKeyMap) Add(id string, key []byte) {
	c.Lock()
	c.value[id] = key
	c.Unlock()
}

func (c *EcdhKeyMap) Get(id string) ([]byte, bool) {
	secretKey, ok := c.value[id]
	if !ok {
		return nil, false
	}

	return secretKey, true
}

func init() {
	ecdhCurve = ecdh.P256()
	ecdhKeyMap.Make()
}

func EcdhKeyExchange(peerPubKeyBytes []byte) (string, []byte, error) {
	peerPemBlock, _ := hex.DecodeString(string(peerPubKeyBytes))
	if peerPemBlock == nil {
		return "", nil, errors.New("invalid public key")
	}

	peerPubKey, err := ecdhCurve.NewPublicKey(peerPemBlock)
	if err != nil {
		return "", nil, err
	}

	privKey, err := ecdhCurve.GenerateKey(rand.Reader)
	if err != nil {
		return "", nil, err
	}

	pubKey := privKey.PublicKey()
	pubKeyBytes := pubKey.Bytes()

	sessionID := uuid.NewString()
	secret, err := privKey.ECDH(peerPubKey)
	if err != nil {
		return "", nil, err
	}

	ecdhKeyMap.Add(sessionID, secret)
	return sessionID, pubKeyBytes, nil
}

func EcdhGetSecretKey(sessionID string) ([]byte, error) {
	secretKey, found := ecdhKeyMap.Get(sessionID)
	if !found {
		return nil, errors.New("invalid session id")
	}

	return secretKey, nil
}
