package utils

func ContainsValue(m map[string]string, value string) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}

	return false
}

func ContainsKey(m map[string]string, value string) bool {
	_, contains := m[value]

	return contains
}

func GetKeyByValue(m map[string]string, value string) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			return k, true
		}
	}

	return "", false
}
