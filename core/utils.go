package core

func tokenize(text string) chan Symbol {
	ch := make(chan Symbol)
	go func() {
		for _, r := range text {
			ch <- Symbol(r)
		}
		close(ch)
	}()
	return ch
}

func CalcOccurrences(text string) map[Symbol]int {
	ps := map[Symbol]int{}
	for sym := range tokenize(text) {
		ps[sym]++
	}
	return ps
}

func textToCodeBits(text string, table map[Symbol]Code) chan uint8 {
	ch := make(chan uint8)
	go func() {
		codes := []Code{}
		for sym := range tokenize(text) {
			code := table[sym]
			codes = append(codes, code)
		}
		// EOF
		codes = append(codes, table[""])
		for _, code := range codes {
			for _, bit := range code {
				if bit == '1' {
					ch <- 1
				} else {
					ch <- 0
				}
			}
		}
		close(ch)
	}()
	return ch
}

func merge(a, b map[Symbol]Code) map[Symbol]Code {
	for k, v := range b {
		a[k] = v
	}
	return a
}
