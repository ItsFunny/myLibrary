package com.charlie.crypt;


import java.util.HashMap;
import java.util.Map;

public enum EnumSymmetryEncryptionType
{
    AES_ECB(1, "AES ECB模式"), AES_CBC(2, "AES CBC模式"), AES_CFB(3, "AES CFB模式"), AES_CFB_256K(4, "AES CFB 256K分片模式");
    private int value;
    private String desc;

    private EnumSymmetryEncryptionType(int value, String desc) {
        this.value = value;
        this.desc = desc;
    }

    public int getValue() {
        return this.value;
    }

    public String getDesc() {
        return this.desc;
    }

    private final static Map<Integer, EnumSymmetryEncryptionType> ENUM_MAP = new HashMap<>();

    static {
        registerEnum(EnumSymmetryEncryptionType.values());
    }

    public static EnumSymmetryEncryptionType fromValue(int valueType) {
        EnumSymmetryEncryptionType enm = ENUM_MAP.get(valueType);
        return enm;
    }

    protected static void registerEnum(EnumSymmetryEncryptionType[] enums) {
        if (enums != null) {
            for (EnumSymmetryEncryptionType enm : enums) {
                int key = enm.getValue();
                EnumSymmetryEncryptionType old = ENUM_MAP.put(key, enm);
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