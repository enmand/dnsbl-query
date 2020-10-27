package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func MarshalUUID(u uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(u.String()))
	})
}

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
