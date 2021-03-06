// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: container.proto

package container

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Persisted represents configured items for a pod that are persisted.
type Persisted struct {
	// id is identifier of pod
	ID string `protobuf:"bytes,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	// podName
	PodName      string `protobuf:"bytes,2,opt,name=podName,proto3" json:"podName,omitempty"`
	PodNamespace string `protobuf:"bytes,3,opt,name=podNamespace,proto3" json:"podNamespace,omitempty"`
	// Veth1 is name of one end end of veth pair that is in the given container namespace.
	// Nil if TAPs are used instead.
	Veth1Name string `protobuf:"bytes,4,opt,name=Veth1Name,json=veth1Name,proto3" json:"Veth1Name,omitempty"`
	// Veth2 is name the other end of veth pair in the default namespace
	// Nil if TAPs are used instead.
	Veth2Name string `protobuf:"bytes,5,opt,name=Veth2Name,json=veth2Name,proto3" json:"Veth2Name,omitempty"`
	// VppIf is name AF_PACKET/TAP interface connecting pod to VPP
	VppIfName string `protobuf:"bytes,6,opt,name=VppIfName,json=vppIfName,proto3" json:"VppIfName,omitempty"`
	// PodTap is name ofthe host end of the tap connecting pod to VPP
	// Nil if TAPs are not used
	PodTapName string `protobuf:"bytes,7,opt,name=PodTapName,json=podTapName,proto3" json:"PodTapName,omitempty"`
	// Loopback is interface name associated with the pod.
	// Nil if VPP TCP stack is disabled.
	LoopbackName string `protobuf:"bytes,8,opt,name=LoopbackName,json=loopbackName,proto3" json:"LoopbackName,omitempty"`
	// StnRule is name of STN rule used to "punt" any traffic via VETHs/TAPs with no match in VPP TCP stack.
	// Nil if VPP TCP stack is disabled.
	StnRuleName string `protobuf:"bytes,9,opt,name=StnRuleName,json=stnRuleName,proto3" json:"StnRuleName,omitempty"`
	// AppNamespace is id of the application namespace associated with the pod.
	// Nil if VPP TCP stack is disabled.
	AppNamespaceID string `protobuf:"bytes,10,opt,name=AppNamespaceID,json=appNamespaceID,proto3" json:"AppNamespaceID,omitempty"`
	// VppARPEntryInterface is name of the Interface associated ARP entry configured in VPP to route traffic from VPP to pod.
	VppARPEntryInterface string `protobuf:"bytes,11,opt,name=VppARPEntryInterface,json=vppARPEntryInterface,proto3" json:"VppARPEntryInterface,omitempty"`
	// VppARPEntryIP is IP associated ARP entry configured in VPP to route traffic from VPP to pod.
	VppARPEntryIP string `protobuf:"bytes,12,opt,name=VppARPEntryIP,json=vppARPEntryIP,proto3" json:"VppARPEntryIP,omitempty"`
	// PodARPEntry is name of ARP entry configured in the pod to route traffic from pod to VPP.
	PodARPEntryName string `protobuf:"bytes,13,opt,name=PodARPEntryName,json=podARPEntryName,proto3" json:"PodARPEntryName,omitempty"`
	// VppRouteVrf is vrf of the route from VPP to the container
	VppRouteVrf uint32 `protobuf:"varint,15,opt,name=VppRouteVrf,json=vppRouteVrf,proto3" json:"VppRouteVrf,omitempty"`
	// VppRouteDest is destination of the route from VPP to the container
	VppRouteDest string `protobuf:"bytes,16,opt,name=VppRouteDest,json=vppRouteDest,proto3" json:"VppRouteDest,omitempty"`
	// VppRouteNextHop is next hop of the route from VPP to the container
	VppRouteNextHop string `protobuf:"bytes,17,opt,name=VppRouteNextHop,json=vppRouteNextHop,proto3" json:"VppRouteNextHop,omitempty"`
	// PodLinkRoute is name of the route from pod to the default gateway.
	PodLinkRouteName string `protobuf:"bytes,18,opt,name=PodLinkRouteName,json=podLinkRouteName,proto3" json:"PodLinkRouteName,omitempty"`
	// PodDefaultRoute is name of the default gateway for the pod.
	PodDefaultRouteName  string   `protobuf:"bytes,19,opt,name=PodDefaultRouteName,json=podDefaultRouteName,proto3" json:"PodDefaultRouteName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Persisted) Reset()         { *m = Persisted{} }
func (m *Persisted) String() string { return proto.CompactTextString(m) }
func (*Persisted) ProtoMessage()    {}
func (*Persisted) Descriptor() ([]byte, []int) {
	return fileDescriptor_container_9b14729bd5c5166a, []int{0}
}
func (m *Persisted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Persisted.Unmarshal(m, b)
}
func (m *Persisted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Persisted.Marshal(b, m, deterministic)
}
func (dst *Persisted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Persisted.Merge(dst, src)
}
func (m *Persisted) XXX_Size() int {
	return xxx_messageInfo_Persisted.Size(m)
}
func (m *Persisted) XXX_DiscardUnknown() {
	xxx_messageInfo_Persisted.DiscardUnknown(m)
}

var xxx_messageInfo_Persisted proto.InternalMessageInfo

func (m *Persisted) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Persisted) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *Persisted) GetPodNamespace() string {
	if m != nil {
		return m.PodNamespace
	}
	return ""
}

func (m *Persisted) GetVeth1Name() string {
	if m != nil {
		return m.Veth1Name
	}
	return ""
}

func (m *Persisted) GetVeth2Name() string {
	if m != nil {
		return m.Veth2Name
	}
	return ""
}

func (m *Persisted) GetVppIfName() string {
	if m != nil {
		return m.VppIfName
	}
	return ""
}

func (m *Persisted) GetPodTapName() string {
	if m != nil {
		return m.PodTapName
	}
	return ""
}

func (m *Persisted) GetLoopbackName() string {
	if m != nil {
		return m.LoopbackName
	}
	return ""
}

func (m *Persisted) GetStnRuleName() string {
	if m != nil {
		return m.StnRuleName
	}
	return ""
}

func (m *Persisted) GetAppNamespaceID() string {
	if m != nil {
		return m.AppNamespaceID
	}
	return ""
}

func (m *Persisted) GetVppARPEntryInterface() string {
	if m != nil {
		return m.VppARPEntryInterface
	}
	return ""
}

func (m *Persisted) GetVppARPEntryIP() string {
	if m != nil {
		return m.VppARPEntryIP
	}
	return ""
}

func (m *Persisted) GetPodARPEntryName() string {
	if m != nil {
		return m.PodARPEntryName
	}
	return ""
}

func (m *Persisted) GetVppRouteVrf() uint32 {
	if m != nil {
		return m.VppRouteVrf
	}
	return 0
}

func (m *Persisted) GetVppRouteDest() string {
	if m != nil {
		return m.VppRouteDest
	}
	return ""
}

func (m *Persisted) GetVppRouteNextHop() string {
	if m != nil {
		return m.VppRouteNextHop
	}
	return ""
}

func (m *Persisted) GetPodLinkRouteName() string {
	if m != nil {
		return m.PodLinkRouteName
	}
	return ""
}

func (m *Persisted) GetPodDefaultRouteName() string {
	if m != nil {
		return m.PodDefaultRouteName
	}
	return ""
}

func init() {
	proto.RegisterType((*Persisted)(nil), "container.Persisted")
}

func init() { proto.RegisterFile("container.proto", fileDescriptor_container_9b14729bd5c5166a) }

var fileDescriptor_container_9b14729bd5c5166a = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdd, 0x6a, 0xe2, 0x40,
	0x14, 0x80, 0xd1, 0x5d, 0x75, 0x73, 0x62, 0x8c, 0x3b, 0x7a, 0x31, 0x17, 0xcb, 0x22, 0xb2, 0x2c,
	0xd2, 0x8b, 0xd2, 0xda, 0x27, 0x10, 0x52, 0x68, 0x40, 0x64, 0x48, 0x4b, 0xee, 0xa3, 0x99, 0xd0,
	0xa0, 0xcd, 0x1c, 0x92, 0x31, 0xb4, 0xcf, 0xd1, 0x17, 0x2e, 0x39, 0x31, 0x7f, 0x6d, 0xef, 0xf4,
	0xfb, 0xbe, 0x61, 0xce, 0x1c, 0x02, 0xf6, 0x41, 0x25, 0x3a, 0x88, 0x13, 0x99, 0x5e, 0x63, 0xaa,
	0xb4, 0x62, 0x46, 0x0d, 0x96, 0xef, 0x03, 0x30, 0x84, 0x4c, 0xb3, 0x38, 0xd3, 0x32, 0x64, 0x13,
	0xe8, 0xbb, 0x0e, 0xef, 0x2d, 0x7a, 0x2b, 0xc3, 0xeb, 0xc7, 0x0e, 0xe3, 0x30, 0x42, 0x15, 0xee,
	0x82, 0x17, 0xc9, 0xfb, 0x04, 0xab, 0xbf, 0x6c, 0x09, 0xe3, 0xcb, 0xcf, 0x0c, 0x83, 0x83, 0xe4,
	0x3f, 0x48, 0x77, 0x18, 0xfb, 0x03, 0x86, 0x2f, 0xf5, 0xf3, 0x2d, 0x9d, 0xff, 0x49, 0x81, 0x91,
	0x57, 0xa0, 0xb2, 0x6b, 0xb2, 0x83, 0xc6, 0xae, 0x6b, 0x8b, 0xe8, 0x46, 0x64, 0x87, 0x17, 0x5b,
	0x01, 0xf6, 0x17, 0x40, 0xa8, 0xf0, 0x29, 0x40, 0xd2, 0x23, 0xd2, 0x80, 0x35, 0x29, 0xa6, 0xdb,
	0x2a, 0x85, 0xfb, 0xe0, 0x70, 0xa4, 0xe2, 0x57, 0x39, 0xdd, 0xa9, 0xc5, 0xd8, 0x02, 0xcc, 0x47,
	0x9d, 0x78, 0xe7, 0x93, 0xa4, 0xc4, 0xa0, 0xc4, 0xcc, 0x1a, 0xc4, 0xfe, 0xc3, 0x64, 0x83, 0x58,
	0xbf, 0xc7, 0x75, 0x38, 0x50, 0x34, 0x09, 0x3a, 0x94, 0xad, 0x61, 0xee, 0x23, 0x6e, 0x3c, 0x71,
	0x9f, 0xe8, 0xf4, 0xcd, 0x4d, 0xb4, 0x4c, 0xa3, 0x62, 0x27, 0x26, 0xd5, 0xf3, 0xfc, 0x1b, 0xc7,
	0xfe, 0x81, 0xd5, 0x3e, 0x23, 0xf8, 0x98, 0x62, 0xab, 0x1d, 0x0b, 0xb6, 0x02, 0x5b, 0xa8, 0xb0,
	0x02, 0x34, 0xa7, 0x45, 0x9d, 0x8d, 0x5d, 0x5c, 0xbc, 0xc6, 0x47, 0xf4, 0xd4, 0x59, 0x4b, 0x3f,
	0x8d, 0xb8, 0xbd, 0xe8, 0xad, 0x2c, 0xcf, 0xcc, 0x1b, 0x54, 0xec, 0xa4, 0x2a, 0x1c, 0x99, 0x69,
	0x3e, 0x2d, 0x77, 0x92, 0xb7, 0x58, 0x71, 0x5f, 0xd5, 0xec, 0xe4, 0xab, 0x7e, 0x50, 0xc8, 0x7f,
	0x97, 0xf7, 0xe5, 0x5d, 0xcc, 0xae, 0x60, 0x2a, 0x54, 0xb8, 0x8d, 0x93, 0x63, 0x89, 0x8b, 0xd1,
	0x18, 0xa5, 0x53, 0xfc, 0xc4, 0xd9, 0x0d, 0xcc, 0x84, 0x0a, 0x1d, 0x19, 0x05, 0xe7, 0x93, 0x6e,
	0xf2, 0x19, 0xe5, 0x33, 0xfc, 0xaa, 0xf6, 0x43, 0xfa, 0x4e, 0xef, 0x3e, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xd1, 0xae, 0x4d, 0x18, 0xba, 0x02, 0x00, 0x00,
}
