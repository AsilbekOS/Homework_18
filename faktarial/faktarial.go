package faktarial

func Fact(son int) int {
	if son <= 1 {
		return son
	}

	return son * Fact(son-1)
}
