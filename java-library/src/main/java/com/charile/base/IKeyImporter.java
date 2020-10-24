package com.charile.base;

import com.charile.utils.BCECUtil;
import com.charile.utils.GMUtil;
import org.bouncycastle.asn1.pkcs.PrivateKeyInfo;
import org.bouncycastle.crypto.params.ECPrivateKeyParameters;
import org.bouncycastle.jcajce.provider.asymmetric.ec.BCECPrivateKey;
import org.bouncycastle.jcajce.provider.config.ProviderConfiguration;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.openssl.PEMParser;
import org.bouncycastle.openssl.jcajce.JcaPEMKeyConverter;

import java.io.ByteArrayInputStream;
import java.io.InputStreamReader;
import java.io.Reader;
import java.io.StringReader;
import java.math.BigInteger;
import java.nio.charset.StandardCharsets;
import java.security.PrivateKey;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-24 15:25
 */
public interface IKeyImporter
{
    PrivateKey bytes2PrivateKey(byte[] origin) throws Exception;


    IKeyImporter ANS1_PEM_KEY_IMPORTER = (penKey) ->
    {
        InputStreamReader inputStreamReader = new InputStreamReader(new ByteArrayInputStream(penKey));
        final Reader pemReader = inputStreamReader;
        final PrivateKeyInfo pemPair;
        PEMParser pemParse = new PEMParser(pemReader);
        pemPair = (PrivateKeyInfo) pemParse.readObject();
        PrivateKey pk = new JcaPEMKeyConverter().getPrivateKey(pemPair);
        pemParse.close();
        return pk;
    };
    IKeyImporter COMMON_PEM_KEY_IMPORTER = (pemKey) ->
    {
        Reader pemReader = new StringReader(new String(pemKey, StandardCharsets.UTF_8));
        PEMParser pemParser = new PEMParser(pemReader);
        JcaPEMKeyConverter converter = (new JcaPEMKeyConverter());
        Object object = pemParser.readObject();
        return converter.getPrivateKey((PrivateKeyInfo) object);
    };
    IKeyImporter STANDARD_SM2_KEY_IMPORTER = (penKey) ->
    {
        byte[] prvBytes = GMUtil.parseSM2PrvK(new String(penKey));
        BigInteger bigInteger = new BigInteger(prvBytes);
        ECPrivateKeyParameters ecPrivateKeyParameters = BCECUtil.createECPrivateKeyParameters(bigInteger, BCECUtil.DOMAIN_PARAMS);
        ProviderConfiguration configuration = BouncyCastleProvider.CONFIGURATION;
        return new BCECPrivateKey("EC", ecPrivateKeyParameters, configuration);
    };
}
