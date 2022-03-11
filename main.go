package main

import "fmt"

func main(){
	var n,k int
	fmt.Scan(&k)
	for k>0{
		fmt.Scan(&n)
		if n == 1 {
			fmt.Println(1)
		} else if n == 2 {
			fmt.Println(2)
		} else if n % 2 == 0 {
			i:=1
			for n > 0 {
				fmt.Print(i)
				n = n-i
				if(i==2){
					i=1
				}else{
					i=2
				}
			}
			fmt.Println()
		} else {
			i:=2
			for n > 0 {
				fmt.Print(i)
				n = n-i
				if(i==2){
					i=1
				}else{
					i=2
				}
			}
			fmt.Println()
		}
		k--
	}
}

