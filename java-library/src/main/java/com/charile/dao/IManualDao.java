package com.charile.dao;

import com.charile.dto.BatchInsertReq;
import org.apache.ibatis.annotations.*;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Map;
import java.util.TreeMap;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-04-06 14:21
 */
public interface IManualDao
{
    // for循环遍历插入
    @Insert("INSERT INTO ${tableName} ${columns} VALUES ${values}")
    void batchInsert(BatchInsertReq req);
    @Select("${sql}")
    Map<String, Object> selectByRawSql(@Param("sql") String sql);
    @Select(("${sql}"))
    List<TreeMap<String, Object>> selectMultiRecords(@Param("sql") String sql);
    @Delete(("${sql}"))
    void deleteByRawSql(@Param("sql") String sql);
}