package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func init() {
	//3.error to output file log. setoutput
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	fmt.Println("hi")
	/*
			_, err := os.Open("myfile.txt")
			if err != nil {
				//1. error to window
				fmt.Println("error happened", err)
				//2.error to log with date and time
				log.Println("err happened", err)

				//4. fatal error this will exit the program
				log.Fatalln(err)
				//5. panic the program.
				panic(err)

			}

		//6. custom errors
		ou, err := sq(0)
		if err != nil {
			fmt.Println(err)
		}
	    fmt.Println(ou)

	    //7. errors.New in new variable
	    //var ErrNorgateMath = errors.New("norgate math: square root of negative number")
	    //8.error fmt.errorf
	    //		ErrNorgateMath := fmt.Errorf("norgate math again: square root of negative number: %v", f)

	    //9. custom error
	*/
	_, err := Sqrt(-10.23)
	if err != nil {
		fmt.Println(err)
	}

}

type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, &NorgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	// implementation
	return 42, nil
}
func sq(number int) (int, error) {
	if number < 1 {
		return 0, errors.New("hey you entered wrong number")
	}
	return number * number, nil
}
