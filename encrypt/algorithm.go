package encrypt

func Nimbus(msg string) string {
	encrypted := ""
	for _, char := range msg {
		encrypted += string(char + 3)

	}
	return encrypted
}
