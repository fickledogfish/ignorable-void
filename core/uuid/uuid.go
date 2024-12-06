package uuid

import (
	"database/sql/driver"

	guuid "github.com/google/uuid"
)

type UUID guuid.UUID

var (
	// Empty UUID, all zeros.
	Nil UUID
)

func (uuid UUID) String() string {
	return guuid.UUID(uuid).String()
}

// Implements encoding.BinaryMarshaler.
func (uuid UUID) MarshalBinary() ([]byte, error) {
	return guuid.UUID(uuid).MarshalBinary()
}

// Implements encoding.TextMarshaler.
func (uuid UUID) MarshalText() ([]byte, error) {
	return guuid.UUID(uuid).MarshalText()
}

// Implements sql.Scanner.
func (uuid *UUID) Scan(src any) error {
	ptr := guuid.UUID(*uuid)

	return ptr.Scan(src)
}

// Implements encoding.BinaryUnmarshaler.
func (uuid *UUID) UnmarshalBinary(data []byte) error {
	ptr := guuid.UUID(*uuid)

	return ptr.UnmarshalBinary(data)
}

// Implements encoding.TextUnmarshaler.
func (uuid *UUID) UnmarshalText(data []byte) error {
	ptr := guuid.UUID(*uuid)

	return ptr.UnmarshalText(data)
}

// Implements sql.Valuer.
func (uuid UUID) Value() (driver.Value, error) {
	return guuid.UUID(uuid).Value()
}
