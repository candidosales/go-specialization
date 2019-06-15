package main

import (
   "fmt"
   "bufio"
   "strings"
   "os"
   )
type Animal interface{
   Eat()
   Speak()
   Move()
}
type Cow struct{
   food string
   locomotion string
   noise string
}
type Bird struct{
   food string
   locomotion string
   noise string
}
type Snake struct{
   food string
   locomotion string
   noise string
}

func (c Cow) Eat() {
   fmt.Println(c.food)
}
func (c Cow) Move() {
   fmt.Println(c.locomotion)
}
func (c Cow) Speak() {
   fmt.Println(c.noise)
}


func (s Snake) Eat() {
   fmt.Println(s.food)
}
func (s Snake) Move() {
   fmt.Println(s.locomotion)
}
func (s Snake) Speak() {
   fmt.Println(s.noise)
}

func (b Bird) Eat() {
   fmt.Println(b.food)
}
func (b Bird) Move() {
   fmt.Println(b.locomotion)
}
func (b Bird) Speak() {
   fmt.Println(b.noise)
}
func main(){
   // example 1
   animals := map[string]Animal{}
   scanner := bufio.NewScanner(os.Stdin)
   for{
      fmt.Print("> ")
      scanner.Scan()
      requestArray := strings.Fields(strings.ToLower(strings.TrimSpace(scanner.Text())))
      // fmt.Println(requestArray)
      if 3 != len(requestArray){
         fmt.Println("Invalid Command!")
         continue
      }
      switch requestArray[0]{
         case "newanimal":
            CreateNewAnimal(requestArray[1:], animals)
         case "query":
            QueryAnimal(requestArray[1:], animals)
         default:
            fmt.Println("Invalid Command")
      }
   }
}

func CreateNewAnimal(requestSli []string, animals map[string]Animal){
   switch requestSli[1]{
   case "cow":
      animals[requestSli[0]] = Cow{"grass", "walk", "moo" }
      fmt.Println("Created It!")
   case "bird":
      animals[requestSli[0]] = Bird{"worms", "fly", "peep"}
      fmt.Println("Created It!")
   case "snake":
      animals[requestSli[0]] = Snake{"mice", "slither", "hsss"}
      fmt.Println("Created It!")
   default:
      fmt.Println("Invalid Animal")

   }
}

func QueryAnimal(requestSli []string, animals map[string]Animal){
   if a,exists := animals[requestSli[0]]; exists{
      switch requestSli[1]{
         case "speak":
            a.Speak()
         case "move":
            a.Move()
         case "eat":
            a.Eat()
         default:
            fmt.Println("Invalid Request")
         }
   }else{
      fmt.Println("Animal Not Found!")
   }
}
