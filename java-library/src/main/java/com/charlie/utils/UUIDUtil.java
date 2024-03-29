package com.charlie.utils;

import org.apache.logging.log4j.core.util.UuidUtil;

import java.security.SecureRandom;

public class UUIDUtil
{

	private static SecureRandom random = new SecureRandom();

	/**
	 * UUID 包含mac地址与随机数，中间有-分割.
	 */
	public static String uuid() {
		return UuidUtil.getTimeBasedUuid().toString();
	}

	/**
	 * UUID 包含mac地址与随机数， 中间无-分割.
	 */
	public static String uuid2() {
		return UuidUtil.getTimeBasedUuid().toString().replaceAll("-", "");
	}

	/**
	 * 使用SecureRandom随机生成Long.
	 */
	public static long randomLong() {
		return Math.abs(random.nextLong());
	}

	public static String random22LengthString() {
		return RandomUtils.randomString(3) + randomLong();
	}

	public static void main(String[] args){
		String str = random22LengthString();
		System.out.println("str===" + str);
	}
}
