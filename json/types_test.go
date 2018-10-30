package json_test

import (
	"testing"
	"time"

	"github.com/novln/soba/json"
)

// Source forked from https://github.com/uber-go/zap and from https://github.com/rs/zerolog

func TestJSON_Encoder_Escaping(t *testing.T) {

	scenarios := []struct {
		input  string
		output string
	}{
		{"", `""`},
		{"\\", `"\\"`},
		{"\x00", `"\u0000"`},
		{"\x01", `"\u0001"`},
		{"\x02", `"\u0002"`},
		{"\x03", `"\u0003"`},
		{"\x04", `"\u0004"`},
		{"\x05", `"\u0005"`},
		{"\x06", `"\u0006"`},
		{"\x07", `"\u0007"`},
		{"\x08", `"\u0008"`},
		{"\x09", `"\t"`},
		{"\x0a", `"\n"`},
		{"\x0b", `"\u000b"`},
		{"\x0c", `"\u000c"`},
		{"\x0d", `"\r"`},
		{"\x0e", `"\u000e"`},
		{"\x0f", `"\u000f"`},
		{"\x10", `"\u0010"`},
		{"\x11", `"\u0011"`},
		{"\x12", `"\u0012"`},
		{"\x13", `"\u0013"`},
		{"\x14", `"\u0014"`},
		{"\x15", `"\u0015"`},
		{"\x16", `"\u0016"`},
		{"\x17", `"\u0017"`},
		{"\x18", `"\u0018"`},
		{"\x19", `"\u0019"`},
		{"\x1a", `"\u001a"`},
		{"\x1b", `"\u001b"`},
		{"\x1c", `"\u001c"`},
		{"\x1d", `"\u001d"`},
		{"\x1e", `"\u001e"`},
		{"\x1f", `"\u001f"`},
		{"✭", `"✭"`},
		{"\xed\xa0\x80", `"\ufffd\ufffd\ufffd"`},
		{"xyz\xed\xa0\x80", `"xyz\ufffd\ufffd\ufffd"`},
		{"foobar", `"foobar"`},
		{"\"a", `"\"a"`},
		{"\\\\hello//", `"\\\\hello//"`},
		{"0x\nabcd\tijk", `"0x\nabcd\tijk"`},
		{"\x1fa", `"\u001fa"`},
		{"foo\"bar\"baz", `"foo\"bar\"baz"`},
		{"\x1ffoo\x1fbar\x1fbaz", `"\u001ffoo\u001fbar\u001fbaz"`},
		{"I \u2764\ufe0f go!", `"I ❤️ go!"`},
		{"<", `"<"`},
		{">", `">"`},
		{"&", `"&"`},
	}

	for i, scenario := range scenarios {
		encoder := json.NewEncoder()
		defer encoder.Close()

		encoder.AppendString(scenario.input)
		buffer := encoder.Bytes()

		if scenario.output != string(buffer) {
			t.Fatalf("Unexpected result for scenario #%d: '%s' should be '%s'", i, string(buffer), scenario.output)
		}
	}

}

func TestJSON_Encoder_AppendBool(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `true,false,true`

		encoder.AppendBool(true)
		encoder.AppendBool(false)
		encoder.AppendBool(true)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendString(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo","bar","xyz"`

		encoder.AppendString("foo")
		encoder.AppendString("bar")
		encoder.AppendString("xyz")
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendTime(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		timezone, err := time.LoadLocation("Europe/Paris")
		if err != nil {
			t.Fatalf("Unexpected error to load timezone: %s", err)
		}
		when := time.Unix(1540912844, 0).In(timezone)
		expected := `"2018-10-30T16:20:44+01:00"`

		encoder.AppendTime(when)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		timezone, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			t.Fatalf("Unexpected error to load timezone: %s", err)
		}
		when := time.Unix(1605090000, 908060454).In(timezone)
		expected := `"2020-11-11T19:20:00.908060454+09:00"`

		encoder.AppendTime(when)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}
