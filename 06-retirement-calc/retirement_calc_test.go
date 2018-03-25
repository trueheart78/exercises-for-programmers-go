package main

import (
	"testing"
	"time"
)

var subject = person{age, retirement_age}
var age = 30
var retirement_age = 65
var current_year, _, _ = time.Now().Date()

func TestYearsToRetirement(t *testing.T) {
	if subject.yearsToRetirement() != retirement_age-age {
		t.Fatal("Fail: person.yearsToRetirement()")
	}
	t.Log("OK: person.yearsToRetirement()")
}

func TestYearOfRetirement(t *testing.T) {
	retirement_year := current_year + retirement_age - age
	if subject.yearOfRetirement() != retirement_year {
		t.Fatalf("Fail: person.yearOfRetirement() [%d, expected %d]", subject.yearOfRetirement(), retirement_year)
	}
	t.Log("OK: person.yearOfRetirement()")
}

func TestYear(t *testing.T) {
	if year() != current_year {
		t.Fatalf("Fail: year(), expected %d, got %d", current_year, year())
	}
	t.Log("OK: year()")
}
