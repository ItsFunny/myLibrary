package com.charlie.file;

import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-10 08:02
 */
@Data
public class UploadResponse
{
    // 存储路径
    private String storePath;
    // 映射路径
    private String mappingPath;

}
