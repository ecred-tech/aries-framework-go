/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

// NewRS256Signer creates a new RS256 signer with generated key.
func NewRS256Signer() (*RS256Signer, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return &RS256Signer{privateKey: privKey, PublicKey: &privKey.PublicKey}, nil
}

// GetRS256Signer creates a new RS256 signer with provided RSA private key.
func GetRS256Signer(privKey *rsa.PrivateKey) *RS256Signer {
	return &RS256Signer{privateKey: privKey, PublicKey: &privKey.PublicKey}
}

// RS256Signer makes RS256 based signatures.
type RS256Signer struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// Sign signs a message.
func (s RS256Signer) Sign(msg []byte) ([]byte, error) {
	hasher := crypto.SHA256.New()
	_, _ = hasher.Write(msg) //nolint:errcheck
	hashed := hasher.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, hashed)
}

// NewPS256Signer creates a new PS256 signer with generated key.
func NewPS256Signer() (*PS256Signer, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return &PS256Signer{privateKey: privKey, PublicKey: &privKey.PublicKey}, nil
}

// GetPS256Signer creates a new PS256 signer with provided RSA private key.
func GetPS256Signer(privKey *rsa.PrivateKey) *PS256Signer {
	return &PS256Signer{privateKey: privKey, PublicKey: &privKey.PublicKey}
}

// PS256Signer makes PS256 based signatures.
type PS256Signer struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// Sign signs a message.
func (s PS256Signer) Sign(msg []byte) ([]byte, error) {
	hasher := crypto.SHA256.New()

	_, _ = hasher.Write(msg) //nolint:errcheck

	hashed := hasher.Sum(nil)

	return rsa.SignPSS(rand.Reader, s.privateKey, crypto.SHA256, hashed, &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthEqualsHash,
	})
}
