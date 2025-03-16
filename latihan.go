// Terurut
package main
import "fmt"
func terurut(a, b *int){
    for *a!= *b{
        fmt.Print(*b, *a)
        fmt.Scan(&*a, &*b)
    }
}

func main(){
    var n, m int
    fmt.Scan(&n, &m)
    terurut(&n, &m)
}

