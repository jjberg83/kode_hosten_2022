package cipher

import(
	"math/rand"
	"strings"
)

/*
Task 4: Scrambling text

In this task the objective is to transform a given text such that all letters of each word
are randomly shuffled except for the first and last letter.

For example, given the word "scramble", the result could be "srmacble" or "sbcamrle",
or any other permutation as long as the first and last letters stay the same.

An entire sentence scrambled like this should still be readable:
"it deosn't mttaer in waht oredr the ltteers in a wrod are,
the olny iprmoetnt tihng is taht the frist and lsat ltteer be at the rghit pclae"
See https://www.mrc-cbu.cam.ac.uk/people/matt.davis/cmabridge/ for more
information and examples.

Implementation:
The task is to implement the scramble function, which takes a text in the form of a string and a seed.
A seed is given so the output from your solution should match the test cases if it is correct.
The seed should be applied at the start of the function.
Remember that the implementation should keep any punctuation and spacing intact, and all numbers should be untouched.

Shuffling the letters and applying the seed can be done using the math/rand package (https://golang.org/pkg/math/rand/).
Use the Shuffle function to ensure you reach the same values as given in the tests (scramble_test.go).
*/


	// for å teste filen, skriv følgende i Terminalen:
	// go test -run TestScramble

func scramble(text string, seed int64) string {
	// gjør at output blir deterministisk
	rand.Seed(seed)
	// tokenize funksjonen ligger i filen scramble_split_func.go
	// input er en string, output er en []string (altså en slice med strings i)	
	minSlice := tokenize(text)
	
	for k,element := range minSlice {
		// hvis nåværende element ikke inneholder noen tall med dobbel 
		// quotes rundt seg, f.eks  "a4t5423",
		// og det i tillegg er større enn 3 bokstaver, gjøres følgende...
		// (første og siste bokstav beholdes, men alt i midten shuffles)
		if !strings.ContainsAny(element, "0123456789") && len(element) > 3 {
			// gjør hvert ord om til en slice
			elementSlice := strings.Split(element, "")
			// gir indeksen til siste element
			indeks := len(element)-1
			firstLetter := elementSlice[0]
			lastLetter := elementSlice[indeks]
			midten := elementSlice[1:indeks]	
			// funksjonen under ligger i math/rand pakken 
			// den bytter om på rekkefølgen på bokstavene i midten av ordet
			rand.Shuffle(len(midten), func(i, j int) {
				midten[i], midten[j] = midten[j], midten[i]
			})
			// vi oppretter en blank slice
			ord := []string{}
			// og fyller den med første bokstav, (muligens) endret midtparti
			// og siste bokstav til slutt
			ord = append(ord, firstLetter)
			ord = append(ord, midten...)
			ord = append(ord, lastLetter)
			// så omgjør vi det til en string...
			element = strings.Join(ord, "")
			// og sier at elementet som loopen begynte med blir lik denne
			minSlice[k] = element
		} 
	}
	// minSlice gjøres om til en string, siden funksjonen skal returnere en string
	endelig := strings.Join(minSlice,"")
	return endelig
}
