package com.charlie.bigfile;

import java.io.IOException;

/**
 * @author Charlie
 * @When
 * @Description chunk处理者
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-28 14:30
 */
public interface FileChunkProcessor
{
    ProcessResp process(FileChunkDecorator decorator) throws IOException;



}
