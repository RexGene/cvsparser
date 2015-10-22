package csvparser

import (
	"testing"
)

func TestParser(t *testing.T) {
	records, err := Parse("test.csv")
	if err != nil {
		t.Log(err)
		t.Fatal()
		return
	}

	for key, feilds := range records {
		println("key", key)
		for name, value := range feilds {
			println("name:", name, "value:", value)
		}
	}
}
