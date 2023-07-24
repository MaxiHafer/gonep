package pointer

func Of[T any](v T) *T {
	pV := &v
	return pV
}
