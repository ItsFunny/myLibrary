package com.charlie;

import com.charlie.blockchain.App;
import com.charlie.utils.JSONUtil;
import com.charlie.utils.RandomUtils;
import com.charlie.utils.UUIDUtil;
import lombok.Data;
import org.apache.commons.configuration.resolver.CatalogResolver;
import org.bouncycastle.pqc.math.linearalgebra.RandUtils;

import javax.xml.ws.Response;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.UUID;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2021-01-26 18:46
 */
@Data
public class CataLog
{
    public static void main(String[] args)
    {
        CataLog cataLog = new CataLog();
        cataLog.setId(UUIDUtil.uuid());
        cataLog.setName(UUIDUtil.uuid());
        cataLog.setEditVersion("editVersion_" + UUIDUtil.uuid());
        cataLog.setShowVersion("showVersion_" + UUIDUtil.uuid());

        Node node = new Node();
        node.setId("nodeId_");
        node.setName("nodeName");
        node.setRegion("nodeRegion");
        cataLog.setOwnerNodeId(node);


        cataLog.setDataItemList(buildItems(1));

        cataLog.setGlossaryList(buildGlossaries(3));

        cataLog.setTags(buildTAGS(3));

        cataLog.setConstraintList(buildContsas(3));

        cataLog.setState(CataLogState.abolishState);

        cataLog.setPublishTime(System.currentTimeMillis());



        System.out.println(JSONUtil.toFormattedJson(cataLog));
    }

    private static List<ForeignKeyConstraint> buildContsas(int limit)
    {
        List<ForeignKeyConstraint> res = new ArrayList<>();
        for (int i = 0; i < limit; i++)
        {
            ForeignKeyConstraint c = new ForeignKeyConstraint();
            c.setCataLogId(RandomUtils.randomString(3));
            c.setDataItemId(RandomUtils.randomString(10));
            c.setEditVersion(RandomUtils.randomString(3));
            c.setId(RandomUtils.randomString(5));
            res.add(c);
        }
        return res;
    }

    private static List<Tag> buildTAGS(int limit)
    {
        List<Tag> ttt = new ArrayList<>();
        for (int i = 0; i < limit; i++)
        {
            Tag tag = new Tag();
            tag.setCreateTagNodeId("123");
            tag.setDeleted(false);
            tag.setId(RandomUtils.randomString(5));
            tag.setName(RandomUtils.randomString(3));
            tag.setToWhichDataItemId(RandomUtils.randomString(8));
            tag.setType(RandomUtils.randomString(3));
            ttt.add(tag);
        }

        return ttt;
    }

    private static List<Glossary> buildGlossaries(int limit)
    {
        List<Glossary> glossaries = new ArrayList<>();
        for (int i = 0; i < limit; i++)
        {
            Glossary glossary = new Glossary();
            glossary.setEnglishName(RandomUtils.randomString(3));
            glossary.setName("name_" + RandomUtils.randomString(4));
            glossary.setValueList(Arrays.asList("123", "346"));
            glossaries.add(glossary);
        }

        return glossaries;
    }

    public static List<DataItem> buildItems(int limit)
    {
        List<DataItem> dataItems = new ArrayList<>();
        for (int i = 0; i < limit; i++)
        {
            dataItems.add(buildDataItem());
        }
        return dataItems;
    }

    private static DataItem buildDataItem()
    {
        DataItem dataItem = new DataItem();
        dataItem.setId("dataItemId_" + RandomUtils.randomString(5));
        dataItem.setName("dataname_" + RandomUtils.randomString(5));
        dataItem.setWeight(RandomUtils.randomInt(1, 10) + "");
        dataItem.setApiDown(buildAPI());
        dataItem.setApiUp(buildAPI());
        return dataItem;
    }

    private static DataItemAPI buildAPI()
    {
        DataItemAPI api = new DataItemAPI();
        api.setCodeList(buildCodes());
        api.setId(RandomUtils.randomString(6));
        api.setMethod("post");
        api.setReqList(buildFiles(RandomUtils.randomInt(1, 10)));
        api.setRespList(buildFiles(RandomUtils.randomInt(1, 10)));
        api.setUrl("asddd");
        api.setEncryptedData("asddddddddd");
        return api;
    }

    private static List<DataField> buildFiles(int count)
    {
        List<DataField> fields = new ArrayList<>();

        for (int i = 0; i < count; i++)
        {
            DataField field = new DataField();
            int l = RandomUtils.randomInt(1, 10);
            field.setEnglishName("eng_" + RandomUtils.randomString(l));
            field.setLength(l + 4 + "");
            field.setName("name_" + RandomUtils.randomString(l));
            field.setType("type");
            field.setWeight("8");
            fields.add(field);
        }


        return fields;
    }

    private static List<ResponseCode> buildCodes()
    {
        List<ResponseCode> codes = new ArrayList<>();
        ResponseCode code1 = new ResponseCode();
        code1.setCode(RandomUtils.randomInt(1, 10) + "");
        code1.setDescription(RandomUtils.randomString(5));
        codes.add(code1);

        ResponseCode code2 = new ResponseCode();
        code2.setCode(RandomUtils.randomInt(1, 10) + "");
        code2.setDescription(RandomUtils.randomString(6));
        codes.add(code2);

        return codes;
    }


    private String id;
    private String name;
    private String editVersion;
    private String showVersion;
    private Node ownerNodeId;
    // 数据项
    private List<DataItem> dataItemList;
    // 编码集
    private List<Glossary> glossaryList;

    private List<Tag> tags;
    private List<ForeignKeyConstraint> constraintList;
    private CataLogState state;
    private Long publishTime;

    @Data
    static class ForeignKeyConstraint
    {
        private String id;
        private String cataLogId;
        private String editVersion;
        private String dataItemId;
    }

    enum CataLogState
    {
        //草稿状态 链上不记录草稿
        draftState,
        //公示中状态
        publicityState,
        //现行标准状态
        workingStandardState,
        //废止状态
        abolishState,
        ;
    }

    @Data
    static class Node
    {
        private String id;
        private String name;
        private String region;
    }

    @Data
    static class DataItem
    {
        private String id;
        private String name;
        private DataItemAPI apiUp;
        private DataItemAPI apiDown;
        private String weight;
    }

    @Data
    static class DataItemAPI
    {
        private String id;
        private String method;
        private String url;
        private List<DataField> reqList;
        private Object encryptedData;
        private List<DataField> respList;
        private List<ResponseCode> codeList;

    }

    @Data
    static class ResponseCode
    {
        private String code;
        private String description;
    }

    @Data
    static class DataField
    {
        private String name;
        private String englishName;
        private boolean isRequired;
        private String weight;
        private String type;
        private String length;
    }

    @Data
    static class Tag
    {
        private String id;
        private String name;
        private String type;
        private String createTagNodeId;
        private String toWhichDataItemId;
        private boolean isDeleted;
    }

    @Data
    static class Glossary
    {
        private String name;
        private String englishName;
        private List<String> valueList;
    }
}
