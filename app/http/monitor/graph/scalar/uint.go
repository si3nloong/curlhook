package scalar

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalUint :
func MarshalUint(v uint) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		fmt.Fprint(w, strconv.FormatUint(uint64(v), 10))
	})
}

// UnmarshalUint :
func UnmarshalUint(v interface{}) (uint, error) {
	switch vi := v.(type) {
	case string:
		u64, err := strconv.ParseUint(vi, 10, 64)
		if err != nil {
			return 0, err
		}
		return uint(u64), nil

	case int:
		if vi < 0 {
			return 0, errors.New("unsigned integer cannot be negative")
		}
		return uint(vi), nil

	case int64:
		if vi < 0 {
			return 0, errors.New("unsigned integer cannot be negative")
		}
		return uint(vi), nil

	case json.Number:
		i64, err := vi.Int64()
		if err != nil {
			return 0, err
		}
		if i64 < 0 {
			return 0, errors.New("unsigned integer cannot be negative")
		}
		return uint(i64), nil

	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}
