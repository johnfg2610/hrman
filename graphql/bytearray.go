package graphql

import (
	"fmt"
	"io"
)

type ByteArray []byte

func (ba *ByteArray) UnmarshalGQL(v interface{}) error {
	arr, ok := v.([]byte)
	if !ok {
		return fmt.Errorf("ByteArray must be a array of bytes")
	}

	*ba = arr
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (ba ByteArray) MarshalGQL(w io.Writer) {
	w.Write(ba)
}
