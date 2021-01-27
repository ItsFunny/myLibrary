package com.charlie;

import com.charlie.utils.Base64Utils;

import javax.crypto.KeyGenerator;
import java.security.NoSuchAlgorithmException;
import java.security.SecureRandom;

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

    public static void main(String[] args)
    {
        int maxPeerSize = 10;
        int minPeerSize =3;
        double percent = 0.6;
        int checkPeerCount = getCheckPeerCount(4, minPeerSize, maxPeerSize, percent);
        System.out.println(checkPeerCount);
//        KeyGenerator gen = null;
//        try
//        {
//            gen = KeyGenerator.getInstance("AES");
//        } catch (NoSuchAlgorithmException e)
//        {
//            throw new RuntimeException("not likely:" + e.getMessage());
//        }
//        gen.init(256, new SecureRandom());
//        byte[] trueAesKey = gen.generateKey().getEncoded();
//        System.out.println(Base64Utils.encode(trueAesKey));
//        Calendar cal = Calendar.getInstance();
//        cal.set(2020, 10, 11);
//        Date time = cal.getTime();
//        int timeTillSeconds = DateUtils.getTimeTillSeconds(time);
//        System.out.println(timeTillSeconds);
    }
    public static int getCheckPeerCount(int peerCount, int minPeerCount, int maxPeerCount, double participatePercent) {
        int result = Double.valueOf(Math.ceil((double)peerCount * participatePercent)).intValue();
        result = result > maxPeerCount ? maxPeerCount : result;
        result = result < minPeerCount ? minPeerCount : result;
        result = result > peerCount ? peerCount : result;
        return result;
    }

//    public static void main(String[] args) throws Exception
//    {
//        Calendar instance = Calendar.getInstance();
//        instance.add(Calendar.MONTH,3);
//        long time = instance.getTime().getTime();
//        System.out.println(time);
//
////        System.out.println(Thread.currentThread().getId());
////        final  CompletableFuture<String> completeFuture=new CompletableFuture<>();
////        new Thread(()->{
////            try
////            {
////                TimeUnit.SECONDS.sleep(5);
////                completeFuture.complete("ok");
////            } catch (InterruptedException e)
////            {
////            }
////        }).start();
////
////        String s = completeFuture.get();
////        System.out.println(s);
////        System.out.println(Thread.currentThread().getId());
//    }


}
