package canonical

import (
	"bytes"
	"testing"

	"github.com/koinos/koinos-proto-golang/koinos"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestCanonicalSerialization(t *testing.T) {
	expectedBytes := []byte{0x0A, 0x03, 0x66, 0x6F, 0x6F, 0x10, 0x01, 0x1a, 0x03, 0x62, 0x61, 0x72}
	bt := koinos.BlockTopology{Id: []byte("foo"), Height: 1, Previous: []byte("bar")}

	// Canonically serialize and test the bytes
	b, err := Marshal(&bt)
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(expectedBytes, b))

	// Ensure we can deserialize it
	bt2 := koinos.BlockTopology{}
	proto.Unmarshal(b, &bt2)
	assert.Equal(t, bt.Height, bt2.Height)
	assert.True(t, bytes.Equal(bt.Id, bt2.Id))
	assert.True(t, bytes.Equal(bt.Previous, bt2.Previous))
}

func TestMap(t *testing.T) {
	tm := testMap{}
	b, err := Marshal(&tm)
	assert.Nil(t, b)
	assert.ErrorIs(t, err, ErrNoCanonical)
}
