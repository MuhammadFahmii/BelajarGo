package Task4

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var tmp int
	for _, val := range *&s.name {
		if int(val) <= 109 {
			tmp = int(val) + ((109-int(val))*2 + 1)
		} else {
			tmp = int(val) - ((int(val)-110)*2 + 1)
		}
		*&s.nameEncode += string(tmp)
	}
	return *&s.nameEncode
}

func (s *student) Decode() string {
	var tmp int
	var nameDecode string
	for _, val := range *&s.nameEncode {
		if int(val) <= 109 {
			tmp = int(val) + ((109-int(val))*2 + 1)
		} else {
			tmp = int(val) - ((int(val)-110)*2 + 1)
		}
		nameDecode += string(tmp)
	}
	return nameDecode
}
