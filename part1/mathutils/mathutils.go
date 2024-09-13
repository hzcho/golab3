package mathutils

func Mathutils(num int) int {
	fac := 1

	for i := 2; i < num+1; i++ {
		fac *= i
	}

	return fac
}
