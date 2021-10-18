package json

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	gojson "encoding/json"
	"reflect"
	"testing"

	"github.com/btcsuite/btcutil/base58"
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

	jsonMap := make(map[string]string)
	expectedMap := make(map[string]string)
	expectedMap["base_64"] = base64Str
	expectedMap["base_58"] = base58Str
	expectedMap["hex"] = hexStr
	expectedMap["block_id"] = hexStr
	expectedMap["transaction_id"] = hexStr
	expectedMap["contract_id"] = base58Str
	expectedMap["address"] = base58Str
	expectedMap["default"] = base64Str

	jsonBytes, err := Marshal(testObj)
	assert.NoError(t, err)
	err = gojson.Unmarshal(jsonBytes, &jsonMap)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(jsonMap, expectedMap))

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
	jsonMap["base_64"] = hexStr
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap["base_64"] = base64Str
	jsonMap["base_58"] = hexStr
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap["base_58"] = base58Str
	jsonMap["hex"] = base64Str
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap["hex"] = hexStr
	jsonMap["block_id"] = base64Str
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap["block_id"] = hexStr
	jsonMap["transaction_id"] = base64Str
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap["transaction_id"] = hexStr
	jsonMap["contract_id"] = hexStr
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap["contract_id"] = base58Str
	jsonMap["address"] = hexStr
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)

	jsonMap["address"] = base58Str
	jsonMap["default"] = hexStr
	jsonBytes, _ = gojson.Marshal(jsonMap)
	err = Unmarshal(jsonBytes, unmarshalledObject)
	assert.Error(t, err)
}
