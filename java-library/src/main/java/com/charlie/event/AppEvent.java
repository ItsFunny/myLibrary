package com.charlie.event;

import lombok.Data;

import java.util.EventObject;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-06-08 17:02
 */
@Data
public class AppEvent extends EventObject
{
    public AppEvent(Object source)
    {
        super(source);
    }
}
