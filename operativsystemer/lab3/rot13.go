package cipher

import (
	"io"
	//"os"
	//"fmt"
	//"strings"
)

/*
Task 3: Rot 13

This task is taken from http://tour.golang.org.

A common pattern is an io.Reader that wraps another io.Reader, modifying the
stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of
compressed data) and returns a *gzip.Reader that also implements io.Reader (a
stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
modifying the stream by applying the rot13 substitution cipher to all
alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing
its Read method.
*/

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	// i motsetning til script der jeg leser input fra bruker (ReadString)
	// vil Read gjengi antallet bokstaver pluss en feilmelding
	n, err = rot.r.Read(p) 
	// i stedet for i < len(p) kunne vi også brukt i < n (begge er jo 21)
	for i := 0; i < len(p); i++ { 
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >= 'a' && p[i] < 'n') {
			p[i] += 13	
		} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z') {	
			p[i] -= 13
		}
	}
	return
}
