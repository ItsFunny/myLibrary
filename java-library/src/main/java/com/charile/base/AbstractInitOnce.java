package com.charile.base;

import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 23:13
 */
@Data
public abstract class AbstractInitOnce implements IInitOnce
{
    private boolean init;

    @Override
    public void initOnce() throws Exception
    {
        if (this.init)
        {
            return;
        }
        this.init();
        this.init = true;
    }

    protected abstract void init() throws Exception;

}
