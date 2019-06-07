package main
import "fmt"
import "strings"
import "strconv"


func main() {
  
  var x[3] int
  for  i := 0; i<=2 ; i++ {
   var v int
   fmt.Println("Ingrese un numero");
   fmt.Scan(&v);
   
   x[i]= v
   if strings.EqualFold(strconv.Itoa(v), "0") {
     fmt.Println("error, input not is integer valid")
     return;
   }
  }
  for count :=0 ;count < len(x) ; count ++  {
  fmt.Println(x[count]);
  }
}