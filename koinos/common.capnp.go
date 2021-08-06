// Code generated by capnpc-go. DO NOT EDIT.

package koinos

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	strconv "strconv"
)

type BlockTopology struct{ capnp.Struct }

// BlockTopology_TypeID is the unique identifier for the type BlockTopology.
const BlockTopology_TypeID = 0xffc7ef2185762c70

func NewBlockTopology(s *capnp.Segment) (BlockTopology, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return BlockTopology{st}, err
}

func NewRootBlockTopology(s *capnp.Segment) (BlockTopology, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return BlockTopology{st}, err
}

func ReadRootBlockTopology(msg *capnp.Message) (BlockTopology, error) {
	root, err := msg.Root()
	return BlockTopology{root.Struct()}, err
}

func (s BlockTopology) String() string {
	str, _ := text.Marshal(0xffc7ef2185762c70, s.Struct)
	return str
}

func (s BlockTopology) Id() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s BlockTopology) HasId() bool {
	return s.Struct.HasPtr(0)
}

func (s BlockTopology) SetId(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s BlockTopology) Height() uint64 {
	return s.Struct.Uint64(0)
}

func (s BlockTopology) SetHeight(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s BlockTopology) Previous() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s BlockTopology) HasPrevious() bool {
	return s.Struct.HasPtr(1)
}

func (s BlockTopology) SetPrevious(v []byte) error {
	return s.Struct.SetData(1, v)
}

// BlockTopology_List is a list of BlockTopology.
type BlockTopology_List struct{ capnp.List }

// NewBlockTopology creates a new list of BlockTopology.
func NewBlockTopology_List(s *capnp.Segment, sz int32) (BlockTopology_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return BlockTopology_List{l}, err
}

func (s BlockTopology_List) At(i int) BlockTopology { return BlockTopology{s.List.Struct(i)} }

func (s BlockTopology_List) Set(i int, v BlockTopology) error { return s.List.SetStruct(i, v.Struct) }

func (s BlockTopology_List) String() string {
	str, _ := text.MarshalList(0xffc7ef2185762c70, s.List)
	return str
}

// BlockTopology_Future is a wrapper for a BlockTopology promised by a client call.
type BlockTopology_Future struct{ *capnp.Future }

func (p BlockTopology_Future) Struct() (BlockTopology, error) {
	s, err := p.Future.Struct()
	return BlockTopology{s}, err
}

type ContractCallBundle struct{ capnp.Struct }

// ContractCallBundle_TypeID is the unique identifier for the type ContractCallBundle.
const ContractCallBundle_TypeID = 0xa11cdc7f34b73147

func NewContractCallBundle(s *capnp.Segment) (ContractCallBundle, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return ContractCallBundle{st}, err
}

func NewRootContractCallBundle(s *capnp.Segment) (ContractCallBundle, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return ContractCallBundle{st}, err
}

func ReadRootContractCallBundle(msg *capnp.Message) (ContractCallBundle, error) {
	root, err := msg.Root()
	return ContractCallBundle{root.Struct()}, err
}

func (s ContractCallBundle) String() string {
	str, _ := text.Marshal(0xa11cdc7f34b73147, s.Struct)
	return str
}

func (s ContractCallBundle) ContractID() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s ContractCallBundle) HasContractID() bool {
	return s.Struct.HasPtr(0)
}

func (s ContractCallBundle) SetContractID(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s ContractCallBundle) EntryPoint() uint32 {
	return s.Struct.Uint32(0)
}

func (s ContractCallBundle) SetEntryPoint(v uint32) {
	s.Struct.SetUint32(0, v)
}

// ContractCallBundle_List is a list of ContractCallBundle.
type ContractCallBundle_List struct{ capnp.List }

// NewContractCallBundle creates a new list of ContractCallBundle.
func NewContractCallBundle_List(s *capnp.Segment, sz int32) (ContractCallBundle_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return ContractCallBundle_List{l}, err
}

func (s ContractCallBundle_List) At(i int) ContractCallBundle {
	return ContractCallBundle{s.List.Struct(i)}
}

func (s ContractCallBundle_List) Set(i int, v ContractCallBundle) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s ContractCallBundle_List) String() string {
	str, _ := text.MarshalList(0xa11cdc7f34b73147, s.List)
	return str
}

// ContractCallBundle_Future is a wrapper for a ContractCallBundle promised by a client call.
type ContractCallBundle_Future struct{ *capnp.Future }

func (p ContractCallBundle_Future) Struct() (ContractCallBundle, error) {
	s, err := p.Future.Struct()
	return ContractCallBundle{s}, err
}

type SystemCallTarget struct{ capnp.Struct }
type SystemCallTarget_Which uint16

const (
	SystemCallTarget_Which_thunkID          SystemCallTarget_Which = 0
	SystemCallTarget_Which_systemCallBundle SystemCallTarget_Which = 1
)

func (w SystemCallTarget_Which) String() string {
	const s = "thunkIDsystemCallBundle"
	switch w {
	case SystemCallTarget_Which_thunkID:
		return s[0:7]
	case SystemCallTarget_Which_systemCallBundle:
		return s[7:23]

	}
	return "SystemCallTarget_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// SystemCallTarget_TypeID is the unique identifier for the type SystemCallTarget.
const SystemCallTarget_TypeID = 0xbc2f08e5acdea3aa

func NewSystemCallTarget(s *capnp.Segment) (SystemCallTarget, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return SystemCallTarget{st}, err
}

func NewRootSystemCallTarget(s *capnp.Segment) (SystemCallTarget, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return SystemCallTarget{st}, err
}

func ReadRootSystemCallTarget(msg *capnp.Message) (SystemCallTarget, error) {
	root, err := msg.Root()
	return SystemCallTarget{root.Struct()}, err
}

func (s SystemCallTarget) String() string {
	str, _ := text.Marshal(0xbc2f08e5acdea3aa, s.Struct)
	return str
}

func (s SystemCallTarget) Which() SystemCallTarget_Which {
	return SystemCallTarget_Which(s.Struct.Uint16(4))
}
func (s SystemCallTarget) ThunkID() uint32 {
	if s.Struct.Uint16(4) != 0 {
		panic("Which() != thunkID")
	}
	return s.Struct.Uint32(0)
}

func (s SystemCallTarget) SetThunkID(v uint32) {
	s.Struct.SetUint16(4, 0)
	s.Struct.SetUint32(0, v)
}

func (s SystemCallTarget) SystemCallBundle() (ContractCallBundle, error) {
	if s.Struct.Uint16(4) != 1 {
		panic("Which() != systemCallBundle")
	}
	p, err := s.Struct.Ptr(0)
	return ContractCallBundle{Struct: p.Struct()}, err
}

func (s SystemCallTarget) HasSystemCallBundle() bool {
	if s.Struct.Uint16(4) != 1 {
		return false
	}
	return s.Struct.HasPtr(0)
}

func (s SystemCallTarget) SetSystemCallBundle(v ContractCallBundle) error {
	s.Struct.SetUint16(4, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSystemCallBundle sets the systemCallBundle field to a newly
// allocated ContractCallBundle struct, preferring placement in s's segment.
func (s SystemCallTarget) NewSystemCallBundle() (ContractCallBundle, error) {
	s.Struct.SetUint16(4, 1)
	ss, err := NewContractCallBundle(s.Struct.Segment())
	if err != nil {
		return ContractCallBundle{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// SystemCallTarget_List is a list of SystemCallTarget.
type SystemCallTarget_List struct{ capnp.List }

// NewSystemCallTarget creates a new list of SystemCallTarget.
func NewSystemCallTarget_List(s *capnp.Segment, sz int32) (SystemCallTarget_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return SystemCallTarget_List{l}, err
}

func (s SystemCallTarget_List) At(i int) SystemCallTarget { return SystemCallTarget{s.List.Struct(i)} }

func (s SystemCallTarget_List) Set(i int, v SystemCallTarget) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s SystemCallTarget_List) String() string {
	str, _ := text.MarshalList(0xbc2f08e5acdea3aa, s.List)
	return str
}

// SystemCallTarget_Future is a wrapper for a SystemCallTarget promised by a client call.
type SystemCallTarget_Future struct{ *capnp.Future }

func (p SystemCallTarget_Future) Struct() (SystemCallTarget, error) {
	s, err := p.Future.Struct()
	return SystemCallTarget{s}, err
}

func (p SystemCallTarget_Future) SystemCallBundle() ContractCallBundle_Future {
	return ContractCallBundle_Future{Future: p.Future.Field(0, nil)}
}

type Opaque struct{ capnp.Struct }
type Opaque_Which uint16

const (
	Opaque_Which_native Opaque_Which = 0
	Opaque_Which_bytes  Opaque_Which = 1
)

func (w Opaque_Which) String() string {
	const s = "nativebytes"
	switch w {
	case Opaque_Which_native:
		return s[0:6]
	case Opaque_Which_bytes:
		return s[6:11]

	}
	return "Opaque_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Opaque_TypeID is the unique identifier for the type Opaque.
const Opaque_TypeID = 0x997dbb766f2696fa

func NewOpaque(s *capnp.Segment) (Opaque, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Opaque{st}, err
}

func NewRootOpaque(s *capnp.Segment) (Opaque, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Opaque{st}, err
}

func ReadRootOpaque(msg *capnp.Message) (Opaque, error) {
	root, err := msg.Root()
	return Opaque{root.Struct()}, err
}

func (s Opaque) String() string {
	str, _ := text.Marshal(0x997dbb766f2696fa, s.Struct)
	return str
}

func (s Opaque) Which() Opaque_Which {
	return Opaque_Which(s.Struct.Uint16(0))
}
func (s Opaque) Native() (capnp.Ptr, error) {
	if s.Struct.Uint16(0) != 0 {
		panic("Which() != native")
	}
	return s.Struct.Ptr(0)
}

func (s Opaque) HasNative() bool {
	if s.Struct.Uint16(0) != 0 {
		return false
	}
	return s.Struct.HasPtr(0)
}

func (s Opaque) SetNative(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 0)
	return s.Struct.SetPtr(0, v)
}

func (s Opaque) Bytes() ([]byte, error) {
	if s.Struct.Uint16(0) != 1 {
		panic("Which() != bytes")
	}
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Opaque) HasBytes() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	return s.Struct.HasPtr(0)
}

func (s Opaque) SetBytes(v []byte) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetData(0, v)
}

// Opaque_List is a list of Opaque.
type Opaque_List struct{ capnp.List }

// NewOpaque creates a new list of Opaque.
func NewOpaque_List(s *capnp.Segment, sz int32) (Opaque_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Opaque_List{l}, err
}

func (s Opaque_List) At(i int) Opaque { return Opaque{s.List.Struct(i)} }

func (s Opaque_List) Set(i int, v Opaque) error { return s.List.SetStruct(i, v.Struct) }

func (s Opaque_List) String() string {
	str, _ := text.MarshalList(0x997dbb766f2696fa, s.List)
	return str
}

// Opaque_Future is a wrapper for a Opaque promised by a client call.
type Opaque_Future struct{ *capnp.Future }

func (p Opaque_Future) Struct() (Opaque, error) {
	s, err := p.Future.Struct()
	return Opaque{s}, err
}

func (p Opaque_Future) Native() *capnp.Future {
	return p.Future.Field(0, nil)
}

type Optional struct{ capnp.Struct }
type Optional_Which uint16

const (
	Optional_Which_null  Optional_Which = 0
	Optional_Which_value Optional_Which = 1
)

func (w Optional_Which) String() string {
	const s = "nullvalue"
	switch w {
	case Optional_Which_null:
		return s[0:4]
	case Optional_Which_value:
		return s[4:9]

	}
	return "Optional_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Optional_TypeID is the unique identifier for the type Optional.
const Optional_TypeID = 0xb6084849c378a152

func NewOptional(s *capnp.Segment) (Optional, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Optional{st}, err
}

func NewRootOptional(s *capnp.Segment) (Optional, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Optional{st}, err
}

func ReadRootOptional(msg *capnp.Message) (Optional, error) {
	root, err := msg.Root()
	return Optional{root.Struct()}, err
}

func (s Optional) String() string {
	str, _ := text.Marshal(0xb6084849c378a152, s.Struct)
	return str
}

func (s Optional) Which() Optional_Which {
	return Optional_Which(s.Struct.Uint16(0))
}
func (s Optional) SetNull() {
	s.Struct.SetUint16(0, 0)

}

func (s Optional) Value() (capnp.Ptr, error) {
	if s.Struct.Uint16(0) != 1 {
		panic("Which() != value")
	}
	return s.Struct.Ptr(0)
}

func (s Optional) HasValue() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	return s.Struct.HasPtr(0)
}

func (s Optional) SetValue(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v)
}

// Optional_List is a list of Optional.
type Optional_List struct{ capnp.List }

// NewOptional creates a new list of Optional.
func NewOptional_List(s *capnp.Segment, sz int32) (Optional_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Optional_List{l}, err
}

func (s Optional_List) At(i int) Optional { return Optional{s.List.Struct(i)} }

func (s Optional_List) Set(i int, v Optional) error { return s.List.SetStruct(i, v.Struct) }

func (s Optional_List) String() string {
	str, _ := text.MarshalList(0xb6084849c378a152, s.List)
	return str
}

// Optional_Future is a wrapper for a Optional promised by a client call.
type Optional_Future struct{ *capnp.Future }

func (p Optional_Future) Struct() (Optional, error) {
	s, err := p.Future.Struct()
	return Optional{s}, err
}

func (p Optional_Future) Value() *capnp.Future {
	return p.Future.Field(0, nil)
}

const schema_8aa958e2a5aa42b1 = "x\xda\x84\x92\xc1K\x14o\x18\xc7\x9f\xef\xf3\x8e\xee*" +
	".\xeb8\x9e~\xfc`\x08*J2\xb5<\x84D\xe6" +
	"\xea\x92[D\xbe\xb5AI\x87\xa6\xddA\x07gg\xc6" +
	"\xdd\xd9\xb5=\x84\xa7.\xd29\xeah\xd1!\x09\xa2\xa0" +
	":\xd4\xad\x83\x7fA\x18\x91\x81\x82\x07OE\xc7\x0e\xbd" +
	"1\xbb\xb2+\xeb\x82\xb7e\x99\xef\xf7\xf9\xbc\x9f\xe7\x19" +
	"^\xc1E\x1e\xe9\xd8\xd0\x88\xe4\xb9\x8eN\xf5\xe7\xf1q" +
	"\xbf\xf2\xf1\xc1S\x92\x09@\xbdI\xad\xbd\xd8\xbe\xf5r" +
	"\x85\xd2\x88q/\x0c\xc9\xcf\x8c\xdb\x1c#2n\xf2\x12" +
	"\x91\xb1\xcbIui\xe4\xc3\xe8\xf2\xf7\xffW[\x12\x1d" +
	"\x88\x11\x9d\xddd\x86\xb1[K\xec\xf0\x12A]_\xbd" +
	"\xff93\x1d\x7f\xdf\xbe\xdf\x11o\x8dE\x11}]\x10" +
	"Q\xff\x96H\xaa\xb5\xe7?^\xed\xc4\x87>\xb5I\x80" +
	"\x8d/\xe2\xb7\xb1UKl\x8a\xa8?8Uyx\xe4" +
	"\xe7\xbaj\xa5\xa9!Lh\xdf\x8c\xabZ\xf4+\xa3\xbd" +
	"\xa6\xbc\xca\xf9\x85\x82\xef\x9d\xce\xc1\x0a\xbc`\xecZ`" +
	"\xc5\x16\xcb\xf6\x0c \xe3B\xebQJ\x03\x91~r\x8c" +
	"H\x1e\x15\x90\xc3\x8c\x04\xfe\xaa~D\xff\x0e\x9e\xd1\x07" +
	"M9- \xb3\x8cq\xcf\x0a\x9d\x8a\x8d>\xa0\xe9\x8f" +
	"\x08}\x04\xf3^5\xb4KR\x03\xab+\x17\x8eu\xa7" +
	"\xab_7Hj\x8c\x89\x1e\x90@\x82\x18\x09\xc2\x84\x06" +
	"\x1d\x03\xc9l5\xb0\x1bH\xa2\x864\xe9{a\xd1\xca" +
	"\x85\x93\x96\xeb\xa6\xca^\xde\xb5i\x0f\x8f\xa8N7\xbb" +
	"\x8f\x03\xa8\xc3\xc9Y\"9# \xef0Tn\xaf\x82" +
	"Df\xaa\xc6\xf1n}\xa7{h\xdb\xfau\x90C\xd9" +
	"^X\xac\xce\xf8\x0e\x09/D\x9c\x18q\xc2\x01I\xa1" +
	"\xe9\xf8\x9e\xe5\xb6j\x1ah\xaf\x89H\x9e\x10\x90\xa3\x8c" +
	"\xa4Wv]\xea4+\x96[\xae\xbbj\xdcB\xddU" +
	";\x0d\\\x1bz\xa3Z\x0a\xedB$!kZ\xc59" +
	";l\x1d\x9ej;|\x85H\x0e\x0b\xc8\xf3\x8c\xe5p" +
	"\xbe\xec-d\xa6\x1a\xaf*\xedU\xa2!\x96\xd0\xdb\xbc" +
	"e\x02z\xf7\xbd\xbd\x8e\x91r\xfd\xdcB\xd6\x0f|\xd7" +
	"\x17s\xd5\x88\xa1\xa7\xb1\x88\xf4\x7fz\xda\x94w\x05\xa4" +
	"\xdb\\\x84\x13\xddN^@\x06\x0c\x9d\xd1\x0f&\xd2\x0b" +
	"\x97\xf5ES>\x12\x90O\x18\xc2\xc9\x1f\xb2\x94\xf1y" +
	"\xdb\x99\x9b\x0f\xd1E\x8c\xae\xe8\xc0\x8bv\xc5\xf1\xcb%" +
	"\":$\xf9/\x00\x00\xff\xff\xbf3\xfe\x15"

func init() {
	schemas.Register(schema_8aa958e2a5aa42b1,
		0x997dbb766f2696fa,
		0xa11cdc7f34b73147,
		0xb6084849c378a152,
		0xbc2f08e5acdea3aa,
		0xffc7ef2185762c70)
}
