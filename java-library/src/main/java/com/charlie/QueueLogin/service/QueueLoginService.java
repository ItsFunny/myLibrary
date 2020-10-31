/**
*
* @author joker 
* @date 创建时间：2018年2月2日 下午3:54:17
* 
*/
package com.charlie.QueueLogin.service;

import com.charlie.QueueLogin.QueueConfig;
import com.charlie.QueueLogin.Exception.QueueLoginRangeException;
import com.charlie.QueueLogin.model.QueueLoginResultInfo;

/**
 * 
 * @author joker
 * @date 创建时间：2018年2月2日 下午3:54:17
 */
public interface QueueLoginService
{
	Integer SUCESS=1;
	Integer WAITLOGIN=2;
	<T> QueueLoginResultInfo login(Object key, T t) throws QueueLoginRangeException;

	<T> void setQueueConfig(QueueConfig<T> config);

	<T> QueueConfig<T> getQueueConfig();

}
