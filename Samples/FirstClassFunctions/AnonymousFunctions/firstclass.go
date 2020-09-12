package firstclass

//Student  struct
type Student struct {
	FirstName string
	LastName  string
	Grade     string
	Country   string
}

//Filter func
func Filter(s []Student, f func(s Student) bool) []Student {

	var r []Student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}
