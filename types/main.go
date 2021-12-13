package main

import "fmt"


func main(){
a:=true
b:=false
fmt.Println("a",a,"b",b)

c:=a&&b
fmt.Println("c:",c)

d:=a||b
fmt.Println("d:",d)

var num int=80
num2:=85
fmt.Println("num is:",num," num2 is",num2)

fmt.Printf("type of num is %T",num)

num3:=5.55
fmt.Printf("type of num3 is %T \n",num3)

firstname:="uthman"
lastname:="rahimi"
fmt.Println("My name is",firstname,lastname)

i:=12
j:=12.5
sum:=i+int(j)
 fmt.Println("sum :",sum)
}