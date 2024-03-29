package com.charlie.base;

import com.charlie.constants.PemConstant;
import com.charlie.utils.Base64Utils;
import org.bouncycastle.openssl.jcajce.JcaPEMWriter;
import org.bouncycastle.util.io.pem.PemObject;
import org.bouncycastle.util.io.pem.PemWriter;

import java.io.IOException;
import java.io.StringWriter;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-05 14:35
 */
public interface IDataDecorator<T>
{
    byte[] decorate(T data) throws IOException;

    IDataDecorator<byte[]> CERTIFICATE_DECORATOR = (data) ->
    {
        String str = new String(data);
        data = Base64Utils.decode(str);

        StringWriter stringWriter = new StringWriter();
        JcaPEMWriter pemWriter = new JcaPEMWriter(stringWriter);
        pemWriter.writeObject(new PemObject(PemConstant.CERTIFICATE, data));
        pemWriter.close();
        return stringWriter.toString().getBytes();
    };
    IDataDecorator<byte[]> PRIVATEKEY_DECORATOR = (data) ->
    {
//        String str = new String(data);
//        data = Base64Util.decode(str);
        StringWriter stringWriter = new StringWriter();
        PemWriter pemWriter = new PemWriter(stringWriter);
        pemWriter.writeObject(new PemObject(PemConstant.PRIVATEKEY, data));
        pemWriter.close();
        return stringWriter.toString().getBytes();
    };

    IDataDecorator<byte[]> PUBLICKEY_DECORATOR = (data) ->
    {
//        String str = new String(data);
//        data = Base64Util.decode(str);

        StringWriter stringWriter = new StringWriter();
        PemWriter pemWriter = new PemWriter(stringWriter);
        pemWriter.writeObject(new PemObject(PemConstant.PUBLICKEY, data));
        pemWriter.close();
        return stringWriter.toString().getBytes();
    };

    IDataDecorator<byte[]> BASE64_DECORATOR = (data) -> Base64Utils.encode(data).getBytes();

    IDataDecorator<byte[]> ORIGIN_BYTES = (data) ->
    {
        StringWriter stringWriter = new StringWriter();
        PemWriter pemWriter = new PemWriter(stringWriter);
        pemWriter.writeObject(new PemObject(PemConstant.PUBLICKEY, data));
        pemWriter.close();
        return stringWriter.toString().getBytes();
    };


}
