package com.charile.bigfile;

import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description 文件信息
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-28 14:46
 */
@Data
public class FileInfo
{
    // 文件存储文件夹
    private String fileStoreDirectory;

    // 是否已经上传成功
    private boolean uploaded;
}
