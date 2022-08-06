package main

import "testing"

func TestMain(t *testing.T) {
 output := Greet()

 if output != "Hello test" {
	t.Errorf("Wanted -'Hello test' \n got - %s",output)
 }
}
