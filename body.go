package umeng

// 结构参考 https://developer.umeng.com/docs/67966/detail/68343#h1-u6D88u606Fu53D1u90014

type Body map[string]interface{}

func (b Body) Set(key string, value interface{}) Body {
	b[key] = value
	return b
}

func (b Body) Get(key string) interface{} {
	if b == nil {
		return nil
	}
	return b[key]
}

func (b Body) Child(key string, fn func(body Body)) Body {
	_b := make(Body)
	fn(_b)
	b[key] = _b
	return b
}
