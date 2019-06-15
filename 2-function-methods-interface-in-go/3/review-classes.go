package main

import (
   "fmt"
   "bufio"
   "strings"
   "os"
   )

type Animal struct{
   food string
   locomotion string
   noise string
}

func (a *Animal) Eat() {
   fmt.Println(a.food)
}
func (a *Animal) Move() {
   fmt.Println(a.locomotion)
}
func (a *Animal) Speak() {
   fmt.Println(a.noise)
}


func main(){
   // example 1
   cow := Animal{"grass", "walk", "moo" }
   bird := Animal{"worms", "fly", "peep"}
   snake := Animal{"mice", "slither", "hsss"}

   scanner := bufio.NewScanner(os.Stdin)
   var a Animal
   for{
      fmt.Print(">")
      scanner.Scan()
      requestArray := strings.Fields(strings.TrimSpace(scanner.Text()))
      // fmt.Println(requestArray)
      switch requestArray[0]{
      case "cow":
         a = cow
      case "bird":
         a = bird
      case "snake":
         a = snake
      default:
         fmt.Println("Invalid Animal")

      }

      if (Animal{}) != a{
         switch requestArray[1]{
         case "speak":
            a.Speak()
         case "move":
            a.Move()
         case "eat":
            a.Eat()
         default:
            fmt.Println("Invalid Request")
         }
      }
   }
}