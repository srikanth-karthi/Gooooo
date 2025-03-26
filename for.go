package main

import "fmt"

func main(){
	i:=1
	for i<=3{
		fmt.Println(i)
		i = i+1
	}
	for j:=7; j<=9; j++{	
		fmt.Println(j)
	}
	for i:=2; i<=4; i++{
		fmt.Println("loop")
	}
	for n:= range 6{
		fmt.Println(n)
	}
	for o:=9;o<=14;o++{
		if o%2==0{
			break
		}
		fmt.Println(o)
	}

}