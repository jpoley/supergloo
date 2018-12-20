// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/api/v1/install.proto

package v1 // import "github.com/solo-io/supergloo/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
// @solo-kit:resource.short_name=install
// @solo-kit:resource.plural_name=installs
// @solo-kit:resource.resource_groups=install.supergloo.solo.io
type Install struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,1,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,2,opt,name=metadata" json:"metadata"`
	// mesh-specific configuration
	//
	// Types that are valid to be assigned to MeshType:
	//	*Install_Istio
	//	*Install_Linkerd2
	//	*Install_Consul
	MeshType     isInstall_MeshType `protobuf_oneof:"mesh_type"`
	ChartLocator *HelmChartLocator  `protobuf:"bytes,6,opt,name=chartLocator" json:"chartLocator,omitempty"`
	Encryption   *Encryption        `protobuf:"bytes,7,opt,name=encryption" json:"encryption,omitempty"`
	// whether or not this install should be enabled
	// if disabled, corresponding resources will be uninstalled
	// defaults to true
	Enabled              *types.BoolValue `protobuf:"bytes,12,opt,name=enabled" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Install) Reset()         { *m = Install{} }
func (m *Install) String() string { return proto.CompactTextString(m) }
func (*Install) ProtoMessage()    {}
func (*Install) Descriptor() ([]byte, []int) {
	return fileDescriptor_install_7a348ed6828049b4, []int{0}
}
func (m *Install) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install.Unmarshal(m, b)
}
func (m *Install) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install.Marshal(b, m, deterministic)
}
func (dst *Install) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install.Merge(dst, src)
}
func (m *Install) XXX_Size() int {
	return xxx_messageInfo_Install.Size(m)
}
func (m *Install) XXX_DiscardUnknown() {
	xxx_messageInfo_Install.DiscardUnknown(m)
}

var xxx_messageInfo_Install proto.InternalMessageInfo

type isInstall_MeshType interface {
	isInstall_MeshType()
	Equal(interface{}) bool
}

type Install_Istio struct {
	Istio *Istio `protobuf:"bytes,10,opt,name=istio,oneof"`
}
type Install_Linkerd2 struct {
	Linkerd2 *Linkerd2 `protobuf:"bytes,20,opt,name=linkerd2,oneof"`
}
type Install_Consul struct {
	Consul *Consul `protobuf:"bytes,30,opt,name=consul,oneof"`
}

func (*Install_Istio) isInstall_MeshType()    {}
func (*Install_Linkerd2) isInstall_MeshType() {}
func (*Install_Consul) isInstall_MeshType()   {}

func (m *Install) GetMeshType() isInstall_MeshType {
	if m != nil {
		return m.MeshType
	}
	return nil
}

func (m *Install) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Install) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Install) GetIstio() *Istio {
	if x, ok := m.GetMeshType().(*Install_Istio); ok {
		return x.Istio
	}
	return nil
}

func (m *Install) GetLinkerd2() *Linkerd2 {
	if x, ok := m.GetMeshType().(*Install_Linkerd2); ok {
		return x.Linkerd2
	}
	return nil
}

func (m *Install) GetConsul() *Consul {
	if x, ok := m.GetMeshType().(*Install_Consul); ok {
		return x.Consul
	}
	return nil
}

func (m *Install) GetChartLocator() *HelmChartLocator {
	if m != nil {
		return m.ChartLocator
	}
	return nil
}

func (m *Install) GetEncryption() *Encryption {
	if m != nil {
		return m.Encryption
	}
	return nil
}

func (m *Install) GetEnabled() *types.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Install) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Install_OneofMarshaler, _Install_OneofUnmarshaler, _Install_OneofSizer, []interface{}{
		(*Install_Istio)(nil),
		(*Install_Linkerd2)(nil),
		(*Install_Consul)(nil),
	}
}

func _Install_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Install)
	// mesh_type
	switch x := m.MeshType.(type) {
	case *Install_Istio:
		_ = b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Istio); err != nil {
			return err
		}
	case *Install_Linkerd2:
		_ = b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Linkerd2); err != nil {
			return err
		}
	case *Install_Consul:
		_ = b.EncodeVarint(30<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Consul); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Install.MeshType has unexpected type %T", x)
	}
	return nil
}

func _Install_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Install)
	switch tag {
	case 10: // mesh_type.istio
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Istio)
		err := b.DecodeMessage(msg)
		m.MeshType = &Install_Istio{msg}
		return true, err
	case 20: // mesh_type.linkerd2
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Linkerd2)
		err := b.DecodeMessage(msg)
		m.MeshType = &Install_Linkerd2{msg}
		return true, err
	case 30: // mesh_type.consul
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Consul)
		err := b.DecodeMessage(msg)
		m.MeshType = &Install_Consul{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Install_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Install)
	// mesh_type
	switch x := m.MeshType.(type) {
	case *Install_Istio:
		s := proto.Size(x.Istio)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Install_Linkerd2:
		s := proto.Size(x.Linkerd2)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Install_Consul:
		s := proto.Size(x.Consul)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HelmChartLocator struct {
	// Types that are valid to be assigned to Kind:
	//	*HelmChartLocator_ChartPath
	Kind                 isHelmChartLocator_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *HelmChartLocator) Reset()         { *m = HelmChartLocator{} }
func (m *HelmChartLocator) String() string { return proto.CompactTextString(m) }
func (*HelmChartLocator) ProtoMessage()    {}
func (*HelmChartLocator) Descriptor() ([]byte, []int) {
	return fileDescriptor_install_7a348ed6828049b4, []int{1}
}
func (m *HelmChartLocator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelmChartLocator.Unmarshal(m, b)
}
func (m *HelmChartLocator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelmChartLocator.Marshal(b, m, deterministic)
}
func (dst *HelmChartLocator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelmChartLocator.Merge(dst, src)
}
func (m *HelmChartLocator) XXX_Size() int {
	return xxx_messageInfo_HelmChartLocator.Size(m)
}
func (m *HelmChartLocator) XXX_DiscardUnknown() {
	xxx_messageInfo_HelmChartLocator.DiscardUnknown(m)
}

var xxx_messageInfo_HelmChartLocator proto.InternalMessageInfo

type isHelmChartLocator_Kind interface {
	isHelmChartLocator_Kind()
	Equal(interface{}) bool
}

type HelmChartLocator_ChartPath struct {
	ChartPath *HelmChartPath `protobuf:"bytes,1,opt,name=chartPath,oneof"`
}

func (*HelmChartLocator_ChartPath) isHelmChartLocator_Kind() {}

func (m *HelmChartLocator) GetKind() isHelmChartLocator_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *HelmChartLocator) GetChartPath() *HelmChartPath {
	if x, ok := m.GetKind().(*HelmChartLocator_ChartPath); ok {
		return x.ChartPath
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HelmChartLocator) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HelmChartLocator_OneofMarshaler, _HelmChartLocator_OneofUnmarshaler, _HelmChartLocator_OneofSizer, []interface{}{
		(*HelmChartLocator_ChartPath)(nil),
	}
}

func _HelmChartLocator_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HelmChartLocator)
	// kind
	switch x := m.Kind.(type) {
	case *HelmChartLocator_ChartPath:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ChartPath); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HelmChartLocator.Kind has unexpected type %T", x)
	}
	return nil
}

func _HelmChartLocator_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HelmChartLocator)
	switch tag {
	case 1: // kind.chartPath
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HelmChartPath)
		err := b.DecodeMessage(msg)
		m.Kind = &HelmChartLocator_ChartPath{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HelmChartLocator_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HelmChartLocator)
	// kind
	switch x := m.Kind.(type) {
	case *HelmChartLocator_ChartPath:
		s := proto.Size(x.ChartPath)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HelmChartPath struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelmChartPath) Reset()         { *m = HelmChartPath{} }
func (m *HelmChartPath) String() string { return proto.CompactTextString(m) }
func (*HelmChartPath) ProtoMessage()    {}
func (*HelmChartPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_install_7a348ed6828049b4, []int{2}
}
func (m *HelmChartPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelmChartPath.Unmarshal(m, b)
}
func (m *HelmChartPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelmChartPath.Marshal(b, m, deterministic)
}
func (dst *HelmChartPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelmChartPath.Merge(dst, src)
}
func (m *HelmChartPath) XXX_Size() int {
	return xxx_messageInfo_HelmChartPath.Size(m)
}
func (m *HelmChartPath) XXX_DiscardUnknown() {
	xxx_messageInfo_HelmChartPath.DiscardUnknown(m)
}

var xxx_messageInfo_HelmChartPath proto.InternalMessageInfo

func (m *HelmChartPath) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*Install)(nil), "supergloo.solo.io.Install")
	proto.RegisterType((*HelmChartLocator)(nil), "supergloo.solo.io.HelmChartLocator")
	proto.RegisterType((*HelmChartPath)(nil), "supergloo.solo.io.HelmChartPath")
}
func (this *Install) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Install)
	if !ok {
		that2, ok := that.(Install)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if that1.MeshType == nil {
		if this.MeshType != nil {
			return false
		}
	} else if this.MeshType == nil {
		return false
	} else if !this.MeshType.Equal(that1.MeshType) {
		return false
	}
	if !this.ChartLocator.Equal(that1.ChartLocator) {
		return false
	}
	if !this.Encryption.Equal(that1.Encryption) {
		return false
	}
	if !this.Enabled.Equal(that1.Enabled) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Install_Istio) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Install_Istio)
	if !ok {
		that2, ok := that.(Install_Istio)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Istio.Equal(that1.Istio) {
		return false
	}
	return true
}
func (this *Install_Linkerd2) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Install_Linkerd2)
	if !ok {
		that2, ok := that.(Install_Linkerd2)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Linkerd2.Equal(that1.Linkerd2) {
		return false
	}
	return true
}
func (this *Install_Consul) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Install_Consul)
	if !ok {
		that2, ok := that.(Install_Consul)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Consul.Equal(that1.Consul) {
		return false
	}
	return true
}
func (this *HelmChartLocator) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HelmChartLocator)
	if !ok {
		that2, ok := that.(HelmChartLocator)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Kind == nil {
		if this.Kind != nil {
			return false
		}
	} else if this.Kind == nil {
		return false
	} else if !this.Kind.Equal(that1.Kind) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HelmChartLocator_ChartPath) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HelmChartLocator_ChartPath)
	if !ok {
		that2, ok := that.(HelmChartLocator_ChartPath)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ChartPath.Equal(that1.ChartPath) {
		return false
	}
	return true
}
func (this *HelmChartPath) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HelmChartPath)
	if !ok {
		that2, ok := that.(HelmChartPath)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/api/v1/install.proto", fileDescriptor_install_7a348ed6828049b4)
}

var fileDescriptor_install_7a348ed6828049b4 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0xb3, 0xba, 0x66, 0xbb, 0xaf, 0x15, 0xda, 0x61, 0x91, 0x74, 0xc5, 0x6d, 0x49, 0x0f,
	0x7a, 0xe9, 0xc4, 0xb6, 0x1e, 0x54, 0x10, 0x24, 0x45, 0x9a, 0x42, 0x05, 0x19, 0xc1, 0x83, 0x08,
	0x32, 0x9b, 0xcc, 0x66, 0x87, 0x9d, 0xcd, 0x0b, 0x93, 0x89, 0xd2, 0x6f, 0xe4, 0x47, 0xf1, 0x53,
	0xf4, 0xe0, 0x47, 0xf0, 0xe6, 0x4d, 0x32, 0x49, 0xb6, 0x5d, 0x0d, 0xa5, 0xa7, 0x9d, 0xcd, 0xfb,
	0xfd, 0xfe, 0x33, 0x2f, 0xf3, 0x02, 0x47, 0xa9, 0x34, 0xf3, 0x72, 0x4a, 0x63, 0x5c, 0x06, 0x05,
	0x2a, 0x3c, 0x94, 0x18, 0x14, 0x65, 0x2e, 0x74, 0xaa, 0x10, 0x03, 0x9e, 0xcb, 0xe0, 0xdb, 0x51,
	0x20, 0xb3, 0xc2, 0x70, 0xa5, 0x68, 0xae, 0xd1, 0x20, 0xd9, 0x59, 0xd5, 0x69, 0x65, 0x50, 0x89,
	0xe3, 0x51, 0x8a, 0x29, 0xda, 0x6a, 0x50, 0xad, 0x6a, 0x70, 0x3c, 0x49, 0x11, 0x53, 0x25, 0x02,
	0xfb, 0x6f, 0x5a, 0xce, 0x82, 0xef, 0x9a, 0xe7, 0xb9, 0xd0, 0x45, 0x53, 0xef, 0xdc, 0xbb, 0xfa,
	0x5d, 0x48, 0xd3, 0x6e, 0xbd, 0x14, 0x86, 0x27, 0xdc, 0xf0, 0x46, 0x09, 0xee, 0xa0, 0x14, 0x86,
	0x9b, 0xb2, 0xdd, 0x63, 0x5b, 0x64, 0xb1, 0xbe, 0xcc, 0x8d, 0xc4, 0xac, 0x79, 0x02, 0x4b, 0x51,
	0xcc, 0xeb, 0xb5, 0xff, 0xe7, 0x3e, 0x0c, 0xce, 0xeb, 0xe6, 0xc8, 0x19, 0xb8, 0xb5, 0xe9, 0xf5,
	0xf6, 0x7b, 0xcf, 0x36, 0x8f, 0x47, 0x34, 0x46, 0x2d, 0xda, 0x16, 0xe9, 0x47, 0x5b, 0x0b, 0x77,
	0x7f, 0x5e, 0xed, 0x39, 0xbf, 0xaf, 0xf6, 0x76, 0x8c, 0x28, 0x4c, 0x22, 0x67, 0xb3, 0xd7, 0xbe,
	0x4c, 0x33, 0xd4, 0xc2, 0x67, 0x8d, 0x4e, 0x5e, 0xc2, 0x46, 0x7b, 0x6a, 0xef, 0x9e, 0x8d, 0x7a,
	0xb4, 0x1e, 0xf5, 0xbe, 0xa9, 0x86, 0xfd, 0x2a, 0x8c, 0xad, 0x68, 0xf2, 0x1c, 0x1e, 0xc8, 0xc2,
	0x48, 0xf4, 0xc0, 0x6a, 0x1e, 0xfd, 0xef, 0x4d, 0xd3, 0xf3, 0xaa, 0x1e, 0x39, 0xac, 0x06, 0xc9,
	0x2b, 0xd8, 0x50, 0x32, 0x5b, 0x08, 0x9d, 0x1c, 0x7b, 0x23, 0x2b, 0x3d, 0xee, 0x90, 0x2e, 0x1a,
	0x24, 0x72, 0xd8, 0x0a, 0x27, 0x27, 0xe0, 0xc6, 0x98, 0x15, 0xa5, 0xf2, 0x26, 0x56, 0xdc, 0xed,
	0x10, 0x4f, 0x2d, 0x10, 0x39, 0xac, 0x41, 0xc9, 0x19, 0x6c, 0xc5, 0x73, 0xae, 0xcd, 0x05, 0xc6,
	0xdc, 0xa0, 0xf6, 0x5c, 0xab, 0x1e, 0x74, 0xa8, 0x91, 0x50, 0xcb, 0xd3, 0x1b, 0x28, 0x5b, 0x13,
	0xc9, 0x1b, 0x80, 0xeb, 0x9b, 0xf1, 0x06, 0x36, 0xe6, 0x49, 0x47, 0xcc, 0xbb, 0x15, 0xc4, 0x6e,
	0x08, 0xe4, 0x05, 0x0c, 0x44, 0xc6, 0xa7, 0x4a, 0x24, 0xde, 0x96, 0x75, 0xc7, 0xb4, 0x1e, 0x36,
	0xda, 0x0e, 0x1b, 0x0d, 0x11, 0xd5, 0x27, 0xae, 0x4a, 0xc1, 0x5a, 0x34, 0xdc, 0x84, 0x61, 0x75,
	0xf9, 0x5f, 0xcd, 0x65, 0x2e, 0xfc, 0x2f, 0xb0, 0xfd, 0xef, 0x19, 0xc9, 0x5b, 0x18, 0xda, 0x53,
	0x7e, 0xe0, 0x66, 0xde, 0x8c, 0xc1, 0xfe, 0x6d, 0xbd, 0x55, 0x5c, 0xe4, 0xb0, 0x6b, 0x29, 0x74,
	0xa1, 0xbf, 0x90, 0x59, 0xe2, 0x1f, 0xc0, 0xc3, 0x35, 0x8a, 0x10, 0xe8, 0xe7, 0x6d, 0xea, 0x90,
	0xd9, 0x75, 0x78, 0xf8, 0xe3, 0xd7, 0xa4, 0xf7, 0xf9, 0xe9, 0xad, 0x9f, 0x60, 0xbe, 0x48, 0x9b,
	0xc1, 0x9e, 0xba, 0xb6, 0xb7, 0x93, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x11, 0x3d, 0xe2, 0x33,
	0xb4, 0x03, 0x00, 0x00,
}
