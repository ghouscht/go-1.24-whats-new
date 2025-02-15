package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type MyTime struct {
	time.Time
}

// IsZero can be used to customize the behavior of omitzero tag.
func (m MyTime) IsZero() bool {
	return m.Time.IsZero()
	//return false
}

type Person struct {
	Name      string `json:"name"`
	BirthDate MyTime `json:"birth_date,omitempty"`
	DeathDate MyTime `json:"death_date,omitzero"`
}

func main() {
	alice := Person{Name: "Alice"}
	b, err := json.Marshal(alice)
	fmt.Println(string(b), err)
}
