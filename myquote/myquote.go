package myquote

import "rsc.io/quote"

func Quote() string {
    return quote.Glass() + "\n" + quote.Go() + "\n" + quote.Opt() + "\n" + quote.Hello()
}