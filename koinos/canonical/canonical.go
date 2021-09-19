package canonical

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

// Marshal ensures the object is serialized in a canonical way
func Marshal(m proto.Message) ([]byte, error) {
	if containsMap(m.ProtoReflect()) {
		return nil, fmt.Errorf("%w: the given object contains a map", ErrNoCanonical)
	}

	return proto.Marshal(m)
}

func containsMap(pr protoreflect.Message) bool {
	fields := pr.Descriptor().Fields()

	// Iterate over the fields
	for i := 0; i < fields.Len(); i++ {
		fdesc := fields.Get(i)

		// If the field is a map we're done here
		if fdesc.IsMap() {
			return true
		}

		// Recurse into the field if it is a message
		if fdesc.Kind() == protoreflect.MessageKind {
			message := pr.Get(fdesc).Message()
			if containsMap(message) {
				return true
			}
		}
	}

	return false
}
