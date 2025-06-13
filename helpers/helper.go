package helpers

func CheckRole(userID string) string {
	if userID[0] != '1' {
		return "mahasiswa"
	} else if userID == "121121121" {
		return "admin"
	} else {
		return "dosen"
	}
}
