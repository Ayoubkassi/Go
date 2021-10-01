package main

import "fmt"
import "os"

func main(){
  var name string
  name = "Ayoub Kassi"
  var slary int = 10
  var is_loading bool
  var grade = 100
  var num1,num2,num3,num4 int; //multi variable line
  var a,b,c byte
  var x1,x2,x3 float32 ; //we have float32 and float 64
  //inferance method
  number_3 := 30
  // :operator is a short hand for a variable without datatype or var keyword
  



  fmt.Println("Hello Kassi")
  fmt.Printf(" %T %T %T %T",name , is_loading , salary, grade)
  //T -> TYPE
  //V -> VALUE
  fmt.Printf(" %v %v %v %v",name , is_loading , salary, grade)
  //for string we can use %q -> quote

}

//comments just like C language
/*and we can use ; at the end of each line
  comment (skiped by compiler)
*/
