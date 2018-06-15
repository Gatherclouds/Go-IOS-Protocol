package transaction

func (this *TransactionIndex) create(constructor func(*TransactionObject)) *TransactionObject {

	var obj TransactionObject
	obj.id = this.get_next_available_id()
	constructor(&obj)

	result := this._index.insert(move(obj))
	return &(*result.first)
}





