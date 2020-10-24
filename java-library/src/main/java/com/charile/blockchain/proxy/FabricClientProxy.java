package com.charile.blockchain.proxy;

import com.charile.base.AbstractInitOnce;
import com.charile.blockchain.exception.InvokeException;
import com.charile.blockchain.model.InstallChaincodeReq;
import com.charile.blockchain.model.InstallChaincodeResp;
import com.charile.blockchain.model.InvokeReq;
import com.charile.blockchain.model.InvokeResp;
import com.charile.exception.ConfigException;
import lombok.Data;

import java.util.List;
import java.util.concurrent.Future;

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

}
