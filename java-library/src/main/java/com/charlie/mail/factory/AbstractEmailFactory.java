/**
*
* @author joker 
* @date 创建时间：2018年6月10日 下午5:38:11
* 
*/
package com.charlie.mail.factory;

import com.charlie.mail.IEmailService;
import com.charlie.mail.property.EmailProperty;

/**
* 
* @author joker 
* @date 创建时间：2018年6月10日 下午5:38:11
*/
public abstract class AbstractEmailFactory
{
	public abstract IEmailService genereateEmailService(EmailProperty emailProperty);
	
}
