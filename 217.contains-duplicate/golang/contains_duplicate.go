package contains_duplicate

func containsDuplicate(nums []int) bool {
	visited := MySet{container: make(map[int]struct{})}
	for _, num := range nums {
		if visited.Exists(num) {
			return true
		}
		visited.Add(num)
	}
	return false
}

type MySet struct {
	container map[int]struct{}
}

func (s *MySet) Exists(key int) bool {
	_, exists := s.container[key]
	return exists
}

func (s *MySet) Add(key int) {
	s.container[key] = struct{}{}
}
