package model

import (
	"math/rand"
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
	Birthday    time.Time
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
func RandPerson() PersonalInfo {
	fake.SetLang("ru")
	pi := PersonalInfo{}
	pi.FirstName = fake.MaleFirstName()
	pi.MiddleName = fake.MalePatronymic()
	pi.LastName = fake.MaleLastName()
	pi.BDFormat = time.DateOnly
	pi.Birthday = pi.randBD()
	return pi
}

func RandEntry() RegistryEntry {
	re := RegistryEntry{}
	re.PersonData = RandPerson()
	re.DTFormat = time.DateTime
	re.DateTime = time.Now().AddDate(0, -1*rand.Intn(3), -1*rand.Intn(30))
	return re
}

func GenEntries() Entries {
	es := Entries{}
	for i := 0; i < 10; i++ {
		es = append(es, RandEntry())
	}
	return es
}
