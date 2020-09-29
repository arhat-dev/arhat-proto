/* Automatically generated nanopb constant definitions */
/* Generated by nanopb-0.4.2 */

#include "device.pb.h"
#if PB_PROTO_HEADER_VERSION != 40
#error Regenerate this file with the current version of nanopb generator.
#endif

PB_BIND(arhat_DeviceCmd, arhat_DeviceCmd, AUTO)


PB_BIND(arhat_TLSConfig, arhat_TLSConfig, AUTO)


PB_BIND(arhat_DeviceConnectCmd, arhat_DeviceConnectCmd, AUTO)


PB_BIND(arhat_DeviceConnectCmd_ParamsEntry, arhat_DeviceConnectCmd_ParamsEntry, AUTO)


PB_BIND(arhat_DeviceOperateCmd, arhat_DeviceOperateCmd, AUTO)


PB_BIND(arhat_DeviceOperateCmd_ParamsEntry, arhat_DeviceOperateCmd_ParamsEntry, AUTO)


PB_BIND(arhat_DeviceMetricsCollectCmd, arhat_DeviceMetricsCollectCmd, AUTO)


PB_BIND(arhat_DeviceMetricsCollectCmd_ParamsEntry, arhat_DeviceMetricsCollectCmd_ParamsEntry, AUTO)


PB_BIND(arhat_DeviceCloseCmd, arhat_DeviceCloseCmd, AUTO)


PB_BIND(arhat_DeviceMsg, arhat_DeviceMsg, AUTO)


PB_BIND(arhat_DeviceDoneMsg, arhat_DeviceDoneMsg, AUTO)


PB_BIND(arhat_DeviceRegisterMsg, arhat_DeviceRegisterMsg, AUTO)


PB_BIND(arhat_DeviceOperateResultMsg, arhat_DeviceOperateResultMsg, AUTO)


PB_BIND(arhat_DeviceMetricsMsg, arhat_DeviceMetricsMsg, AUTO)


PB_BIND(arhat_DeviceMetricsMsg_Value, arhat_DeviceMetricsMsg_Value, AUTO)


PB_BIND(arhat_DeviceEventMsg, arhat_DeviceEventMsg, AUTO)


PB_BIND(arhat_ErrorMsg, arhat_ErrorMsg, AUTO)






#ifndef PB_CONVERT_DOUBLE_FLOAT
/* On some platforms (such as AVR), double is really float.
 * To be able to encode/decode double on these platforms, you need.
 * to define PB_CONVERT_DOUBLE_FLOAT in pb.h or compiler command line.
 */
PB_STATIC_ASSERT(sizeof(double) == 8, DOUBLE_MUST_BE_8_BYTES)
#endif

