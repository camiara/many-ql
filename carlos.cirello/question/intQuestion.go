package question

import "strconv"

// IntQuestion stores the answer of question which type is integer numeric
type IntQuestion int

// FromString takes the input from Frontend and stores locally - Int
func (s *IntQuestion) FromString(str string) error {
	val, err := strconv.Atoi(str)
	*s = IntQuestion(val)
	return err
}

// String prints in human form the content of the question - Int
func (s IntQuestion) String() string {
	return strconv.Itoa(int(s))
}
