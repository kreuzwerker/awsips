package awsips

type Node map[string]interface{}

func N() Node {
	return make(map[string]interface{})
}

func (b Node) L(key string, value interface{}) Node {
	b[key] = value
	return b
}

func (b Node) N(key string) Node {
	b[key] = N()
	return b[key].(Node)
}
