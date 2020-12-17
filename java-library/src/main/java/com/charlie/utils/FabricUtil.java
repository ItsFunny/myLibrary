package com.charlie.utils;

import com.charlie.constants.TransactionValidCode;
import com.google.protobuf.InvalidProtocolBufferException;
import org.hyperledger.fabric.protos.ledger.rwset.kvrwset.KvRwset;
import org.hyperledger.fabric.sdk.BlockInfo;
import org.hyperledger.fabric.sdk.TxReadWriteSetInfo;

import java.io.UnsupportedEncodingException;
import java.util.*;

import static org.hyperledger.fabric.sdk.BlockInfo.EnvelopeType.TRANSACTION_ENVELOPE;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-03 15:02
 */
public class FabricUtil
{
    public static List<Map> getRWSetFromBlock(BlockInfo blockInfo) throws InvalidProtocolBufferException, UnsupportedEncodingException
    {
        List<Map> transactionList = new ArrayList<>();
        for (BlockInfo.EnvelopeInfo envelopeInfo : blockInfo.getEnvelopeInfos())
        {
            String id = envelopeInfo.getCreator().getId();
            String mspid = envelopeInfo.getCreator().getMspid();

            if (envelopeInfo.getType() == TRANSACTION_ENVELOPE)
            {
                Date timestamp = envelopeInfo.getTimestamp();
                BlockInfo.TransactionEnvelopeInfo transactionEnvelopeInfo = (BlockInfo.TransactionEnvelopeInfo) envelopeInfo;
                String transactionID = transactionEnvelopeInfo.getTransactionID();
                boolean valid = transactionEnvelopeInfo.isValid();
                byte validationCode = transactionEnvelopeInfo.getValidationCode();
                for (BlockInfo.TransactionEnvelopeInfo.TransactionActionInfo transactionActionInfo : transactionEnvelopeInfo.getTransactionActionInfos())
                {
                    Map<String, Object> transactionMap = new HashMap<>();
                    transactionMap.put("transactionID", transactionID);
                    transactionMap.put("timestamp", timestamp.getTime());
                    transactionMap.put("isValid", valid);
                    transactionMap.put("MSPID", mspid);
                    transactionMap.put("usercert", id);
                    transactionMap.put("validationCode", validationCode);
                    int index = validationCode;
                    if (index <= 24)
                    {
                        transactionMap.put("validationCodeName", TransactionValidCode.transactionValidCode[index]);
                    } else
                    {
                        transactionMap.put("validationCodeName", TransactionValidCode.transactionValidCode[index - 229]);
                    }
                    int chaincodeInputArgsCount = transactionActionInfo.getChaincodeInputArgsCount();
                    String[] argus = new String[chaincodeInputArgsCount];
                    for (int i = 0; i < chaincodeInputArgsCount; i++)
                    {
                        argus[i] = new String(transactionActionInfo.getChaincodeInputArgs(i));
                    }
                    transactionMap.put("argus", argus);
                    transactionMap.put("status", transactionActionInfo.getResponseStatus());
                    transactionMap.put("endorsementsCount", transactionActionInfo.getEndorsementsCount());
                    String chaincodeIDName = transactionActionInfo.getChaincodeIDName();
                    transactionMap.put("chaincodeName", chaincodeIDName);
                    String chaincodeIDVersion = transactionActionInfo.getChaincodeIDVersion();
                    transactionMap.put("chaincodeVersion", chaincodeIDVersion);
                    TxReadWriteSetInfo rwsetInfo = transactionActionInfo.getTxReadWriteSet();
                    if (null != rwsetInfo)
                    {
                        List<Map> rwList = new ArrayList<Map>();

                        for (TxReadWriteSetInfo.NsRwsetInfo nsRwsetInfo : rwsetInfo.getNsRwsetInfos())
                        {

                            Map<String, Object> rwMap = new HashMap<>();
                            Map<String, String> writeMap = new HashMap<>();
                            KvRwset.KVRWSet rws = nsRwsetInfo.getRwset();
                            String[] readSet = new String[rws.getReadsCount()];
                            int i = 0;
                            for (KvRwset.KVRead readList : rws.getReadsList())
                            {
                                String key = readList.getKey();
                                readSet[i++] = key;
                            }
                            rwMap.put("read", readSet);
                            for (KvRwset.KVWrite writeList : rws.getWritesList())
                            {
                                String valAsString = printableString(new String(writeList.getValue().toByteArray(), "UTF-8"));
                                writeList.getKey();
                                writeMap.put(writeList.getKey(), valAsString);
                            }
                            rwMap.put("write", writeMap);
                            rwList.add(rwMap);

                        }
                        transactionMap.put("RWSet", rwList);
                    }
                    transactionList.add(transactionMap);
                }
            }

        }
        return transactionList;
    }


    static String printableString(String string)
    {
        int maxLogStringLength = 64;
        if (string == null || string.length() == 0)
        {
            return string;
        }
        String ret = string.replaceAll("[^\\p{Print}]", "?");
        ret = ret.substring(0, Math.min(ret.length(), maxLogStringLength)) + (ret.length() > maxLogStringLength ? "..." : "");
        return ret;

    }




}
