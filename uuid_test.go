package uuid

import (
	"testing"
	"time"
)

var tests = []struct {
	s  string
	id UUID
}{
	{
		// v1
		"c232ab00-9414-11ec-b3c8-9e6bdeced846",
		UUID{0xc2, 0x32, 0xab, 0x00, 0x94, 0x14, 0x11, 0xec,
			0xb3, 0xc8, 0x9e, 0x6b, 0xde, 0xce, 0xd8, 0x46},
	},
	{
		// v3
		"5df41881-3aed-3515-88a7-2f4a814cf09e",
		UUID{0x5d, 0xf4, 0x18, 0x81, 0x3a, 0xed, 0x35, 0x15,
			0x88, 0xa7, 0x2f, 0x4a, 0x81, 0x4c, 0xf0, 0x9e},
	},
	{
		// v4
		"919108f7-52d1-4320-9bac-f847db4148a8",
		UUID{0x91, 0x91, 0x08, 0xf7, 0x52, 0xd1, 0x43, 0x20,
			0x9b, 0xac, 0xf8, 0x47, 0xdb, 0x41, 0x48, 0xa8},
	},
	{
		// v5
		"2ed6657d-e927-568b-95e1-2665a8aea6a2",
		UUID{0x2e, 0xd6, 0x65, 0x7d, 0xe9, 0x27, 0x56, 0x8b,
			0x95, 0xe1, 0x26, 0x65, 0xa8, 0xae, 0xa6, 0xa2},
	},
	{
		// v6
		"1ec9414c-232a-6b00-b3c8-9e6bdeced846",
		UUID{0x1e, 0xc9, 0x41, 0x4c, 0x23, 0x2a, 0x6b, 0x00,
			0xb3, 0xc8, 0x9e, 0x6b, 0xde, 0xce, 0xd8, 0x46},
	},
	{
		// v7
		"017f22e2-79b0-7cc3-98c4-dc0c0c07398f",
		UUID{0x01, 0x7f, 0x22, 0xe2, 0x79, 0xb0, 0x7c, 0xc3,
			0x98, 0xc4, 0xdc, 0x0c, 0x0c, 0x07, 0x39, 0x8f},
	},
	{
		// v8
		"320c3d4d-cc00-875b-8ec9-32d5f69181c0",
		UUID{0x32, 0x0c, 0x3d, 0x4d, 0xcc, 0x00, 0x87, 0x5b,
			0x8e, 0xc9, 0x32, 0xd5, 0xf6, 0x91, 0x81, 0xc0},
	},
}

func TestString(t *testing.T) {
	for _, test := range tests {
		s := test.id.String()
		if s != test.s {
			t.Errorf("%q was serialized to %q instead of %q",
				test.id, s, test.s)
		}
	}
}

func TestParse(t *testing.T) {
	for _, test := range tests {
		var id UUID
		if err := id.Parse(test.s); err == nil {
			if !id.Equal(test.id) {
				t.Errorf("%q was parsed as %q instead of %q",
					test.s, id, test.id)
			}
		} else {
			t.Errorf("cannot parse %q: %v", test.s, err)
		}
	}
}

func TestGenerateV7Zero(t *testing.T) {
	idTime := time.UnixMilli(1717090333787)
	id := GenerateV7Zero(idTime)

	expectedIdString := "018fca8f-345b-7000-8000-000000000000"

	if id.String() != expectedIdString {
		t.Errorf("id is %q but should be %q", id.String(), expectedIdString)
	}
}

func TestV7Time(t *testing.T) {
	id := MustParse("018fca8f-345b-711a-838a-a276340388e7")
	expectedTime := time.UnixMilli(1717090333787)

	if idTime := id.V7Time(); idTime != expectedTime {
		t.Errorf("time is %q but should be %q",
			idTime.Format(time.RFC3339Nano),
			expectedTime.Format(time.RFC3339Nano))
	}
}
