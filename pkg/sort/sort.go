package sort

import (
	"interview.go/models"
)

type Stats []models.DomainStats
type CustomerStats []models.Customer

func (s Stats) Len() int {
	return len(s)
}
func (s Stats) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (c CustomerStats) Len() int {
	return len(c)
}
func (c CustomerStats) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type ByDomain struct {
	Stats
}

func (s ByDomain) Less(i, j int) bool {
	return s.Stats[i].Domain < s.Stats[j].Domain
}

type ByEmail struct {
	CustomerStats
}

func (c ByEmail) Less(i, j int) bool {
	return c.CustomerStats[i].Email < c.CustomerStats[j].Email
}
