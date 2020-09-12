package employee

import (
	"fmt"
)

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

//New fun
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

//LeavesRemaining func
func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
