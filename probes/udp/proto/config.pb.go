// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/udp/proto/config.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/probes/udp/proto/config.proto

It has these top-level messages:
	ProbeConf
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ProbeConf struct {
	// Export stats after these many milliseconds
	StatsExportIntervalMsec *int32 `protobuf:"varint,2,opt,name=stats_export_interval_msec,json=statsExportIntervalMsec,def=10000" json:"stats_export_interval_msec,omitempty"`
	// Port to send UDP Ping to (UDP Echo).  Should be same as
	// ProberConfig.udp_echo_server_port.
	// TODO(): Can we just read this from ProberConfig?
	Port *int32 `protobuf:"varint,3,opt,name=port,def=31122" json:"port,omitempty"`
	// Number of sending side ports to use.
	NumTxPorts *int32 `protobuf:"varint,4,opt,name=num_tx_ports,json=numTxPorts,def=16" json:"num_tx_ports,omitempty"`
	// message max to account for MTU.
	MaxLength *int32 `protobuf:"varint,5,opt,name=max_length,json=maxLength,def=1300" json:"max_length,omitempty"`
	// IP proto: v4|v6. Temporary workaround till we add support in base probe proto.
	IpVersion        *int32 `protobuf:"varint,6,opt,name=ip_version,json=ipVersion,def=4" json:"ip_version,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ProbeConf) Reset()                    { *m = ProbeConf{} }
func (m *ProbeConf) String() string            { return proto1.CompactTextString(m) }
func (*ProbeConf) ProtoMessage()               {}
func (*ProbeConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_ProbeConf_StatsExportIntervalMsec int32 = 10000
const Default_ProbeConf_Port int32 = 31122
const Default_ProbeConf_NumTxPorts int32 = 16
const Default_ProbeConf_MaxLength int32 = 1300
const Default_ProbeConf_IpVersion int32 = 4

func (m *ProbeConf) GetStatsExportIntervalMsec() int32 {
	if m != nil && m.StatsExportIntervalMsec != nil {
		return *m.StatsExportIntervalMsec
	}
	return Default_ProbeConf_StatsExportIntervalMsec
}

func (m *ProbeConf) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return Default_ProbeConf_Port
}

func (m *ProbeConf) GetNumTxPorts() int32 {
	if m != nil && m.NumTxPorts != nil {
		return *m.NumTxPorts
	}
	return Default_ProbeConf_NumTxPorts
}

func (m *ProbeConf) GetMaxLength() int32 {
	if m != nil && m.MaxLength != nil {
		return *m.MaxLength
	}
	return Default_ProbeConf_MaxLength
}

func (m *ProbeConf) GetIpVersion() int32 {
	if m != nil && m.IpVersion != nil {
		return *m.IpVersion
	}
	return Default_ProbeConf_IpVersion
}

func init() {
	proto1.RegisterType((*ProbeConf)(nil), "cloudprober.probes.udp.ProbeConf")
}

func init() {
	proto1.RegisterFile("github.com/google/cloudprober/probes/udp/proto/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xca, 0x4f, 0x4b, 0xf4, 0x30,
	0x10, 0xc7, 0x71, 0xba, 0x4f, 0xfb, 0x40, 0x07, 0x4f, 0x39, 0x68, 0xf5, 0xb4, 0xa8, 0x07, 0x4f,
	0x6d, 0xba, 0x2b, 0x1e, 0xf4, 0xa6, 0x78, 0x10, 0x14, 0x96, 0x22, 0x5e, 0x43, 0xb7, 0xcd, 0x76,
	0x03, 0x4d, 0x26, 0xe4, 0xcf, 0x92, 0x77, 0xea, 0xdb, 0x91, 0xa4, 0x1e, 0x3c, 0xcd, 0xfc, 0xbe,
	0x7c, 0xe0, 0x69, 0x12, 0xee, 0xe8, 0xf7, 0xf5, 0x80, 0xb2, 0x99, 0x10, 0xa7, 0x99, 0x37, 0xc3,
	0x8c, 0x7e, 0xd4, 0x06, 0xf7, 0xdc, 0x34, 0xe9, 0xd8, 0xc6, 0x8f, 0x3a, 0xbe, 0x0e, 0x9b, 0x01,
	0xd5, 0x41, 0x4c, 0x75, 0x1a, 0xe4, 0xfc, 0x0f, 0xad, 0x17, 0x5a, 0xfb, 0x51, 0x5f, 0x7f, 0x67,
	0x50, 0xee, 0xe2, 0x7c, 0x41, 0x75, 0x20, 0xcf, 0x70, 0x65, 0x5d, 0xef, 0x2c, 0xe3, 0x41, 0xa3,
	0x71, 0x4c, 0x28, 0xc7, 0xcd, 0xa9, 0x9f, 0x99, 0xb4, 0x7c, 0xa8, 0x56, 0xeb, 0xec, 0xae, 0x78,
	0x2c, 0x5a, 0x4a, 0x29, 0xed, 0x2e, 0x12, 0x7c, 0x4d, 0xee, 0xed, 0x97, 0x7d, 0x58, 0x3e, 0x90,
	0x4b, 0xc8, 0x63, 0xab, 0xfe, 0x2d, 0x7a, 0xdb, 0xb6, 0x9b, 0x4d, 0x97, 0x12, 0xb9, 0x85, 0x33,
	0xe5, 0x25, 0x73, 0x81, 0xc5, 0x69, 0xab, 0x3c, 0x91, 0x55, 0xfb, 0xd0, 0x81, 0xf2, 0xf2, 0x33,
	0xec, 0x62, 0x25, 0x37, 0x00, 0xb2, 0x0f, 0x6c, 0xe6, 0x6a, 0x72, 0xc7, 0xaa, 0x48, 0x26, 0x6f,
	0xb7, 0x94, 0x76, 0xa5, 0xec, 0xc3, 0x7b, 0xca, 0x64, 0x0d, 0x20, 0x34, 0x3b, 0x71, 0x63, 0x05,
	0xaa, 0xea, 0x7f, 0x42, 0xd9, 0x7d, 0x57, 0x0a, 0xfd, 0xb5, 0xb4, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xec, 0x40, 0xdc, 0x7a, 0x2f, 0x01, 0x00, 0x00,
}
