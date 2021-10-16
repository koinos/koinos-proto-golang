package json

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	gojson "encoding/json"
	"fmt"
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/iancoleman/orderedmap"
	"github.com/stretchr/testify/assert"
)

func TestByteEncodings(t *testing.T) {
	testData := []byte{0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72}
	testObj := &BytesSerializationTestObject{}
	testObj.Base_64 = testData
	testObj.Base_58 = testData
	testObj.Hex = testData
	testObj.BlockId = testData
	testObj.TransactionId = testData
	testObj.ContractId = testData
	testObj.Address = testData
	testObj.Default = testData

	hexStr := "0x" + hex.EncodeToString(testData)
	base58Str := base58.Encode(testData)
	base64Str := base64.URLEncoding.EncodeToString(testData)

	jsonMap := orderedmap.New()
	jsonMap.Set("base_64", base64Str)
	jsonMap.Set("base_58", base58Str)
	jsonMap.Set("hex", hexStr)
	jsonMap.Set("block_id", hexStr)
	jsonMap.Set("transaction_id", hexStr)
	jsonMap.Set("contract_id", base58Str)
	jsonMap.Set("address", base58Str)
	jsonMap.Set("default", base64Str)

	expectedBytes, _ := gojson.Marshal(jsonMap)

	jsonBytes, err := Marshal(testObj)
	assert.NoError(t, err)
	assert.True(t, string(expectedBytes) == string(jsonBytes))

	fmt.Printf(string(expectedBytes) + "\n")
	fmt.Printf(string(jsonBytes) + "\n")

	unmarshalledObject := &BytesSerializationTestObject{}
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(testObj.Base_64, unmarshalledObject.Base_64))
	assert.True(t, bytes.Equal(testObj.Base_58, unmarshalledObject.Base_58))
	assert.True(t, bytes.Equal(testObj.Hex, unmarshalledObject.Hex))
	assert.True(t, bytes.Equal(testObj.BlockId, unmarshalledObject.BlockId))
	assert.True(t, bytes.Equal(testObj.TransactionId, unmarshalledObject.TransactionId))
	assert.True(t, bytes.Equal(testObj.ContractId, unmarshalledObject.ContractId))
	assert.True(t, bytes.Equal(testObj.Address, unmarshalledObject.Address))

	// Test bad encodings
	jsonMap.Set("base_64", hexStr)
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap.Set("base_64", base64Str)
	jsonMap.Set("base_58", hexStr)
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap.Set("base_58", base58Str)
	jsonMap.Set("hex", base64Str)
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap.Set("hex", hexStr)
	jsonMap.Set("block_id", base64Str)
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap.Set("block_id", hexStr)
	jsonMap.Set("transaction_id", base64Str)
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap.Set("transaction_id", hexStr)
	jsonMap.Set("contract_id", hexStr)
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap.Set("contract_id", base58Str)
	jsonMap.Set("address", hexStr)
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap.Set("address", base58Str)
	jsonMap.Set("default", hexStr)
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)
}
