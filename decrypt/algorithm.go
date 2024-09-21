package decrypt

func Nimbus(msg string) string {
	decrypted := ""
	for _, char := range msg {
		decrypted += string(char - 3)

	}
	return decrypted
}
