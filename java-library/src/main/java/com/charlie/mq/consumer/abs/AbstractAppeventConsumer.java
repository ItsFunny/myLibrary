/**
*
* @author joker 
* @date 创建时间：2018年9月8日 下午4:24:35
* 
*/
package com.charlie.mq.consumer.abs;

import com.charlie.mq.AppEventConsumer;

import lombok.Data;

/**
* 
* @author joker 
* @date 创建时间：2018年9月8日 下午4:24:35
*/
@Data
public abstract class AbstractAppeventConsumer implements AppEventConsumer
{
	private String type;
	

}
