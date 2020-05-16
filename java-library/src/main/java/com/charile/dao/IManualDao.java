package com.charile.dao;

import com.charile.dto.BatchInsertReq;
import org.apache.ibatis.annotations.Delete;
import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;

import java.util.Map;

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

    @Delete(("${sql}"))
    void deleteByRawSql(@Param("sql") String sql);
}