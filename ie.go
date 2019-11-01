package gobr

import (
	"fmt"
	"regexp"
	"strings"
)

func modulo11(str string, pesomax int) (int, int) {
	sum := 0
	for i, c := range str {
		sum += int(c-'0') * (((len(str) - 1 - i) % (pesomax - 1)) + 2)
	}
	dv := ((sum * 10) % 11) % 10
	return dv, sum
}

func modulo10(str string, pesomax int) (int, int) {
	sum := 0
	for i, c := range str {
		sum += int(c-'0') * (((len(str) - 1 - i) % (pesomax - 1)) + 2)
	}
	dv := (sum * 9) % 10
	return dv, sum
}

// Identidades matemáticas úteis para facilitar/uniformizar a codificação das regras, já que as Sefazes decidiram escrever (quase) a mesma coisa de um jeito, cada uma de um jeito:
// (a/b)*b + a%b = a
// (b - a%b)%b = (a*b - a)%b
func ValidaIE(ie string, uf string) bool {
	switch strings.ToUpper(uf) {
	case "AC":
		if r, err := regexp.MatchString(`^01\.\d{3}\.\d{3}\/\d{3}-\d{2}$|^01\d{11}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, ".", "", -1)
			ie = strings.Replace(ie, "/", "", -1)
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-2], 9)
			if dv != int(ie[len(ie)-2]-'0') {
				return false
			}

			dv, _ = modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "AL":
		if r, err := regexp.MatchString(`^24[03578]\d{6}$`, ie); err != nil || !r {
			return false
		} else {
			dv, _ := modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "AP":
		if r, err := regexp.MatchString(`^03\d{7}$`, ie); err != nil || !r {
			return false
		} else {
			var p, d int
			if ie <= "03017000" {
				p = 5
				d = 0
			} else if ie <= "03019022" {
				p = 9
				d = 1
			} else {
				p = 0
				d = 0
			}
			_, sum := modulo11(ie[:len(ie)-1], 9)
			sum += p
			dv := 11 - (sum % 11)
			if dv == 11 {
				dv = d
			} else {
				dv = dv % 10
			}
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "AM":
		if r, err := regexp.MatchString(`^\d{2}\.\d{3}\.\d{3}-\d$|^\d{9}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, ".", "", -1)
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "BA":
		if r, err := regexp.MatchString(`^\d{6}-\d{2}$|^\d{8}$`, ie); err != nil || !r {
			if r, err := regexp.MatchString(`^\d{7}-\d{2}$|^\d{9}$`, ie); err != nil || !r {
				return false
			} else {
				ie = strings.Replace(ie, "-", "", -1)

				if strings.ContainsAny(ie[1:2], "0123458") {
					dv, _ := modulo10(ie[:len(ie)-2], 9)
					if dv != int(ie[len(ie)-1]-'0') {
						return false
					}

					dv, _ = modulo10(ie[:len(ie)-2]+ie[len(ie)-1:], 9)
					return dv == int(ie[len(ie)-2]-'0')
				} else {
					dv, _ := modulo11(ie[:len(ie)-2], 9)
					if dv != int(ie[len(ie)-1]-'0') {
						return false
					}

					dv, _ = modulo11(ie[:len(ie)-2]+ie[len(ie)-1:], 9)
					return dv == int(ie[len(ie)-2]-'0')
				}
			}
		} else {
			ie = strings.Replace(ie, "-", "", -1)

			sum := 0
			for i, c := range ie[:len(ie)-2] {
				sum += int(c-'0') * (((len(ie) - 3 - i) % 8) + 2)
			}
			if strings.ContainsAny(ie[0:1], "0123458") {
				dv, _ := modulo10(ie[:len(ie)-2], 9)
				if dv != int(ie[len(ie)-1]-'0') {
					return false
				}

				dv, _ = modulo10(ie[:len(ie)-2]+ie[len(ie)-1:], 9)
				return dv == int(ie[len(ie)-2]-'0')
			} else {
				dv, _ := modulo11(ie[:len(ie)-2], 9)
				if dv != int(ie[len(ie)-1]-'0') {
					return false
				}

				dv, _ = modulo11(ie[:len(ie)-2]+ie[len(ie)-1:], 9)
				return dv == int(ie[len(ie)-2]-'0')
			}
		}
	case "CE":
		if r, err := regexp.MatchString(`^\d{8}-\d$|^\d{9}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "DF":
		if r, err := regexp.MatchString(`^\d{11}-\d{2}$|^\d{13}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-2], 9)
			if dv != int(ie[len(ie)-2]-'0') {
				return false
			}

			dv, _ = modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "ES":
		if r, err := regexp.MatchString(`^\d{8}-\d$|^\d{9}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "GO":
		if r, err := regexp.MatchString(`^(10|11|15)\.\d{3}\.\d{3}-\d$|^(10|11|15)\d{7}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, "-", "", -1)
			ie = strings.Replace(ie, ".", "", -1)

			dv, sum := modulo11(ie[:len(ie)-1], 9)
			if (ie >= "10103105") && (ie <= "10119997") && ((sum % 11) == 1) {
				dv = 1
			}
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "MA":
		if r, err := regexp.MatchString(`^12\d{6}-\d$|^12\d{7}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "MT":
		if r, err := regexp.MatchString(`^\d{10}-\d$|^\d{11}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "MS":
		if r, err := regexp.MatchString(`^28\d{6}-\d$|^28\d{7}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "MG":
		if r, err := regexp.MatchString(`^\d{3}\.\d{3}\.\d{3}\/\d{4}|^\d{13}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, ".", "", -1)
			ie = strings.Replace(ie, "/", "", -1)

			sum := 0
			for i, c := range ie[:3] {
				sum += (int(c-'0') * (i%2 + 1)) / 10
				sum += (int(c-'0') * (i%2 + 1)) % 10
			}
			for i, c := range ie[3 : len(ie)-2] {
				sum += (int(c-'0') * (i%2 + 1)) / 10
				sum += (int(c-'0') * (i%2 + 1)) % 10
			}
			dv := (((sum/10)+1)*10 - sum) % 10
			if dv != int(ie[len(ie)-2]-'0') {
				return false
			}

			dv, _ = modulo11(ie[:len(ie)-1], 11)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "PA":
		if r, err := regexp.MatchString(`^15-\d{6}-\d$|^15\d{7}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "PB":
		if r, err := regexp.MatchString(`^\d{8}-\d$|^\d{9}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "PR":
		if r, err := regexp.MatchString(`^\d{3}\.\d{5}-\d{2}$|^\d{10}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, ".", "", -1)
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-2], 7)
			if dv != int(ie[len(ie)-2]-'0') {
				return false
			}

			dv, _ = modulo11(ie[:len(ie)-1], 7)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "PE":
		if r, err := regexp.MatchString(`^\d{7}-\d{2}$|^\d{9}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-2], 9)
			if dv != int(ie[len(ie)-2]-'0') {
				return false
			}

			dv, _ = modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "PI":
		if r, err := regexp.MatchString(`^\d{8}-\d$|^\d{9}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "RJ":
		if r, err := regexp.MatchString(`^\d{2}\.\d{3}\.\d{2}-\d$|^\d{8}$`, ie); err != nil || !r {
			return false
		} else {
			ie = strings.Replace(ie, ".", "", -1)
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[:len(ie)-1], 7)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "RN":
		if r, err := regexp.MatchString(`^20\.\d{3}\.\d{3}-\d$|^\d{9}$`, ie); err != nil || !r {
			if r, err := regexp.MatchString(`^20\.\d\.\d{3}\.\d{3}-\d$|^\d{10}$`, ie); err != nil || !r {
				return false
			}
		}
		ie = strings.Replace(ie, ".", "", -1)
		ie = strings.Replace(ie, "-", "", -1)

		dv, _ := modulo11(ie[:len(ie)-1], 10)
		return dv == int(ie[len(ie)-1]-'0')
	case "RS":
		if r, err := regexp.MatchString(`^\d{3}\/\d{7}$|^\d{10}$`, ie); err != nil || !r {
			return false
		}
		ie = strings.Replace(ie, "/", "", -1)

		dv, _ := modulo11(ie[:len(ie)-1], 9)
		return dv == int(ie[len(ie)-1]-'0')
	case "RO":
		if r, err := regexp.MatchString(`^\d{3}\.\d{5}-\d$|^\d{9}$`, ie); err != nil || !r {
			if r, err := regexp.MatchString(`^\d{13}-\d$|^\d{14}$`, ie); err != nil || !r {
				return false
			} else {
				ie = strings.Replace(ie, "-", "", -1)

				dv, _ := modulo11(ie[:len(ie)-1], 9)
				return dv == int(ie[len(ie)-1]-'0')
			}
		} else {
			ie = strings.Replace(ie, ".", "", -1)
			ie = strings.Replace(ie, "-", "", -1)

			dv, _ := modulo11(ie[3:len(ie)-1], 9)
			return dv == int(ie[len(ie)-1]-'0')
		}
	case "RR":
		if r, err := regexp.MatchString(`^24\d{6}-\d$|^24\d{7}$`, ie); err != nil || !r {
			return false
		}
		ie = strings.Replace(ie, "-", "", -1)

		sum := 0
		for i, c := range ie[:len(ie)-1] {
			sum += int(c-'0') * (i + 1)
		}
		dv := (sum % 9)
		return dv == int(ie[len(ie)-1]-'0')
	case "SC":
		if r, err := regexp.MatchString(`^\d{3}\.\d{3}\.\d{3}$|^\d{9}$`, ie); err != nil || !r {
			return false
		}
		ie = strings.Replace(ie, ".", "", -1)

		dv, _ := modulo11(ie[:len(ie)-1], 9)
		return dv == int(ie[len(ie)-1]-'0')
	case "SP":
		if r, err := regexp.MatchString(`^\d{3}\.\d{3}\.\d{3}\.\d{3}$|^\d{12}$`, ie); err != nil || !r {
			if r, err := regexp.MatchString(`^P-\d{8}\.\d\/\d{3}$|^P\d{12}$`, ie); err != nil || !r {
				return false
			} else {
				//Inscrição Estadual de Produtor Rural
				ie = strings.Replace(ie, ".", "", -1)
				ie = strings.Replace(ie, "/", "", -1)
				ie = strings.Replace(ie, "-", "", -1)
				ie = strings.Replace(ie, "P", "", -1)

				sum := int(ie[0]-'0')*1 + int(ie[7]-'0')*10
				for i, c := range ie[1:7] {
					sum += int(c-'0') * (i + 3)
				}
				dv := (sum % 11) % 10
				if dv != int(ie[8]-'0') {
					return false
				}

				return true
			}
		} else {
			//Inscrição Estadual de Industriais e Comerciantes
			ie = strings.Replace(ie, ".", "", -1)

			sum := int(ie[0]-'0')*1 + int(ie[7]-'0')*10
			for i, c := range ie[1:7] {
				sum += int(c-'0') * (i + 3)
			}
			dv := (sum % 11) % 10
			if dv != int(ie[8]-'0') {
				return false
			}

			dv, _ = modulo11(ie[:len(ie)-2], 10)
			return dv == int(ie[11]-'0')
		}
	case "SE":
		if r, err := regexp.MatchString(`^\d{8}-\d$|^\d{9}$`, ie); err != nil || !r {
			return false
		}
		ie = strings.Replace(ie, "-", "", -1)

		dv, _ := modulo11(ie[:len(ie)-1], 9)
		return dv == int(ie[len(ie)-1]-'0')
	case "TO":
		if r, err := regexp.MatchString(`^\d{2}(01|02|03|99)\d{6}-\d$|^\d{2}(01|02|03|99)\d{7}$`, ie); err != nil || !r {
			return false
		}
		ie = strings.Replace(ie, "-", "", -1)

		dv, _ := modulo11(ie[:2]+ie[4:len(ie)-1], 9)
		return dv == int(ie[len(ie)-1]-'0')
	default:
		return false
	}
}
