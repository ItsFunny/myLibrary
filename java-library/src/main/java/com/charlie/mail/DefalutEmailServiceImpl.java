/**
 * 
 */
package com.charlie.mail;

import com.charlie.mail.property.EmailProperty;

/**
 * @author Administrator
 *
 */
public class DefalutEmailServiceImpl extends AbstractEmailService 
{

	@Override
	public void config(String smtp, String host, boolean auth,String sendMailAccount,String sendMailAccountPWD)
	{
		EmailProperty emailProperty=new EmailProperty();
		emailProperty.setSendEmailAccount(sendMailAccount);
		emailProperty.setSendEmailPwd(sendMailAccountPWD);
		emailProperty.setProtocol(smtp);
		emailProperty.setHost(host);
		emailProperty.setAuth(auth);
		super.setEmailProperty(emailProperty);
		super.config();
	}
	@Override
	public void config(EmailProperty emailProperty)
	{
		super.setEmailProperty(emailProperty);
		super.config();
	}
}
