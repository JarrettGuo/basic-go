package set

func Intersection[T comparable](a, b []T) []T {
	m := make(map[T]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}
	var result []T
	for _, v := range b {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}
	return result
}

func Union[T comparable](a, b []T) []T {
	m := make(map[T]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		m[v] = struct{}{}
	}
	var result []T
	for v := range m {
		result = append(result, v)
	}
	return result
}

func Difference[T comparable](a, b []T) []T {
	m := make(map[T]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}
	var result []T
	for _, v := range b {
		if _, exists := m[v]; !exists {
			result = append(result, v)
		}
	}
	return result
}
