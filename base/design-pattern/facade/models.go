package facade

type deduplication struct {
	hashMap map[interface{}]int
}

// NewDeduplication 获取去重对象
func NewDeduplication() Deduplication {
	d := new(deduplication)
	d.hashMap = make(map[interface{}]int)
	return d
}
func (d *deduplication) Put(data interface{}) bool {
	ret := true
	defer throwErr(&ret)
	d.hashMap[data] = 1
	return ret
}

func (d *deduplication) Exist(data interface{}) bool {
	ret := true
	defer throwErr(&ret)
	value, ok := d.hashMap[data]
	if ok && value == 1 {
		return ret
	}
	return false
}

func (d *deduplication) Delete(data interface{}) bool {
	ret := true
	defer throwErr(&ret)
	if !d.Exist(data) {
		return false
	}
	delete(d.hashMap, data)
	return ret
}

func throwErr(ret *bool) {
	err := recover()
	if err != nil {
		*ret = false
	}
}
