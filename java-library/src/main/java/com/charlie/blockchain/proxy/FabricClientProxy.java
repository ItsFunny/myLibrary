package com.charlie.blockchain.proxy;

import com.charlie.base.AbstractInitOnce;
import com.charlie.exception.ConfigException;
import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description fabric client的代理类
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-23 22:55
 */
@Data
public class FabricClientProxy<T> extends AbstractInitOnce
{
    /*
       这个是全局唯一的对象,所以,这个如果不是内部持有一个INSTANCE的话,则需要一个工厂类,然后再有一个holder类
       如 FabricClientProxyFactory 所示,需要先获取到PRoxyFactory => 再获取到 FabricClientProxy => 再内部接口
       如果此处内部持有对象的话: 就只需要getInstance即可
       2种方式都是可以的,不同在于 factory的形式 FabricClientProxy的构造方法不能私有化
       除非有多种proxy,则现在直接使用即可
     */

    private static FabricClientProxy INSTANTCE;

    private AbstractFabricClient handler;

    @Override
    protected void init() throws ConfigException
    {
    }


    public static void setInstance(FabricClientProxy proxy, boolean force)
    {
        if (INSTANTCE != null)
        {
            if (force)
            {
                INSTANTCE = proxy;
            }
        } else
        {
            INSTANTCE = proxy;
        }
    }

    public static FabricClientProxy getInstance()
    {
        return INSTANTCE;
    }

    public static FabricClientProxy getINSTANTCE()
    {
        return INSTANTCE;
    }

    public static void setINSTANTCE(FabricClientProxy INSTANTCE)
    {
        FabricClientProxy.INSTANTCE = INSTANTCE;
    }

    public AbstractFabricClient getHandler()
    {
        return handler;
    }

    public void setHandler(AbstractFabricClient handler)
    {
        this.handler = handler;
    }
}
