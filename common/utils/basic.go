package utils

func IsNull(value interface{}) bool {
	return nil == value
}

func IsTrue(value *bool, defaultValue bool) bool {
	if nil == value {
		return defaultValue
	}

	return *value
}

func IsFalse(value *bool, defaultValue bool) bool {
	return !IsTrue(value, defaultValue)
}

func Strcpy(value *string) *string {
	if nil == value {
		return nil
	}

	newValue := *value
	return &newValue
}
