package main

import(
  "fmt"
  "time"
)

func sleeptime(n int){
  select {
  case <-time.After(time.Duration(n)*time.Second):
    fmt.Println("End of Sleep")
  }
}

func main(){
  fmt.Println("Hi I am going to sleep")
  sleeptime(6)
  fmt.Println("Hi i am awake")

  var input string
  fmt.Scanln(&input)
}
