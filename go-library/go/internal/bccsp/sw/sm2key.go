/*
Copyright Suzhou Tongji Fintech Research Institute 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package sw

import (
	"crypto"
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"myLibrary/go-library/go/internal/bccsp"

	"github.com/tjfoc/gmsm/sm2"
)

type sm2PrivateKey struct {
	privKey *sm2.PrivateKey
}

func (k *sm2PrivateKey) GetOrigin() interface{} {
	return k.privKey
}

func (k *sm2PrivateKey) ToCryptoSigner() crypto.Signer {
	return k.privKey
}

func (k *sm2PrivateKey) GetPublicKey() (bccsp.IPublicKey, error) {
	return &sm2PublicKey{&k.privKey.PublicKey}, nil
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *sm2PrivateKey) Bytes() ([]byte, error) {
	return k.privKey.D.Bytes(),nil
}

// SKI returns the subject key identifier of this key.
func (k *sm2PrivateKey) SKI() []byte {
	if k.privKey == nil {
		return nil
	}

	// Marshall the public key
	raw := elliptic.Marshal(k.privKey.Curve, k.privKey.PublicKey.X, k.privKey.PublicKey.Y)

	// Hash it
	// TODO 调整SKI计算
	//hash := sm3.New()
	hash := sha256.New()
	hash.Write(raw)
	return hash.Sum(nil)
}

// Symmetric returns true if this key is a symmetric key,
// false if this key is asymmetric
func (k *sm2PrivateKey) Symmetric() bool {
	return false
}

// Private returns true if this key is a private key,
// false otherwise.
func (k *sm2PrivateKey) Private() bool {
	return true
}

// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// This method returns an error in symmetric key schemes.
func (k *sm2PrivateKey) PublicKey() (bccsp.Key, error) {
	return &sm2PublicKey{&k.privKey.PublicKey}, nil
}

type sm2PublicKey struct {
	pubKey *sm2.PublicKey
}

func (k *sm2PublicKey) GetOrigin() interface{} {
	return k.pubKey
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *sm2PublicKey) Bytes() (raw []byte, err error) {
	raw, err = sm2.MarshalPKIXPublicKey(k.pubKey)
	if err != nil {
		return nil, fmt.Errorf("Failed marshalling key [%s]", err)
	}
	return
}

// SKI returns the subject key identifier of this key.
func (k *sm2PublicKey) SKI() []byte {
	if k.pubKey == nil {
		return nil
	}

	// Marshall the public key
	raw := elliptic.Marshal(k.pubKey.Curve, k.pubKey.X, k.pubKey.Y)

	// Hash it
	// TODO 计算hash 修改为SM3
	hash := sha256.New()
	hash.Write(raw)
	return hash.Sum(nil)
}

// Symmetric returns true if this key is a symmetric key,
// false if this key is asymmetric
func (k *sm2PublicKey) Symmetric() bool {
	return false
}

// Private returns true if this key is a private key,
// false otherwise.
func (k *sm2PublicKey) Private() bool {
	return false
}

// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// This method returns an error in symmetric key schemes.
func (k *sm2PublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}
