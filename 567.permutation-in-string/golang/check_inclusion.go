package check_inclusion

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	left_index := 0
	right_index := len(s1) - 1

	s1_chars := makeDefaultDictFromWord(s1)
	substring := s2[left_index : right_index+1]
	substring_chars := makeDefaultDictFromWord(substring)

	for right_index < len(s2)-1 {
		if s1_chars.Equals(substring_chars) {
			return true
		}
		substring_chars.Remove(s2[left_index])
		left_index++
		right_index++
		substring_chars.Add(s2[right_index])
	}

	return s1_chars.Equals(substring_chars)
}

type DefaultDict struct {
	container map[byte]int
}

func makeDefaultDict() *DefaultDict {
	return &DefaultDict{
		container: make(map[byte]int),
	}
}

func makeDefaultDictFromWord(word string) *DefaultDict {
	default_dict := makeDefaultDict()
	for index := range word {
		default_dict.Add(word[index])
	}
	return default_dict
}

func (d *DefaultDict) Add(char byte) {
	_, exists := d.container[char]
	if exists {
		d.container[char]++
		return
	}

	d.container[char] = 1
}

func (d *DefaultDict) Remove(char byte) {
	_, exists := d.container[char]
	if exists {
		d.container[char]--
		return
	}

	d.container[char] = -1
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
