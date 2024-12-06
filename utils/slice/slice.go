package slice

func Remove[T comparable](s []T, elem T) []T {
	result := s[:0]
	for _, v := range s {
		if v != elem {
			result = append(result, v)
		}
	}
	return result
}

func Contains[T comparable](s []T, elem T) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}

func Filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[K, V any](s []K, f func(K) V) []V {
	var result []V
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

func Reduce[K, V any](s []K, f func(V, K) V, init V) V {
	result := init
	for _, v := range s {
		result = f(result, v)
	}
	return result
}
