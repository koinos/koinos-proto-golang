package json

import (
	"encoding/base64"
	"encoding/hex"
	gojson "encoding/json"
	"errors"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/btcsuite/btcutil/base58"
	"github.com/koinos/koinos-proto-golang/koinos"
)

var (
	KoinosFieldOverrides = map[protoreflect.Kind]protojson.FieldOverride{
		protoreflect.BytesKind: protojson.FieldOverride{
			MarshalField: func(val protoreflect.Value, fd protoreflect.FieldDescriptor, opt protojson.MarshalOptions) ([]byte, error) {
				var encodedStr string
				opts := fd.Options()

				if opts != nil {
					fieldOpts := opts.(*descriptorpb.FieldOptions)
					ext := koinos.E_KoinosBytesType.TypeDescriptor()
					enum := fieldOpts.ProtoReflect().Get(ext).Enum()

					switch koinos.BytesType(enum) {
					case koinos.BytesType_HEX:
						fallthrough
					case koinos.BytesType_BLOCK_ID:
						fallthrough
					case koinos.BytesType_TRANSACTION_ID:
						encodedStr = "0x" + hex.EncodeToString(val.Bytes())
					case koinos.BytesType_BASE58:
						fallthrough
					case koinos.BytesType_CONTRACT_ID:
						fallthrough
					case koinos.BytesType_ADDRESS:
						encodedStr = base58.Encode(val.Bytes())
					case koinos.BytesType_BASE64:
						fallthrough
					default:
						encodedStr = base64.URLEncoding.EncodeToString(val.Bytes())
					}
				} else {
					encodedStr = base64.URLEncoding.EncodeToString(val.Bytes())
				}

				return gojson.Marshal(encodedStr)
			},
			UnmarshalField: func(b []byte, fd protoreflect.FieldDescriptor, opt protojson.UnmarshalOptions) (protoreflect.Value, error) {
				val := protoreflect.Value{}

				var decodedBytes []byte
				var encodedStr string
				err := gojson.Unmarshal(b, &encodedStr)
				if err != nil {
					return val, err
				}

				opts := fd.Options()
				if opts != nil {
					fieldOpts := opts.(*descriptorpb.FieldOptions)
					ext := koinos.E_KoinosBytesType.TypeDescriptor()
					enum := fieldOpts.ProtoReflect().Get(ext).Enum()

					switch koinos.BytesType(enum) {
					case koinos.BytesType_HEX:
						fallthrough
					case koinos.BytesType_BLOCK_ID:
						fallthrough
					case koinos.BytesType_TRANSACTION_ID:
						{
							if len(encodedStr) < 2 || encodedStr[:2] != "0x" {
								return val, errors.New("hex string not prepended with '0x'")
							}

							decodedBytes, err = hex.DecodeString(encodedStr[2:])
						}
					case koinos.BytesType_BASE58:
						fallthrough
					case koinos.BytesType_CONTRACT_ID:
						fallthrough
					case koinos.BytesType_ADDRESS:
						{
							decodedBytes = base58.Decode(encodedStr)
							if len(decodedBytes) == 0 && len(encodedStr) != 0 {
								err = errors.New("error decoding base58")
							}
						}
					case koinos.BytesType_BASE64:
						fallthrough
					default:
						{
							decodedBytes, err = base64.URLEncoding.DecodeString(encodedStr)
						}
					}
				} else {
					decodedBytes, err = base64.URLEncoding.DecodeString(encodedStr)
				}

				return protoreflect.ValueOfBytes(decodedBytes), err
			},
		},
	}

	KoinosMarshalOptions = protojson.MarshalOptions{
		Multiline:      false,
		Indent:         "",
		UseProtoNames:  true,
		FieldOverrides: KoinosFieldOverrides,
	}

	KoinosUnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
		FieldOverrides: KoinosFieldOverrides,
	}
)

func Marshal(m proto.Message) ([]byte, error) {
	return KoinosMarshalOptions.Marshal(m)
}

func Unmarshal(b []byte, m proto.Message) error {
	return KoinosUnmarshalOptions.Unmarshal(b, m)
}
