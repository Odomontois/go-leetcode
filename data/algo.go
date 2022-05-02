package data

func MapSlice[In any, Out any](in []In, f func(In) Out) []Out {
	res := make([]Out, len(in))
	for i := range res {
		res[i] = f(in[i])
	}
	return res
}
