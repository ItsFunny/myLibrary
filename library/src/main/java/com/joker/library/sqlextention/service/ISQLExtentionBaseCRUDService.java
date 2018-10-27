/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月27日 下午1:00:27
* 
*/
package com.joker.library.sqlextention.service;

import java.util.List;

import com.joker.library.page.PageBaseService;

/**
* 
* @When
* @Description	service层的公共接口
* @Detail
* @author joker 
* @date 创建时间：2018年10月27日 下午1:00:27
*/
public interface ISQLExtentionBaseCRUDService<T> extends PageBaseService<List<T>>
{
	int insert(T t);
}
