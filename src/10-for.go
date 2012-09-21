package main                    // `for` is Go's only looping construct. Below are
                                // two common forms.
import "fmt"                    
                                
func main() {                   
	i := 1                      // We initialize `i` with `1` and loop until it's 10.
	for i <= 10 {               
		fmt.Println(i)          
		i = i + 1               
	}                           
                                
	for j := 1; j <= 10; j++ {  // That type of loop is common. We can do it on one
		fmt.Println(j)          // line.
	}                           
}                               // There are other forms of `for`; we'll see them
                                // later.

