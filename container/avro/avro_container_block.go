// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCES:
 *     block.avsc
 *     header.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

type AvroContainerBlock struct {
	NumRecords int64

	RecordBytes []byte

	Sync Sync
}

const AvroContainerBlockAvroCRC64Fingerprint = "\x0e\xecj@ٔ\xe14"

func NewAvroContainerBlock() *AvroContainerBlock {
	return &AvroContainerBlock{}
}

func DeserializeAvroContainerBlock(r io.Reader) (*AvroContainerBlock, error) {
	t := NewAvroContainerBlock()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeAvroContainerBlockFromSchema(r io.Reader, schema string) (*AvroContainerBlock, error) {
	t := NewAvroContainerBlock()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeAvroContainerBlock(r *AvroContainerBlock, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.NumRecords, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.RecordBytes, w)
	if err != nil {
		return err
	}
	err = writeSync(r.Sync, w)
	if err != nil {
		return err
	}
	return err
}

func (r *AvroContainerBlock) Serialize(w io.Writer) error {
	return writeAvroContainerBlock(r, w)
}

func (r *AvroContainerBlock) Schema() string {
	return "{\"fields\":[{\"name\":\"numRecords\",\"type\":\"long\"},{\"name\":\"recordBytes\",\"type\":\"bytes\"},{\"name\":\"sync\",\"type\":{\"name\":\"sync\",\"size\":16,\"type\":\"fixed\"}}],\"name\":\"AvroContainerBlock\",\"type\":\"record\"}"
}

func (r *AvroContainerBlock) SchemaName() string {
	return "AvroContainerBlock"
}

func (_ *AvroContainerBlock) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *AvroContainerBlock) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *AvroContainerBlock) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *AvroContainerBlock) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *AvroContainerBlock) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *AvroContainerBlock) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *AvroContainerBlock) SetString(v string)   { panic("Unsupported operation") }
func (_ *AvroContainerBlock) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AvroContainerBlock) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Long{Target: &r.NumRecords}
	case 1:
		return &types.Bytes{Target: &r.RecordBytes}
	case 2:
		return &SyncWrapper{Target: &r.Sync}
	}
	panic("Unknown field index")
}

func (r *AvroContainerBlock) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (_ *AvroContainerBlock) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *AvroContainerBlock) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *AvroContainerBlock) Finalize()                        {}

func (_ *AvroContainerBlock) AvroCRC64Fingerprint() []byte {
	return []byte(AvroContainerBlockAvroCRC64Fingerprint)
}
