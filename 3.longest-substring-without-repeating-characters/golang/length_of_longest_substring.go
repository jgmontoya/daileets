package length_of_longest_substring

func LengthOfLongestSubstring(s string) int {
	left_index := 0
	right_index := 0
	max_length := 0

	substring_letters := makeSet()

	for right_index < len(s) {
		right_letter := s[right_index]
		if !substring_letters.Exists(right_letter) {
			substring_letters.Add(right_letter)
			number_of_letters := substring_letters.Size()
			if number_of_letters > max_length {
				max_length = number_of_letters
			}
			right_index++
		} else {
			for substring_letters.Exists(right_letter) {
				left_letter := s[left_index]
				substring_letters.Remove(left_letter)
				left_index++
			}
		}

	}

	return max_length
}

type MySet struct {
	container map[byte]struct{}
	size      int
}

func makeSet() *MySet {
	return &MySet{
		container: make(map[byte]struct{}),
	}
}

func (s *MySet) Exists(key byte) bool {
	_, exists := s.container[key]
	return exists
}

func (s *MySet) Add(key byte) {
	if !s.Exists(key) {
		s.size++
	}
	s.container[key] = struct{}{}
}

func (s *MySet) Remove(key byte) {
	if s.Exists(key) {
		s.size--
	}
	delete(s.container, key)
}

func (s *MySet) Size() int {
	return s.size
}
