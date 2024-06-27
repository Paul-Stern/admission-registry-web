package model

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/icrowley/fake"
)

type RegistryEntry struct {
	ID         int
	DateTime   time.Time
	DTFormat   string
	PersonData PersonalInfo
}

type PersonalInfo struct {
	ID          int
	FirstName   string
	MiddleName  string
	LastName    string
	Birthdate   time.Time
	BDFormat    string
	Address     string
	InsuranceID string
}

type Entries []RegistryEntry

func (pi *PersonalInfo) randBD() time.Time {
	// return rand.Intn(100)
	return time.Date(
		2024-rand.Intn(100),
		time.Month(rand.Intn(11)+1),
		rand.Intn(27)+1,
		0, 0, 0, 0, time.Now().Location(),
	)
}

// RandPerson creates random person
func RandPerson() PersonalInfo {
	fake.SetLang("ru")
	pi := NewPerson(
		fake.MaleFirstName(),
		fake.MalePatronymic(),
		fake.MaleLastName(),
		*new(time.Time),
	)
	pi.Birthdate = pi.randBD()
	pi.InsuranceID = RandIns()
	return pi
}

// NewPerson creates and returns new Person.
func NewPerson(fname, mname, lname string, bd time.Time) PersonalInfo {
	return PersonalInfo{
		FirstName:  fname,
		MiddleName: mname,
		LastName:   lname,
		BDFormat:   time.DateOnly,
		Birthdate:  bd,
	}
}

// RandEntry generates random RegistryEntry.
func RandEntry(id int) RegistryEntry {
	// Generate random date and time in range of last 3 months and 30 days.
	rt := time.Now().AddDate(0, -1*rand.Intn(3), -1*rand.Intn(30))
	return NewEntry(id, RandPerson(), rt)
}

func NewEntry(id int, pd PersonalInfo, dt time.Time) RegistryEntry {
	return RegistryEntry{
		ID:         id,
		PersonData: pd,
		DTFormat:   time.DateTime,
		DateTime:   dt,
	}
}

// GenEntries generates a slice of n random entries.
func GenEntries(n int) Entries {
	es := Entries{}
	for i := 0; i < n; i++ {
		es = append(es, RandEntry(i))
	}
	return es
}

func RandIns() string {
	// Allocate memory for the slice in advance
	islice := make([]string, 20)
	// Insurance ID contains 16 numerals so we take 16 steps
	for i := 0; i < 16; i++ {
		if i%4 == 0 {
			islice = append(islice, " ")
		}
		n := fmt.Sprint(rand.Intn(10))
		islice = append(islice, n)
	}
	return strings.Join(islice, "")
}
