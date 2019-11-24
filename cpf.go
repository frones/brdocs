package brdocs

import (
	"github.com/frones/strmask"
	"regexp"
	"strings"
)

func ValidaCPF(cpf string) bool {
	if r, err := regexp.MatchString(`^\d{3}\.\d{3}\.\d{3}-\d{2}$|^\d{11}$`, cpf); err != nil || !r {
		return false
	}

	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)

	if r, err := regexp.MatchString(`^0{11}|1{11}|2{11}|3{11}|4{11}|5{11}|6{11}|7{11}|8{11}|9{11}$`, cpf); err != nil || r {
		return false
	}

	sum := 0
	for i, c := range cpf[:len(cpf)-2] {
		sum += int(c-'0') * (len(cpf) - 1 - i)
	}
	dv := ((sum * 10) % 11) % 10
	if dv != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i, c := range cpf[:len(cpf)-1] {
		sum += int(c-'0') * (len(cpf) - i)
	}
	dv = ((sum * 10) % 11) % 10
	if dv != int(cpf[10]-'0') {
		return false
	}

	return true
}

func FormataCPF(cpf string) string {
	return strmask.FormatMask("000.000.000-00;0", cpf)
}
