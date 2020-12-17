package constants

type Algorithm byte

const GM_SIGNATURE_STRING = "1.2.156.10197.1.501"
const GM_SIGNATURE_SM2WITHSM3="SM2-SM3"

var (
	algorithDescMap = map[Algorithm]string{
		ECDSA: "ecdsa",
		SM2:   "sm2",
		RSA:   "rsa",
		SHA:   "sha",
	}
)
// base 算法
const (
	ECDSA Algorithm = 1 << 0
	SM2   Algorithm = 1 << 1
	RSA   Algorithm = 1 << 2
	AES   Algorithm = 1 << 3
	SHA   Algorithm = 1 << 4
	NONE  Algorithm = 1 << 7
)
// hash算法
const (
	HASH_SHA256 Algorithm=100
	HASH_SM3_WITH_NO Algorithm=101
)

// 前开后闭
const (
	// 15个字节为magicWord
	MAGIC_WORD="magicWord"
	// magicWord起始下标
	EXTENSION_MAGICWORD_START_INDEX Algorithm=0
	// magicWord截止下标
	EXTENSION_MAGICWORD_END_INDEX Algorithm=10
	// 算法下标
	EXTENSION_ALGORITHM_BASE_INDEX Algorithm=10

	EXTENSION_MIN_LENGTH Algorithm=12
)

func GetAlgorithmDescription(alg Algorithm) string {
	return algorithDescMap[alg]
}
