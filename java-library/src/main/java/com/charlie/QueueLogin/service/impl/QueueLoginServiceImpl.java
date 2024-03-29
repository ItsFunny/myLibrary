/**
*
* @author joker 
* @date 创建时间：2018年2月2日 下午4:48:38
* 
*/
package com.charlie.QueueLogin.service.impl;

import java.util.Collections;
import java.util.HashMap;
import java.util.Map;

import com.charlie.QueueLogin.AbstractLoginServiceImpl;
import com.charlie.QueueLogin.QueueConfig;
import com.charlie.QueueLogin.Exception.QueueLoginRangeException;
import com.charlie.QueueLogin.model.QueueLoginResultInfo;

/**
* 
* @author joker 
 * @param <T>
* @date 创建时间：2018年2月2日 下午4:48:38
*/
@SuppressWarnings("rawtypes")
public class QueueLoginServiceImpl extends AbstractLoginServiceImpl
{
	@SuppressWarnings("unchecked")
	@Override
	public <T> QueueLoginResultInfo login(Object key, T t) throws QueueLoginRangeException
	{
		return super.enter(key, t);
	}
	@SuppressWarnings("unchecked")
	@Override
	public <T> void setQueueConfig(QueueConfig<T> config)
	{
		Map<String, String> synchronizedMap = Collections.synchronizedMap(new HashMap<String,String>());
		synchronizedMap.put("", "");
		super.config=config;
	}
	@SuppressWarnings("unchecked")
	@Override
	public <T> QueueConfig<T> getQueueConfig()
	{
		return super.config;
	}







	


}
