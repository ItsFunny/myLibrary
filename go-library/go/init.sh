#!/bin/bash

rm -rf vendor
#govendor init
#govendor add +e
go mod vendor

cp -r \
   "${GOPATH}/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1" \
   "${GOPATH}/src/myLibrary/go-library/go/vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/"
