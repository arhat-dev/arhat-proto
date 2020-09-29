/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.2 */

#ifndef PB_ARHAT_DEVICE_PB_H_INCLUDED
#define PB_ARHAT_DEVICE_PB_H_INCLUDED
#include <pb.h>

#if PB_PROTO_HEADER_VERSION != 40
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Enum definitions */
typedef enum _arhat_DeviceCmdType {
    arhat_DeviceCmdType__INVALID_DEV_CMD = 0,
    arhat_DeviceCmdType_CMD_DEV_CONNECT = 1,
    arhat_DeviceCmdType_CMD_DEV_OPERATE = 2,
    arhat_DeviceCmdType_CMD_DEV_COLLECT_METRICS = 3,
    arhat_DeviceCmdType_CMD_DEV_CLOSE = 4
} arhat_DeviceCmdType;

typedef enum _arhat_DeviceMsgType {
    arhat_DeviceMsgType__INVALID_DEV_MSG = 0,
    arhat_DeviceMsgType_MSG_DEV_ERROR = 1,
    arhat_DeviceMsgType_MSG_DEV_DONE = 2,
    arhat_DeviceMsgType_MSG_DEV_REGISTER = 3,
    arhat_DeviceMsgType_MSG_DEV_OPERATION_RESULT = 4,
    arhat_DeviceMsgType_MSG_DEV_METRICS = 5,
    arhat_DeviceMsgType_MSG_DEV_EVENTS = 6
} arhat_DeviceMsgType;

typedef enum _arhat_DeviceEventType {
    arhat_DeviceEventType__INVALID_DEV_EVENT = 0
} arhat_DeviceEventType;

/* Struct definitions */
typedef struct _arhat_DeviceCloseCmd {
    char dummy_field;
} arhat_DeviceCloseCmd;

typedef struct _arhat_DeviceConnectCmd_ParamsEntry {
    pb_callback_t key;
    pb_callback_t value;
} arhat_DeviceConnectCmd_ParamsEntry;

typedef struct _arhat_DeviceDoneMsg {
    char dummy_field;
} arhat_DeviceDoneMsg;

typedef struct _arhat_DeviceMetricsCollectCmd {
    pb_callback_t params;
} arhat_DeviceMetricsCollectCmd;

typedef struct _arhat_DeviceMetricsCollectCmd_ParamsEntry {
    pb_callback_t key;
    pb_callback_t value;
} arhat_DeviceMetricsCollectCmd_ParamsEntry;

typedef struct _arhat_DeviceMetricsMsg {
    pb_callback_t values;
} arhat_DeviceMetricsMsg;

typedef struct _arhat_DeviceOperateCmd {
    pb_callback_t params;
    pb_callback_t data;
} arhat_DeviceOperateCmd;

typedef struct _arhat_DeviceOperateCmd_ParamsEntry {
    pb_callback_t key;
    pb_callback_t value;
} arhat_DeviceOperateCmd_ParamsEntry;

typedef struct _arhat_DeviceOperateResultMsg {
    pb_callback_t result;
} arhat_DeviceOperateResultMsg;

typedef struct _arhat_DeviceRegisterMsg {
    pb_callback_t name;
} arhat_DeviceRegisterMsg;

typedef struct _arhat_ErrorMsg {
    pb_callback_t description;
} arhat_ErrorMsg;

typedef struct _arhat_DeviceCmd {
    arhat_DeviceCmdType kind;
    uint64_t device_id;
    uint64_t seq;
    pb_callback_t payload;
} arhat_DeviceCmd;

typedef struct _arhat_DeviceEventMsg {
    arhat_DeviceEventType kind;
} arhat_DeviceEventMsg;

typedef struct _arhat_DeviceMetricsMsg_Value {
    double value;
    int64_t timestamp;
} arhat_DeviceMetricsMsg_Value;

typedef struct _arhat_DeviceMsg {
    arhat_DeviceMsgType kind;
    uint64_t device_id;
    uint64_t ack;
    pb_callback_t payload;
} arhat_DeviceMsg;

typedef struct _arhat_TLSConfig {
    pb_callback_t server_name;
    bool insecure_skip_verify;
    uint32_t min_version;
    uint32_t max_version;
    pb_callback_t ca_cert;
    pb_callback_t cert;
    pb_callback_t key;
    pb_callback_t cipher_suites;
    pb_callback_t next_protos;
} arhat_TLSConfig;

typedef struct _arhat_DeviceConnectCmd {
    pb_callback_t target;
    pb_callback_t params;
    bool has_tls;
    arhat_TLSConfig tls;
} arhat_DeviceConnectCmd;


/* Helper constants for enums */
#define _arhat_DeviceCmdType_MIN arhat_DeviceCmdType__INVALID_DEV_CMD
#define _arhat_DeviceCmdType_MAX arhat_DeviceCmdType_CMD_DEV_CLOSE
#define _arhat_DeviceCmdType_ARRAYSIZE ((arhat_DeviceCmdType)(arhat_DeviceCmdType_CMD_DEV_CLOSE+1))

#define _arhat_DeviceMsgType_MIN arhat_DeviceMsgType__INVALID_DEV_MSG
#define _arhat_DeviceMsgType_MAX arhat_DeviceMsgType_MSG_DEV_EVENTS
#define _arhat_DeviceMsgType_ARRAYSIZE ((arhat_DeviceMsgType)(arhat_DeviceMsgType_MSG_DEV_EVENTS+1))

#define _arhat_DeviceEventType_MIN arhat_DeviceEventType__INVALID_DEV_EVENT
#define _arhat_DeviceEventType_MAX arhat_DeviceEventType__INVALID_DEV_EVENT
#define _arhat_DeviceEventType_ARRAYSIZE ((arhat_DeviceEventType)(arhat_DeviceEventType__INVALID_DEV_EVENT+1))


/* Initializer values for message structs */
#define arhat_DeviceCmd_init_default             {_arhat_DeviceCmdType_MIN, 0, 0, {{NULL}, NULL}}
#define arhat_TLSConfig_init_default             {{{NULL}, NULL}, 0, 0, 0, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define arhat_DeviceConnectCmd_init_default      {{{NULL}, NULL}, {{NULL}, NULL}, false, arhat_TLSConfig_init_default}
#define arhat_DeviceConnectCmd_ParamsEntry_init_default {{{NULL}, NULL}, {{NULL}, NULL}}
#define arhat_DeviceOperateCmd_init_default      {{{NULL}, NULL}, {{NULL}, NULL}}
#define arhat_DeviceOperateCmd_ParamsEntry_init_default {{{NULL}, NULL}, {{NULL}, NULL}}
#define arhat_DeviceMetricsCollectCmd_init_default {{{NULL}, NULL}}
#define arhat_DeviceMetricsCollectCmd_ParamsEntry_init_default {{{NULL}, NULL}, {{NULL}, NULL}}
#define arhat_DeviceCloseCmd_init_default        {0}
#define arhat_DeviceMsg_init_default             {_arhat_DeviceMsgType_MIN, 0, 0, {{NULL}, NULL}}
#define arhat_DeviceDoneMsg_init_default         {0}
#define arhat_DeviceRegisterMsg_init_default     {{{NULL}, NULL}}
#define arhat_DeviceOperateResultMsg_init_default {{{NULL}, NULL}}
#define arhat_DeviceMetricsMsg_init_default      {{{NULL}, NULL}}
#define arhat_DeviceMetricsMsg_Value_init_default {0, 0}
#define arhat_DeviceEventMsg_init_default        {_arhat_DeviceEventType_MIN}
#define arhat_ErrorMsg_init_default              {{{NULL}, NULL}}
#define arhat_DeviceCmd_init_zero                {_arhat_DeviceCmdType_MIN, 0, 0, {{NULL}, NULL}}
#define arhat_TLSConfig_init_zero                {{{NULL}, NULL}, 0, 0, 0, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}, {{NULL}, NULL}}
#define arhat_DeviceConnectCmd_init_zero         {{{NULL}, NULL}, {{NULL}, NULL}, false, arhat_TLSConfig_init_zero}
#define arhat_DeviceConnectCmd_ParamsEntry_init_zero {{{NULL}, NULL}, {{NULL}, NULL}}
#define arhat_DeviceOperateCmd_init_zero         {{{NULL}, NULL}, {{NULL}, NULL}}
#define arhat_DeviceOperateCmd_ParamsEntry_init_zero {{{NULL}, NULL}, {{NULL}, NULL}}
#define arhat_DeviceMetricsCollectCmd_init_zero  {{{NULL}, NULL}}
#define arhat_DeviceMetricsCollectCmd_ParamsEntry_init_zero {{{NULL}, NULL}, {{NULL}, NULL}}
#define arhat_DeviceCloseCmd_init_zero           {0}
#define arhat_DeviceMsg_init_zero                {_arhat_DeviceMsgType_MIN, 0, 0, {{NULL}, NULL}}
#define arhat_DeviceDoneMsg_init_zero            {0}
#define arhat_DeviceRegisterMsg_init_zero        {{{NULL}, NULL}}
#define arhat_DeviceOperateResultMsg_init_zero   {{{NULL}, NULL}}
#define arhat_DeviceMetricsMsg_init_zero         {{{NULL}, NULL}}
#define arhat_DeviceMetricsMsg_Value_init_zero   {0, 0}
#define arhat_DeviceEventMsg_init_zero           {_arhat_DeviceEventType_MIN}
#define arhat_ErrorMsg_init_zero                 {{{NULL}, NULL}}

/* Field tags (for use in manual encoding/decoding) */
#define arhat_DeviceConnectCmd_ParamsEntry_key_tag 1
#define arhat_DeviceConnectCmd_ParamsEntry_value_tag 2
#define arhat_DeviceMetricsCollectCmd_params_tag 1
#define arhat_DeviceMetricsCollectCmd_ParamsEntry_key_tag 1
#define arhat_DeviceMetricsCollectCmd_ParamsEntry_value_tag 2
#define arhat_DeviceMetricsMsg_values_tag        1
#define arhat_DeviceOperateCmd_params_tag        1
#define arhat_DeviceOperateCmd_data_tag          2
#define arhat_DeviceOperateCmd_ParamsEntry_key_tag 1
#define arhat_DeviceOperateCmd_ParamsEntry_value_tag 2
#define arhat_DeviceOperateResultMsg_result_tag  1
#define arhat_DeviceRegisterMsg_name_tag         1
#define arhat_ErrorMsg_description_tag           1
#define arhat_DeviceCmd_kind_tag                 1
#define arhat_DeviceCmd_device_id_tag            2
#define arhat_DeviceCmd_seq_tag                  3
#define arhat_DeviceCmd_payload_tag              4
#define arhat_DeviceEventMsg_kind_tag            1
#define arhat_DeviceMetricsMsg_Value_value_tag   1
#define arhat_DeviceMetricsMsg_Value_timestamp_tag 2
#define arhat_DeviceMsg_kind_tag                 1
#define arhat_DeviceMsg_device_id_tag            2
#define arhat_DeviceMsg_ack_tag                  3
#define arhat_DeviceMsg_payload_tag              4
#define arhat_TLSConfig_server_name_tag          1
#define arhat_TLSConfig_insecure_skip_verify_tag 2
#define arhat_TLSConfig_min_version_tag          3
#define arhat_TLSConfig_max_version_tag          4
#define arhat_TLSConfig_ca_cert_tag              5
#define arhat_TLSConfig_cert_tag                 6
#define arhat_TLSConfig_key_tag                  7
#define arhat_TLSConfig_cipher_suites_tag        8
#define arhat_TLSConfig_next_protos_tag          9
#define arhat_DeviceConnectCmd_target_tag        1
#define arhat_DeviceConnectCmd_params_tag        2
#define arhat_DeviceConnectCmd_tls_tag           3

/* Struct field encoding specification for nanopb */
#define arhat_DeviceCmd_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, UENUM,    kind,              1) \
X(a, STATIC,   SINGULAR, UINT64,   device_id,         2) \
X(a, STATIC,   SINGULAR, UINT64,   seq,               3) \
X(a, CALLBACK, SINGULAR, BYTES,    payload,           4)
#define arhat_DeviceCmd_CALLBACK pb_default_field_callback
#define arhat_DeviceCmd_DEFAULT NULL

#define arhat_TLSConfig_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   server_name,       1) \
X(a, STATIC,   SINGULAR, BOOL,     insecure_skip_verify,   2) \
X(a, STATIC,   SINGULAR, UINT32,   min_version,       3) \
X(a, STATIC,   SINGULAR, UINT32,   max_version,       4) \
X(a, CALLBACK, SINGULAR, BYTES,    ca_cert,           5) \
X(a, CALLBACK, SINGULAR, BYTES,    cert,              6) \
X(a, CALLBACK, SINGULAR, BYTES,    key,               7) \
X(a, CALLBACK, REPEATED, UINT32,   cipher_suites,     8) \
X(a, CALLBACK, REPEATED, STRING,   next_protos,       9)
#define arhat_TLSConfig_CALLBACK pb_default_field_callback
#define arhat_TLSConfig_DEFAULT NULL

#define arhat_DeviceConnectCmd_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   target,            1) \
X(a, CALLBACK, REPEATED, MESSAGE,  params,            2) \
X(a, STATIC,   OPTIONAL, MESSAGE,  tls,               3)
#define arhat_DeviceConnectCmd_CALLBACK pb_default_field_callback
#define arhat_DeviceConnectCmd_DEFAULT NULL
#define arhat_DeviceConnectCmd_params_MSGTYPE arhat_DeviceConnectCmd_ParamsEntry
#define arhat_DeviceConnectCmd_tls_MSGTYPE arhat_TLSConfig

#define arhat_DeviceConnectCmd_ParamsEntry_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   key,               1) \
X(a, CALLBACK, SINGULAR, STRING,   value,             2)
#define arhat_DeviceConnectCmd_ParamsEntry_CALLBACK pb_default_field_callback
#define arhat_DeviceConnectCmd_ParamsEntry_DEFAULT NULL

#define arhat_DeviceOperateCmd_FIELDLIST(X, a) \
X(a, CALLBACK, REPEATED, MESSAGE,  params,            1) \
X(a, CALLBACK, SINGULAR, BYTES,    data,              2)
#define arhat_DeviceOperateCmd_CALLBACK pb_default_field_callback
#define arhat_DeviceOperateCmd_DEFAULT NULL
#define arhat_DeviceOperateCmd_params_MSGTYPE arhat_DeviceOperateCmd_ParamsEntry

#define arhat_DeviceOperateCmd_ParamsEntry_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   key,               1) \
X(a, CALLBACK, SINGULAR, STRING,   value,             2)
#define arhat_DeviceOperateCmd_ParamsEntry_CALLBACK pb_default_field_callback
#define arhat_DeviceOperateCmd_ParamsEntry_DEFAULT NULL

#define arhat_DeviceMetricsCollectCmd_FIELDLIST(X, a) \
X(a, CALLBACK, REPEATED, MESSAGE,  params,            1)
#define arhat_DeviceMetricsCollectCmd_CALLBACK pb_default_field_callback
#define arhat_DeviceMetricsCollectCmd_DEFAULT NULL
#define arhat_DeviceMetricsCollectCmd_params_MSGTYPE arhat_DeviceMetricsCollectCmd_ParamsEntry

#define arhat_DeviceMetricsCollectCmd_ParamsEntry_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   key,               1) \
X(a, CALLBACK, SINGULAR, STRING,   value,             2)
#define arhat_DeviceMetricsCollectCmd_ParamsEntry_CALLBACK pb_default_field_callback
#define arhat_DeviceMetricsCollectCmd_ParamsEntry_DEFAULT NULL

#define arhat_DeviceCloseCmd_FIELDLIST(X, a) \

#define arhat_DeviceCloseCmd_CALLBACK NULL
#define arhat_DeviceCloseCmd_DEFAULT NULL

#define arhat_DeviceMsg_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, UENUM,    kind,              1) \
X(a, STATIC,   SINGULAR, UINT64,   device_id,         2) \
X(a, STATIC,   SINGULAR, UINT64,   ack,               3) \
X(a, CALLBACK, SINGULAR, BYTES,    payload,           4)
#define arhat_DeviceMsg_CALLBACK pb_default_field_callback
#define arhat_DeviceMsg_DEFAULT NULL

#define arhat_DeviceDoneMsg_FIELDLIST(X, a) \

#define arhat_DeviceDoneMsg_CALLBACK NULL
#define arhat_DeviceDoneMsg_DEFAULT NULL

#define arhat_DeviceRegisterMsg_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   name,              1)
#define arhat_DeviceRegisterMsg_CALLBACK pb_default_field_callback
#define arhat_DeviceRegisterMsg_DEFAULT NULL

#define arhat_DeviceOperateResultMsg_FIELDLIST(X, a) \
X(a, CALLBACK, REPEATED, BYTES,    result,            1)
#define arhat_DeviceOperateResultMsg_CALLBACK pb_default_field_callback
#define arhat_DeviceOperateResultMsg_DEFAULT NULL

#define arhat_DeviceMetricsMsg_FIELDLIST(X, a) \
X(a, CALLBACK, REPEATED, MESSAGE,  values,            1)
#define arhat_DeviceMetricsMsg_CALLBACK pb_default_field_callback
#define arhat_DeviceMetricsMsg_DEFAULT NULL
#define arhat_DeviceMetricsMsg_values_MSGTYPE arhat_DeviceMetricsMsg_Value

#define arhat_DeviceMetricsMsg_Value_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, DOUBLE,   value,             1) \
X(a, STATIC,   SINGULAR, INT64,    timestamp,         2)
#define arhat_DeviceMetricsMsg_Value_CALLBACK NULL
#define arhat_DeviceMetricsMsg_Value_DEFAULT NULL

#define arhat_DeviceEventMsg_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, UENUM,    kind,              1)
#define arhat_DeviceEventMsg_CALLBACK NULL
#define arhat_DeviceEventMsg_DEFAULT NULL

#define arhat_ErrorMsg_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   description,       1)
#define arhat_ErrorMsg_CALLBACK pb_default_field_callback
#define arhat_ErrorMsg_DEFAULT NULL

extern const pb_msgdesc_t arhat_DeviceCmd_msg;
extern const pb_msgdesc_t arhat_TLSConfig_msg;
extern const pb_msgdesc_t arhat_DeviceConnectCmd_msg;
extern const pb_msgdesc_t arhat_DeviceConnectCmd_ParamsEntry_msg;
extern const pb_msgdesc_t arhat_DeviceOperateCmd_msg;
extern const pb_msgdesc_t arhat_DeviceOperateCmd_ParamsEntry_msg;
extern const pb_msgdesc_t arhat_DeviceMetricsCollectCmd_msg;
extern const pb_msgdesc_t arhat_DeviceMetricsCollectCmd_ParamsEntry_msg;
extern const pb_msgdesc_t arhat_DeviceCloseCmd_msg;
extern const pb_msgdesc_t arhat_DeviceMsg_msg;
extern const pb_msgdesc_t arhat_DeviceDoneMsg_msg;
extern const pb_msgdesc_t arhat_DeviceRegisterMsg_msg;
extern const pb_msgdesc_t arhat_DeviceOperateResultMsg_msg;
extern const pb_msgdesc_t arhat_DeviceMetricsMsg_msg;
extern const pb_msgdesc_t arhat_DeviceMetricsMsg_Value_msg;
extern const pb_msgdesc_t arhat_DeviceEventMsg_msg;
extern const pb_msgdesc_t arhat_ErrorMsg_msg;

/* Defines for backwards compatibility with code written before nanopb-0.4.0 */
#define arhat_DeviceCmd_fields &arhat_DeviceCmd_msg
#define arhat_TLSConfig_fields &arhat_TLSConfig_msg
#define arhat_DeviceConnectCmd_fields &arhat_DeviceConnectCmd_msg
#define arhat_DeviceConnectCmd_ParamsEntry_fields &arhat_DeviceConnectCmd_ParamsEntry_msg
#define arhat_DeviceOperateCmd_fields &arhat_DeviceOperateCmd_msg
#define arhat_DeviceOperateCmd_ParamsEntry_fields &arhat_DeviceOperateCmd_ParamsEntry_msg
#define arhat_DeviceMetricsCollectCmd_fields &arhat_DeviceMetricsCollectCmd_msg
#define arhat_DeviceMetricsCollectCmd_ParamsEntry_fields &arhat_DeviceMetricsCollectCmd_ParamsEntry_msg
#define arhat_DeviceCloseCmd_fields &arhat_DeviceCloseCmd_msg
#define arhat_DeviceMsg_fields &arhat_DeviceMsg_msg
#define arhat_DeviceDoneMsg_fields &arhat_DeviceDoneMsg_msg
#define arhat_DeviceRegisterMsg_fields &arhat_DeviceRegisterMsg_msg
#define arhat_DeviceOperateResultMsg_fields &arhat_DeviceOperateResultMsg_msg
#define arhat_DeviceMetricsMsg_fields &arhat_DeviceMetricsMsg_msg
#define arhat_DeviceMetricsMsg_Value_fields &arhat_DeviceMetricsMsg_Value_msg
#define arhat_DeviceEventMsg_fields &arhat_DeviceEventMsg_msg
#define arhat_ErrorMsg_fields &arhat_ErrorMsg_msg

/* Maximum encoded size of messages (where known) */
/* arhat_DeviceCmd_size depends on runtime parameters */
/* arhat_TLSConfig_size depends on runtime parameters */
/* arhat_DeviceConnectCmd_size depends on runtime parameters */
/* arhat_DeviceConnectCmd_ParamsEntry_size depends on runtime parameters */
/* arhat_DeviceOperateCmd_size depends on runtime parameters */
/* arhat_DeviceOperateCmd_ParamsEntry_size depends on runtime parameters */
/* arhat_DeviceMetricsCollectCmd_size depends on runtime parameters */
/* arhat_DeviceMetricsCollectCmd_ParamsEntry_size depends on runtime parameters */
#define arhat_DeviceCloseCmd_size                0
/* arhat_DeviceMsg_size depends on runtime parameters */
#define arhat_DeviceDoneMsg_size                 0
/* arhat_DeviceRegisterMsg_size depends on runtime parameters */
/* arhat_DeviceOperateResultMsg_size depends on runtime parameters */
/* arhat_DeviceMetricsMsg_size depends on runtime parameters */
#define arhat_DeviceMetricsMsg_Value_size        20
#define arhat_DeviceEventMsg_size                2
/* arhat_ErrorMsg_size depends on runtime parameters */

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif
