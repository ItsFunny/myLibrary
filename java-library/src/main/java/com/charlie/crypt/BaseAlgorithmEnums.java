package com.charlie.crypt;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 09:53
 */
public enum BaseAlgorithmEnums
{

    ECDSA(1<<0,"ECDSA"),
    ;
    private  int baseType;
    private String   description;

    BaseAlgorithmEnums(int baseType, String description)
    {
        this.baseType = baseType;
        this.description = description;
    }}
