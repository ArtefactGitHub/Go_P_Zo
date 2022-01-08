package common

// QueryMap
type QueryMap map[string]string

func (qm QueryMap) Get(name string) string {
	if _, ok := qm[name]; ok {
		return qm[name]
	}

	return ""
}
