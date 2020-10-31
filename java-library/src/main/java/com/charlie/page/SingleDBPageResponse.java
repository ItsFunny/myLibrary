package com.charlie.page;

import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-02-21 15:48
 */
@Data
public class SingleDBPageResponse
{
    private Object data;
    private Integer pageSize;
    private Integer pageNum;
    private Long totalCount;
    private Integer maxPage;

    public SingleDBPageResponse()
    {
        super();
    }

    public SingleDBPageResponse(Object data, Integer pageSize, Integer pageNum, Long totalCount)
    {
        this.data = data;
        this.pageSize = pageSize;
        this.pageNum = pageNum;
        this.totalCount = totalCount;
        this.maxPage = (int) ((totalCount % pageSize) == 0 ? totalCount / pageSize : (totalCount / pageSize) + 1);
    }

    public void setPageInfo(Integer pageSize, Integer pageNum, Long totalCount)
    {
        this.pageSize = pageSize;
        this.pageNum = pageNum;
        this.totalCount = totalCount;
        this.maxPage = (int) ((totalCount % pageSize) == 0 ? totalCount / pageSize : (totalCount / pageSize) + 1);
    }


}
