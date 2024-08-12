package main

import "fmt"

func checkEvenOdd(){
	listOfInt:=[]int{1,2,3,4,5,6,7}
	for num:= range listOfInt{
		if num % 2 == 0{
			fmt.Printf("Number is even. num %v \n", num)
		}else{
			fmt.Printf("Number is odd. num %v \n", num)
		}
	}
}