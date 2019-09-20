/**
*
* @author joker 
* @date 创建时间：2018年2月2日 下午3:54:17
* 
*/
package com.charile.QueueLogin.service;

import com.charile.QueueLogin.QueueConfig;
import com.charile.QueueLogin.Exception.QueueLoginRangeException;
import com.charile.QueueLogin.model.QueueLoginResultInfo;

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
