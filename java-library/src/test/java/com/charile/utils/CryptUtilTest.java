package com.charile.utils;

import org.junit.Test;

import java.io.IOException;
import java.security.NoSuchAlgorithmException;

import static org.junit.Assert.*;

public class CryptUtilTest
{

    @Test
    public void md5File() throws IOException, NoSuchAlgorithmException
    {
//        426d1ad84581dac49a5ed203a417825e
        String path = "/Users/joker/Desktop/a.jpg";
        String s = CryptUtil.md5File(path);
        System.out.println(s);
        String s1 = CryptUtil.Sha256(s);
        System.out.println();
        String s2 = "62cae8002fc1285cb3552602fb52f2badac7eda6b4047e0434c624887e9f0bd4";
        System.out.println(s1.equalsIgnoreCase(s2));
    }
}