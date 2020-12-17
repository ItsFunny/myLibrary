package com.charlie.utils;

import com.charlie.model.CertInfo;
import org.junit.Test;

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

    @Test
    public void testSign() throws Exception
    {
//        String prvStr = "-----BEGIN PRIVATE KEY-----\n" +
//                "JtXTdE0BfnqkRxLw3CX89rtRzCjAum+3ITBid7hFFDw=\n" +
//                "-----END PRIVATE KEY-----\n";
//        String prvStr = "VCJoKMbCL6xaKn1ijlrYVY49s2dJyQDvvZKEjA51Tys=";
        String prvStr = "-----BEGIN PRIVATE KEY-----\n" +
                "JyWCdi/UlfKjkO0IXfxx861EWO/3EOJfzpDqykUPO1g=\n" +
                "-----END PRIVATE KEY-----\n";
//        String certStr = "-----BEGIN CERTIFICATE-----\n" +
//                "MIIB/zCCAaSgAwIBAgIUZz4YRV73/IeOHAvi92BdWHsY3s0wCgYIKoEcz1UBg3Uw\n" +
//                "XjELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxEDAOBgNVBAcTB0JlaWpp\n" +
//                "bmcxEzARBgNVBAoTCmJpZHN1bi5jb20xFjAUBgNVBAMTDWNhLmJpZHN1bi5jb20w\n" +
//                "HhcNMjAxMTE3MTM0NDA1WhcNMjAxMjI5MDU0NDA1WjBdMQswCQYDVQQGEwJVUzEX\n" +
//                "MBUGA1UECBMOTm9ydGggQ2Fyb2xpbmExFDASBgNVBAoTC0h5cGVybGVkZ2VyMQ8w\n" +
//                "DQYDVQQLEwZGYWJyaWMxDjAMBgNVBAMTBWFkbWluMFkwEwYHKoZIzj0CAQYIKoEc\n" +
//                "z1UBgi0DQgAEObVjMg0TmVJHvh0FsU0pyvyRJF47zQAI6eL2SMHV85KuUDBBneed\n" +
//                "c37mv1Y1KBvNQlCCJOggU+eTfMPZ4iX8waNBMD8wKwYDVR0jBCQwIoAgKkvUUjAw\n" +
//                "DAZPYIWEnrIXiEQO4+lc0MPwptU/dyh4HSMwEAYDVR0RBAkwB4IFam9rZXIwCgYI\n" +
//                "KoEcz1UBg3UDSQAwRgIhAKKD5hoBpaPcYjtFd10tHrShBe8G1Fi1mgV3oJuL0F54\n" +
//                "AiEA9braJbBi0GSm9yCqESkC6SdgS0ze4QEISOq96Fi9tNw=\n" +
//                "-----END CERTIFICATE-----\n";
//        String certStr = "MIICJzCCAc2gAwIBAgIQZdx33rUcvc+bsvzhN81lIzAKBggqgRzPVQGDdTBeMQswCQYDVQQGEwJDTjEQMA4GA1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzETMBEGA1UEChMKYmlkc3VuLmNvbTEWMBQGA1UEAxMNY2EuYmlkc3VuLmNvbTAeFw0yMDExMDUwOTE4MDBaFw0zMDExMDMwOTE4MDBaMF4xCzAJBgNVBAYTAkNOMRAwDgYDVQQIEwdCZWlqaW5nMRAwDgYDVQQHEwdCZWlqaW5nMRMwEQYDVQQKEwpiaWRzdW4uY29tMRYwFAYDVQQDEw1jYS5iaWRzdW4uY29tMFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEkLyoI/sYQuzljuFFOS3QG+GeNqr0pu1qN5TisyXyelWEQvro7PInxtRca4lqco69xfIO9kCphnF8+LBqOeFFiKNtMGswDgYDVR0PAQH/BAQDAgGmMB0GA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1UdDgQiBCAqS9RSMDAMBk9ghYSesheIRA7j6VzQw/Cm1T93KHgdIzAKBggqgRzPVQGDdQNIADBFAiEAwmQ5thd3WuMcFp0fWdDlqZzgiMLHXqzQWEBuNFVHtf4CIC7oigK3UZV7T/nSR2pm78ZWx8gJzQrLwfbvk2G/j5XU";
        String certStr = "-----BEGIN CERTIFICATE-----\n" +
                "MIICBzCCAa2gAwIBAgIRAP87V6AFe+vj7vaChirtSRUwCgYIKoEcz1UBg3UwXjEL\n" +
                "MAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxEDAOBgNVBAcTB0JlaWppbmcx\n" +
                "EzARBgNVBAoTCmJpZHN1bi5jb20xFjAUBgNVBAMTDWNhLmJpZHN1bi5jb20wHhcN\n" +
                "MjAxMTA1MDkxODAwWhcNMzAxMTAzMDkxODAwWjBdMQswCQYDVQQGEwJDTjEQMA4G\n" +
                "A1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzEPMA0GA1UECxMGY2xpZW50\n" +
                "MRkwFwYDVQQDDBBBZG1pbkBiaWRzdW4uY29tMFkwEwYHKoZIzj0CAQYIKoEcz1UB\n" +
                "gi0DQgAEMGyVhtrZbT4LLpka/Ao6qqpz0A9iRIIKWXEjvLh/Ok4k53ctk8EpM2sG\n" +
                "zc5hn9o+kC3Dqpe6C2YFi3wapKf/JaNNMEswDgYDVR0PAQH/BAQDAgeAMAwGA1Ud\n" +
                "EwEB/wQCMAAwKwYDVR0jBCQwIoAgKkvUUjAwDAZPYIWEnrIXiEQO4+lc0MPwptU/\n" +
                "dyh4HSMwCgYIKoEcz1UBg3UDSAAwRQIgZAZK27Lvinw6VvTz5okVhtPWKQDhHted\n" +
                "+fYRzvvm2OMCIQCp0T897Joz8Pv5L6R3AsFbPsay0zIaZdU0Q1qIljOt/Q==\n" +
                "-----END CERTIFICATE-----\n";
        String bidsunCert = "-----BEGIN CERTIFICATE-----\n" +
                "MIICJzCCAc2gAwIBAgIQZdx33rUcvc+bsvzhN81lIzAKBggqgRzPVQGDdTBeMQsw\n" +
                "CQYDVQQGEwJDTjEQMA4GA1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzET\n" +
                "MBEGA1UEChMKYmlkc3VuLmNvbTEWMBQGA1UEAxMNY2EuYmlkc3VuLmNvbTAeFw0y\n" +
                "MDExMDUwOTE4MDBaFw0zMDExMDMwOTE4MDBaMF4xCzAJBgNVBAYTAkNOMRAwDgYD\n" +
                "VQQIEwdCZWlqaW5nMRAwDgYDVQQHEwdCZWlqaW5nMRMwEQYDVQQKEwpiaWRzdW4u\n" +
                "Y29tMRYwFAYDVQQDEw1jYS5iaWRzdW4uY29tMFkwEwYHKoZIzj0CAQYIKoEcz1UB\n" +
                "gi0DQgAEkLyoI/sYQuzljuFFOS3QG+GeNqr0pu1qN5TisyXyelWEQvro7PInxtRc\n" +
                "a4lqco69xfIO9kCphnF8+LBqOeFFiKNtMGswDgYDVR0PAQH/BAQDAgGmMB0GA1Ud\n" +
                "JQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1Ud\n" +
                "DgQiBCAqS9RSMDAMBk9ghYSesheIRA7j6VzQw/Cm1T93KHgdIzAKBggqgRzPVQGD\n" +
                "dQNIADBFAiEAwmQ5thd3WuMcFp0fWdDlqZzgiMLHXqzQWEBuNFVHtf4CIC7oigK3\n" +
                "UZV7T/nSR2pm78ZWx8gJzQrLwfbvk2G/j5XU\n" +
                "-----END CERTIFICATE-----\n";
        // ca 生成的
        String newCert = "-----BEGIN CERTIFICATE-----\n" +
                "MIIB/jCCAaSgAwIBAgIUC3cFGGHw2IWVS6VnUOt/q89bwsAwCgYIKoEcz1UBg3Uw\n" +
                "XjELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxEDAOBgNVBAcTB0JlaWpp\n" +
                "bmcxEzARBgNVBAoTCmJpZHN1bi5jb20xFjAUBgNVBAMTDWNhLmJpZHN1bi5jb20w\n" +
                "HhcNMjAxMTE4MDI0NzQwWhcNMjAxMjI5MTg0NzQwWjBdMQswCQYDVQQGEwJVUzEX\n" +
                "MBUGA1UECBMOTm9ydGggQ2Fyb2xpbmExFDASBgNVBAoTC0h5cGVybGVkZ2VyMQ8w\n" +
                "DQYDVQQLEwZGYWJyaWMxDjAMBgNVBAMTBWFkbWluMFkwEwYHKoZIzj0CAQYIKoEc\n" +
                "z1UBgi0DQgAEXm0oi0+2ZidPGfSXrtKxr8uf6emg+F1v+CvjTXHLw+emf5/L6SAk\n" +
                "U0NTsnNFh7TXgCJv7o35XiEf2Pho96WbQ6NBMD8wKwYDVR0jBCQwIoAgKkvUUjAw\n" +
                "DAZPYIWEnrIXiEQO4+lc0MPwptU/dyh4HSMwEAYDVR0RBAkwB4IFam9rZXIwCgYI\n" +
                "KoEcz1UBg3UDSAAwRQIhAILPaJWR0yjCoez6xROFbWekkCf4pPI4cj4wNpRqKQa1\n" +
                "AiB1YxwmIcREDxu9GMpn4N/TOx6b9BqU+nFHK3Itc9JcIQ==\n" +
                "-----END CERTIFICATE-----\n";
        // 自带的: 既crypro-gen 生成的amdin证书
        String adminCerts = "-----BEGIN CERTIFICATE-----\n" +
                "MIICBzCCAa2gAwIBAgIRAP87V6AFe+vj7vaChirtSRUwCgYIKoEcz1UBg3UwXjEL\n" +
                "MAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxEDAOBgNVBAcTB0JlaWppbmcx\n" +
                "EzARBgNVBAoTCmJpZHN1bi5jb20xFjAUBgNVBAMTDWNhLmJpZHN1bi5jb20wHhcN\n" +
                "MjAxMTA1MDkxODAwWhcNMzAxMTAzMDkxODAwWjBdMQswCQYDVQQGEwJDTjEQMA4G\n" +
                "A1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzEPMA0GA1UECxMGY2xpZW50\n" +
                "MRkwFwYDVQQDDBBBZG1pbkBiaWRzdW4uY29tMFkwEwYHKoZIzj0CAQYIKoEcz1UB\n" +
                "gi0DQgAEMGyVhtrZbT4LLpka/Ao6qqpz0A9iRIIKWXEjvLh/Ok4k53ctk8EpM2sG\n" +
                "zc5hn9o+kC3Dqpe6C2YFi3wapKf/JaNNMEswDgYDVR0PAQH/BAQDAgeAMAwGA1Ud\n" +
                "EwEB/wQCMAAwKwYDVR0jBCQwIoAgKkvUUjAwDAZPYIWEnrIXiEQO4+lc0MPwptU/\n" +
                "dyh4HSMwCgYIKoEcz1UBg3UDSAAwRQIgZAZK27Lvinw6VvTz5okVhtPWKQDhHted\n" +
                "+fYRzvvm2OMCIQCp0T897Joz8Pv5L6R3AsFbPsay0zIaZdU0Q1qIljOt/Q==\n" +
                "-----END CERTIFICATE-----\n";
//        String certPath = "/Users/joker/go/src/java_go_web/testdata/a.pem";
        byte[] prvByes = GMUtil.parseSM2PrvK(prvStr);
        CertInfo certInfo = GMUtil.parseSM2CertStr(certStr);
        CertInfo bisunCert = GMUtil.parseSM2CertStr(bidsunCert);
        CertInfo certInfo1 = GMUtil.parseSM2CertStr(newCert);
        CertInfo certInfo2 = GMUtil.parseSM2CertStr(adminCerts);
        System.out.println(1);
//        CertInfo certInfo = GMUtil.parseSM2CertStrFromFile(certPath);
//        byte[] bs = "123".getBytes();
//        byte[] sign = GMUtil.sign(prvByes, bs);
//        byte[] pubBytes = GMUtil.formatPubKey(certInfo.getPubKey());
//        boolean verify = GMUtil.verifyOwner(pubBytes, bs, sign);
//        System.out.println(verify);
    }

    @Test
    public void test1111() throws Exception
    {
        String sm2Crt = "-----BEGIN CERTIFICATE-----\n" +
                "MIICBTCCAaygAwIBAgIQQZ9Hhxz4aTZBkg2SizyXkjAKBggqgRzPVQGDdTBeMQsw\n" +
                "CQYDVQQGEwJDTjEQMA4GA1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzET\n" +
                "MBEGA1UEChMKYmlkc3VuLmNvbTEWMBQGA1UEAxMNY2EuYmlkc3VuLmNvbTAeFw0y\n" +
                "MDEwMjMwMjMxMDBaFw0zMDEwMjEwMjMxMDBaMF0xCzAJBgNVBAYTAkNOMRAwDgYD\n" +
                "VQQIEwdCZWlqaW5nMRAwDgYDVQQHEwdCZWlqaW5nMQ8wDQYDVQQLEwZjbGllbnQx\n" +
                "GTAXBgNVBAMMEEFkbWluQGJpZHN1bi5jb20wWTATBgcqhkjOPQIBBggqgRzPVQGC\n" +
                "LQNCAATNgwM63FUDn/DQWhIzY5G654/E+i3rvMpeVNvRc1dgnjC6nNiZbNdxnaS4\n" +
                "kLFsLl3qorMEytOoJHrX6xV7+B8Ao00wSzAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0T\n" +
                "AQH/BAIwADArBgNVHSMEJDAigCAwskzdALYN4KxAhHoIH5NvU679BOS7FcUGdWt3\n" +
                "SXfqATAKBggqgRzPVQGDdQNHADBEAiAF5UjfxSAaxEUrVTNkbo6jrC53iIbYQVMw\n" +
                "ZdCydE8GTAIgbTc1NSwUeugaJtzpMfPdAL/v7sXXzBFOFQDdpcbigDA=\n" +
                "-----END CERTIFICATE-----\n";
        CertInfo certInfo = GMUtil.parseSM2CertStr(sm2Crt);
        byte[] prvK = GMUtil.parseSM2PrvK("-----BEGIN PRIVATE KEY-----\n" +
                "NcRv3UwxSYfBDYs+qhMuZSuLF+qYDGQSBWMuVx9XU7w=\n" +
                "-----END PRIVATE KEY-----\n");
        byte[] bytes = "123".getBytes();
        byte[] sign = GMUtil.sign(prvK, bytes);
        System.out.println("====================");
        System.out.println(new String(sign));
        System.out.println("====================");
        boolean b = GMUtil.verifyOwner(GMUtil.formatPubKey(certInfo.getPubKey()), bytes, sign);
        System.out.println(b);
    }

    @Test
    public void test22222() throws Exception
    {
        String sig = "3df2c2b6a064895a8ba08dfa2d043e998b62e9b249d156cad449e6b82719ccac6f453d04599676cf3244fc53f670cff96d82f774b06d641cb035c45b09ae7bde";
        String sm2Crt = "-----BEGIN CERTIFICATE-----\n" +
                "MIICBTCCAaygAwIBAgIQQZ9Hhxz4aTZBkg2SizyXkjAKBggqgRzPVQGDdTBeMQsw\n" +
                "CQYDVQQGEwJDTjEQMA4GA1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzET\n" +
                "MBEGA1UEChMKYmlkc3VuLmNvbTEWMBQGA1UEAxMNY2EuYmlkc3VuLmNvbTAeFw0y\n" +
                "MDEwMjMwMjMxMDBaFw0zMDEwMjEwMjMxMDBaMF0xCzAJBgNVBAYTAkNOMRAwDgYD\n" +
                "VQQIEwdCZWlqaW5nMRAwDgYDVQQHEwdCZWlqaW5nMQ8wDQYDVQQLEwZjbGllbnQx\n" +
                "GTAXBgNVBAMMEEFkbWluQGJpZHN1bi5jb20wWTATBgcqhkjOPQIBBggqgRzPVQGC\n" +
                "LQNCAATNgwM63FUDn/DQWhIzY5G654/E+i3rvMpeVNvRc1dgnjC6nNiZbNdxnaS4\n" +
                "kLFsLl3qorMEytOoJHrX6xV7+B8Ao00wSzAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0T\n" +
                "AQH/BAIwADArBgNVHSMEJDAigCAwskzdALYN4KxAhHoIH5NvU679BOS7FcUGdWt3\n" +
                "SXfqATAKBggqgRzPVQGDdQNHADBEAiAF5UjfxSAaxEUrVTNkbo6jrC53iIbYQVMw\n" +
                "ZdCydE8GTAIgbTc1NSwUeugaJtzpMfPdAL/v7sXXzBFOFQDdpcbigDA=\n" +
                "-----END CERTIFICATE-----\n";
        CertInfo certInfo = GMUtil.parseSM2CertStr(sm2Crt);
        byte[] content = "123".getBytes();
        boolean b = GMUtil.verifyOwner(certInfo.getPubKey(), content, sig.getBytes());
        System.out.println(b);
    }

    @Test
    public void test2333() throws Exception
    {
        String sm2PrvK = "-----BEGIN PRIVATE KEY-----\\nS1NOngg/JG1e316Jdgae48P46U3N0zULqlpP6F+Qm+k=\\n-----END PRIVATE KEY-----\\n";
        
    }
}