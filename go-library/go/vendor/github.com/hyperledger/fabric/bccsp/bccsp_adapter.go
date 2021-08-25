package bccsp

import (
	"crypto"
	"github.com/hyperledger/fabric/base"
	"hash"
)

/*
	存在的作用:
	所有接口都是用责任链来解决,通过opts来确定具体使用哪一个,除了opts外,所以都需要有一个基类,
*/

type SignerAdapter interface {
}

type IBaseKey interface {
	Bytes() ([]byte, error)
	SKI() []byte
	Symmetric() bool
	// // 获取最底层的私钥或者公钥,如ecdsa.PrivateKey等
	GetOrigin() interface{}
}

// FIXME ,名字不对,应该是IKeyAdapter 而非PrivateKey
// type IPrivateKey interface {
// 	IBaseKey
// 	GetPublicKey() (IPublicKey, error)
// 	ToCryptoSigner() crypto.Signer
// 	ToFabricKey() Key
// 	GetBaseType()base.BaseType
// 	GetDetailType()base.DetailType
// }
type IPublicKey interface {
	IBaseKey
	// ToFabricKey() Key
}

type KeyAdapter interface {
	Key
	GetOrigin() interface{}
	GetPublicKey() (IPublicKey, error)
	ToCryptoSigner() crypto.Signer
	GetBaseType() base.BaseType
	GetDetailType() base.DetailType
	ToPem() ([]byte, error)
}

func ConvtKey2CommonOpts(key KeyAdapter) base.CommonBaseOpts {
	return base.CommonBaseOpts{
		AlgorithmType: key.GetBaseType(),
		DetailType:    key.GetDetailType(),
		MinExcepted:   key.GetBaseType,
	}
}

type BCCSPAdapter interface {
	// KeyGen generates a key using opts.
	KeyGen(opts KeyGenOptsAdapter) (k KeyAdapter, err error)

	// KeyDeriv derives a key from k using opts.
	// The opts argument should be appropriate for the primitive used.
	KeyDeriv(k Key, opts KeyDerivOptsAdapter) (dk KeyAdapter, err error)

	// KeyImport imports a key from its raw representation using opts.
	// The opts argument should be appropriate for the primitive used.
	KeyImport(raw interface{}, opts KeyImportOptsAdapter) (k KeyAdapter, err error)

	// GetKey returns the key this CSP associates to
	// the Subject Key Identifier ski.
	GetKey(ski []byte) (k KeyAdapter, err error)

	// Hash hashes messages msg using options opts.
	// If opts is nil, the default hash function will be used.
	Hash(msg []byte, opts HashOptsAdapter) (hash []byte, err error)

	// GetHash returns and instance of hash.Hash using options opts.
	// If opts is nil, the default hash function will be returned.
	GetHash(opts HashOptsAdapter) (h hash.Hash, err error)

	// Sign signs digest using key k.
	// The opts argument should be appropriate for the algorithm used.
	//
	// Note that when a signature of a hash of a larger message is needed,
	// the caller is responsible for hashing the larger message and passing
	// the hash (as digest).
	Sign(k KeyAdapter, digest []byte, opts SignerOptsAdapter) (signature []byte, err error)

	// Verify verifies signature against key k and digest
	// The opts argument should be appropriate for the algorithm used.
	Verify(k KeyAdapter, signature, digest []byte, opts SignerOptsAdapter) (valid bool, err error)

	// Encrypt encrypts plaintext using key k.
	// The opts argument should be appropriate for the algorithm used.
	Encrypt(k KeyAdapter, plaintext []byte, opts EncrypterOptsAdapter) (ciphertext []byte, err error)

	// Decrypt decrypts ciphertext using key k.
	// The opts argument should be appropriate for the algorithm used.
	Decrypt(k KeyAdapter, ciphertext []byte, opts DecrypterOptsAdapter) (plaintext []byte, err error)
}
