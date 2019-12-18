package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	//s := []byte{'a','b','c'}
	s := []byte{'A','m','a','n','a','p','l','a','n','a','c','a','m','e','o','Z','e','n','a','B','i','r','d','M','o','c','h','a','P','r','o','w','e','l','a','r','a','v','e','U','g','a','n','d','a','W','a','i','t','a','l','o','b','o','l','a','A','r','g','o','G','o','t','o','K','o','s','e','r','I','h','a','b','U','d','a','l','l','a','r','e','v','o','c','a','t','i','o','n','E','b','a','r','t','a','M','u','s','c','a','t','e','y','e','s','R','e','h','m','a','c','e','s','s','i','o','n','U','d','e','l','l','a','E','-','b','o','a','t','O','A','S','a','m','i','r','a','g','e','I','P','B','M','a','c','a','r','e','s','s','E','t','a','m','F','C','A','a','m','i','c','a','O','j','a','i','L','e','b','o','w','a','Y','a','e','g','e','r','a','b','a','r','g','e','R','a','b','A','l','s','a','t','i','a','n','a','m','o','d','A','d','v','a','r','p','s','I','l','e','a','n','e','V','a','l','e','t','a','G','r','e','n','a','d','a','H','e','t','t','y','F','a','y','m','e','R','E','M','E','H','C','M','T','s','a','n','O','w','e','n','a','T','a','m','a','r','Y','o','m','p','u','r','I','s','a','N','i','l','L','o','r','r','i','n','R','i','k','s','d','a','g','M','o','n','a','R','o','n','n','O','C','o','n','n','e','r','K','i','r','k','a','n','o','k','a','y','N','a','f','l','L','i','r','a','R','o','b','i','R','a','m','e','F','I','F','A','a','n','e','e','d','R','o','d','i','M','u','h','a','r','r','a','m','O','b','e','r','a','c','o','m','a','c','a','r','r','i','H','w','a','n','g','T','a','o','s','S','a','l','a','d','o','O','l','f','e','C','a','m','a','g','K','d','a','r','a','h','d','k','f','J','e','m','i','n','a','N','a','d','l','e','r','E','h','u','d','R','u','t','a','n','a','b','a','s','t','e','r','E','b','n','a','e','d','e','g','i','a','g','a','l','s','I','r','a','T','e','p','p','e','r','a','m','i','n','i','m','a','g','o','w','d','U','l','d','a','O','g','a','w','a','T','S','g','t','C','a','l','l','i','d','a','R','o','d','l','E','w','a','r','t','s','e','r','a','p','h','s','A','i','n','N','e','w','g','a','t','e','V','a','d','e','n','n','a','v','e','t','t','e','s','S','a','b','a','h','S','w','a','t',','}
	reverseString(&s)
	elapsed := time.Since(start)
	fmt.Printf("time taken %s", elapsed)
	fmt.Print(s)
}

func reverseString(s *[]byte) {
	left, right := 0, len(*s)-1
	for i := 0; i < len(*s)/2; i++ {
		(*s)[left], (*s)[right] = (*s)[right], (*s)[left]
		left++
		right--

	}
}
