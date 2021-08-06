// Code generated by capnpc-go. DO NOT EDIT.

package broadcast

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	koinos "koinos"
	protocol "koinos/protocol"
)

type TransactionAccepted struct{ capnp.Struct }

// TransactionAccepted_TypeID is the unique identifier for the type TransactionAccepted.
const TransactionAccepted_TypeID = 0x926c186ab23617c0

func NewTransactionAccepted(s *capnp.Segment) (TransactionAccepted, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return TransactionAccepted{st}, err
}

func NewRootTransactionAccepted(s *capnp.Segment) (TransactionAccepted, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return TransactionAccepted{st}, err
}

func ReadRootTransactionAccepted(msg *capnp.Message) (TransactionAccepted, error) {
	root, err := msg.Root()
	return TransactionAccepted{root.Struct()}, err
}

func (s TransactionAccepted) String() string {
	str, _ := text.Marshal(0x926c186ab23617c0, s.Struct)
	return str
}

func (s TransactionAccepted) Transaction() (protocol.Transaction, error) {
	p, err := s.Struct.Ptr(0)
	return protocol.Transaction{Struct: p.Struct()}, err
}

func (s TransactionAccepted) HasTransaction() bool {
	return s.Struct.HasPtr(0)
}

func (s TransactionAccepted) SetTransaction(v protocol.Transaction) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTransaction sets the transaction field to a newly
// allocated protocol.Transaction struct, preferring placement in s's segment.
func (s TransactionAccepted) NewTransaction() (protocol.Transaction, error) {
	ss, err := protocol.NewTransaction(s.Struct.Segment())
	if err != nil {
		return protocol.Transaction{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s TransactionAccepted) Payer() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s TransactionAccepted) HasPayer() bool {
	return s.Struct.HasPtr(1)
}

func (s TransactionAccepted) SetPayer(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s TransactionAccepted) MaxPayerResources() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s TransactionAccepted) HasMaxPayerResources() bool {
	return s.Struct.HasPtr(2)
}

func (s TransactionAccepted) SetMaxPayerResources(v []byte) error {
	return s.Struct.SetData(2, v)
}

func (s TransactionAccepted) TrxResourceLimit() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s TransactionAccepted) HasTrxResourceLimit() bool {
	return s.Struct.HasPtr(3)
}

func (s TransactionAccepted) SetTrxResourceLimit(v []byte) error {
	return s.Struct.SetData(3, v)
}

func (s TransactionAccepted) Height() uint64 {
	return s.Struct.Uint64(0)
}

func (s TransactionAccepted) SetHeight(v uint64) {
	s.Struct.SetUint64(0, v)
}

// TransactionAccepted_List is a list of TransactionAccepted.
type TransactionAccepted_List struct{ capnp.List }

// NewTransactionAccepted creates a new list of TransactionAccepted.
func NewTransactionAccepted_List(s *capnp.Segment, sz int32) (TransactionAccepted_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return TransactionAccepted_List{l}, err
}

func (s TransactionAccepted_List) At(i int) TransactionAccepted {
	return TransactionAccepted{s.List.Struct(i)}
}

func (s TransactionAccepted_List) Set(i int, v TransactionAccepted) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s TransactionAccepted_List) String() string {
	str, _ := text.MarshalList(0x926c186ab23617c0, s.List)
	return str
}

// TransactionAccepted_Future is a wrapper for a TransactionAccepted promised by a client call.
type TransactionAccepted_Future struct{ *capnp.Future }

func (p TransactionAccepted_Future) Struct() (TransactionAccepted, error) {
	s, err := p.Future.Struct()
	return TransactionAccepted{s}, err
}

func (p TransactionAccepted_Future) Transaction() protocol.Transaction_Future {
	return protocol.Transaction_Future{Future: p.Future.Field(0, nil)}
}

type BlockAccepted struct{ capnp.Struct }

// BlockAccepted_TypeID is the unique identifier for the type BlockAccepted.
const BlockAccepted_TypeID = 0xc180f754a3d036b2

func NewBlockAccepted(s *capnp.Segment) (BlockAccepted, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return BlockAccepted{st}, err
}

func NewRootBlockAccepted(s *capnp.Segment) (BlockAccepted, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return BlockAccepted{st}, err
}

func ReadRootBlockAccepted(msg *capnp.Message) (BlockAccepted, error) {
	root, err := msg.Root()
	return BlockAccepted{root.Struct()}, err
}

func (s BlockAccepted) String() string {
	str, _ := text.Marshal(0xc180f754a3d036b2, s.Struct)
	return str
}

func (s BlockAccepted) Block() (protocol.Block, error) {
	p, err := s.Struct.Ptr(0)
	return protocol.Block{Struct: p.Struct()}, err
}

func (s BlockAccepted) HasBlock() bool {
	return s.Struct.HasPtr(0)
}

func (s BlockAccepted) SetBlock(v protocol.Block) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBlock sets the block field to a newly
// allocated protocol.Block struct, preferring placement in s's segment.
func (s BlockAccepted) NewBlock() (protocol.Block, error) {
	ss, err := protocol.NewBlock(s.Struct.Segment())
	if err != nil {
		return protocol.Block{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// BlockAccepted_List is a list of BlockAccepted.
type BlockAccepted_List struct{ capnp.List }

// NewBlockAccepted creates a new list of BlockAccepted.
func NewBlockAccepted_List(s *capnp.Segment, sz int32) (BlockAccepted_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return BlockAccepted_List{l}, err
}

func (s BlockAccepted_List) At(i int) BlockAccepted { return BlockAccepted{s.List.Struct(i)} }

func (s BlockAccepted_List) Set(i int, v BlockAccepted) error { return s.List.SetStruct(i, v.Struct) }

func (s BlockAccepted_List) String() string {
	str, _ := text.MarshalList(0xc180f754a3d036b2, s.List)
	return str
}

// BlockAccepted_Future is a wrapper for a BlockAccepted promised by a client call.
type BlockAccepted_Future struct{ *capnp.Future }

func (p BlockAccepted_Future) Struct() (BlockAccepted, error) {
	s, err := p.Future.Struct()
	return BlockAccepted{s}, err
}

func (p BlockAccepted_Future) Block() protocol.Block_Future {
	return protocol.Block_Future{Future: p.Future.Field(0, nil)}
}

type BlockIrreversible struct{ capnp.Struct }

// BlockIrreversible_TypeID is the unique identifier for the type BlockIrreversible.
const BlockIrreversible_TypeID = 0xfb289a76e96e2384

func NewBlockIrreversible(s *capnp.Segment) (BlockIrreversible, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return BlockIrreversible{st}, err
}

func NewRootBlockIrreversible(s *capnp.Segment) (BlockIrreversible, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return BlockIrreversible{st}, err
}

func ReadRootBlockIrreversible(msg *capnp.Message) (BlockIrreversible, error) {
	root, err := msg.Root()
	return BlockIrreversible{root.Struct()}, err
}

func (s BlockIrreversible) String() string {
	str, _ := text.Marshal(0xfb289a76e96e2384, s.Struct)
	return str
}

func (s BlockIrreversible) Topology() (koinos.BlockTopology, error) {
	p, err := s.Struct.Ptr(0)
	return koinos.BlockTopology{Struct: p.Struct()}, err
}

func (s BlockIrreversible) HasTopology() bool {
	return s.Struct.HasPtr(0)
}

func (s BlockIrreversible) SetTopology(v koinos.BlockTopology) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTopology sets the topology field to a newly
// allocated koinos.BlockTopology struct, preferring placement in s's segment.
func (s BlockIrreversible) NewTopology() (koinos.BlockTopology, error) {
	ss, err := koinos.NewBlockTopology(s.Struct.Segment())
	if err != nil {
		return koinos.BlockTopology{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// BlockIrreversible_List is a list of BlockIrreversible.
type BlockIrreversible_List struct{ capnp.List }

// NewBlockIrreversible creates a new list of BlockIrreversible.
func NewBlockIrreversible_List(s *capnp.Segment, sz int32) (BlockIrreversible_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return BlockIrreversible_List{l}, err
}

func (s BlockIrreversible_List) At(i int) BlockIrreversible {
	return BlockIrreversible{s.List.Struct(i)}
}

func (s BlockIrreversible_List) Set(i int, v BlockIrreversible) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s BlockIrreversible_List) String() string {
	str, _ := text.MarshalList(0xfb289a76e96e2384, s.List)
	return str
}

// BlockIrreversible_Future is a wrapper for a BlockIrreversible promised by a client call.
type BlockIrreversible_Future struct{ *capnp.Future }

func (p BlockIrreversible_Future) Struct() (BlockIrreversible, error) {
	s, err := p.Future.Struct()
	return BlockIrreversible{s}, err
}

func (p BlockIrreversible_Future) Topology() koinos.BlockTopology_Future {
	return koinos.BlockTopology_Future{Future: p.Future.Field(0, nil)}
}

type ForkHeads struct{ capnp.Struct }

// ForkHeads_TypeID is the unique identifier for the type ForkHeads.
const ForkHeads_TypeID = 0xc344d482fcf3e3e4

func NewForkHeads(s *capnp.Segment) (ForkHeads, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return ForkHeads{st}, err
}

func NewRootForkHeads(s *capnp.Segment) (ForkHeads, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return ForkHeads{st}, err
}

func ReadRootForkHeads(msg *capnp.Message) (ForkHeads, error) {
	root, err := msg.Root()
	return ForkHeads{root.Struct()}, err
}

func (s ForkHeads) String() string {
	str, _ := text.Marshal(0xc344d482fcf3e3e4, s.Struct)
	return str
}

func (s ForkHeads) ForkHeads() (koinos.BlockTopology_List, error) {
	p, err := s.Struct.Ptr(0)
	return koinos.BlockTopology_List{List: p.List()}, err
}

func (s ForkHeads) HasForkHeads() bool {
	return s.Struct.HasPtr(0)
}

func (s ForkHeads) SetForkHeads(v koinos.BlockTopology_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewForkHeads sets the forkHeads field to a newly
// allocated koinos.BlockTopology_List, preferring placement in s's segment.
func (s ForkHeads) NewForkHeads(n int32) (koinos.BlockTopology_List, error) {
	l, err := koinos.NewBlockTopology_List(s.Struct.Segment(), n)
	if err != nil {
		return koinos.BlockTopology_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s ForkHeads) LastIrreversibleBlock() (koinos.BlockTopology, error) {
	p, err := s.Struct.Ptr(1)
	return koinos.BlockTopology{Struct: p.Struct()}, err
}

func (s ForkHeads) HasLastIrreversibleBlock() bool {
	return s.Struct.HasPtr(1)
}

func (s ForkHeads) SetLastIrreversibleBlock(v koinos.BlockTopology) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewLastIrreversibleBlock sets the lastIrreversibleBlock field to a newly
// allocated koinos.BlockTopology struct, preferring placement in s's segment.
func (s ForkHeads) NewLastIrreversibleBlock() (koinos.BlockTopology, error) {
	ss, err := koinos.NewBlockTopology(s.Struct.Segment())
	if err != nil {
		return koinos.BlockTopology{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// ForkHeads_List is a list of ForkHeads.
type ForkHeads_List struct{ capnp.List }

// NewForkHeads creates a new list of ForkHeads.
func NewForkHeads_List(s *capnp.Segment, sz int32) (ForkHeads_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return ForkHeads_List{l}, err
}

func (s ForkHeads_List) At(i int) ForkHeads { return ForkHeads{s.List.Struct(i)} }

func (s ForkHeads_List) Set(i int, v ForkHeads) error { return s.List.SetStruct(i, v.Struct) }

func (s ForkHeads_List) String() string {
	str, _ := text.MarshalList(0xc344d482fcf3e3e4, s.List)
	return str
}

// ForkHeads_Future is a wrapper for a ForkHeads promised by a client call.
type ForkHeads_Future struct{ *capnp.Future }

func (p ForkHeads_Future) Struct() (ForkHeads, error) {
	s, err := p.Future.Struct()
	return ForkHeads{s}, err
}

func (p ForkHeads_Future) LastIrreversibleBlock() koinos.BlockTopology_Future {
	return koinos.BlockTopology_Future{Future: p.Future.Field(1, nil)}
}

const schema_8bc2b44efe5ccdc1 = "x\xda\x8c\x92OH\x14o\x18\xc7\xbf\xdf\xf7\x9du\xd6" +
	"\x1f\xeao\x87]H\xbc\x0c\xd5!\x83\xc84\x10\xf2b" +
	"+\x16\x15Q\xfb\xaat\xaa\xc3\xec\xec\xa4\xab\xeb\xce2" +
	"3\x99RAIB\xf4\xe7\x90\xc7 \x08\xea\x12\xd4!" +
	")\xa8C\x90z\xf0R\xd0\xa1\xa0K\x97\xfe\x1c\x0a\xa2" +
	"\xa0CB\xd6\x1b\xb3\x98,\x1a\xd9\xeda\xf8<\x9f\xe7" +
	"y\xbe\xf3\xee8,v\x8b\xf6\xc4\xff\x0d\x80\xba\x9c\xa8" +
	"\xd3O6t\xce\x0c7\x97\xa6\xa1ZH=\xfb\xec\xe8" +
	"\xcfC\xf7\xe7.!a\x98\xc0\xceO\x89\x80i\xd6\xc5" +
	"\xe5\x8f\x84MP\xcft>\xbf9\xf0\xed\xec,\xac\x96" +
	"Z\x9a1r\xce\xeca\xfa\xaai\x02\xe9+f7\xa8" +
	"\xdf\xbd\xfd\xba4\xf9\xa2w~\x15,b\xf8\xae\xb9\x89" +
	"\xe9\xc7U\xf8\x91y\x12\xd4\xe77\x97?\x8c]k\xfd" +
	"\xfe's{2\xcf\xf4\x9ed\x0cg\x93\xdd\xd08\xa2" +
	"\xf3\x81\xef\x14\\'4\xa2\xb6\xdfe\xb4\xddu*\xe5" +
	"J\xd7@\xe0\x94C\xc7\x8d\x8a~9\xeb\xba\xb6W\x89" +
	"\xbcB\x8eT\x19i\x00\x06\x01\xebL\x1eP\xa7%\xd5" +
	"\x05A\x8b\xcc0\xfe8\xd5aM\xd9\xea\x86\xa4\xba#" +
	"h\x09\x91\xa1\x00\xac\xdb\xd3\xd6=[=\x95T\xaf\x04" +
	"-)3\x94\x80\xf5\xf2\xa2\xf5\xdaV\x8b\x92\xfd\x06\x05" +
	"idh\x00i\xb2\x0bPK\x92\xfdI\x0a\xeahy" +
	"\x0f\x98E\xbf\xcc\x94>\xb6\xe5\xd4\xae\xc9\x8f\xc3\xb7\x00" +
	"2\x05\xda\x15g\xc2\x0b\x94A\xa1\x1f,\xbc\xff\xaf\xed" +
	"\x8d\xf3\x05\xca\x10\xcc6\x10\x92\x8d\x10l\x04\xf5\xa83" +
	"\x9e\x8b9\xf6y\xa1\x7f\"p=\x86\xeb\xb5D\xc1x" +
	"\x15f\xe0z\x07\x8b\xa3\xc5\x08X\xa7\xa5{\xc8+\x0e" +
	"\x0eE\xac\x87`=\xf8\xb7p{J\xbe;\x92u\xdd" +
	"j\xaa\x88c5Vbm\xec\x00TRRe\x04\xed" +
	"|\x0c2\xa5\xaf\x0f\xf6\xcd\xcdw<\\\\\xbez\xc5" +
	"-\xd7\xba\xf7\xfa\xc1\xc8>\xcft\x0aa\xecM\xaex" +
	"\xb7\xf6\x01\xaaUR\xf5\xd6\xfc\xae\xec\x0c\xa0z%U" +
	"NP\x1f\xafv:\x050d\x13\x98\x93dJW\xb6" +
	"\x8dMm\xfc\xbc\xa0\xe3\xc9M\xa0.9a\xb4?\x08" +
	"<\x8eyAX\xcc\x97<\xbbgy\xc7Z2\xf5\x0f" +
	"\xf7\xc7\x96\xaa\xc4\xcc\x97\xbcU\x19\x1c\x00T\x83\xa4j" +
	"\x8e_\x80_\xf1K\xfe\xe0\x04\x80\xb5S~\x05\x00\x00" +
	"\xff\xffEH\xfc%"

func init() {
	schemas.Register(schema_8bc2b44efe5ccdc1,
		0x926c186ab23617c0,
		0xc180f754a3d036b2,
		0xc344d482fcf3e3e4,
		0xfb289a76e96e2384)
}
