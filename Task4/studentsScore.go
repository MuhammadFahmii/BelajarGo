package Task4

type Student struct {
	name  []string
	nilai []int
}

func (s Student) Average() (hasil float32) {
	for _, val := range s.nilai {
		hasil += float32(val)
	}
	hasil = hasil / float32(len(s.nilai))
	return
}

func (s Student) Min() (min int, name string) {
	for _, val := range s.nilai {
		for j, val1 := range s.nilai {
			if val1 == 0 {
				min, name = val1, s.name[j]
			} else if val1 < val {
				min, name = val1, s.name[j]
			}
		}
	}
	return
}

func (s Student) Max() (max int, name string) {
	for _, val := range s.nilai {
		for j, val1 := range s.nilai {
			if val1 > val {
				max, name = val1, s.name[j]
			}
		}
	}
	return
}
