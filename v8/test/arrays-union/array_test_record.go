// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v8/compiler"
	"github.com/actgardner/gogen-avro/v8/vm"
	"github.com/actgardner/gogen-avro/v8/vm/types"
)

var _ = fmt.Printf

type ArrayTestRecord struct {
	IntField []*UnionNullInt `json:"IntField"`
}

const ArrayTestRecordAvroCRC64Fingerprint = "t\x06\x9e\xc8\u0088\xa0\xcb"

func NewArrayTestRecord() ArrayTestRecord {
	r := ArrayTestRecord{}
	r.IntField = make([]*UnionNullInt, 0)

	return r
}

func DeserializeArrayTestRecord(r io.Reader) (ArrayTestRecord, error) {
	t := NewArrayTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeArrayTestRecordFromSchema(r io.Reader, schema string) (ArrayTestRecord, error) {
	t := NewArrayTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeArrayTestRecord(r ArrayTestRecord, w io.Writer) error {
	var err error
	err = writeArrayUnionNullInt(r.IntField, w)
	if err != nil {
		return err
	}
	return err
}

func (r ArrayTestRecord) Serialize(w io.Writer) error {
	return writeArrayTestRecord(r, w)
}

func (r ArrayTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":{\"items\":[\"null\",\"int\"],\"type\":\"array\"}}],\"name\":\"ArrayTestRecord\",\"type\":\"record\"}"
}

func (r ArrayTestRecord) SchemaName() string {
	return "ArrayTestRecord"
}

func (_ ArrayTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ ArrayTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ArrayTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		r.IntField = make([]*UnionNullInt, 0)

		return &ArrayUnionNullIntWrapper{Target: &r.IntField}
	}
	panic("Unknown field index")
}

func (r *ArrayTestRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ArrayTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ArrayTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ArrayTestRecord) Finalize()                        {}

func (_ ArrayTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(ArrayTestRecordAvroCRC64Fingerprint)
}

func (r ArrayTestRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IntField"], err = json.Marshal(r.IntField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ArrayTestRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IntField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IntField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IntField")
	}
	return nil
}
