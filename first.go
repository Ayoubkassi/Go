package main

import "fmt"
import "os"

func main(){
  /////////////////////////////// First Lesson
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


  ///////////////////////////////////// Second Lesson
  ////////////////// OPerator , Relational , Logical and Assignment


  var num1 int = 1
  var num2 int = 2
  total = num1+num2

  fmt.Printf("%d",num1+num2)   //%d -> digit -> number
  fmt.Printf("%d",total)

  var sub = num1-num2

  var divide = num1 / num2

  var multi = num1 * num2

  num1++
  num2--  //we have increment and decrement

  if(num1 == num2){
    fmt.Println("Equal")
  }else{
    fmt.Println("NO doesn\'t match")
  }
  //same for relational for > < >= <= !=

  //logical ops

  if(num1 && num2){
    fmt.Printf("Oook")
  }fmt.Printf("Niice")

  // also with ||


  //Assignment

  num1+=num2
  num1*=num2
  num1/=num2
  num1-=num2

  // modulo operatot %

  ////////////////////////////// Go Package & import

  //package -> to get better design for go language
  //the main package is neccessary in our go project
  // import for librarries
  //import ("fmt" "os")
  //all librarirs must be in small letters


  //Example for if else

  var num1 int = 10
  var num2 int = 20

  if(num1 == num2){
    fmt.Pritln("Equal")
  }else if (num1 > num2){
    fmt.Println("Grt")
  }else{
    fmt.Pritln("lt")
  }

  //Switch statemnet

  switch num2{
  case 1:
    fmt.Println("num1 = 1")
    break
  case 2:
    fmt.Println("num1 = 2")
    break
  case 3:
    fmt.Println("num1 = 2")
    break
  default :
  fmt.Println("No one mutch")
  break
  }


  //loops
  //for loop

  for a:=0;a<10;a++{
    //fmt.Printf("Num is : %d",a)
    fmt.Println("Num : ",a)
  }

  var limit = 10
  for limit < 30{
    limit+= limit
    fmt.Println("Num : ",limit)
  }

  //json_handling

  var my_nums = []int{1,2,3,4,5,6}
  var num1 = 0

  for a,num := range my_nums{
    num1 += num
    //a just for index we can replace it by _ for blank (no index)
  }


  fmt.Println("Values : ",num1)

  var my_data = map[string]string{
    "name":"Ayoub kassi",
    "age" : "21",
    "major" : "Computer Science"
  }

  for key,value := range my_data{
    fmt.Println("key : ",key)
    fmt.Println("Value : ",value)
  }



  //goto statemnet

  var num = 18;

  Loop:
  fmt.Println("Sorry , +18")


    if(num1 <= 17){
      goto Loop
    }else{
      fmt.Println("Yes , approved")
    }



}

//comments just like C language
/*and we can use ; at the end of each line
  comment (skiped by compiler)
*/
