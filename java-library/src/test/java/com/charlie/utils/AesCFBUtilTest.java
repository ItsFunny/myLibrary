package com.charlie.utils;

import org.junit.Test;

public class AesCFBUtilTest
{

    @Test
    public void decryWithGo() throws Exception
    {
        String guangzhouKey="WbUBp5dzlv7gHeST94SqiLhj/IQU0Kw9EN1hmypmyig=";
        String encStr="uw6f8GLTSAHLGyazX4gdFez1Cg==";
        byte[] bytes = AesCFBUtil.decryptWithGo( encStr,guangzhouKey);
        System.out.println(new String(bytes));
    }

    @Test
    public void test111()throws Exception{
        String str="MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAErKhaCg4U/v0wJnkBGrtto9IFcIpLXveIy0oMwXPqQw7RSxxyXS2Y0UC2huuQWlqaF5Sz5LseIE/SH8/CI28HUA==";
        byte[] decode = Base64Utils.decode(str);
        System.out.println(decode.length);
    }
}