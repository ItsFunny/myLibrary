package com.charlie.utils;

import com.charlie.model.CertInfo;
import org.junit.Test;

import static org.junit.Assert.*;

public class KeyUtilsTest
{

    @Test
    public void test11() throws Exception
    {
        String base64Prv = "KafaemcohSPpwu978UHkHDSUGoNXuRm0T1OSOnG/13Y=";
        String base64Crt = "-----BEGIN CERTIFICATE-----\n" +
                "MIIBqDCCAU6gAwIBAgIUAoY6vLw6CI99CSaWPvT4qHtkUy8wCgYIKoEcz1UBg3Uw\n" +
                "cTElMCMGA1UECh4cAFoANwByAEkAcwBfAGIAaQBkAHMAcwBzAHUAbjElMCMGA1UE\n" +
                "Cx4cAFoANwByAEkAcwBfAGIAaQBkAHMAcwBzAHUAbjEhMB8GA1UEAx4YAFoANwBy\n" +
                "AEkAcwBvAHIAZwBOAGEAbQBlMB4XDTIwMTAzMTA1MzEyOFoXDTIwMTIxMTIxMzEy\n" +
                "OFowEjEQMA4GA1UEAxMHb3JnNDY2NDBZMBMGByqGSM49AgEGCCqBHM9VAYItA0IA\n" +
                "BEVNHF6sd86mVrMZW3HUlXpUVJVbNIQ6RgH19x/7aQGLZgCKIhrqFomNtBTKyyey\n" +
                "jkMO5vomyt/UaSy3JA10jVCjIzAhMB8GA1UdIwQYMBaAFHAOluR6oZJudbClaffb\n" +
                "lbsinbp/MAoGCCqBHM9VAYN1A0gAMEUCIQDEBy54EfTMxnRlKFhpWeMneb46xq7y\n" +
                "3PPxvqdFlouE/wIgTLt1c3Cs2mXlLJOZrSmtNe4Hum1dU7dZeAZNf3+qf0g=\n" +
                "-----END CERTIFICATE-----";
        byte[] prvBytes = KeyUtils.parseSM2PrvK(base64Prv);
        CertInfo certInfo = CertificateUtils.parseCertStr2CertInfo(base64Crt);
        String pubKey = certInfo.getPubKey();
        byte[] pubBytes = Base64Utils.decode(pubKey);
        byte[] bytes = "123".getBytes();
        String s = KeyUtils.standardSMSignWithHexReturn(prvBytes, bytes);
        boolean b = KeyUtils.standardSM2VerifyWithHexString(pubBytes, bytes, s);
        System.out.println(b);
    }

}