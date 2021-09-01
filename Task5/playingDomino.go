package Task5

/*
	Input: PlayingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	Output: [3,4]

	Input: PlayingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))
	Output: "tutup kartu"
*/
func PlayingDomino(cards [][]int, deck []int) (hasil interface{}) {
	n := len(cards)
	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			if cards[i][j] == deck[0] || cards[i][j] == deck[1] {
				if cards[i][0]+cards[i][1] > max {
					max = cards[i][0] + cards[i][1]
					hasil = []int{cards[i][0], cards[i][1]}
				}
			}
		}
	}
	if hasil == nil {
		return "tutup kartu"
	}
	return hasil
}
