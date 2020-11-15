package com.charlie.blockchain.util;

import org.bouncycastle.asn1.pkcs.PrivateKeyInfo;
import org.bouncycastle.asn1.sec.ECPrivateKey;
import org.bouncycastle.asn1.x509.SubjectPublicKeyInfo;
import org.bouncycastle.crypto.params.ECPrivateKeyParameters;
import org.bouncycastle.jcajce.provider.asymmetric.ec.BCECPrivateKey;
import org.bouncycastle.jcajce.provider.config.ProviderConfiguration;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.openssl.jcajce.JcaPEMWriter;
import org.bouncycastle.util.Arrays;
import org.bouncycastle.util.BigIntegers;
import org.hyperledger.fabric.util.BCECUtil;
import org.hyperledger.fabric.util.SMUtil;

import java.io.IOException;
import java.io.StringWriter;
import java.math.BigInteger;
import java.security.PrivateKey;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-30 13:17
 */
public class KeyUtils
{

    public static String convKey2TPem(PrivateKey privateKey) throws IOException
    {
        StringWriter pemStrWriter = new StringWriter();
        JcaPEMWriter pemWriter = new JcaPEMWriter(pemStrWriter);

        pemWriter.writeObject(privateKey);

        pemWriter.close();

        return pemStrWriter.toString();

    }


    public static PrivateKey convPemKey2TPrv(byte[] origin)
    {
        byte[] prvBytes = new byte[0];
        try
        {
            prvBytes = SMUtil.parseSM2PrvK(new String(origin));
        } catch (Exception e)
        {
            throw new RuntimeException(e);
        }
        BigInteger bigInteger = new BigInteger(prvBytes);
        ECPrivateKeyParameters ecPrivateKeyParameters = org.hyperledger.fabric.util.BCECUtil.createECPrivateKeyParameters(bigInteger, BCECUtil.DOMAIN_PARAMS);
        ProviderConfiguration configuration = BouncyCastleProvider.CONFIGURATION;
        return new BCECPrivateKey("EC", ecPrivateKeyParameters, configuration);
    }

    public static enum EnumSm2ProducerType{
        BIAOXIN,
        SANWEI;

        private EnumSm2ProducerType() {
        }
    }

    public static  byte[] formatPrvKey(byte[] prvKey){
        if (prvKey.length == 32)
        {
            return prvKey;
        }
        if (prvKey.length == 33)
        {
            return java.util.Arrays.copyOfRange(prvKey, 1, 33);
        }
        PrivateKeyInfo bb = PrivateKeyInfo.getInstance(prvKey);
        byte[] keyDer = new byte[0];
        try
        {
            keyDer = bb.parsePrivateKey().toASN1Primitive().getEncoded("DER");
        } catch (IOException e)
        {
            throw new RuntimeException(e);
        }
        ECPrivateKey ecPrivateKey = ECPrivateKey.getInstance(keyDer);
        return BigIntegers.asUnsignedByteArray(32, ecPrivateKey.getKey());
    }
    public static byte[] formatPubKey(org.hyperledger.fabric.util.KeyUtils.EnumSm2ProducerType producerType, byte[] pubKey)
    {
        byte[] result = pubKey;
        if (result.length > 65)
        {
            result = SubjectPublicKeyInfo.getInstance(result).getPublicKeyData().getBytes();
        }
        if (result.length == 65)
        {
            result = Arrays.copyOfRange(result, 1, 65);
        }

        switch (producerType)
        {
            case BIAOXIN:
                result = Arrays.concatenate(new byte[]{0x04}, result);
                break;
            default:
                break;
        }
        return result;
    }
}
