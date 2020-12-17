package com.charlie.constants;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-03 15:02
 */
public class TransactionValidCode
{
    /**
     * <code>VALID = 0;</code>
     */
    public static final int VALID_VALUE = 0;
    /**
     * <code>NIL_ENVELOPE = 1;</code>
     */
    public static final int NIL_ENVELOPE_VALUE = 1;
    /**
     * <code>BAD_PAYLOAD = 2;</code>
     */
    public static final int BAD_PAYLOAD_VALUE = 2;
    /**
     * <code>BAD_COMMON_HEADER = 3;</code>
     */
    public static final int BAD_COMMON_HEADER_VALUE = 3;
    /**
     * <code>BAD_CREATOR_SIGNATURE = 4;</code>
     */
    public static final int BAD_CREATOR_SIGNATURE_VALUE = 4;
    /**
     * <code>INVALID_ENDORSER_TRANSACTION = 5;</code>
     */
    public static final int INVALID_ENDORSER_TRANSACTION_VALUE = 5;
    /**
     * <code>INVALID_CONFIG_TRANSACTION = 6;</code>
     */
    public static final int INVALID_CONFIG_TRANSACTION_VALUE = 6;
    /**
     * <code>UNSUPPORTED_TX_PAYLOAD = 7;</code>
     */
    public static final int UNSUPPORTED_TX_PAYLOAD_VALUE = 7;
    /**
     * <code>BAD_PROPOSAL_TXID = 8;</code>
     */
    public static final int BAD_PROPOSAL_TXID_VALUE = 8;
    /**
     * <code>DUPLICATE_TXID = 9;</code>
     */
    public static final int DUPLICATE_TXID_VALUE = 9;
    /**
     * <code>ENDORSEMENT_POLICY_FAILURE = 10;</code>
     */
    public static final int ENDORSEMENT_POLICY_FAILURE_VALUE = 10;
    /**
     * <code>MVCC_READ_CONFLICT = 11;</code>
     */
    public static final int MVCC_READ_CONFLICT_VALUE = 11;
    /**
     * <code>PHANTOM_READ_CONFLICT = 12;</code>
     */
    public static final int PHANTOM_READ_CONFLICT_VALUE = 12;
    /**
     * <code>UNKNOWN_TX_TYPE = 13;</code>
     */
    public static final int UNKNOWN_TX_TYPE_VALUE = 13;
    /**
     * <code>TARGET_CHAIN_NOT_FOUND = 14;</code>
     */
    public static final int TARGET_CHAIN_NOT_FOUND_VALUE = 14;
    /**
     * <code>MARSHAL_TX_ERROR = 15;</code>
     */
    public static final int MARSHAL_TX_ERROR_VALUE = 15;
    /**
     * <code>NIL_TXACTION = 16;</code>
     */
    public static final int NIL_TXACTION_VALUE = 16;
    /**
     * <code>EXPIRED_CHAINCODE = 17;</code>
     */
    public static final int EXPIRED_CHAINCODE_VALUE = 17;
    /**
     * <code>CHAINCODE_VERSION_CONFLICT = 18;</code>
     */
    public static final int CHAINCODE_VERSION_CONFLICT_VALUE = 18;
    /**
     * <code>BAD_HEADER_EXTENSION = 19;</code>
     */
    public static final int BAD_HEADER_EXTENSION_VALUE = 19;
    /**
     * <code>BAD_CHANNEL_HEADER = 20;</code>
     */
    public static final int BAD_CHANNEL_HEADER_VALUE = 20;
    /**
     * <code>BAD_RESPONSE_PAYLOAD = 21;</code>
     */
    public static final int BAD_RESPONSE_PAYLOAD_VALUE = 21;
    /**
     * <code>BAD_RWSET = 22;</code>
     */
    public static final int BAD_RWSET_VALUE = 22;
    /**
     * <code>ILLEGAL_WRITESET = 23;</code>
     */
    public static final int ILLEGAL_WRITESET_VALUE = 23;
    /**
     * <code>INVALID_WRITESET = 24;</code>
     */
    public static final int INVALID_WRITESET_VALUE = 24;
    /**
     * <code>NOT_VALIDATED = 254;</code>
     */
    public static final int NOT_VALIDATED_VALUE = 254;
    /**
     * <code>INVALID_OTHER_REASON = 255;</code>
     */
    public static final int INVALID_OTHER_REASON_VALUE = 255;

    public static String[] transactionValidCode = new String[]
            {"VALID", "NIL_ENVELOPE_VALUE", "BAD_PAYLOAD_VALUE", "BAD_COMMON_HEADER_VALUE",
                    "BAD_CREATOR_SIGNATURE_VALUE", "INVALID_ENDORSER_TRANSACTION_VALUE", "INVALID_CONFIG_TRANSACTION_VALUE", "UNSUPPORTED_TX_PAYLOAD_VALUE",
                    "BAD_PROPOSAL_TXID_VALUE", "DUPLICATE_TXID_VALUE", "ENDORSEMENT_POLICY_FAILURE_VALUE", "MVCC_READ_CONFLICT_VALUE",
                    "PHANTOM_READ_CONFLICT_VALUE", "UNKNOWN_TX_TYPE_VALUE", "TARGET_CHAIN_NOT_FOUND_VALUE", "MARSHAL_TX_ERROR_VALUE",
                    "NIL_TXACTION_VALUE", "EXPIRED_CHAINCODE_VALUE", "CHAINCODE_VERSION_CONFLICT_VALUE", "BAD_HEADER_EXTENSION_VALUE",
                    "BAD_CHANNEL_HEADER_VALUE", "BAD_RESPONSE_PAYLOAD_VALUE", "BAD_RWSET_VALUE", "ILLEGAL_WRITESET_VALUE",
                    "INVALID_WRITESET_VALUE", "NOT_VALIDATED_VALUE", "INVALID_OTHER_REASON_VALUE"
            };
}
