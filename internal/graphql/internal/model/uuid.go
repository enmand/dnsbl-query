package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

// MarshalUUID returns a graphql.Marshaler that can marshal uuid.UUIDs
func MarshalUUID(u uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(u.String()))
	})
}

// UnmarshalUUID returns a uuid.UUID if it can successfull parse one out,
// otherwise it returns an error
func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	if v, ok := v.(string); ok {
		id, err := uuid.Parse(v)
		if err != nil {
			return uuid.Nil, fmt.Errorf("unable to parse ID: %w", err)
		}

		return id, nil
	}

	return uuid.Nil, fmt.Errorf("invalid type for ID")
}
