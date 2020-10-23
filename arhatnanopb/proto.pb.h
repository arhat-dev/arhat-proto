/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.2 */

#ifndef PB_ARHAT_PROTO_PB_H_INCLUDED
#define PB_ARHAT_PROTO_PB_H_INCLUDED
#include <pb.h>

#if PB_PROTO_HEADER_VERSION != 40
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Enum definitions */
typedef enum _arhat_CmdType {
    arhat_CmdType__INVALID_CMD = 0,
    arhat_CmdType_CMD_PERIPHERAL_CONNECT = 11,
    arhat_CmdType_CMD_PERIPHERAL_OPERATE = 12,
    arhat_CmdType_CMD_PERIPHERAL_COLLECT_METRICS = 13,
    arhat_CmdType_CMD_PERIPHERAL_CLOSE = 14
} arhat_CmdType;

typedef enum _arhat_MsgType {
    arhat_MsgType__INVALID_MSG = 0,
    arhat_MsgType_MSG_ERROR = 1,
    arhat_MsgType_MSG_DONE = 2,
    arhat_MsgType_MSG_REGISTER = 3,
    arhat_MsgType_MSG_PERIPHERAL_OPERATION_RESULT = 11,
    arhat_MsgType_MSG_PERIPHERAL_METRICS = 12,
    arhat_MsgType_MSG_PERIPHERAL_EVENTS = 13
} arhat_MsgType;

/* Struct definitions */
typedef struct _arhat_DoneMsg {
    char dummy_field;
} arhat_DoneMsg;

typedef struct _arhat_ErrorMsg {
    pb_callback_t description;
} arhat_ErrorMsg;

typedef struct _arhat_RegisterMsg {
    pb_callback_t name;
} arhat_RegisterMsg;

typedef struct _arhat_Cmd {
    arhat_CmdType kind;
    uint64_t id;
    uint64_t seq;
    pb_callback_t payload;
} arhat_Cmd;

typedef struct _arhat_Msg {
    arhat_MsgType kind;
    uint64_t id;
    uint64_t ack;
    pb_callback_t payload;
} arhat_Msg;


/* Helper constants for enums */
#define _arhat_CmdType_MIN arhat_CmdType__INVALID_CMD
#define _arhat_CmdType_MAX arhat_CmdType_CMD_PERIPHERAL_CLOSE
#define _arhat_CmdType_ARRAYSIZE ((arhat_CmdType)(arhat_CmdType_CMD_PERIPHERAL_CLOSE+1))

#define _arhat_MsgType_MIN arhat_MsgType__INVALID_MSG
#define _arhat_MsgType_MAX arhat_MsgType_MSG_PERIPHERAL_EVENTS
#define _arhat_MsgType_ARRAYSIZE ((arhat_MsgType)(arhat_MsgType_MSG_PERIPHERAL_EVENTS+1))


/* Initializer values for message structs */
#define arhat_Cmd_init_default                   {_arhat_CmdType_MIN, 0, 0, {{NULL}, NULL}}
#define arhat_Msg_init_default                   {_arhat_MsgType_MIN, 0, 0, {{NULL}, NULL}}
#define arhat_DoneMsg_init_default               {0}
#define arhat_ErrorMsg_init_default              {{{NULL}, NULL}}
#define arhat_RegisterMsg_init_default           {{{NULL}, NULL}}
#define arhat_Cmd_init_zero                      {_arhat_CmdType_MIN, 0, 0, {{NULL}, NULL}}
#define arhat_Msg_init_zero                      {_arhat_MsgType_MIN, 0, 0, {{NULL}, NULL}}
#define arhat_DoneMsg_init_zero                  {0}
#define arhat_ErrorMsg_init_zero                 {{{NULL}, NULL}}
#define arhat_RegisterMsg_init_zero              {{{NULL}, NULL}}

/* Field tags (for use in manual encoding/decoding) */
#define arhat_ErrorMsg_description_tag           1
#define arhat_RegisterMsg_name_tag               1
#define arhat_Cmd_kind_tag                       1
#define arhat_Cmd_id_tag                         2
#define arhat_Cmd_seq_tag                        3
#define arhat_Cmd_payload_tag                    4
#define arhat_Msg_kind_tag                       1
#define arhat_Msg_id_tag                         2
#define arhat_Msg_ack_tag                        3
#define arhat_Msg_payload_tag                    4

/* Struct field encoding specification for nanopb */
#define arhat_Cmd_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, UENUM,    kind,              1) \
X(a, STATIC,   SINGULAR, UINT64,   id,                2) \
X(a, STATIC,   SINGULAR, UINT64,   seq,               3) \
X(a, CALLBACK, SINGULAR, BYTES,    payload,           4)
#define arhat_Cmd_CALLBACK pb_default_field_callback
#define arhat_Cmd_DEFAULT NULL

#define arhat_Msg_FIELDLIST(X, a) \
X(a, STATIC,   SINGULAR, UENUM,    kind,              1) \
X(a, STATIC,   SINGULAR, UINT64,   id,                2) \
X(a, STATIC,   SINGULAR, UINT64,   ack,               3) \
X(a, CALLBACK, SINGULAR, BYTES,    payload,           4)
#define arhat_Msg_CALLBACK pb_default_field_callback
#define arhat_Msg_DEFAULT NULL

#define arhat_DoneMsg_FIELDLIST(X, a) \

#define arhat_DoneMsg_CALLBACK NULL
#define arhat_DoneMsg_DEFAULT NULL

#define arhat_ErrorMsg_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   description,       1)
#define arhat_ErrorMsg_CALLBACK pb_default_field_callback
#define arhat_ErrorMsg_DEFAULT NULL

#define arhat_RegisterMsg_FIELDLIST(X, a) \
X(a, CALLBACK, SINGULAR, STRING,   name,              1)
#define arhat_RegisterMsg_CALLBACK pb_default_field_callback
#define arhat_RegisterMsg_DEFAULT NULL

extern const pb_msgdesc_t arhat_Cmd_msg;
extern const pb_msgdesc_t arhat_Msg_msg;
extern const pb_msgdesc_t arhat_DoneMsg_msg;
extern const pb_msgdesc_t arhat_ErrorMsg_msg;
extern const pb_msgdesc_t arhat_RegisterMsg_msg;

/* Defines for backwards compatibility with code written before nanopb-0.4.0 */
#define arhat_Cmd_fields &arhat_Cmd_msg
#define arhat_Msg_fields &arhat_Msg_msg
#define arhat_DoneMsg_fields &arhat_DoneMsg_msg
#define arhat_ErrorMsg_fields &arhat_ErrorMsg_msg
#define arhat_RegisterMsg_fields &arhat_RegisterMsg_msg

/* Maximum encoded size of messages (where known) */
/* arhat_Cmd_size depends on runtime parameters */
/* arhat_Msg_size depends on runtime parameters */
#define arhat_DoneMsg_size                       0
/* arhat_ErrorMsg_size depends on runtime parameters */
/* arhat_RegisterMsg_size depends on runtime parameters */

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif
