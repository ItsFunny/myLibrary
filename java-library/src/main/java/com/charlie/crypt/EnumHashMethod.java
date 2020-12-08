package com.charlie.crypt;


import java.util.HashMap;
import java.util.Map;

public enum EnumHashMethod
{
    SM3(1, "SM3算法"), MD5(2, "MD5算法"), SHA1(3, "SHA1算法"), SHA256(4, "SHA256算法"), SHA384(5, "SHA384算法"), SHA512(6, "SHA512算法"), SHA256_256K(7, "SHA256_256K分片算法");
    private int value;
    private String desc;

    private EnumHashMethod(int value, String desc) {
        this.value = value;
        this.desc = desc;
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