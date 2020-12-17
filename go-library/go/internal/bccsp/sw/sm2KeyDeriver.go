/*
Copyright IBM Corp. 2016 All Rights Reserved.

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
	"errors"
	"fmt"
	"myLibrary/go-library/go/internal/bccsp"

	"github.com/tjfoc/gmsm/sm2"
	"math/big"
)

type sm2PublicKeyKeyDeriver struct{}

func (kd *sm2PublicKeyKeyDeriver) KeyDeriv(k bccsp.Key, opts bccsp.KeyDerivOpts) (bccsp.Key, error) {
	// Validate opts
	if opts == nil {
		return nil, errors.New("Invalid opts parameter. It must not be nil.")
	}

	pubKey := k.(*sm2PublicKey)

	switch opts.(type) {
	// Re-randomized an ECDSA private key
	case *bccsp.SM2ReRandKeyOpts:
		reRandOpts := opts.(*bccsp.SM2ReRandKeyOpts)
		tempSK := &sm2.PublicKey{
			Curve: pubKey.pubKey.Curve,
			X:     new(big.Int),
			Y:     new(big.Int),
		}

		var k = new(big.Int).SetBytes(reRandOpts.ExpansionValue())
		var one = new(big.Int).SetInt64(1)
		n := new(big.Int).Sub(pubKey.pubKey.Params().N, one)
		k.Mod(k, n)
		k.Add(k, one)

		// Compute temporary public key
		tempX, tempY := pubKey.pubKey.ScalarBaseMult(k.Bytes())
		tempSK.X, tempSK.Y = tempSK.Add(
			pubKey.pubKey.X, pubKey.pubKey.Y,
			tempX, tempY,
		)

		// Verify temporary public key is a valid point on the reference curve
		isOn := tempSK.Curve.IsOnCurve(tempSK.X, tempSK.Y)
		if !isOn {
			return nil, errors.New("Failed temporary public key IsOnCurve check.")
		}

		return &sm2PublicKey{tempSK}, nil
	default:
		return nil, fmt.Errorf("Unsupported 'KeyDerivOpts' provided [%v]", opts)
	}
}

type sm2PrivateKeyKeyDeriver struct{}

func (kd *sm2PrivateKeyKeyDeriver) KeyDeriv(k bccsp.Key, opts bccsp.KeyDerivOpts) (bccsp.Key, error) {
	// Validate opts
	if opts == nil {
		return nil, errors.New("Invalid opts parameter. It must not be nil.")
	}

	sm2K := k.(*sm2PrivateKey)

	switch opts.(type) {
	// Re-randomized an ECDSA private key
	case *bccsp.SM2ReRandKeyOpts:
		reRandOpts := opts.(*bccsp.SM2ReRandKeyOpts)
		tempSK := &sm2.PrivateKey{
			PublicKey: sm2.PublicKey{
				Curve: sm2K.privKey.Curve,
				X:     new(big.Int),
				Y:     new(big.Int),
			},
			D: new(big.Int),
		}

		var k = new(big.Int).SetBytes(reRandOpts.ExpansionValue())
		var one = new(big.Int).SetInt64(1)
		n := new(big.Int).Sub(sm2K.privKey.Params().N, one)
		k.Mod(k, n)
		k.Add(k, one)

		tempSK.D.Add(sm2K.privKey.D, k)
		tempSK.D.Mod(tempSK.D, sm2K.privKey.PublicKey.Params().N)

		// Compute temporary public key
		tempX, tempY := sm2K.privKey.PublicKey.ScalarBaseMult(k.Bytes())
		tempSK.PublicKey.X, tempSK.PublicKey.Y =
			tempSK.PublicKey.Add(
				sm2K.privKey.PublicKey.X, sm2K.privKey.PublicKey.Y,
				tempX, tempY,
			)

		// Verify temporary public key is a valid point on the reference curve
		isOn := tempSK.Curve.IsOnCurve(tempSK.PublicKey.X, tempSK.PublicKey.Y)
		if !isOn {
			return nil, errors.New("Failed temporary public key IsOnCurve check.")
		}

		return &sm2PrivateKey{tempSK}, nil
	default:
		return nil, fmt.Errorf("Unsupported 'KeyDerivOpts' provided [%v]", opts)
	}
}
