package main

import "testing"

func TestMymarshal(t *testing.T) {
	group := []Person{
		{Name: "Samir Kape", ID: 50},
		{Name: "xxx yyy", ID: 60},
		{Name: "zzz yyy", ID: 70},
		{Name: "yyy xxx", ID: 80},
	}
	expected := []string{
		`{"name":"Samir Kape","id":50}`,
		`{"name":"xxx yyy","id":60}`,
		`{"name":"zzz yyy","id":70}`,
		`{"name":"yyy xxx","id":80}`,
	}
	for index, person := range group {
		t.Run(person.Name, func(t *testing.T) {
			got := Mymarshal(person)
			if expected[index] != got {
				t.Errorf("Failed for %s", person.Name)
			}
		})
	}
}
