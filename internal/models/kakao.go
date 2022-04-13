package kakao

type Kakao []interface{}

type K map[string]interface{}

// Add : 모든 객체를 Kakao 객체로 add
func (k *Kakao) Add(s ...interface{}) {
	for _, inter := range s {
		*k = append(*k, inter)
	}
}
