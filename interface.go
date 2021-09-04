package utils

type IDataMap interface {
	ToMap() map[string]string
}

func TryToMap(x interface{}) interface{} {
	switch x.(type) {
	case IDataMap:
		return x.(IDataMap).ToMap()
	default:
		return x
	}
}
