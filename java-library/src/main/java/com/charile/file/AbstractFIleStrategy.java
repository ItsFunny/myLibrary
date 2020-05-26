/**
 * @author joker
 * @date 创建时间：2018年8月9日 上午9:12:41
 */
package com.charile.file;

import java.io.File;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;

import com.charile.exception.FTPConfigException;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;

/**
 * @author joker
 * @date 创建时间：2018年8月9日 上午9:12:41
 */
@Data
public abstract class AbstractFIleStrategy implements IFileStrategy
{

    //访问方式的前缀,如默认/imgs 等 ftp形式则是ip地址或者域名
    /*
     * 访问资源的前缀
     * 默认的单系统形式:/imgs  imgs在webapp下
     * 而ftp: /asdd   通常是nginx 做了反向代理之后的地址
     */

    private Map<String, String> basePathMap;

    private Map<String, String> visitPrefixMap;

    @Override
    public BatchFileUploadWrapperResp batchUpload(BatchFileUploadWrapper wrapper)
    {
        if (wrapper == null) throw new RuntimeException("参数不可为空");
        wrapper.valid();

        return this.doBatchUpload(wrapper);
    }

    protected abstract BatchFileUploadWrapperResp doBatchUpload(BatchFileUploadWrapper wrapper);

    // 如果末尾以 / 结尾,则截断
    public void setBasePathMap(Map<String, String> basePathMap)
    {
        if (basePathMap == null)
        {
            throw new FTPConfigException("路径定义不可为空");
        }
        validMap(basePathMap);
        this.basePathMap = basePathMap;
    }


    public void setVisitPrefixMap(Map<String, String> visitPrefixMap)
    {
        if (visitPrefixMap == null)
        {
            throw new FTPConfigException("mapping路径定义不可为空");
        }
        validMap(visitPrefixMap);
        this.visitPrefixMap = visitPrefixMap;
    }

    private void validMap(Map<String, String> map)
    {
        Set<String> strings = map.keySet();
        for (String key : strings)
        {
            String value = map.get(key);
            if (File.separator.equals(value.charAt(value.length() - 1)))
            {
                map.put(key, value.substring(0, value.length() - 1));
            }
        }
    }

    public AbstractFIleStrategy()
    {
        this.basePathMap = new HashMap<>();
        this.visitPrefixMap = new HashMap<>();
    }


    @Override
    public String getStoreBasePath(String key)
    {
        String res = this.basePathMap.get(key);
        return res;
    }


    @Override
    public String getVisitPrefix(String key)
    {
        String prefix = this.visitPrefixMap.get(key);
        if (StringUtils.isAllBlank(prefix))
        {
            throw new RuntimeException(key + "对应的访问前缀地址不存在");
        }
        return prefix;
    }
}
