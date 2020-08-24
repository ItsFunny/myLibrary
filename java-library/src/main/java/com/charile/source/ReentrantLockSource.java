package com.charile.source;

import java.util.HashMap;
import java.util.concurrent.locks.ReentrantLock;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-08-19 10:45
 */
public class ReentrantLockSource
{
    public static void main(String[] args)
    {
        HashMap<String,String>stringStringHashMap=new HashMap<>();
        stringStringHashMap.put("","");
        ReentrantLock lock = new ReentrantLock();
        lock.lock();
    }

}
