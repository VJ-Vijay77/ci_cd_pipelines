package response

import(
	"testing"
)

func TestGreet(t *testing.T) {
	result := Greet()
	expected := "He,this is testing..."

	if result != expected {
		t.Errorf("want = %s \n got = %s",expected,result)
	}

}