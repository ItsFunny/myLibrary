/**
 * @author joker
 * @date 创建时间：2018年8月8日 下午4:45:59
 */
package com.charile.file;


import com.charile.service.IValidater;
import lombok.Data;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.io.InputStream;
import java.util.List;

/**
 * The interface File strategy.
 *
 * @author joker
 * @date 创建时间 ：2018年8月8日 下午4:45:59
 */
public interface IFileStrategy
{
    /**
     * The constant IMG_TYPE.
     */
    String IMG_TYPE = "img";


    /**
     * Upload upload resp.
     *
     * @param file        the file
     * @param storePath   the store path  不可以以 / 起始,并且当只有1个目录的时候,要删除后面的/
     * @param newFileName the new file name,新的名字必须包含后缀
     * @param key         the key
     * @return the upload resp: 包含绝对存储路径和映射路径
     * @throws IOException the io exception
     */
    UploadResponse upload(MultipartFile file, String storePath, String newFileName, String key) throws IOException;

    UploadResponse upload(InputStream inputStream, String storePath, String newFileName, String key) throws IOException;

    @Data
    class BatchFileUploadWrapper implements IValidater
    {
        private Integer uniqueKey;
        private List<InputStream> streams;
        private List<String> pathList;
        private List<String> newNames;
        private List<String> keys;

        @Override
        public void valid()
        {

        }
    }

    @Data
    class BatchFileUploadWrapperResp
    {

    }

    BatchFileUploadWrapperResp batchUpload(BatchFileUploadWrapper wrapper);

    /**
     * Delete boolean.
     *
     * @param filePathName the file path name
     * @return the boolean
     */
    Boolean delete(String filePathName);


    /**
     * Gets store base path.
     *
     * @param key the key
     * @return the store base path
     */
    String getStoreBasePath(String key);

    /**
     * Gets visit prefix.
     *
     * @param key the key
     * @return the visit prefix
     */
    String getVisitPrefix(String key);


    /**
     * The type Upload resp.
     */
}
