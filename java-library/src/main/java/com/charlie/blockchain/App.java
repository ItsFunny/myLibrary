package com.charlie.blockchain;

import lombok.Data;
import org.apache.commons.configuration.resolver.CatalogResolver;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-24 06:26
 */
public class App
{
    private static final Logger LOGGER = LoggerFactory.getLogger(App.class);

    public static void main(String[] args)
    {
        CatalogResolver.Catalog catalog=new CatalogResolver.Catalog();
    }



}
