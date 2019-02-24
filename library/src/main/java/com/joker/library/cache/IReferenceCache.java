package com.joker.library.cache;

/**
 * @author joker
 * @When
 * @Description 高速缓存
 * 提供了不同的选择:
 * 1.可以基于软引用创建缓存,
 * 2.或者是通过弱引用创建缓存,
 * 底层也提供了不同的选择,
 * 可以通过CHM,也可以通过RingBuffer,
 * 也可以自定义 自定义只需要继承AbstractReferenceCache 实现自定义的底层结构即可
 *
 * 原理就是用某种容器存储数据,
 * 然后当内存不足的时候,gc roots为0的对象会被放入到ReferenceQueue中同时那个对象会被置为null
 * 这样get的时候判断是否为空,为空则表明要么本来就没有,要么是原先的过期了,
 * 然后判断队列是否为空,不为空则一个一个取出那么就会先将队列中的取出,然后从map中移除
 * @Detail
 * @date 创建时间 ：2019-02-01 06:16
 */

/*
    TODO:
    []  底层采用无锁的RingBuffer
 */
public interface IReferenceCache<K, V>
{
    /**
     * 当值不存在的时候会自动添加默认值
     * 并且应当提供一种创建对象的方式,提供一种key-value映射的关系
     * 如涉及业务的情况下:通过key查询对应的值findById等
     * @param key   the key
     * @param value the value
     */
    void put(K key, V value);

    /**
     * Get v.
     *
     * @param key the key
     * @return the v
     */
    V get(K key);


    /***
     * 当值不存在的时候不会再添加
     * @param key the key
     * @return v
     */
    V getIfAbsent(K key);
}

