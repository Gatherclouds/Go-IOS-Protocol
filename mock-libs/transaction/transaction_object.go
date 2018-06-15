package transaction

import "reflect"

func (this *TransactionIndex) create(constructor func(*TransactionObject)) *TransactionObject {

	var obj TransactionObject
	obj.id = this.get_next_available_id()
	constructor(&obj)

	result := this._index.insert(move(obj))
	return &(*result.first)
}

func (this *TransactionIndex) remove(id ObjectIdType) {
	index := this._index.get()
	itr := index.find(id.instance())

	if itr == index.end() {
		return
	}
	inde.erase(itr)
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

