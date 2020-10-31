package com.charlie.cache;

import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-04-06 14:49
 */
@Data
public class ObjectInfoThreadLocalCache
{
    private static final ThreadLocal<Object> objectContext = new ThreadLocal<>();

    public static void set(Object o)
    {
        objectContext.set(o);
    }

    public static Object get()
    {
        return objectContext.get();
    }

}
