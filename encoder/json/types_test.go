package json_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	libencoder "github.com/novln/soba/encoder"
	"github.com/novln/soba/encoder/json"
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

func TestJSON_Encoder_AppendInt(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `1,-15,30`

		encoder.AppendInt(1)
		encoder.AppendInt(-15)
		encoder.AppendInt(30)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendInt8(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `1,-15,30`

		encoder.AppendInt8(1)
		encoder.AppendInt8(-15)
		encoder.AppendInt8(30)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendInt16(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `1,-15,30`

		encoder.AppendInt16(1)
		encoder.AppendInt16(-15)
		encoder.AppendInt16(30)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendInt32(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `1,-15,30`

		encoder.AppendInt32(1)
		encoder.AppendInt32(-15)
		encoder.AppendInt32(30)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendInt64(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `1,-15,30`

		encoder.AppendInt64(1)
		encoder.AppendInt64(-15)
		encoder.AppendInt64(30)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendUint(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `1,15,30`

		encoder.AppendUint(1)
		encoder.AppendUint(15)
		encoder.AppendUint(30)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendUint8(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `1,15,30`

		encoder.AppendUint8(1)
		encoder.AppendUint8(15)
		encoder.AppendUint8(30)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendUint16(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `1,15,30`

		encoder.AppendUint16(1)
		encoder.AppendUint16(15)
		encoder.AppendUint16(30)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendUint32(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `1,15,30`

		encoder.AppendUint32(1)
		encoder.AppendUint32(15)
		encoder.AppendUint32(30)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendUint64(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `1,15,30`

		encoder.AppendUint64(1)
		encoder.AppendUint64(15)
		encoder.AppendUint64(30)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendFloat32(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `0.01,0.001,0.0001`

		encoder.AppendFloat32(0.01)
		encoder.AppendFloat32(0.001)
		encoder.AppendFloat32(0.0001)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `0.1,0.1,0.1`

		encoder.AppendFloat32(0.1)
		encoder.AppendFloat32(0.10)
		encoder.AppendFloat32(0.1000)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `0.00000000000000000000001`

		encoder.AppendFloat32(0.00000000000000000000001)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendFloat64(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `0.01,0.001,0.0001`

		encoder.AppendFloat64(0.01)
		encoder.AppendFloat64(0.001)
		encoder.AppendFloat64(0.0001)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `0.1,0.1,0.1`

		encoder.AppendFloat64(0.1)
		encoder.AppendFloat64(0.10)
		encoder.AppendFloat64(0.1000)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `0.00000000000000000000001`

		encoder.AppendFloat64(0.00000000000000000000001)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"-Inf","NaN","+Inf"`

		encoder.AppendFloat64(math.Inf(-1))
		encoder.AppendFloat64(math.Log(-1))
		encoder.AppendFloat64(math.Inf(+1))
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
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

func TestJSON_Encoder_AppendDuration(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		latency := (2 * time.Millisecond) + (523 * time.Microsecond)
		expected := `"2.523ms"`

		encoder.AppendDuration(latency)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		latency1 := (2 * time.Millisecond) + (523 * time.Microsecond)
		latency2 := (34 * time.Millisecond) + (127 * time.Microsecond) + (65 * time.Nanosecond)
		expected := `"2.523ms","34.127065ms"`

		encoder.AppendDuration(latency1)
		encoder.AppendDuration(latency2)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendBinary(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"SGVsbG8gd29ybGQ="`

		encoder.AppendBinary([]byte("Hello world"))
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"SGVsbG8=","d29ybGQ="`

		encoder.AppendBinary([]byte("Hello"))
		encoder.AppendBinary([]byte("world"))
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendArray(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		array := TestArray{
			Enabled: true,
			Status:  "published",
			ID:      145,
		}

		expected := `[true,"published",145]`

		encoder.AppendArray(array)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		array1 := TestArray{
			Enabled: true,
			Status:  "published",
			ID:      145,
		}
		array2 := TestArray{
			Enabled: true,
			Status:  "closed",
			ID:      247,
		}

		expected := `[true,"published",145],[true,"closed",247]`

		encoder.AppendArray(array1)
		encoder.AppendArray(array2)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AppendObject(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		object := TestObject{
			Enabled: true,
			Status:  "published",
			ID:      145,
		}

		expected := `{"enabled":true,"status":"published","id":145}`

		encoder.AppendObject(object)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		object1 := TestObject{
			Enabled: true,
			Status:  "published",
			ID:      145,
		}
		object2 := TestObject{
			Enabled: true,
			Status:  "closed",
			ID:      247,
		}

		expected := `{"enabled":true,"status":"published","id":145},{"enabled":true,"status":"closed","id":247}`

		encoder.AppendObject(object1)
		encoder.AppendObject(object2)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddArray(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		array := TestArray{
			Enabled: true,
			Status:  "published",
			ID:      145,
		}

		expected := `"buffer":[true,"published",145]`

		encoder.AddArray("buffer", array)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		array1 := TestArray{
			Enabled: true,
			Status:  "published",
			ID:      145,
		}
		array2 := TestArray{
			Enabled: true,
			Status:  "closed",
			ID:      247,
		}

		expected := `"pbuffer":[true,"published",145],"ubuffer":[true,"closed",247]`

		encoder.AddArray("pbuffer", array1)
		encoder.AddArray("ubuffer", array2)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddObject(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		object := TestObject{
			Enabled: true,
			Status:  "published",
			ID:      145,
		}

		expected := `"project":{"enabled":true,"status":"published","id":145}`

		encoder.AddObject("project", object)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		object1 := TestObject{
			Enabled: true,
			Status:  "published",
			ID:      145,
		}
		object2 := TestObject{
			Enabled: true,
			Status:  "closed",
			ID:      247,
		}

		expected := fmt.Sprint(
			`"project":{"enabled":true,"status":"published","id":145},`,
			`"user":{"enabled":true,"status":"closed","id":247}`,
		)

		encoder.AddObject("project", object1)
		encoder.AddObject("user", object2)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddObjects(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		object := TestObject{
			Enabled: true,
			Status:  "published",
			ID:      145,
		}

		expected := `"foo":[{"enabled":true,"status":"published","id":145}]`

		encoder.AddObjects("foo", []libencoder.ObjectMarshaler{object})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		object1 := TestObject{
			Enabled: true,
			Status:  "published",
			ID:      145,
		}
		object2 := TestObject{
			Enabled: true,
			Status:  "closed",
			ID:      247,
		}

		expected := fmt.Sprint(
			`"foo":[{"enabled":true,"status":"published","id":145}],`,
			`"bar":[{"enabled":true,"status":"closed","id":247}]`,
		)

		encoder.AddObjects("foo", []libencoder.ObjectMarshaler{object1})
		encoder.AddObjects("bar", []libencoder.ObjectMarshaler{object2})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		object1 := TestObject{
			Enabled: true,
			Status:  "published",
			ID:      145,
		}
		object2 := TestObject{
			Enabled: true,
			Status:  "closed",
			ID:      247,
		}
		object3 := TestObject{
			Enabled: false,
			Status:  "waiting",
			ID:      358,
		}

		expected := fmt.Sprint(
			`"foo":[{"enabled":true,"status":"published","id":145},`,
			`{"enabled":true,"status":"closed","id":247},`,
			`{"enabled":false,"status":"waiting","id":358}]`,
		)

		encoder.AddObjects("foo", []libencoder.ObjectMarshaler{object1, object2, object3})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		object1 := TestObject{
			Enabled: true,
			Status:  "published",
			ID:      145,
		}
		object2 := TestObject{
			Enabled: true,
			Status:  "closed",
			ID:      247,
		}
		object3 := TestObject{
			Enabled: false,
			Status:  "waiting",
			ID:      358,
		}
		object4 := TestObject{
			Enabled: false,
			Status:  "waiting",
			ID:      400,
		}
		object5 := TestObject{
			Enabled: false,
			Status:  "waiting",
			ID:      401,
		}
		object6 := TestObject{
			Enabled: false,
			Status:  "waiting",
			ID:      402,
		}
		object7 := TestObject{
			Enabled: false,
			Status:  "waiting",
			ID:      403,
		}

		expected := fmt.Sprint(
			`"foo":[{"enabled":true,"status":"published","id":145},`,
			`{"enabled":true,"status":"closed","id":247},`,
			`{"enabled":false,"status":"waiting","id":358}],`,
			`"bar":[],`,
			`"xyz":[{"enabled":false,"status":"waiting","id":400},`,
			`{"enabled":false,"status":"waiting","id":401},`,
			`{"enabled":false,"status":"waiting","id":402},`,
			`{"enabled":false,"status":"waiting","id":403}]`,
		)
		encoder.AddObjects("foo", []libencoder.ObjectMarshaler{object1, object2, object3})
		encoder.AddObjects("bar", []libencoder.ObjectMarshaler{})
		encoder.AddObjects("xyz", []libencoder.ObjectMarshaler{object4, object5, object6, object7})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddInt(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42`

		encoder.AddInt("foo", 42)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26`

		encoder.AddInt("foo", 42)
		encoder.AddInt("bar", 26)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26,"mask":777`

		encoder.AddInt("foo", 42)
		encoder.AddInt("bar", 26)
		encoder.AddInt("mask", 777)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddInts(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42]`

		encoder.AddInts("foo", []int{42})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42],"bar":[26]`

		encoder.AddInts("foo", []int{42})
		encoder.AddInts("bar", []int{26})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44]`

		encoder.AddInts("foo", []int{42, 43, 44})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44],"bar":[],"mask":[1,2,3,4]`

		encoder.AddInts("foo", []int{42, 43, 44})
		encoder.AddInts("bar", []int{})
		encoder.AddInts("mask", []int{1, 2, 3, 4})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddInt8(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42`

		encoder.AddInt8("foo", 42)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26`

		encoder.AddInt8("foo", 42)
		encoder.AddInt8("bar", 26)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26,"mask":126`

		encoder.AddInt8("foo", 42)
		encoder.AddInt8("bar", 26)
		encoder.AddInt8("mask", 126)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddInt8s(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42]`

		encoder.AddInt8s("foo", []int8{42})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42],"bar":[26]`

		encoder.AddInt8s("foo", []int8{42})
		encoder.AddInt8s("bar", []int8{26})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44]`

		encoder.AddInt8s("foo", []int8{42, 43, 44})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44],"bar":[],"mask":[1,2,3,4]`

		encoder.AddInt8s("foo", []int8{42, 43, 44})
		encoder.AddInt8s("bar", []int8{})
		encoder.AddInt8s("mask", []int8{1, 2, 3, 4})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddInt16(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42`

		encoder.AddInt16("foo", 42)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26`

		encoder.AddInt16("foo", 42)
		encoder.AddInt16("bar", 26)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26,"mask":777`

		encoder.AddInt16("foo", 42)
		encoder.AddInt16("bar", 26)
		encoder.AddInt16("mask", 777)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddInt16s(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42]`

		encoder.AddInt16s("foo", []int16{42})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42],"bar":[26]`

		encoder.AddInt16s("foo", []int16{42})
		encoder.AddInt16s("bar", []int16{26})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44]`

		encoder.AddInt16s("foo", []int16{42, 43, 44})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44],"bar":[],"mask":[1,2,3,4]`

		encoder.AddInt16s("foo", []int16{42, 43, 44})
		encoder.AddInt16s("bar", []int16{})
		encoder.AddInt16s("mask", []int16{1, 2, 3, 4})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddInt32(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42`

		encoder.AddInt32("foo", 42)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26`

		encoder.AddInt32("foo", 42)
		encoder.AddInt32("bar", 26)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26,"mask":777`

		encoder.AddInt32("foo", 42)
		encoder.AddInt32("bar", 26)
		encoder.AddInt32("mask", 777)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddInt32s(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42]`

		encoder.AddInt32s("foo", []int32{42})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42],"bar":[26]`

		encoder.AddInt32s("foo", []int32{42})
		encoder.AddInt32s("bar", []int32{26})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44]`

		encoder.AddInt32s("foo", []int32{42, 43, 44})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44],"bar":[],"mask":[1,2,3,4]`

		encoder.AddInt32s("foo", []int32{42, 43, 44})
		encoder.AddInt32s("bar", []int32{})
		encoder.AddInt32s("mask", []int32{1, 2, 3, 4})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddInt64(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42`

		encoder.AddInt64("foo", 42)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26`

		encoder.AddInt64("foo", 42)
		encoder.AddInt64("bar", 26)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26,"mask":777`

		encoder.AddInt64("foo", 42)
		encoder.AddInt64("bar", 26)
		encoder.AddInt64("mask", 777)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddInt64s(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42]`

		encoder.AddInt64s("foo", []int64{42})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42],"bar":[26]`

		encoder.AddInt64s("foo", []int64{42})
		encoder.AddInt64s("bar", []int64{26})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44]`

		encoder.AddInt64s("foo", []int64{42, 43, 44})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44],"bar":[],"mask":[1,2,3,4]`

		encoder.AddInt64s("foo", []int64{42, 43, 44})
		encoder.AddInt64s("bar", []int64{})
		encoder.AddInt64s("mask", []int64{1, 2, 3, 4})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddUint(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42`

		encoder.AddUint("foo", 42)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26`

		encoder.AddUint("foo", 42)
		encoder.AddUint("bar", 26)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26,"mask":777`

		encoder.AddUint("foo", 42)
		encoder.AddUint("bar", 26)
		encoder.AddUint("mask", 777)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddUints(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42]`

		encoder.AddUints("foo", []uint{42})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42],"bar":[26]`

		encoder.AddUints("foo", []uint{42})
		encoder.AddUints("bar", []uint{26})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44]`

		encoder.AddUints("foo", []uint{42, 43, 44})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44],"bar":[],"mask":[1,2,3,4]`

		encoder.AddUints("foo", []uint{42, 43, 44})
		encoder.AddUints("bar", []uint{})
		encoder.AddUints("mask", []uint{1, 2, 3, 4})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddUint8(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42`

		encoder.AddUint8("foo", 42)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26`

		encoder.AddUint8("foo", 42)
		encoder.AddUint8("bar", 26)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26,"mask":255`

		encoder.AddUint8("foo", 42)
		encoder.AddUint8("bar", 26)
		encoder.AddUint8("mask", 255)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddUint8s(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42]`

		encoder.AddUint8s("foo", []uint8{42})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42],"bar":[26]`

		encoder.AddUint8s("foo", []uint8{42})
		encoder.AddUint8s("bar", []uint8{26})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44]`

		encoder.AddUint8s("foo", []uint8{42, 43, 44})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44],"bar":[],"mask":[1,2,3,4]`

		encoder.AddUint8s("foo", []uint8{42, 43, 44})
		encoder.AddUint8s("bar", []uint8{})
		encoder.AddUint8s("mask", []uint8{1, 2, 3, 4})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddUint16(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42`

		encoder.AddUint16("foo", 42)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26`

		encoder.AddUint16("foo", 42)
		encoder.AddUint16("bar", 26)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26,"mask":777`

		encoder.AddUint16("foo", 42)
		encoder.AddUint16("bar", 26)
		encoder.AddUint16("mask", 777)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddUint16s(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42]`

		encoder.AddUint16s("foo", []uint16{42})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42],"bar":[26]`

		encoder.AddUint16s("foo", []uint16{42})
		encoder.AddUint16s("bar", []uint16{26})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44]`

		encoder.AddUint16s("foo", []uint16{42, 43, 44})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44],"bar":[],"mask":[1,2,3,4]`

		encoder.AddUint16s("foo", []uint16{42, 43, 44})
		encoder.AddUint16s("bar", []uint16{})
		encoder.AddUint16s("mask", []uint16{1, 2, 3, 4})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddUint32(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42`

		encoder.AddUint32("foo", 42)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26`

		encoder.AddUint32("foo", 42)
		encoder.AddUint32("bar", 26)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26,"mask":777`

		encoder.AddUint32("foo", 42)
		encoder.AddUint32("bar", 26)
		encoder.AddUint32("mask", 777)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddUint32s(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42]`

		encoder.AddUint32s("foo", []uint32{42})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42],"bar":[26]`

		encoder.AddUint32s("foo", []uint32{42})
		encoder.AddUint32s("bar", []uint32{26})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44]`

		encoder.AddUint32s("foo", []uint32{42, 43, 44})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44],"bar":[],"mask":[1,2,3,4]`

		encoder.AddUint32s("foo", []uint32{42, 43, 44})
		encoder.AddUint32s("bar", []uint32{})
		encoder.AddUint32s("mask", []uint32{1, 2, 3, 4})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddUint64(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42`

		encoder.AddUint64("foo", 42)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26`

		encoder.AddUint64("foo", 42)
		encoder.AddUint64("bar", 26)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":42,"bar":26,"mask":777`

		encoder.AddUint64("foo", 42)
		encoder.AddUint64("bar", 26)
		encoder.AddUint64("mask", 777)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddUint64s(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42]`

		encoder.AddUint64s("foo", []uint64{42})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42],"bar":[26]`

		encoder.AddUint64s("foo", []uint64{42})
		encoder.AddUint64s("bar", []uint64{26})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44]`

		encoder.AddUint64s("foo", []uint64{42, 43, 44})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[42,43,44],"bar":[],"mask":[1,2,3,4]`

		encoder.AddUint64s("foo", []uint64{42, 43, 44})
		encoder.AddUint64s("bar", []uint64{})
		encoder.AddUint64s("mask", []uint64{1, 2, 3, 4})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddFloat32(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":0.01`

		encoder.AddFloat32("foo", 0.01)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":0.01,"bar":3.3333`

		encoder.AddFloat32("foo", 0.01)
		encoder.AddFloat32("bar", 3.3333)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":0.01,"bar":3.3333,"xyz":50`

		encoder.AddFloat32("foo", 0.01)
		encoder.AddFloat32("bar", 3.3333)
		encoder.AddFloat32("xyz", 50)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddFloat32s(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[0.01]`

		encoder.AddFloat32s("foo", []float32{0.01})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[0.01],"bar":[3.3333]`

		encoder.AddFloat32s("foo", []float32{0.01})
		encoder.AddFloat32s("bar", []float32{3.3333})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[0.01,0.001,0.0001]`

		encoder.AddFloat32s("foo", []float32{0.01, 0.001, 0.0001})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[0.01,0.001,0.0001],"bar":[],"xyz":[1,1.2,1.4,1.6]`

		encoder.AddFloat32s("foo", []float32{0.01, 0.001, 0.0001})
		encoder.AddFloat32s("bar", []float32{})
		encoder.AddFloat32s("xyz", []float32{1.0, 1.2, 1.4, 1.6})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddFloat64(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":0.01`

		encoder.AddFloat64("foo", 0.01)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":0.01,"bar":3.3333`

		encoder.AddFloat64("foo", 0.01)
		encoder.AddFloat64("bar", 3.3333)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":0.01,"bar":3.3333,"xyz":50`

		encoder.AddFloat64("foo", 0.01)
		encoder.AddFloat64("bar", 3.3333)
		encoder.AddFloat64("xyz", 50)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddFloat64s(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[0.01]`

		encoder.AddFloat64s("foo", []float64{0.01})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[0.01],"bar":[3.3333]`

		encoder.AddFloat64s("foo", []float64{0.01})
		encoder.AddFloat64s("bar", []float64{3.3333})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[0.01,0.001,0.0001]`

		encoder.AddFloat64s("foo", []float64{0.01, 0.001, 0.0001})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[0.01,0.001,0.0001],"bar":[],"xyz":[1,1.2,1.4,1.6]`

		encoder.AddFloat64s("foo", []float64{0.01, 0.001, 0.0001})
		encoder.AddFloat64s("bar", []float64{})
		encoder.AddFloat64s("xyz", []float64{1.0, 1.2, 1.4, 1.6})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddBool(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":false`

		encoder.AddBool("foo", false)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":false,"bar":true`

		encoder.AddBool("foo", false)
		encoder.AddBool("bar", true)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":false,"bar":true,"enabled":true`

		encoder.AddBool("foo", false)
		encoder.AddBool("bar", true)
		encoder.AddBool("enabled", true)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddBools(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[false]`

		encoder.AddBools("foo", []bool{false})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[false],"bar":[true]`

		encoder.AddBools("foo", []bool{false})
		encoder.AddBools("bar", []bool{true})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[false,true,false]`

		encoder.AddBools("foo", []bool{false, true, false})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":[false,true,true],"bar":[],"mask":[true,true,false,false]`

		encoder.AddBools("foo", []bool{false, true, true})
		encoder.AddBools("bar", []bool{})
		encoder.AddBools("mask", []bool{true, true, false, false})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddString(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":"bar"`

		encoder.AddString("foo", "bar")
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"user":"01CV2N1Z7BE4AV5NBRMZBRT9CZ","group":"01CV2N2BJP3F6V5DA1QA9N5Q2D"`

		encoder.AddString("user", "01CV2N1Z7BE4AV5NBRMZBRT9CZ")
		encoder.AddString("group", "01CV2N2BJP3F6V5DA1QA9N5Q2D")
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"user":"01CV2N1Z7BE4AV5NBRMZBRT9CZ","group":"01CV2N2BJP3F6V5DA1QA9N5Q2D","description":"foobar"`

		encoder.AddString("user", "01CV2N1Z7BE4AV5NBRMZBRT9CZ")
		encoder.AddString("group", "01CV2N2BJP3F6V5DA1QA9N5Q2D")
		encoder.AddString("description", "foobar")
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := fmt.Sprint(
			`"name":"Les \"artisans\"","description":"\\\\ Communauté d'artisans //\n\n\t`,
			`Découvrez les ateliers d'initiation","website":"https://artisans.io"`,
		)

		encoder.AddString("name", `Les "artisans"`)
		encoder.AddString("description", `\\ Communauté d'artisans //`+"\n\n\t"+"Découvrez les ateliers d'initiation")
		encoder.AddString("website", "https://artisans.io")
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddStrings(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":["bar"]`

		encoder.AddStrings("foo", []string{"bar"})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := fmt.Sprint(
			`"user":["01CV2N1Z7BE4AV5NBRMZBRT9CZ","01CV4WDQ3CN2VMCC7N24WMYB2Z"],`,
			`"group":["01CV2N2BJP3F6V5DA1QA9N5Q2D","01CV4WETCDWMXHA899AE4321MF"]`,
		)

		encoder.AddStrings("user", []string{"01CV2N1Z7BE4AV5NBRMZBRT9CZ", "01CV4WDQ3CN2VMCC7N24WMYB2Z"})
		encoder.AddStrings("group", []string{"01CV2N2BJP3F6V5DA1QA9N5Q2D", "01CV4WETCDWMXHA899AE4321MF"})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"foo":["a","b","c"],"bar":[],"mask":["w","x","y","z"]`

		encoder.AddStrings("foo", []string{"a", "b", "c"})
		encoder.AddStrings("bar", []string{})
		encoder.AddStrings("mask", []string{"w", "x", "y", "z"})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddStringer(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		latency1 := 230 * time.Millisecond
		latency2 := (2 * time.Millisecond) + (523 * time.Microsecond)
		latency3 := (34 * time.Millisecond) + (127 * time.Microsecond) + (65 * time.Nanosecond)

		expected := `"l1":"230ms","l2":"2.523ms","l3":"34.127065ms"`

		encoder.AddStringer("l1", latency1)
		encoder.AddStringer("l2", latency2)
		encoder.AddStringer("l3", latency3)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddStringers(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		latency1 := 230 * time.Millisecond
		latency2 := (2 * time.Millisecond) + (523 * time.Microsecond)
		latency3 := (34 * time.Millisecond) + (127 * time.Microsecond) + (65 * time.Nanosecond)

		expected := `"latencies":["230ms","2.523ms","34.127065ms"]`

		encoder.AddStringers("latencies", []fmt.Stringer{latency1, latency2, latency3})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddTime(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		timezone, err := time.LoadLocation("Europe/Paris")
		if err != nil {
			t.Fatalf("Unexpected error to load timezone: %s", err)
		}
		when := time.Unix(1540912844, 0).In(timezone)

		expected := `"time":"2018-10-30T16:20:44+01:00"`

		encoder.AddTime("time", when)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		tz1, err := time.LoadLocation("Europe/Paris")
		if err != nil {
			t.Fatalf("Unexpected error to load timezone: %s", err)
		}
		tz2, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			t.Fatalf("Unexpected error to load timezone: %s", err)
		}

		t1 := time.Unix(1540912844, 0).In(tz1)
		t2 := time.Unix(1605090000, 0).In(tz2)

		expected := `"created_at":"2018-10-30T16:20:44+01:00","updated_at":"2020-11-11T19:20:00+09:00"`

		encoder.AddTime("created_at", t1)
		encoder.AddTime("updated_at", t2)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddTimes(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		tz1, err := time.LoadLocation("Europe/Paris")
		if err != nil {
			t.Fatalf("Unexpected error to load timezone: %s", err)
		}
		tz2, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			t.Fatalf("Unexpected error to load timezone: %s", err)
		}

		t1 := time.Unix(1540912844, 0).In(tz1)
		t2 := time.Unix(1605090000, 0).In(tz2)

		expected := `"login":["2018-10-30T16:20:44+01:00","2020-11-11T19:20:00+09:00"]`

		encoder.AddTimes("login", []time.Time{t1, t2})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddDuration(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		latency1 := 230 * time.Millisecond
		latency2 := (2 * time.Millisecond) + (523 * time.Microsecond)
		latency3 := (34 * time.Millisecond) + (127 * time.Microsecond) + (65 * time.Nanosecond)

		expected := `"l1":"230ms","l2":"2.523ms","l3":"34.127065ms"`

		encoder.AddDuration("l1", latency1)
		encoder.AddDuration("l2", latency2)
		encoder.AddDuration("l3", latency3)
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddDurations(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		latency1 := 230 * time.Millisecond
		latency2 := (2 * time.Millisecond) + (523 * time.Microsecond)
		latency3 := (34 * time.Millisecond) + (127 * time.Microsecond) + (65 * time.Nanosecond)

		expected := `"latencies":["230ms","2.523ms","34.127065ms"]`

		encoder.AddDurations("latencies", []time.Duration{latency1, latency2, latency3})
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}

func TestJSON_Encoder_AddBinary(t *testing.T) {
	{
		encoder := json.NewEncoder()
		defer encoder.Close()

		expected := `"message":"SGVsbG8gd29ybGQ="`

		encoder.AddBinary("message", []byte("Hello world"))
		buffer := encoder.Bytes()

		if expected != string(buffer) {
			t.Fatalf("Unexpected buffer: '%s' should be '%s'", string(buffer), expected)
		}
	}
}
