package com.charlie.crypt;


import java.util.HashMap;
import java.util.Map;

public enum EnumCertAlgorithm
{
    SM2_256(1, "SM2_256"), RSA_1024(2, "RSA_1024"), RSA_2048(3, "RSA_2048");
    private int value;
    private String desc;

    private EnumCertAlgorithm(int value, String desc) {
        this.value = value;
        this.desc = desc;
    }

    public int getValue() {
        return this.value;
    }

    public String getDesc() {
        return this.desc;
    }

    private final static Map<Integer, EnumCertAlgorithm> ENUM_MAP = new HashMap<>();

    static {
        registerEnum(EnumCertAlgorithm.values());
    }

    public static EnumCertAlgorithm fromValue(int valueType) {
        EnumCertAlgorithm enm = ENUM_MAP.get(valueType);
        return enm;
    }

    protected static void registerEnum(EnumCertAlgorithm[] enums) {
        if (enums != null) {
            for (EnumCertAlgorithm enm : enums) {
                int key = enm.getValue();
                EnumCertAlgorithm old = ENUM_MAP.put(key, enm);
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