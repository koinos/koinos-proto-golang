package canonical

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

// Marshal ensures the object is serialized in a canonical way
func Marshal(m proto.Message) ([]byte, error) {
	if containsMap(m) {
		return nil, fmt.Errorf("%w: the given object contains a map", ErrNoCanonical)
	}

	return proto.Marshal(m)
}

func containsMap(m proto.Message) bool {
	pr := m.ProtoReflect()
	fd := pr.Descriptor().Fields()
	for i := 0; i < fd.Len(); i++ {
		if fd.Get(i).IsMap() {
			return true
		}

		v := pr.Get(fd.Get(i))
		if mv, ok := v.Interface().(proto.Message); ok {
			if containsMap(mv) {
				return true
			}
		}
	}

	return false
}
