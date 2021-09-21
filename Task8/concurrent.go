package task8

type FreqMap map[string]int

func frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[string(r)]++
	}
	return m
}

func concurrentFreq(s []string) FreqMap {
	sum := FreqMap{}
	count := len(s)
	res := make(chan FreqMap, count)

	for _, s := range s {
		go func(s string) {
			res <- frequency(s)
		}(s)
	}

	for i := 0; i < count; i++ {
		for r, req := range <-res {
			sum[string(r)] += req
		}
	}
	return sum
}
