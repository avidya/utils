// the functions are not mandatory to be written in recursive way, iterative way is also allowed.
// but, once talk about recursive way, they are supposed to follow the tail recursion style.
// lazy-evaluation is not supported so far.

type Number interface {
	Int | Uint | Float
}

// a ball-ache Mapper implementation
func Map[E, T any](mapper func(t E) T, ts []E) []T {
	return rMap(mapper, ts, []T{})
}

func rMap[E, T any](mapper func(t E) T, ts []E, result []T) []T {
	if len(ts) == 0 {
		return result
	} else {
		return rMap(mapper, ts[1:], append(result, mapper(ts[0])))
	}
}

// a ball-ache Filter implementation
func Filter[T any](filter func(t T) bool, ts []T) []T {
	return rFilter(filter, ts, []T{})
}

func rFilter[T any](filter func(t T) bool, ts []T, result []T) []T {
	if len(ts) == 0 {
		return result
	} else if filter(ts[0]) {
		return rFilter(filter, ts[1:], append(result, ts[0]))
	} else {
		return rFilter(filter, ts[1:], result)
	}
}

func Reduce[S any, T any](accumulator func(s S, a T) S, ts []T, init interface{}) S {
	if len(ts) == 0 {
		return init.(S)
	}
	if init == nil {
		return Reduce(accumulator, ts[1:], ts[0])
	} else {
		return Reduce(accumulator, ts[1:], accumulator(init.(S), ts[0]))
	}
}

func Max[T Number](ts []T) T {
	return Reduce(func(a1 T, a2 T) T {
		if a1 > a2 {
			return a1
		} else {
			return a2
		}
	}, ts, nil)
}

// stupid implementation...
func InList[T Number](obj T, ts []T) bool {
	return Reduce(func(a1 bool, a2 T) bool {
		return a2 == obj || a1
	}, ts, false)
}

func Sum[T Number](ts []T) T {
	return Reduce(func(t T, a T) T {
		return t + a
	}, ts, nil)
}
