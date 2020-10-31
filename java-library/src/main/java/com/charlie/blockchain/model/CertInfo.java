package com.charlie.blockchain.model;

import lombok.Data;

@Data
public class CertInfo
{
	private String serialNum;
	private Long fromTime;
	private Long toTime;
	private String issuer;
	private String pubKey;
	private String version;
	private String signatureAlgorithm;
	private String subject;

	public CertInfo() {
		super();
	}

	public CertInfo(String serialNum, Long fromTime, Long toTime, String issuer, String pubKey, String version,
                    String signatureAlgorithm, String subject) {
		super();
		this.serialNum = serialNum;
		this.fromTime = fromTime;
		this.toTime = toTime;
		this.issuer = issuer;
		this.pubKey = pubKey;
		this.version = version;
		this.signatureAlgorithm = signatureAlgorithm;
		this.subject = subject;
	}

}
