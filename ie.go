package gobr

func ValidaIE(ie string, uf string) bool {
	switch uf {
	case "SP", "sp":
		return true
	default:
		return false
	}
}
