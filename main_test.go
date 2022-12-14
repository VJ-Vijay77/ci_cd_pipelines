package main

import "testing"

func TestMain(t *testing.T) {
 output := Greet()

 if output != "Hello test" {
	t.Errorf("\n\nWanted -'Hello test' \n got - %s\n\n",output)
 }

 output2 := Add(5,6)
 if output2 != 11 {
	t.Errorf("\n\nWanted -11 \ngot - %d\n\n",output2)
 }
 
 
 output3 := Substract(10,6)
 if output3 != 4 {
	t.Errorf("\n\nWanted -4 \ngot - %d\n\n",output3)
 }
 
 output4 := Tail("test")
 expected := "test-TailAdded"
 if output4 != expected {
	t.Errorf("\n\nWanted - %s \nGot - %s\n\n",expected,output4)
 }

 output5 := SomeFunc(10,4)
 expected5 := TstFunc(10,4)
 if output5 != expected5 {
	t.Errorf("\n\nWanted - %d \nGot - %d\n\n",expected5,output5)
 }

}

func TstFunc(a int,b int)int{
	k := a + b
	k = k + 2
 	q := k/2
	return q  
}
