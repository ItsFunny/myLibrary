package com.charile.event;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.context.event.SpringApplicationEvent;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-06-08 17:02
 */
public class WebAppEvent extends SpringApplicationEvent
{

    public WebAppEvent(SpringApplication application, String[] args)
    {
        super(application, args);
    }
}
