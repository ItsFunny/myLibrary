package com.charlie;

import org.junit.Test;

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

    @Test
    public void test111(){
        String   a="-----BEGIN CERTIFICATE-----" +
                "MIICLzCCAdWgAwIBAgIUUZ5qIB9wFU7BlWQRGN0bVyuUJGAwCgYIKoZIzj0EAwIw" +
                "ZzELMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNh" +
                "biBGcmFuY2lzY28xEzARBgNVBAoTCmJpZHN1bi5jb20xFjAUBgNVBAMTDWNhLmJp" +
                "ZHN1bi5jb20wHhcNMjAxMjE2MDQ0NzAwWhcNMjExMjE2MDQ1MjAwWjAsMR4wDQYD" +
                "VQQLEwZjbGllbnQwDQYDVQQLEwZiaWRzdW4xCjAIBgNVBAMTATEwWTATBgcqhkjO" +
                "PQIBBggqhkjOPQMBBwNCAAQYga8+g/ugwq0qf/PZKVTQwZhXM5pLrw7JerLJVIrx" +
                "Htex98+r0Y03lt1nCLMx2w7wGSYdecTa6zmQ7GZR0aXso4GZMIGWMA4GA1UdDwEB" +
                "/wQEAwIHgDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBSsx6UpQNvwnLOXO7gn11py" +
                "o6LK0DArBgNVHSMEJDAigCBfnOBA73ZtQVbOoERR4UGSPIMoMV7PEgaunxcKzhmI" +
                "GDAqBggqAwQFBgcIAQQeeyJhdHRycyI6eyJpc0VCSURTVU4iOiJ0cnVlIn19MAoG" +
                "CCqGSM49BAMCA0gAMEUCIQDIz+GmFHSnfrD/ArsSvKcXbQ5lGyIoltbkQZ0Re+0L" +
                "VAIgDR/Gt/gVIvcogOg97WxOjCz06fRoTq+vGos+8goPpBQ=" +
                "-----END CERTIFICATE-----";

        String   b="-----BEGIN CERTIFICATE-----" +
                "MIICNDCCAdqgAwIBAgIUD7o7m6iJbsdJtRN7dVwSsPnbvggwCgYIKoZIzj0EAwIw" +
                "ZzELMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNh" +
                "biBGcmFuY2lzY28xEzARBgNVBAoTCmJpZHN1bi5jb20xFjAUBgNVBAMTDWNhLmJp" +
                "ZHN1bi5jb20wHhcNMjAxMjE2MDYxMTAwWhcNMjExMjE2MDYxNjAwWjAxMR4wDQYD" +
                "VQQLEwZjbGllbnQwDQYDVQQLEwZiaWRzdW4xDzANBgNVBAMTBjEyMzEyMzBZMBMG" +
                "ByqGSM49AgEGCCqGSM49AwEHA0IABIuuuIQ1ED0lLwzTa9/t8VVcf7ouv7IG+YAc" +
                "kU8z3d3Du2yqTPSP8ehqUBdmO/5iIxV8H5zLRqW4iPOkP9gwS/ajgZkwgZYwDgYD" +
                "VR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFH4eSwUw+FLSmwon" +
                "rS7EbG6HPONPMCsGA1UdIwQkMCKAIF+c4EDvdm1BVs6gRFHhQZI8gygxXs8SBq6f" +
                "FwrOGYgYMCoGCCoDBAUGBwgBBB57ImF0dHJzIjp7ImlzRUJJRFNVTiI6InRydWUi" +
                "fX0wCgYIKoZIzj0EAwIDSAAwRQIhAN8vdiprlxYH7gkmSslKcmB0zwrQetrHM0Dw" +
                "ZrOH2ZbFAiAHSzp9ItQ56j8w4skWoVisLRDixARHo0rvVHx77MFdwA==" +
                "-----END CERTIFICATE-----";

        String c="-----BEGIN CERTIFICATE-----" +
                "MIIB3zCCAYSgAwIBAgIUcJIPBaGkuJhgcD7U+CXGbxGfWBowCgYIKoEcz1UBg3Uw" +
                "XjELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxEDAOBgNVBAcTB0JlaWpp" +
                "bmcxEzARBgNVBAoTCmJpZHN1bi5jb20xFjAUBgNVBAMTDWNhLmJpZHN1bi5jb20w" +
                "HhcNMjAxMjE2MDgwNzIwWhcNMjEwMTI3MDAwNzIwWjAjMQ8wDQYDVQQLEwZjbGll" +
                "bnQxEDAOBgNVBAMTB0VCSURTVU4wWTATBgcqhkjOPQIBBggqgRzPVQGCLQNCAAQO" +
                "oRSpgCew6AZThPi6k3Dt7n8B8nfkMCtlA8o0RpU1Cs/fbpwCc8HcodGLQyiMGHPs" +
                "1zAY1KfqcQGTL23HpoQJo1swWTArBgNVHSMEJDAigCA+kWBNgv4cliPwV9vTs7l1" +
                "iZIctTmXmd3Jy9vhMo//xDAqBggqAwQFBgcIAQQeeyJhdHRycyI6eyJpc0VCSURT" +
                "VU4iOiJ0cnVlIn19MAoGCCqBHM9VAYN1A0kAMEYCIQCI/sJpS/XMl1mSkWQFr2S5" +
                "7Q2WdzwDvN+lHZxkUUL5hAIhAM8SncnMLqxJGH0a3o/df/I9amgkxNDowg/VwLQz" +
                "kj95-----END CERTIFICATE-----";
    }
}
