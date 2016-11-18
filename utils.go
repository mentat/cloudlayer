package cloudlayer

func SafeString(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}
