package com.charlie.crypt;


import java.util.HashMap;
import java.util.Map;

public enum EnumAsymmetricAlgorithm
{
    SM2_256(1, "SM2_256",EnumBaseType.ENUM_ASYMMETRIC_SM2.getValue()), RSA_1024(2, "RSA_1024",EnumBaseType.ENUM_ASYMMETRIC_RSA.getValue()), RSA_2048(3, "RSA_2048",EnumBaseType.ENUM_ASYMMETRIC_RSA.getValue());
    private int value;
    private String desc;
    private int baseType;


    private EnumAsymmetricAlgorithm(int value, String desc,int baseType) {
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

    private final static Map<Integer, EnumAsymmetricAlgorithm> ENUM_MAP = new HashMap<>();

    static {
        registerEnum(EnumAsymmetricAlgorithm.values());
    }

    public static EnumAsymmetricAlgorithm fromValue(int valueType) {
        EnumAsymmetricAlgorithm enm = ENUM_MAP.get(valueType);
        return enm;
    }

    protected static void registerEnum(EnumAsymmetricAlgorithm[] enums) {
        if (enums != null) {
            for (EnumAsymmetricAlgorithm enm : enums) {
                int key = enm.getValue();
                EnumAsymmetricAlgorithm old = ENUM_MAP.put(key, enm);
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