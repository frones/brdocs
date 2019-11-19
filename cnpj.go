package gobr

import (
	"regexp"
	"strings"
)

func ValidaCNPJ(cnpj string) bool {
	if r, err := regexp.MatchString(`^\d{2}\.\d{3}\.\d{3}\/\d{4}-\d{2}$|^\d{14}$`, cnpj); err != nil || !r {
		return false
	}

	cnpj = strings.Replace(cnpj, ".", "", -1)
	cnpj = strings.Replace(cnpj, "/", "", -1)
	cnpj = strings.Replace(cnpj, "-", "", -1)

	if r, err := regexp.MatchString(`^0{14}$`, cnpj); err != nil || r {
		return false
	}

	sum := 0
	for i, c := range cnpj[:len(cnpj)-2] {
		sum += int(c-'0') * (((len(cnpj) - 3 - i) % 8) + 2)
	}
	dv := ((sum * 10) % 11) % 10
	if dv != int(cnpj[12]-'0') {
		return false
	}

	sum = 0
	for i, c := range cnpj[:len(cnpj)-1] {
		sum += int(c-'0') * (((len(cnpj) - 2 - i) % 8) + 2)
	}
	dv = ((sum * 10) % 11) % 10
	if dv != int(cnpj[13]-'0') {
		return false
	}

	return true
}

func FormataCNPJ(cnpj string) string {
	cnpj = strings.Replace(cnpj, ".", "", -1)
	cnpj = strings.Replace(cnpj, "/", "", -1)
	cnpj = strings.Replace(cnpj, "-", "", -1)

	if len(cnpj) < 14 {
		cnpj = cnpj + strings.Repeat(" ", 14-len(cnpj))
	}

	return cnpj[:2] + "." + cnpj[2:5] + "." + cnpj[5:8] + "/" + cnpj[8:12] + "-" + cnpj[12:]
}
