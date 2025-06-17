package helpers

func CheckRole(userID string) string {
	if userID == "6111111111" {
		return "admin"
	} else if userID[0] != '1' {
		return "mahasiswa"
	} else {
		return "dosen"
	}
}
