package com.charlie;

import io.netty.util.concurrent.CompleteFuture;

import java.lang.ref.PhantomReference;
import java.lang.ref.WeakReference;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeUnit;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-11-05 21:07
 */
public class T
{
    public static class AA
    {
        @Override
        protected void finalize() throws Throwable
        {
            super.finalize();
            System.out.println("aa 回收");
        }
    }

    public static class BB extends AA
    {

    }

    public static void main(String[] args) throws Exception
    {
        System.out.println(Thread.currentThread().getId());
        final  CompletableFuture<String> completeFuture=new CompletableFuture<>();
        new Thread(()->{
            try
            {
                TimeUnit.SECONDS.sleep(5);
                completeFuture.complete("ok");
            } catch (InterruptedException e)
            {
            }
        }).start();

        String s = completeFuture.get();
        System.out.println(s);
        System.out.println(Thread.currentThread().getId());
    }

}
