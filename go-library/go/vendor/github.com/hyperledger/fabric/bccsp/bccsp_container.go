package bccsp

import (
	"crypto/x509"
	"encoding/base64"
	"errors"
	"github.com/cloudflare/cfssl/certdb"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/signer"
	"github.com/hyperledger/fabric/base"
	opts2 "github.com/hyperledger/fabric/bccsp/opts"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/debugutil"
	"math/big"
	"strings"
)

var (
	bccspContainer  *DefaultChainAbleBccspManager
	containerLogger = flogging.MustGetLogger("BccspContainer")
)

type IKeyGeneratorFunc func() (Key, error)

// 返回的是crypto的key
type IKeyImporterFunc func(raw, pwd []byte) (PrivateKeyWrapper, error)
type IPublicKeyImporterFunc func(raw interface{}) (KeyAdapter, error)
type IOriginPublicKeyImporterFunc func(bytes []byte) (interface{}, error)
type ISignerFunc func(prvKInterface interface{}, msg []byte) ([]byte, error)
type IByteSignerFunc func(prvKBytes []byte, msg []byte) ([]byte, error)
type IVerifierFunc func(pubKeyInterface interface{}, signature, msg []byte) (bool, error)
type ICertificateImporterFunc func(bytes []byte) (*x509.Certificate, base.CertificateBaseOpts, error)
type IUnmarshalFunc func(sig []byte) (r, s *big.Int, err error)
type IMarshalFunc func(r, s *big.Int) ([]byte, error)
type IHashFunc func(msg []byte) (hash []byte, err error)
type IPrivateKeyGeneratorFunc func() (KeyAdapter, error)
type IIdentitySerialize func()

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////// ca相关
// 生成csr
type ICSRGenerator func(adapter KeyAdapter, req *csr.CertificateRequest) ([]byte, error)

// 生成csr,并且生成对应的证书
type ICRTWithCsrGenerator func(key KeyAdapter, req *csr.CertificateRequest, CrtCsrWrapper CrtCsrWrapper) (cert []byte, csr []byte, err error)

// 解析csr
type ICSRParser func(raw interface{}) (*x509.CertificateRequest, error)

// 生成签发证书
type ISignCertGenerator func(keyAdapter KeyAdapter, rootCa *x509.Certificate, req signer.SignRequest, others ...interface{}) (*certdb.CertificateRecord, []byte, error)

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////

// FIXME 添加一个生成证书的opts,返回值最好也还是一个接口 ,兼容ecdsa和sm2
type ICertificateGenerator func(params base.ParamBuilder) (*x509.Certificate, []byte, error)

type DefaultDetailTemplateMediator struct {
	Handler base.IDetailTemplate
}

func NewDefaultDetailTemplateMediator(handler base.IDetailTemplate) *DefaultDetailTemplateMediator {
	return &DefaultDetailTemplateMediator{Handler: handler}
}

func (this *DefaultDetailTemplateMediator) String() string {
	return this.Handler.String()
}

func (this *DefaultDetailTemplateMediator) ValidIsMine(detailType base.DetailType) bool {
	return this.Handler.ValidIsMine(detailType)
}

func (this *DefaultDetailTemplateMediator) LinkLast(template base.IDetailTemplate) {
	this.Handler.LinkLast(template)
}

func (this *DefaultDetailTemplateMediator) GetNext() base.IDetailTemplate {
	return this.Handler.GetNext()
}

func (this *DefaultDetailTemplateMediator) SetNext(template base.IDetailTemplate) {
	this.Handler.SetNext(template)
}

type SignCertGeneratorHandler struct {
	*base.AlgorithmDetailTemplate
	ISignCertGenerator ISignCertGenerator
}
type CSRParserHandler struct {
	*base.AlgorithmDetailTemplate
	ICSRParser ICSRParser
}
type GenCRTWithCsrHandler struct {
	*base.AlgorithmDetailTemplate
	ICRTWithCsrGenerator ICRTWithCsrGenerator
}

type GenCSRHandler struct {
	*base.AlgorithmDetailTemplate
	ICSRGenerator ICSRGenerator
}

type CertificateGeneratorHandler struct {
	*base.AlgorithmDetailTemplate
	ICertificateGenerator ICertificateGenerator
}

type MarshalHandler struct {
	*base.AlgorithmDetailTemplate
	IMarshal IMarshalFunc
}

type UnmarshalHandler struct {
	*base.AlgorithmDetailTemplate
	IUnmarshal IUnmarshalFunc
}
type KeyImporterHandler struct {
	*base.AlgorithmDetailTemplate
	IKeyImporter IKeyImporterFunc
}

type KeyGeneratorHandler struct {
	*base.AlgorithmDetailTemplate
	IKeyGenerator IKeyGeneratorFunc
}

type SignerHandler struct {
	*base.AlgorithmDetailTemplate
	ISigner ISignerFunc
}
type VerifierHandler struct {
	*base.AlgorithmDetailTemplate
	IVerifier IVerifierFunc
}

type CertificateImporterHandler struct {
	*base.AlgorithmDetailTemplate
	ICertificateImporter ICertificateImporterFunc
}

type PublicKeyImportHandler struct {
	*base.AlgorithmDetailTemplate
	IPublicKeyImporter IPublicKeyImporterFunc
}

type OriginPublicKeyHandler struct {
	*base.AlgorithmDetailTemplate
	IOriginPublicKeyImporterFunc IOriginPublicKeyImporterFunc
}

type HashHandler struct {
	*base.AlgorithmDetailTemplate
	IHash IHashFunc
}
type PrivateKeyGeneratorHandler struct {
	*base.AlgorithmDetailTemplate
	IPrivateKeyGenerator IPrivateKeyGeneratorFunc
}

func init() {
	bccspContainer = new(DefaultChainAbleBccspManager)
}
func GenSignCert(keyAdapter KeyAdapter, rootCa *x509.Certificate, req signer.SignRequest, opts SignCertGenOpts, others ...interface{}) (*certdb.CertificateRecord, []byte, error) {
	return bccspContainer.GenSignCert(keyAdapter, rootCa, req, opts, others...)
}
func ParseCSR(raw interface{}, failStra base.FailStrategy, opts CSRParserOpts) (*x509.CertificateRequest, error) {
	return bccspContainer.ParseCSR(raw, failStra, opts)
}
func GenCRTWithCsr(adapter KeyAdapter, req *csr.CertificateRequest, opts CRTWithCsrGeneratorOpts, wrapper CrtCsrWrapper) ([]byte, []byte, error) {
	return bccspContainer.GenCRTWithCSR(adapter, req, opts, wrapper)
}
func GenCSR(adapter KeyAdapter, req *csr.CertificateRequest, opts CSRGenerateOpts) ([]byte, error) {
	return bccspContainer.GenCSR(adapter, req, opts)
}
func SignatureMarshal(r, s *big.Int, opts SignatureMarshalOptsAdapter) ([]byte, error) {
	return bccspContainer.SignatureMarshal(r, s, opts)
}

// 反序列化
func SignatureUnmarshal(sig []byte, opts SignatureUnMarshalOptsAdapter) (r, s *big.Int, e error) {
	return bccspContainer.SignatureUnmarshal(sig, opts)
}
func Sign(key Key, digest []byte, opts SignerOptsAdapter) (signature []byte, err error) {
	return bccspContainer.Sign(key, digest, opts)
}
func Verify(key Key, digest, signature []byte, failStrategy base.FailStrategy, opts VerifierOptsAdapter) error {
	verify := bccspContainer.Verify(key, digest, signature, failStrategy, opts)
	if nil != verify {
		msg := strings.Builder{}
		if bytes, e := key.Bytes(); nil != e {
			msg.WriteString("prvK:" + base64.StdEncoding.EncodeToString(bytes))
		}
		if publicKey, e := key.PublicKey(); nil != e {
			if bytes, e := publicKey.Bytes(); nil != e {
				msg.WriteString("pubK:" + base64.StdEncoding.EncodeToString(bytes))
			}
		}

		debugutil.ImportantLogging("校验失败:key:" + msg.String())
		return verify
	}
	return nil
}
func InterfaceKeyImport(raw, pwd []byte, failStrategy base.FailStrategy, opts KeyImportOptsAdapter) (PrivateKeyWrapper, error) {
	return bccspContainer.InterfaceKeyImport(raw, pwd, failStrategy, opts)
}

// 导入原生的fabric key
func PublicKeyImport(raw interface{}, failStrategy base.FailStrategy, opts PublicKeyImportOptsAdapter) (KeyAdapter, error) {
	return bccspContainer.PublicKeyImport(raw, failStrategy, opts)
}

// 导入原先的key,如ecdsa.Publickey 或者是sm2.PublicKey等
func OriginPublicKeyImport(raw []byte, failStrategy base.FailStrategy, opts OriginPublicKeyImportOptsAdapter) (interface{}, error) {
	return bccspContainer.OriginPublicKeyImport(raw, failStrategy, opts)
}
func ParseCertificate(raw []byte, failStrategy base.FailStrategy, opts CertificateImportOptsAdapter) (*x509.Certificate, base.CertificateBaseOpts, error) {
	return bccspContainer.ParseCertificate(raw, failStrategy, opts)
	// certificate, baseOpts, err := bccspContainer.ParseCertificate(raw, failStrategy, opts)
	// if nil != err {
	// 	return nil, nil, err
	// }
	// typeStr := reflect.TypeOf(certificate.PublicKey).String()
	//
	// debugutil.ImportantLogging("证书为:" + string(certificate.Raw) + ",证书中的公钥类型为:" + typeStr + ",opts中的类型:" + strconv.Itoa(int(baseOpts.GetBaseType()())))
	// return certificate,baseOpts,nil
}
func HashByBaseType(msg []byte, baseType base.BaseType) ([]byte, error) {
	var opts HashOptsAdapter
	if baseType.EqualsBaseType(base.SM2) {
		opts = opts2.DefaultSM3HashOptsImpl{}
	} else {
		opts = opts2.DefaultSHA256HashOptsImpl{}
	}
	return Hash(msg, opts)
}
func Hash(msg []byte, opts HashOptsAdapter) ([]byte, error) {
	return bccspContainer.Hash(msg, opts)
}
func GenPrivateKey(opts KeyGenOptsAdapter) (KeyAdapter, error) {
	return bccspContainer.GenPrivateKey(opts)
}
func GenCertificate(opts CertificateGeneratorOpts, params base.ParamBuilder) (*x509.Certificate, []byte, error) {
	return bccspContainer.GenCertificate(opts, params)
}

// FIXME 这里的直接使用 GetBaseType就行,私钥只分ecdsa和sm2,不用分具体的位数
func (this *DefaultDetailTemplateMediator) InterfaceKeyImport(raw, pwd []byte, failStrategy base.FailStrategy, opts KeyImportOptsAdapter) (PrivateKeyWrapper, error) {
	if failStrategy() {
		for tmp := this.Handler.(*KeyImporterHandler); nil != tmp; {
			if !tmp.ValidIsMine(opts.GetDetailType()()) {
				if nil == tmp.GetNext() {
					break
				}
				tmp = tmp.GetNext().(*KeyImporterHandler)
				continue
			}
			if k, e := tmp.IKeyImporter(raw, pwd); nil != e {
				return PrivateKeyWrapper{}, errors.New("KeyImport# 解析失败:" + e.Error())
			} else {
				return k, nil
			}
		}
		return PrivateKeyWrapper{}, errors.New("KeyImport找不到对应的handler")
	}
	for tmp := this.Handler.(*KeyImporterHandler); nil != tmp; tmp = tmp.GetNext().(*KeyImporterHandler) {
		if k, e := tmp.IKeyImporter(raw, pwd); nil != e {
			if tmp.GetNext() == nil {
				break
			}
			continue
		} else {
			return k, nil
		}
	}
	return PrivateKeyWrapper{}, errors.New("遍历也无法import key,失败")
}

func (this *DefaultDetailTemplateMediator) PublicKeyImport(raw interface{}, failStrategy base.FailStrategy, opts PublicKeyImportOptsAdapter) (KeyAdapter, error) {
	if failStrategy() {
		for tmp := this.Handler.(*PublicKeyImportHandler); nil != tmp; {
			if !tmp.ValidIsMine(opts.GetDetailType()()) {
				if tmp.GetNext() == nil {
					break
				}
				tmp = tmp.GetNext().(*PublicKeyImportHandler)
				continue
			}
			if k, e := tmp.IPublicKeyImporter(raw); nil != e {
				return nil, errors.New("KeyImport# 解析失败:" + e.Error())
			} else {
				return k, nil
			}
		}
	}

	for tmp := this.Handler.(*PublicKeyImportHandler); nil != tmp; tmp = tmp.GetNext().(*PublicKeyImportHandler) {
		if k, e := tmp.IPublicKeyImporter(raw); nil != e {
			if tmp.GetNext() == nil {
				break
			}
			continue
		} else {
			return k, nil
		}
	}
	return nil, errors.New("遍历也无法解析公钥")
}
func (this *DefaultDetailTemplateMediator) OriginPublicKeyImport(raw []byte, failStrategy base.FailStrategy, opts OriginPublicKeyImportOptsAdapter) (interface{}, error) {
	if failStrategy() {
		for tmp := this.Handler.(*OriginPublicKeyHandler); nil != tmp; {
			if !tmp.ValidIsMine(opts.GetDetailType()()) {
				if tmp.GetNext() == nil {
					break
				}
				tmp = tmp.GetNext().(*OriginPublicKeyHandler)
				continue
			}
			if k, e := tmp.IOriginPublicKeyImporterFunc(raw); nil != e {
				return nil, errors.New("KeyImport# 解析失败:" + e.Error())
			} else {
				return k, nil
			}
		}
		return nil, errors.New("OriginPublicKeyImport#无法找到对应的handler")
	}
	for tmp := this.Handler.(*OriginPublicKeyHandler); nil != tmp; tmp = tmp.GetNext().(*OriginPublicKeyHandler) {
		if k, e := tmp.IOriginPublicKeyImporterFunc(raw); nil != e {
			if tmp.GetNext() == nil {
				break
			}
			continue
		} else {
			return k, nil
		}
	}
	return nil, errors.New("OriginPublicKeyImport#遍历也无法解析公钥")
}
func (this *DefaultDetailTemplateMediator) ParseCertificate(bytes []byte, failStrategy base.FailStrategy, opts CertificateImportOptsAdapter) (*x509.Certificate, base.CertificateBaseOpts, error) {
	if failStrategy() {
		for tmp := this.Handler.(*CertificateImporterHandler); nil != tmp; {
			if !tmp.ValidIsMine(opts.GetDetailType()()) {
				if tmp.GetNext() == nil {
					break
				}
				tmp = tmp.GetNext().(*CertificateImporterHandler)
				continue
			}
			if certificate, crtOpts, e := tmp.ICertificateImporter(bytes); nil != e {
				return nil, nil, errors.New("ParseCertificate# 解析失败:" + e.Error())
			} else {
				return certificate, crtOpts, nil
			}
		}
		return nil, nil, errors.New("ParseCertificate#无法找到对应的handler")
	}

	for tmp := this.Handler.(*CertificateImporterHandler); nil != tmp; tmp = tmp.GetNext().(*CertificateImporterHandler) {
		if certificate, crtOpts, e := tmp.ICertificateImporter(bytes); nil != e {
			if nil == tmp.GetNext() {
				return nil, nil, e
			}
			continue
		} else {
			return certificate, crtOpts, nil
		}
	}
	return nil, nil, errors.New("ParseCertificate#遍历也无法解析公钥")
}
func (this *DefaultDetailTemplateMediator) SignatureMarshal(r *big.Int, s *big.Int, opts SignatureMarshalOptsAdapter) ([]byte, error) {
	for tmp := this.Handler.(*MarshalHandler); nil != tmp && tmp.ValidIsMine(opts.GetDetailType()()); {
		if !tmp.ValidIsMine(opts.GetDetailType()()) {
			if tmp.GetNext() == nil {
				break
			}
			tmp = tmp.GetNext().(*MarshalHandler)
		}
		if res, e := tmp.IMarshal(r, s); nil != e {
			return nil, errors.New("SignatureMarshal失败:" + e.Error())
		} else {
			return res, nil
		}
	}
	return nil, errors.New("SignatureMarshal#无法找到对应的handler")
}
func (this *DefaultDetailTemplateMediator) SignatureUnmarshal(sig []byte, opts SignatureUnMarshalOptsAdapter) (*big.Int, *big.Int, error) {
	for tmp := this.Handler.(*UnmarshalHandler); nil != tmp; {
		if !tmp.ValidIsMine(opts.GetDetailType()()) {
			if tmp.GetNext() == nil {
				break
			}
			tmp = tmp.GetNext().(*UnmarshalHandler)
		}
		if r, s, e := tmp.IUnmarshal(sig); nil != e {
			return nil, nil, errors.New("SignatureUnmarshal失败:" + e.Error())
		} else {
			return r, s, nil
		}
	}
	return nil, nil, errors.New("SignatureUnmarshal#无法找到对应的handler")
}

// FIXME ,sign 和verify 需要做特殊处理,先通过baseType过滤,然后再通过detailType过滤
func (this *DefaultDetailTemplateMediator) Sign(key Key, digest []byte, opts SignerOptsAdapter) (signature []byte, err error) {
	for tmp := this.Handler.(*SignerHandler); nil != tmp; {
		if !tmp.ValidIsMine(opts.GetDetailType()()) {
			if tmp.GetNext() == nil {
				break
			}
			tmp = tmp.GetNext().(*SignerHandler)
		}
		if res, e := tmp.ISigner(key, digest); nil != e {
			return nil, errors.New("签名失败:" + e.Error())
		} else {
			return res, nil
		}
	}
	return nil, errors.New("Sign#无法找到对应的handler")
}

func (this *DefaultDetailTemplateMediator) Verify(key Key, digest, signature []byte, failStrategy base.FailStrategy, opts VerifierOptsAdapter) error {
	if failStrategy() {
		for tmp := this.Handler.(*VerifierHandler); nil != tmp; {
			if !tmp.ValidIsMine(opts.GetDetailType()()) {
				if tmp.GetNext() == nil {
					break
				}
				tmp = tmp.GetNext().(*VerifierHandler)
				continue
			}
			if k, e := tmp.IVerifier(key, signature, digest); nil != e {
				return errors.New("KeyImport# 解析失败:" + e.Error())
			} else if k {
				return nil
			} else {
				goto loopVerify
			}
		}
		return errors.New("Verify#无法找到对应的handler")
	}
	goto loopVerify

loopVerify:
	for tmp := this.Handler.(*VerifierHandler); nil != tmp; tmp = tmp.GetNext().(*VerifierHandler) {
		if v, e := tmp.IVerifier(key, signature, digest); nil != e || !v {
			if tmp.GetNext() == nil {
				break
			}
			continue
		} else if v {
			return nil
		}
	}
	return errors.New("Verify#遍历也无法通过校验")
}

func (this *DefaultDetailTemplateMediator) Hash(msg []byte, opts HashOptsAdapter) ([]byte, error) {
	for tmp := this.Handler.(*HashHandler); nil != tmp; {
		if !tmp.ValidIsMine(opts.GetDetailType()()) {
			if tmp.GetNext() == nil {
				break
			}
			tmp = tmp.GetNext().(*HashHandler)
		}
		if v, e := tmp.IHash(msg); nil != e {
			return nil, errors.New("hash失败:" + e.Error())
		} else {
			return v, nil
		}
	}
	return nil, errors.New("Hash#遍历也无法解析公钥")
}
func (this *DefaultDetailTemplateMediator) KeyGen(opts KeyGenOptsAdapter) (KeyAdapter, error) {
	for tmp := this.Handler.(*PrivateKeyGeneratorHandler); nil != tmp; {
		if !tmp.ValidIsMine(opts.GetDetailType()()) {
			if tmp.GetNext() == nil {
				break
			}
			tmp = tmp.GetNext().(*PrivateKeyGeneratorHandler)
		}
		if v, e := tmp.IPrivateKeyGenerator(); nil != e {
			return v, errors.New("keyGen失败:" + e.Error())
		} else {
			return v, nil
		}
	}
	return nil, errors.New("KeyGen#遍历也无法解析公钥")
}
func (this *DefaultDetailTemplateMediator) CertificateGen(opts CertificateGeneratorOpts, params base.ParamBuilder) (*x509.Certificate, []byte, error) {
	for tmp := this.Handler.(*CertificateGeneratorHandler); nil != tmp; {
		if !tmp.ValidIsMine(opts.GetDetailType()()) {
			if tmp.GetNext() == nil {
				break
			}
			tmp = tmp.GetNext().(*CertificateGeneratorHandler)
		}
		if crt, bs, e := tmp.ICertificateGenerator(params); nil != e {
			return nil, nil, e
		} else {
			return crt, bs, nil
		}
	}
	return nil, nil, errors.New("CertificateGen#遍历也无法解析公钥")
}

func (this *DefaultDetailTemplateMediator) GenCSR(adapter KeyAdapter, req *csr.CertificateRequest, opts CSRGenerateOpts) ([]byte, error) {
	for tmp := this.Handler.(*GenCSRHandler); nil != tmp; {
		if !tmp.ValidIsMine(opts.GetDetailType()()) {
			if tmp.GetNext() == nil {
				break
			}
			tmp = tmp.GetNext().(*GenCSRHandler)
		}
		if bs, e := tmp.ICSRGenerator(adapter, req); nil != e {
			return nil, e
		} else {
			return bs, nil
		}
	}
	return nil, errors.New("GenCSR#无法生成csr")
}

func (this *DefaultDetailTemplateMediator) GenCRTWithCSR(adapter KeyAdapter, req *csr.CertificateRequest, opts CRTWithCsrGeneratorOpts, wrapper CrtCsrWrapper) ([]byte, []byte, error) {
	for tmp := this.Handler.(*GenCRTWithCsrHandler); nil != tmp; {
		if !tmp.ValidIsMine(opts.GetDetailType()()) {
			if tmp.GetNext() == nil {
				break
			}
			tmp = tmp.GetNext().(*GenCRTWithCsrHandler)
		}
		if crt, csr, e := tmp.ICRTWithCsrGenerator(adapter, req, wrapper); nil != e {
			return nil, nil, e
		} else {
			return crt, csr, nil
		}
	}
	return nil, nil, errors.New("GenCRTWithCSR失败")
}

func (this *DefaultDetailTemplateMediator) ParseCSR(raw interface{}, failStrategy base.FailStrategy, opts CSRParserOpts) (*x509.CertificateRequest, error) {
	if failStrategy() {
		for tmp := this.Handler.(*CSRParserHandler); nil != tmp; {
			if !tmp.ValidIsMine(opts.GetDetailType()()) {
				if nil == tmp.GetNext() {
					break
				}
				tmp = tmp.GetNext().(*CSRParserHandler)
				continue
			}
			if k, e := tmp.ICSRParser(raw); nil != e {
				return nil, errors.New("parsecsr# 解析失败:" + e.Error())
			} else {
				return k, nil
			}
		}
		return nil, errors.New("parseCsr找不到handler")
	}
	for tmp := this.Handler.(*CSRParserHandler); nil != tmp; tmp = tmp.GetNext().(*CSRParserHandler) {
		if k, e := tmp.ICSRParser(raw); nil != e {
			if tmp.GetNext() == nil {
				break
			}
			continue
		} else {
			return k, nil
		}
	}
	return nil, errors.New("遍历也无法parseCSSSR,失败")
}

func (this *DefaultDetailTemplateMediator) GenSignCert(adapter KeyAdapter, ca *x509.Certificate, req signer.SignRequest, opts SignCertGenOpts, others ...interface{}) (*certdb.CertificateRecord, []byte, error) {
	for tmp := this.Handler.(*SignCertGeneratorHandler); nil != tmp; {
		if !tmp.ValidIsMine(opts.GetDetailType()()) {
			if tmp.GetNext() == nil {
				break
			}
			tmp = tmp.GetNext().(*SignCertGeneratorHandler)
		}
		if crt, bs, e := tmp.ISignCertGenerator(adapter, ca, req, others...); nil != e {
			return nil, nil, e
		} else {
			return crt, bs, nil
		}
	}
	return nil, nil, errors.New("GenCRTWithCSR失败")
}
