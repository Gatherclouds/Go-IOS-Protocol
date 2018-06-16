package rlp

import (
	"errors"
	"io"
	"bytes"
	"reflect"
	"fmt"
	"math/big"
)

var (
	errNoPointer     = errors.New("rlp: interface given to Decode must be a pointer")
	errDecodeIntoNil = errors.New("rlp: pointer given to Decode must not be nil")
)

type Decoder interface {
	DecodeRLP(*Stream) error
}

func Decode(r io.Reader, val interface{}) error {
	// TODO: this could use a Stream from a pool.
	return NewStream(r, 0).Decode(val)
}

func DecodeBytes(b []byte, val interface{}) error {
	// TODO: this could use a Stream from a pool.
	r := bytes.NewReader(b)
	if err := NewStream(r, uint64(len(b))).Decode(val); err != nil {
		return err
	}
	if r.Len() > 0 {
		return ErrMoreThanOneValue
	}
	return nil
}
type decodeError struct {
	msg string
	typ reflect.Type
	ctx []string
}

func (err *decodeError) Error() string {
	ctx := ""
	if len(err.ctx) > 0 {
		ctx = ", decoding into "
		for i := len(err.ctx) - 1; i >= 0; i-- {
			ctx += err.ctx[i]
		}
	}
	return fmt.Sprintf("rlp: %s for %v%s", err.msg, err.typ, ctx)
}

func wrapStreamError(err error, typ reflect.Type) error {
	switch err {
	case ErrCanonInt:
		return &decodeError{msg: "non-canonical integer (leading zero bytes)", typ: typ}
	case ErrCanonSize:
		return &decodeError{msg: "non-canonical size information", typ: typ}
	case ErrExpectedList:
		return &decodeError{msg: "expected input list", typ: typ}
	case ErrExpectedString:
		return &decodeError{msg: "expected input string or byte", typ: typ}
	case errUintOverflow:
		return &decodeError{msg: "input string too long", typ: typ}
	case errNotAtEOL:
		return &decodeError{msg: "input list has too many elements", typ: typ}
	}
	return err
}

func addErrorContext(err error, ctx string) error {
	if decErr, ok := err.(*decodeError); ok {
		decErr.ctx = append(decErr.ctx, ctx)
	}
	return err
}

var (
	decoderInterface = reflect.TypeOf(new(Decoder)).Elem()
	bigInt           = reflect.TypeOf(big.Int{})
)

func makeDecoder(typ reflect.Type, tags tags) (dec decoder, err error) {
	kind := typ.Kind()
	switch {
	case typ == rawValueType:
		return decodeRawValue, nil
	case typ.Implements(decoderInterface):
		return decodeDecoder, nil
	case kind != reflect.Ptr && reflect.PtrTo(typ).Implements(decoderInterface):
		return decodeDecoderNoPtr, nil
	case typ.AssignableTo(reflect.PtrTo(bigInt)):
		return decodeBigInt, nil
	case typ.AssignableTo(bigInt):
		return decodeBigIntNoPtr, nil
	case isUint(kind):
		return decodeUint, nil
	case kind == reflect.Bool:
		return decodeBool, nil
	case kind == reflect.String:
		return decodeString, nil
	case kind == reflect.Slice || kind == reflect.Array:
		return makeListDecoder(typ, tags)
	case kind == reflect.Struct:
		return makeStructDecoder(typ)
	case kind == reflect.Ptr:
		if tags.nilOK {
			return makeOptionalPtrDecoder(typ)
		}
		return makePtrDecoder(typ)
	case kind == reflect.Interface:
		return decodeInterface, nil
	default:
		return nil, fmt.Errorf("rlp: type %v is not RLP-serializable", typ)
	}
}

func decodeRawValue(s *Stream, val reflect.Value) error {
	r, err := s.Raw()
	if err != nil {
		return err
	}
	val.SetBytes(r)
	return nil
}

func decodeUint(s *Stream, val reflect.Value) error {
	typ := val.Type()
	num, err := s.uint(typ.Bits())
	if err != nil {
		return wrapStreamError(err, val.Type())
	}
	val.SetUint(num)
	return nil
}

func decodeBool(s *Stream, val reflect.Value) error {
	b, err := s.Bool()
	if err != nil {
		return wrapStreamError(err, val.Type())
	}
	val.SetBool(b)
	return nil
}

func decodeString(s *Stream, val reflect.Value) error {
	b, err := s.Bytes()
	if err != nil {
		return wrapStreamError(err, val.Type())
	}
	val.SetString(string(b))
	return nil
}

func decodeBigIntNoPtr(s *Stream, val reflect.Value) error {
	return decodeBigInt(s, val.Addr())
}

func decodeBigInt(s *Stream, val reflect.Value) error {
	b, err := s.Bytes()
	if err != nil {
		return wrapStreamError(err, val.Type())
	}
	i := val.Interface().(*big.Int)
	if i == nil {
		i = new(big.Int)
		val.Set(reflect.ValueOf(i))
	}
	// Reject leading zero bytes
	if len(b) > 0 && b[0] == 0 {
		return wrapStreamError(ErrCanonInt, val.Type())
	}
	i.SetBytes(b)
	return nil
}

func makeListDecoder(typ reflect.Type, tag tags) (decoder, error) {
	etype := typ.Elem()
	if etype.Kind() == reflect.Uint8 && !reflect.PtrTo(etype).Implements(decoderInterface) {
		if typ.Kind() == reflect.Array {
			return decodeByteArray, nil
		} else {
			return decodeByteSlice, nil
		}
	}
	etypeinfo, err := cachedTypeInfo1(etype, tags{})
	if err != nil {
		return nil, err
	}
	var dec decoder
	switch {
	case typ.Kind() == reflect.Array:
		dec = func(s *Stream, val reflect.Value) error {
			return decodeListArray(s, val, etypeinfo.decoder)
		}
	case tag.tail:
		dec = func(s *Stream, val reflect.Value) error {
			return decodeSliceElems(s, val, etypeinfo.decoder)
		}
	default:
		dec = func(s *Stream, val reflect.Value) error {
			return decodeListSlice(s, val, etypeinfo.decoder)
		}
	}
	return dec, nil
}

func decodeListSlice(s *Stream, val reflect.Value, elemdec decoder) error {
	size, err := s.List()
	if err != nil {
		return wrapStreamError(err, val.Type())
	}
	if size == 0 {
		val.Set(reflect.MakeSlice(val.Type(), 0, 0))
		return s.ListEnd()
	}
	if err := decodeSliceElems(s, val, elemdec); err != nil {
		return err
	}
	return s.ListEnd()
}

func decodeSliceElems(s *Stream, val reflect.Value, elemdec decoder) error {
	i := 0
	for ; ; i++ {
		// grow slice if necessary
		if i >= val.Cap() {
			newcap := val.Cap() + val.Cap()/2
			if newcap < 4 {
				newcap = 4
			}
			newv := reflect.MakeSlice(val.Type(), val.Len(), newcap)
			reflect.Copy(newv, val)
			val.Set(newv)
		}
		if i >= val.Len() {
			val.SetLen(i + 1)
		}
		// decode into element
		if err := elemdec(s, val.Index(i)); err == EOL {
			break
		} else if err != nil {
			return addErrorContext(err, fmt.Sprint("[", i, "]"))
		}
	}
	if i < val.Len() {
		val.SetLen(i)
	}
	return nil
}

func decodeListArray(s *Stream, val reflect.Value, elemdec decoder) error {
	if _, err := s.List(); err != nil {
		return wrapStreamError(err, val.Type())
	}
	vlen := val.Len()
	i := 0
	for ; i < vlen; i++ {
		if err := elemdec(s, val.Index(i)); err == EOL {
			break
		} else if err != nil {
			return addErrorContext(err, fmt.Sprint("[", i, "]"))
		}
	}
	if i < vlen {
		return &decodeError{msg: "input list has too few elements", typ: val.Type()}
	}
	return wrapStreamError(s.ListEnd(), val.Type())
}

func decodeByteSlice(s *Stream, val reflect.Value) error {
	b, err := s.Bytes()
	if err != nil {
		return wrapStreamError(err, val.Type())
	}
	val.SetBytes(b)
	return nil
}

func decodeByteArray(s *Stream, val reflect.Value) error {
	kind, size, err := s.Kind()
	if err != nil {
		return err
	}
	vlen := val.Len()
	switch kind {
	case Byte:
		if vlen == 0 {
			return &decodeError{msg: "input string too long", typ: val.Type()}
		}
		if vlen > 1 {
			return &decodeError{msg: "input string too short", typ: val.Type()}
		}
		bv, _ := s.Uint()
		val.Index(0).SetUint(bv)

	case String:
		if uint64(vlen) < size {
			return &decodeError{msg: "input string too long", typ: val.Type()}
		}
		if uint64(vlen) > size {
			return &decodeError{msg: "input string too short", typ: val.Type()}
		}
		slice := val.Slice(0, vlen).Interface().([]byte)
		if err := s.readFull(slice); err != nil {
			return err
		}
		// Reject cases where single byte encoding should have been used.
		if size == 1 && slice[0] < 128 {
			return wrapStreamError(ErrCanonSize, val.Type())
		}
	case List:
		return wrapStreamError(ErrExpectedString, val.Type())
	}
	return nil
}

func makeStructDecoder(typ reflect.Type) (decoder, error) {
	fields, err := structFields(typ)
	if err != nil {
		return nil, err
	}
	dec := func(s *Stream, val reflect.Value) (err error) {
		if _, err := s.List(); err != nil {
			return wrapStreamError(err, typ)
		}
		for _, f := range fields {
			err := f.info.decoder(s, val.Field(f.index))
			if err == EOL {
				return &decodeError{msg: "too few elements", typ: typ}
			} else if err != nil {
				return addErrorContext(err, "."+typ.Field(f.index).Name)
			}
		}
		return wrapStreamError(s.ListEnd(), typ)
	}
	return dec, nil
}

// makePtrDecoder creates a decoder that decodes into
// the pointer's element type.
func makePtrDecoder(typ reflect.Type) (decoder, error) {
	etype := typ.Elem()
	etypeinfo, err := cachedTypeInfo1(etype, tags{})
	if err != nil {
		return nil, err
	}
	dec := func(s *Stream, val reflect.Value) (err error) {
		newval := val
		if val.IsNil() {
			newval = reflect.New(etype)
		}
		if err = etypeinfo.decoder(s, newval.Elem()); err == nil {
			val.Set(newval)
		}
		return err
	}
	return dec, nil
}

// makeOptionalPtrDecoder creates a decoder that decodes empty values
// as nil. Non-empty values are decoded into a value of the element type,
// just like makePtrDecoder does.
//
// This decoder is used for pointer-typed struct fields with struct tag "nil".
func makeOptionalPtrDecoder(typ reflect.Type) (decoder, error) {
	etype := typ.Elem()
	etypeinfo, err := cachedTypeInfo1(etype, tags{})
	if err != nil {
		return nil, err
	}
	dec := func(s *Stream, val reflect.Value) (err error) {
		kind, size, err := s.Kind()

		if err != nil || size == 0 && kind != Byte {
			// rearm s.Kind. This is important because the input
			// position must advance to the next value even though
			// we don't read anything.
			s.kind = -1
			// set the pointer to nil.
			val.Set(reflect.Zero(typ))
			return err
		}

		newval := val
		if val.IsNil() {
			newval = reflect.New(etype)
		}
		if err = etypeinfo.decoder(s, newval.Elem()); err == nil {
			val.Set(newval)
		}
		return err
	}
	return dec, nil
}

var ifsliceType = reflect.TypeOf([]interface{}{})

func decodeInterface(s *Stream, val reflect.Value) error {
	if val.Type().NumMethod() != 0 {
		return fmt.Errorf("rlp: type %v is not RLP-serializable", val.Type())
	}
	kind, _, err := s.Kind()
	if err != nil {
		return err
	}
	if kind == List {
		slice := reflect.New(ifsliceType).Elem()
		if err := decodeListSlice(s, slice, decodeInterface); err != nil {
			return err
		}
		val.Set(slice)
	} else {
		b, err := s.Bytes()
		if err != nil {
			return err
		}
		val.Set(reflect.ValueOf(b))
	}
	return nil
}

// This decoder is used for non-pointer values of types
// that implement the Decoder interface using a pointer receiver.
func decodeDecoderNoPtr(s *Stream, val reflect.Value) error {
	return val.Addr().Interface().(Decoder).DecodeRLP(s)
}

func decodeDecoder(s *Stream, val reflect.Value) error {
	// Decoder instances are not handled using the pointer rule if the type
	// implements Decoder with pointer receiver (i.e. always)
	// because it might handle empty values specially.
	// We need to allocate one here in this case, like makePtrDecoder does.
	if val.Kind() == reflect.Ptr && val.IsNil() {
		val.Set(reflect.New(val.Type().Elem()))
	}
	return val.Interface().(Decoder).DecodeRLP(s)
}

// Kind represents the kind of value contained in an RLP stream.
type Kind int

const (
	Byte Kind = iota
	String
	List
)

func (k Kind) String() string {
	switch k {
	case Byte:
		return "Byte"
	case String:
		return "String"
	case List:
		return "List"
	default:
		return fmt.Sprintf("Unknown(%d)", k)
	}
}

var (
	// EOL is returned when the end of the current list
	// has been reached during streaming.
	EOL = errors.New("rlp: end of list")

	// Actual Errors
	ErrExpectedString = errors.New("rlp: expected String or Byte")
	ErrExpectedList   = errors.New("rlp: expected List")
	ErrCanonInt       = errors.New("rlp: non-canonical integer format")
	ErrCanonSize      = errors.New("rlp: non-canonical size information")
	ErrElemTooLarge   = errors.New("rlp: element is larger than containing list")
	ErrValueTooLarge  = errors.New("rlp: value size exceeds available input length")

	// This error is reported by DecodeBytes if the slice contains
	// additional data after the first RLP value.
	ErrMoreThanOneValue = errors.New("rlp: input contains more than one value")

	// internal errors
	errNotInList    = errors.New("rlp: call of ListEnd outside of any list")
	errNotAtEOL     = errors.New("rlp: call of ListEnd not positioned at EOL")
	errUintOverflow = errors.New("rlp: uint overflow")
)

// ByteReader must be implemented by any input reader for a Stream. It
// is implemented by e.g. bufio.Reader and bytes.Reader.
type ByteReader interface {
	io.Reader
	io.ByteReader
}


// Stream is not safe for concurrent use.
type Stream struct {
	r ByteReader

	// number of bytes remaining to be read from r.
	remaining uint64
	limited   bool

	// auxiliary buffer for integer decoding
	uintbuf []byte

	kind    Kind   // kind of value ahead
	size    uint64 // size of value ahead
	byteval byte   // value of single byte in type tag
	kinderr error  // error from last readKind
	stack   []listpos
}
