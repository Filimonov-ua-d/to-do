package output

import (
	"fmt"

	"interview.go/models"
)

type Outputer struct {
}

func NewOutputer() *Outputer {
	return &Outputer{}
}

// printStats prints the domain statistics.
func (o Outputer) PrintStats(s []models.DomainStats, errors []error) {
	for _, o := range s {
		fmt.Printf("Domain: %s \nAmount: %v\nCustomers: %v \n\n", o.Domain, o.Count, o.Customers)
	}
	if errors != nil {
		for _, e := range errors {
			fmt.Printf("Error: %v\n", e)
		}
	}
}
