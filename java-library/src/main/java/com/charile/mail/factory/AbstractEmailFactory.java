/**
*
* @author joker 
* @date 创建时间：2018年6月10日 下午5:38:11
* 
*/
package com.charile.mail.factory;

import com.charile.mail.IEmailService;
import com.charile.mail.property.EmailProperty;

/**
* 
* @author joker 
* @date 创建时间：2018年6月10日 下午5:38:11
*/
public abstract class AbstractEmailFactory
{
	public abstract IEmailService genereateEmailService(EmailProperty emailProperty);
	
}
