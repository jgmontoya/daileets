package check_inclusion

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	left_index := 0
	right_index := len(s1) - 1

	s1_chars := makeDefaultDictFromWord(s1)

	for right_index < len(s2) {
		substring := s2[left_index : right_index+1]
		substring_chars := makeDefaultDictFromWord(substring)
		if s1_chars.Equals(substring_chars) {
			return true
		}
		left_index++
		right_index++
	}

	return false
}

type DefaultDict struct {
	container map[rune]int
}

func makeDefaultDict() *DefaultDict {
	return &DefaultDict{
		container: make(map[rune]int),
	}
}

func makeDefaultDictFromWord(word string) *DefaultDict {
	default_dict := makeDefaultDict()
	for _, char := range word {
		default_dict.Add(char)
	}
	return default_dict
}

func (d *DefaultDict) Add(char rune) int {
	counter, exists := d.container[char]
	if exists {
		d.container[char]++
		return counter + 1
	}

	d.container[char] = 1
	return 1
}

func (d *DefaultDict) AllZero() bool {
	for _, counter := range d.container {
		if counter != 0 {
			return false
		}
	}
	return true
}

func (d *DefaultDict) Equals(other *DefaultDict) bool {
	for char, counter := range d.container {
		if counter != other.container[char] {
			return false
		}
	}
	return true
}
