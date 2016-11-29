package wtfhash

import (
	"fmt"
	"testing"
)

func Test_JSONtoMD5(t *testing.T) {

	json := []byte(`
        {
            "c": false,
            "a": "123",
            "b": 456
        }
    `)
	x := JSONtoMD5(json, "testkey")

	if x != "be192dfdbcd37ce00d5b89818cc5bcac" {
		t.Error("Failed")
	}

	fmt.Println(x)
}

func Test_MaptoMD5(t *testing.T) {

	m := map[string]interface{}{
		"c": false,
		"a": "123",
		"b": 456,
	}

	x := MaptoMD5(m, "testkey")

	if x != "be192dfdbcd37ce00d5b89818cc5bcac" {
		t.Error("Failed")
	}

	fmt.Println(x)
}
