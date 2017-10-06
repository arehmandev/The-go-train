package acronym

import (
	"strings"
)

const testVersion = 3

func Abbreviate(Acronyminput string) (output string) {

	test := strings.Split(Acronyminput, " ")

	var accronym []string

	for _, element := range test {
		if strings.ContainsAny(element, "-") {
			seperatewords := strings.Split(element, "-")

			for i, _ := range seperatewords {
				accronym = append(accronym, string(seperatewords[i][0]))
			}
		} else {
			accronym = append(accronym, string(element[0]))
		}
	}
	asword := strings.Join(accronym, "")
	output = strings.ToUpper(asword)

	return
}
