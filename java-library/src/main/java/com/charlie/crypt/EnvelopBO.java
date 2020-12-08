package com.charlie.crypt;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 08:55
 */
public class EnvelopBO
{
    // 原始数据,为json数据,可以是list
    private byte[] originData;
    private String envelopeIdentifier;
    private String extension;
    private String description;
    private EnumCertAlgorithm certAlgorithm;

    public EnumCertAlgorithm getCertAlgorithm()
    {
        return certAlgorithm;
    }

    public void setCertAlgorithm(EnumCertAlgorithm certAlgorithm)
    {
        this.certAlgorithm = certAlgorithm;
    }

    public byte[] getOriginData()
    {
        return originData;
    }

    public void setOriginData(byte[] originData)
    {
        this.originData = originData;
    }

    public String getEnvelopeIdentifier()
    {
        return envelopeIdentifier;
    }

    public void setEnvelopeIdentifier(String envelopeIdentifier)
    {
        this.envelopeIdentifier = envelopeIdentifier;
    }

    public String getExtension()
    {
        return extension;
    }

    public void setExtension(String extension)
    {
        this.extension = extension;
    }

    public String getDescription()
    {
        return description;
    }

    public void setDescription(String description)
    {
        this.description = description;
    }
}
