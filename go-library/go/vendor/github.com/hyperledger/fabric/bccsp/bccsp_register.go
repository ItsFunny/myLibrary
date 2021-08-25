/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-09 16:18
# @File : bccsp_register.go
# @Description :
# @Attention :
*/
package bccsp

import (
	"github.com/hyperledger/fabric/base"
)

func RegisterSignCertGenerator(algInterface base.BaseType, detailType base.DetailType, im ISignCertGenerator) {
	alg := algInterface
	newTemplate := NewDefaultDetailTemplateMediator(&SignCertGeneratorHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		ISignCertGenerator:      im,
	})
	if bccspContainer.ISignCertGenerator == nil {
		bccspContainer.ISignCertGenerator = NewDefaultMediator(alg)
	}
	bccspContainer.ISignCertGenerator.LinkHiera(alg, newTemplate)
}
func RegisterCsrParser(algInterface base.BaseType, detailType base.DetailType, im ICSRParser) {
	alg := algInterface
	newTemplate := NewDefaultDetailTemplateMediator(&CSRParserHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		ICSRParser:              im,
	})
	if bccspContainer.ICSRParser == nil {
		bccspContainer.ICSRParser = NewDefaultMediator(alg)
	}
	bccspContainer.ICSRParser.LinkHiera(alg, newTemplate)
}
func RegisterCreateCrtWithCsrGenerator(algInterface base.BaseType, detailType base.DetailType, im ICRTWithCsrGenerator) {
	alg := algInterface
	newTemplate := NewDefaultDetailTemplateMediator(&GenCRTWithCsrHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		ICRTWithCsrGenerator:    im,
	})
	if bccspContainer.ICRTWithCsrGenerator == nil {
		bccspContainer.ICRTWithCsrGenerator = NewDefaultMediator(alg)
	}
	bccspContainer.ICRTWithCsrGenerator.LinkHiera(alg, newTemplate)
}
func RegisterCreateCsrGenerator(algInterface base.BaseType, detailType base.DetailType, im ICSRGenerator) {
	alg := algInterface
	newTemplate := NewDefaultDetailTemplateMediator(&GenCSRHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		ICSRGenerator:           im,
	})
	if bccspContainer.ICSRGenerator == nil {
		bccspContainer.ICSRGenerator = NewDefaultMediator(alg)
	}
	bccspContainer.ICSRGenerator.LinkHiera(alg, newTemplate)
}

func RegisterCertificateGenerator(algInterface base.BaseType, detailType base.DetailType, im ICertificateGenerator) {
	alg := base.BaseType(algInterface)
	newTemplate := NewDefaultDetailTemplateMediator(&CertificateGeneratorHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		ICertificateGenerator:   im,
	})
	if bccspContainer.ICertificateGenerator == nil {
		bccspContainer.ICertificateGenerator = NewDefaultMediator(alg)
	}
	bccspContainer.ICertificateGenerator.LinkHiera(alg, newTemplate)
}
func RegisterPrivateKeyGenerator(algInterface base.BaseType, detailType base.DetailType, importer IPrivateKeyGeneratorFunc) {
	alg := base.BaseType(algInterface)
	newTemplate := NewDefaultDetailTemplateMediator(&PrivateKeyGeneratorHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		IPrivateKeyGenerator:    importer,
	})
	if bccspContainer.IPrivateKeyGenerator == nil {
		bccspContainer.IPrivateKeyGenerator = NewDefaultMediator(alg)
	}
	bccspContainer.IPrivateKeyGenerator.LinkHiera(alg, newTemplate)
}

func RegisterKeyImporter(algInterface base.BaseType, detailType base.DetailType, importer IKeyImporterFunc) {
	alg := base.BaseType(algInterface)
	newTemplate := NewDefaultDetailTemplateMediator(&KeyImporterHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		IKeyImporter:            importer,
	})
	if bccspContainer.IKeyImporter == nil {
		bccspContainer.IKeyImporter = NewDefaultMediator(alg)
	}
	bccspContainer.IKeyImporter.LinkHiera(alg, newTemplate)
}

func RegisterKeyGenerator(algInterface base.BaseType, detailType base.DetailType, generator IKeyGeneratorFunc) {
	alg := base.BaseType(algInterface)
	newTemplate := NewDefaultDetailTemplateMediator(&KeyGeneratorHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		IKeyGenerator:           generator,
	})
	if bccspContainer.IKeyGenerator == nil {
		bccspContainer.IKeyGenerator = NewDefaultMediator(alg)
	}
	bccspContainer.IKeyGenerator.LinkHiera(alg, newTemplate)
}
func RegisterKeySigner(algInterface base.BaseType, detailType base.DetailType, signer ISignerFunc) {
	alg := base.BaseType(algInterface)
	newTemplate := NewDefaultDetailTemplateMediator(&SignerHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		ISigner:                 signer,
	})
	if bccspContainer.ISigner == nil {
		bccspContainer.ISigner = NewDefaultMediator(alg)
	}
	bccspContainer.ISigner.LinkHiera(alg, newTemplate)
}
func RegisterKeyVerifier(algInterface base.BaseType, detailType base.DetailType, verifier IVerifierFunc) {
	alg := base.BaseType(algInterface)
	newTemplate := NewDefaultDetailTemplateMediator(&VerifierHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		IVerifier:               verifier,
	})

	if bccspContainer.IVerifier == nil {
		bccspContainer.IVerifier = NewDefaultMediator(alg)
	}
	bccspContainer.IVerifier.LinkHiera(alg, newTemplate)
	// }
}
func RegisterCertificateImporter(algInterface base.BaseType, detailType base.DetailType, certificateImporter ICertificateImporterFunc) {
	alg := base.BaseType(algInterface)
	newTemplate := NewDefaultDetailTemplateMediator(&CertificateImporterHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		ICertificateImporter:    certificateImporter,
	})
	if bccspContainer.ICertificateImporter == nil {
		bccspContainer.ICertificateImporter = NewDefaultMediator(alg)
	}
	bccspContainer.ICertificateImporter.LinkHiera(alg, newTemplate)
}
func RegisterPublicKeyImporter(algInterface base.BaseType, detailType base.DetailType, im IPublicKeyImporterFunc) {
	alg := base.BaseType(algInterface)
	newTemplate := NewDefaultDetailTemplateMediator(&PublicKeyImportHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		IPublicKeyImporter:      im,
	})
	if bccspContainer.IPublicKeyImporter == nil {
		bccspContainer.IPublicKeyImporter = NewDefaultMediator(alg)
	} else {
	}
	bccspContainer.IPublicKeyImporter.LinkHiera(alg, newTemplate)
	// }
}
func RegisterOriginPublicKeyImporter(algInterface base.BaseType, detailType base.DetailType, im IOriginPublicKeyImporterFunc) {
	alg := base.BaseType(algInterface)
	newTemplate := NewDefaultDetailTemplateMediator(&OriginPublicKeyHandler{
		AlgorithmDetailTemplate:      base.NewAlgorithmDetailTemplate(detailType),
		IOriginPublicKeyImporterFunc: im,
	})
	if bccspContainer.IOriginPublicKeyImporter == nil {
		bccspContainer.IOriginPublicKeyImporter = NewDefaultMediator(alg)
	}
	bccspContainer.IOriginPublicKeyImporter.LinkHiera(alg, newTemplate)
	// }
}
func RegisterUnmarshal(algInterface base.BaseType, detailType base.DetailType, im IUnmarshalFunc) {
	alg := base.BaseType(algInterface)
	newTemplate := NewDefaultDetailTemplateMediator(&UnmarshalHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		IUnmarshal:              im,
	})
	if bccspContainer.IUnmarshal == nil {
		bccspContainer.IUnmarshal = NewDefaultMediator(alg)
	}
	bccspContainer.IUnmarshal.LinkHiera(alg, newTemplate)
}
func RegisterMarshal(algInterface base.BaseType, detailType base.DetailType, im IMarshalFunc) {
	alg := base.BaseType(algInterface)
	newTemplate := NewDefaultDetailTemplateMediator(&MarshalHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		IMarshal:                im,
	})
	if bccspContainer.IMarshal == nil {
		bccspContainer.IMarshal = NewDefaultMediator(alg)
	}
	bccspContainer.IMarshal.LinkHiera(alg, newTemplate)
}
func RegisterHash(algInterface base.BaseType, detailType base.DetailType, im IHashFunc) {
	alg := base.BaseType(algInterface)
	newTemplate := NewDefaultDetailTemplateMediator(&HashHandler{
		AlgorithmDetailTemplate: base.NewAlgorithmDetailTemplate(detailType),
		IHash:                   im,
	})
	if bccspContainer.IHash == nil {
		bccspContainer.IHash = NewDefaultMediator(alg)
	}
	bccspContainer.IHash.LinkHiera(alg, newTemplate)
}
