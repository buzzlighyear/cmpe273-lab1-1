Go has a special statement called defer which schedules a function call to be run after the function completes. Consider the following example:

package main


import "fmt"


func first(){

    fmt.Println("first")

}


func second() {
 
   fmt.Println("second")

}

func main()
 {

    defer first()

    second()

    
}


This program prints 1st followed by 2nd.
Defer is often used when resources need to be freed in some way.



package main


import "fmt"


func main() {

    defer func() {
 
   
        str := recover()

        fmt.Println(str)

        }()
 
        panic("PANIC")
     
   }
        


We can handle a run-time panic with the built-in recover function
A panic generally indicates a programmer error (for example attempting to access an index of an array that's out of bounds, forgetting to initialize a map, etc.) or an exceptional condition that there's no easy way to recover from.



