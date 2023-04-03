package uuid

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
)

type UUID [16]byte

type Version int

const (
	V1 Version = 1
	V3 Version = 3
	V4 Version = 4
	V5 Version = 5
	V6 Version = 6
	V7 Version = 7
	V8 Version = 8
)

var (
	Nil = [16]byte{}
	Max = [16]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
)

var (
	ErrInvalidFormat      = errors.New("invalid uuid format")
	ErrUnsupportedVersion = errors.New("unsupported uuid version")
)

func (pId *UUID) Generate(v Version) error {
	var id UUID

	switch v {
	case V4:
		if _, err := rand.Read(id[:]); err != nil {
			return fmt.Errorf("cannot read random data: %w", err)
		}

		id[6] = (id[6] & 0x0f) | 0x40 // Version 4
		id[8] = (id[8] & 0x3f) | 0x80 // Variant b10

	default:
		return ErrUnsupportedVersion
	}

	*pId = id

	return nil
}

func MustGenerate(v Version) (id UUID) {
	if err := id.Generate(v); err != nil {
		panic(fmt.Sprintf("cannot generate uuid v%d: %v", v, err))
	}

	return
}

func (id1 UUID) Equal(id2 UUID) bool {
	return bytes.Equal(id1.Bytes(), id2.Bytes())
}

func (id UUID) String() string {
	data, _ := id.MarshalText()
	return string(data)
}

func (id UUID) Bytes() []byte {
	return id[:]
}

func (pId *UUID) Parse(s string) error {
	return pId.UnmarshalText([]byte(s))
}

func MustParse(s string) (id UUID) {
	if err := id.Parse(s); err != nil {
		panic(fmt.Sprintf("invalid uuid %q", s))
	}

	return id
}

// encoding.TextMarshaler
func (id UUID) MarshalText() ([]byte, error) {
	data := make([]byte, 36)

	hex.Encode(data[0:8], id[0:4])
	data[8] = '-'
	hex.Encode(data[9:13], id[4:6])
	data[13] = '-'
	hex.Encode(data[14:18], id[6:8])
	data[18] = '-'
	hex.Encode(data[19:23], id[8:10])
	data[23] = '-'
	hex.Encode(data[24:36], id[10:16])

	return data, nil
}

// encoding.TextUnmarshaler
func (pId *UUID) UnmarshalText(data []byte) error {
	var err error
	var id UUID

	if len(data) != 36 {
		return ErrInvalidFormat
	}

	_, err = hex.Decode(id[0:4], data[0:8])
	_, err = hex.Decode(id[4:6], data[9:13])
	_, err = hex.Decode(id[6:8], data[14:18])
	_, err = hex.Decode(id[8:10], data[19:23])
	_, err = hex.Decode(id[10:16], data[24:36])

	if err != nil || data[8] != '-' || data[13] != '-' ||
		data[18] != '-' || data[23] != '-' {
		return ErrInvalidFormat
	}

	*pId = id

	return nil
}
