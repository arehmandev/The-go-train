package accumulate

const testVersion = 1

func Accumulate(list []string, function func(string) string) []string {

	var array []string

	for _, element := range list {
		value := function(element)
		array = append(array, value)
	}
	return array

}
