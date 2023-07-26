package main

import "fmt"

func main() {
//  i for rows
// j for columns


	var n int = 5

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println()

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	

	fmt.Println()

	
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println()

		for i := 1; i <= n; i++ {

		for j := i; j <= n; j++ {
			fmt.Print("  ")
		}

		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println()

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("  ")
		}

		for j := i; j <= n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println()

			for i := 1; i <= n; i++ {

		for j := i; j <= n; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
