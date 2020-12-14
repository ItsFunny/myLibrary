package com.charlie.crypt;


import java.util.HashMap;
import java.util.Map;

public enum EnumHashMethod
{
    SM3(1, "SM3算法",EnumBaseType.ENUM_HASH_SM3.getValue()), MD5(2, "MD5算法",EnumBaseType.ENUM_HASH_MD5.getValue()), SHA1(3, "SHA1算法",EnumBaseType.ENUM_HASH_SHA.getValue()), SHA256(4, "SHA256算法",EnumBaseType.ENUM_HASH_SHA.getValue()), SHA384(5, "SHA384算法",EnumBaseType.ENUM_HASH_SHA.getValue()), SHA512(6, "SHA512算法",EnumBaseType.ENUM_HASH_SHA.getValue()), SHA256_256K(7, "SHA256_256K分片算法",EnumBaseType.ENUM_HASH_SHA.getValue());
    private int value;
    private String desc;
    private int baseType;

    private EnumHashMethod(int value, String desc,int baseType) {
        this.value = value;
        this.desc = desc;
        this.baseType=baseType;
    }

    public int getValue() {
        return this.value;
    }

    public String getDesc() {
        return this.desc;
    }

    private final static Map<Integer, EnumHashMethod> ENUM_MAP = new HashMap<>();

    static {
        registerEnum(EnumHashMethod.values());
    }

    public static EnumHashMethod fromValue(int valueType) {
        EnumHashMethod enm = ENUM_MAP.get(valueType);
        return enm;
    }

    protected static void registerEnum(EnumHashMethod[] enums) {
        if (enums != null) {
            for (EnumHashMethod enm : enums) {
                int key = enm.getValue();
                EnumHashMethod old = ENUM_MAP.put(key, enm);
                if (old != null) {
                    throw new RuntimeException("Repeated value:" + old.name());
                }
            }
        }
    }

    /********************************************************************************
     * 以下为自定义函数
     ********************************************************************************/

}