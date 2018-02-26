package main

import "fmt"

func main() {
	fmt.Println(weatherCelsius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelsius(34, "Tak sunny"))
	fmt.Println(weatherCelsius(17, "Phuket rainy"))
	fmt.Println(weatherCelsius(9, "Chiang-mai cold"))

}

func weatherCelsius(temp int, desc string) string {
	if temp == 25 {
		return " _  _ \n _||_ \n|_  _| c\n" + desc
		/* _  _
		 _||_
		|_  _| c
		*/
	} else if temp == 34 {
		return " _    \n _||_|\n _|  | c\n" + desc
		/* _
		   _||_|
		   _|  | c
		*/
	} else if temp == 17 {
		return "    _ \n  |  |\n  |  | c\n" + desc
		/*    _
		|  |
		|  | c
		*/
	} else if temp == 9 {
		return " _ \n|_|\n _| c\n" + desc
		/* _
		|_|
		 _| c
		*/
	} else {
		return "none"
	}
}

/*
if num==0 {
	fmt.println(" _ ")
	fmt.println("| |")
	fmt.println("|_|")
} else if num==1 {
	fmt.println("   ")
	fmt.println("  |")
	fmt.println("  |")
} else if num==2 {
	fmt.println(" _ ")
	fmt.println(" _|")
	fmt.println("|_ ")
} else if num==3 {
	fmt.println(" _ ")
	fmt.println(" _|")
	fmt.println(" _|")
} else if num==4 {
	fmt.println("   ")
	fmt.println("|_|")
	fmt.println("  |")
} else if num==5 {
	fmt.println(" _ ")
	fmt.println("|_ ")
	fmt.println(" _|")
} else if num==6 {
	fmt.println(" _ ")
	fmt.println("|_ ")
	fmt.println("|_|")
} else if num==7 {
	fmt.println(" _ ")
	fmt.println("  |")
	fmt.println("  |")
} else if num==8 {
	fmt.println(" _ ")
	fmt.println("|_|")
	fmt.println("|_|")
} else {
	fmt.println(" _ ")
	fmt.println("|_|")
	fmt.println(" _|")
}
*/
