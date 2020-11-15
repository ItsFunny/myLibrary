package com.charlie;

import com.charlie.utils.Base64Utils;
import org.apache.commons.io.FileUtils;
import org.apache.http.util.Asserts;
import org.junit.Test;

import java.io.File;
import java.util.Base64;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-11-13 16:35
 */

public class Other
{
    @Test
    public void testBase64() throws Exception
    {
//        String str = "/Users/joker/company/go/src/github.com/hyperledger/fabric-demo/bidsun-sdk-multi-channels/src/main/resources/crypto-config/peerOrganizations/bidsun.com/users/Admin@bidsun.com/msp/signcerts/a.pem";
//        String s = FileUtils.readFileToString(new File(str), "UTF-8");
//        s = Base64Utils.encode(s.getBytes());
//        String cert2 = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNCVENDQWF5Z0F3SUJBZ0lRUVo5SGh4ejRhVFpCa2cyU2l6eVhrakFLQmdncWdSelBWUUdEZFRCZU1Rc3cKQ1FZRFZRUUdFd0pEVGpFUU1BNEdBMVVFQ0JNSFFtVnBhbWx1WnpFUU1BNEdBMVVFQnhNSFFtVnBhbWx1WnpFVApNQkVHQTFVRUNoTUtZbWxrYzNWdUxtTnZiVEVXTUJRR0ExVUVBeE1OWTJFdVltbGtjM1Z1TG1OdmJUQWVGdzB5Ck1ERXdNak13TWpNeE1EQmFGdzB6TURFd01qRXdNak14TURCYU1GMHhDekFKQmdOVkJBWVRBa05PTVJBd0RnWUQKVlFRSUV3ZENaV2xxYVc1bk1SQXdEZ1lEVlFRSEV3ZENaV2xxYVc1bk1ROHdEUVlEVlFRTEV3WmpiR2xsYm5ReApHVEFYQmdOVkJBTU1FRUZrYldsdVFHSnBaSE4xYmk1amIyMHdXVEFUQmdjcWhrak9QUUlCQmdncWdSelBWUUdDCkxRTkNBQVROZ3dNNjNGVURuL0RRV2hJelk1RzY1NC9FK2kzcnZNcGVWTnZSYzFkZ25qQzZuTmlaYk5keG5hUzQKa0xGc0xsM3Fvck1FeXRPb0pIclg2eFY3K0I4QW8wMHdTekFPQmdOVkhROEJBZjhFQkFNQ0I0QXdEQVlEVlIwVApBUUgvQkFJd0FEQXJCZ05WSFNNRUpEQWlnQ0F3c2t6ZEFMWU40S3hBaEhvSUg1TnZVNjc5Qk9TN0ZjVUdkV3QzClNYZnFBVEFLQmdncWdSelBWUUdEZFFOSEFEQkVBaUFGNVVqZnhTQWF4RVVyVlROa2JvNmpyQzUzaUliWVFWTXcKWmRDeWRFOEdUQUlnYlRjMU5Td1VldWdhSnR6cE1mUGRBTC92N3NYWHpCRk9GUURkcGNiaWdEQT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=";
//        System.out.println(s.equals(cert2));
////        System.exit(-1);
//        String cert = "-----BEGIN CERTIFICATE-----\n" +
//                "MIICBTCCAaygAwIBAgIQQZ9Hhxz4aTZBkg2SizyXkjAKBggqgRzPVQGDdTBeMQsw\n" +
//                "CQYDVQQGEwJDTjEQMA4GA1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzET\n" +
//                "MBEGA1UEChMKYmlkc3VuLmNvbTEWMBQGA1UEAxMNY2EuYmlkc3VuLmNvbTAeFw0y\n" +
//                "MDEwMjMwMjMxMDBaFw0zMDEwMjEwMjMxMDBaMF0xCzAJBgNVBAYTAkNOMRAwDgYD\n" +
//                "VQQIEwdCZWlqaW5nMRAwDgYDVQQHEwdCZWlqaW5nMQ8wDQYDVQQLEwZjbGllbnQx\n" +
//                "GTAXBgNVBAMMEEFkbWluQGJpZHN1bi5jb20wWTATBgcqhkjOPQIBBggqgRzPVQGC\n" +
//                "LQNCAATNgwM63FUDn/DQWhIzY5G654/E+i3rvMpeVNvRc1dgnjC6nNiZbNdxnaS4\n" +
//                "kLFsLl3qorMEytOoJHrX6xV7+B8Ao00wSzAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0T\n" +
//                "AQH/BAIwADArBgNVHSMEJDAigCAwskzdALYN4KxAhHoIH5NvU679BOS7FcUGdWt3\n" +
//                "SXfqATAKBggqgRzPVQGDdQNHADBEAiAF5UjfxSAaxEUrVTNkbo6jrC53iIbYQVMw\n" +
//                "ZdCydE8GTAIgbTc1NSwUeugaJtzpMfPdAL/v7sXXzBFOFQDdpcbigDA=\n" +
//                "-----END CERTIFICATE-----\n";
//        String encode = Base64Utils.encode(cert.getBytes());
//        System.out.println(encode.equals(cert2));

        int compare = Integer.compare(1, 2);
        System.out.println(compare);

    }

}
