package myquote

import (
	"rsc.io/quote"
	"fmt"
 )


func Quote() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())

}
