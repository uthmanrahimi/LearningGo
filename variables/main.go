package main
import "fmt"
import "math"

func main(){
var age int
fmt.Println("My age is ",age)
age=29
fmt.Println("My new age is",age)

var index int =1
fmt.Println("Current index is",index)

var counter=1 // no need to define its type
fmt.Println("counter is",counter)

// mutiple variables in single statement

var width,height =450,300
fmt.Println("width is",width,"and height is ",height)

var (
name ="uthman"
userage =29
userheight int
)
fmt.Println("My name is",name,"I am",userage,"and my height is",userheight)

count:=10
fmt.Println("Count=",count)

x,y:=10.0,15.0
c:=math.Min(x,y)
fmt.Println("Min value between x,y is",c)
}