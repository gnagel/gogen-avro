// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCES:
 *     block.avsc
 *     header.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

func writeMagic(r Magic, w io.Writer) error {
	_, err := w.Write(r[:])
	return err
}

type Magic MagicWrapper
type MagicWrapper [4]byte

func (_ *MagicWrapper) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *MagicWrapper) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *MagicWrapper) SetLong(v int64)     { panic("Unsupported operation") }
func (_ *MagicWrapper) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *MagicWrapper) SetDouble(v float64) { panic("Unsupported operation") }
func (r *MagicWrapper) SetBytes(v []byte) {
	copy((*r)[:], v)
}
func (_ *MagicWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ *MagicWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ *MagicWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ *MagicWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *MagicWrapper) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *MagicWrapper) Finalize()                        {}
func (_ *MagicWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
