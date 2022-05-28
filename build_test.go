package urlBuilder

import (
	"fmt"
	"testing"
)


func TestBuilder_SetBase(t *testing.T) {
	base := "http://localhost:8080"
	expected := "http://localhost:8080"
	actual := URLBuilder().SetBase(base).ToString()
	if actual == expected {
		fmt.Println("SetBase succeed")
	} else {
		fmt.Println("SetBase failed")
	}

}

func TestBuilder_SetPath(t *testing.T) {
	base := "http://localhost:8080"
	path := "/test"
	expected := "http://localhost:8080/test"
	actual := URLBuilder().SetBase(base).SetPath(path).ToString()
	if actual == expected {
		fmt.Println("Setpath succeed")
	} else {
		fmt.Println("SetPath failed")
	}
}

func TestBuilder_SetParameter(t *testing.T) {
	base := "http://localhost:8080"
	path := "/test"
	expected := "http://localhost:8080/test?a=b&c=d&e=f"
	actual := URLBuilder().SetBase(base).SetPath(path).SetParameter("a", "b", "c", "d", "e", "f").ToString()
	if actual == expected {
		fmt.Println("SetParameter succeed")
	} else {
		fmt.Println("SetParameter failed")
	}

	//special characters
	expected_special := "http://localhost:8080/test?a=%23b%26&c=d&e=f"
	actual_special := URLBuilder().SetBase(base).SetPath(path).SetParameter("a", "#b&", "c", "d", "e", "f").ToString()
	if actual_special == expected_special {
		fmt.Println("SetSpecialParameter succeed")
	} else {
		fmt.Println("SetSpecialParameter failed")
	}
}

func BenchmarkBuilder_SetParameter(b *testing.B) {
	base := "http://localhost:8080"
	path := "/test"
	for n := 0; n < b.N; n++ {
		URLBuilder().SetBase(base).SetPath(path).SetParameter("a", "b", "c", "d", "e", "f").ToString()
	}
}


