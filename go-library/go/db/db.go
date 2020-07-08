/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-11-13 17:09 
# @File : db.go
# @Description : 
*/
package db

import (
	"strings"
)

type BaseBatchInsertTableBO struct {
	TableName string
	Colunms   []string
	ParamsLen int
}

func BuildBatchInsertExecSQL(bo BaseBatchInsertTableBO) string {
	sql := "INSERT INTO " + bo.TableName
	col := len(bo.Colunms)
	sql += "( "
	for i := 0; i < col-1; i++ {
		sql += bo.Colunms[i] + ", "
	}
	sql += bo.Colunms[col-1] + " ) VALUES  "
	recordNumbers := bo.ParamsLen / len(bo.Colunms)
	for i := 0; i < recordNumbers-1; i++ {
		sql += " ( "
		sql += strings.Repeat("?,", col-1)
		sql += " ? ) ,"
	}
	sql += " ( "
	sql += strings.Repeat("?,", col-1)
	sql += " ? ) "


	return sql
}

// orAnd  or 或者是and  ,用于条件查询
// judgeStr   like 或者是=  =通常适用于整型
// conditions 是否加入查询条件的主要依据,true才会加入,长度与剩余参数必须一致
// colums  对应数据表的列
// values 对应的查询条件
func CombineOrConditionSQL(orAnd string, judgeStr string, conditions []bool, colums []string, values []interface{}) (string, []interface{}) {
	res := " ( "
	need := false
	l := len(conditions)
	params := make([]interface{}, 0)
	for i := l - 1; i >= 0; i-- {
		if !conditions[i] {
			continue
		}

		// a | a,b | a,b,c
		if need {
			res += " " + orAnd + " "
		} else if i-1 >= 0 {
			need = true
		}

		strInterface := values[i]
		switch strInterface.(type) {
		case string:
			t := strInterface.(string)
			t += "%"
			params = append(params, t)
			res += colums[i] + " " + judgeStr + " ? "
		case int, int32, int64, int8, uint32, uint64, byte:
			// t := strconv.Itoa(strInterface.(int))
			params = append(params, strInterface)
			res += colums[i] + " " + " = " + " ? "
		case []int:
			t := strInterface.([]int)

			spliceInSQL, spliceParams := SpliceInSQL(colums[i], t)
			if len(spliceParams) > 0 {
				params = append(params, spliceParams...)
				res += spliceInSQL
			}
		}
	}
	res += ")"
	return res, params
}

type BuildOrderDeleteSqlReqByPrKey struct {
	TableNames  []string
	Colunms     []string
	PrimKeyList [][]int
}

//  @User 吕聪
// 	@Description  用于线性执行删除任务,现仅支持删除
//                批量线性删除(事务),暂且仅适用于一个表通过匹配一个字段删除
//                参数校验必须确保正常
//  @Param tableNames 数据库表名list
//  @Param colunms 对应的字段匹配名称
//  @Param values 匹配的校验字段,列数代表的是表数
//  @Return []string 返回的多条sql
//  @Return [][]interface 返回的参数集合, 行数代表[]string中的长度
func BuildOrderedDeletedSql(req BuildOrderDeleteSqlReqByPrKey) ([]string, [][]interface{}) {
	tableNames := req.TableNames
	colunms := req.Colunms
	idList2Array := req.PrimKeyList
	l := len(tableNames)
	sqls := make([]string, 0)
	params := make([][]interface{}, 0)
	s := "DELETE FROM "
	for i := 0; i < l; i++ {
		s += tableNames[i] + " WHERE "
		sql, p := SpliceDeleteSQL(tableNames[i], colunms[i], idList2Array[i])
		sqls = append(sqls, s+sql)
		params = append(params, p)
	}
	return sqls, params
}

// 确保不可为空交给上层
func SpliceDeleteSQL(tableName, colunmName string, idList []int) (string, []interface{}) {
	sql := "DELETE FROM " + tableName + " WHERE "
	s, i := SpliceInSQL(colunmName, idList)
	sql += s
	return s, i
}

// 嵌套删除
// func MultileOrderedDeleted(tx services.IBaseDB, log *log.Log, tableColunms map[string]map[string]*models2.NestedDeleteModel) error {
// 	params := make([]interface{}, 0)
// 	sql := "DELETE FROM ? WHERE ? IN (SELECT ? FROM ?  %s)"
// 	t := ""
// 	tParams := make([]interface{}, 0)
// 	for tableName, nestedMap := range tableColunms {
// 		params = append(params, tableName)
// 		for colunmName, model := range nestedMap {
// 			params = append(params, colunmName)
// 			params = append(params, model.TableName)
// 			params = append(params, model.ColunName)
// 			if len(model.IdList) > 0 {
// 				t, tParams = SpliceInSQL(model.ColunName, model.IdList)
// 				t = " WHERE " + t
// 				sql = fmt.Sprintf(sql, t)
// 				params = append(params, tParams)
// 			}
// 		}
// 		if e := tx.Raw(sql).Err(); nil != e {
// 			log.Error(fmt.Sprintf("[OrderedDeleted]批量线性删除%s的时候失败:%v", tableName, e))
// 			return e
// 		}
// 		tp := &params
// 		params = (*tp)[0:0]
// 	}
// 	return nil
// }
