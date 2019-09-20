#!/bin/bash

rm -rf vendor
govendor init
govendor add +e

cp -r \
   "${GOPATH}/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1" \
   "${GOPATH}/src/myLibrary/go-libary/go/vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/"
