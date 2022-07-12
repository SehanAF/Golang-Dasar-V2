package golang-dasarv2

import "fmt"

func main() {
	var name string // ini adalah variable

	name = "Sehan Alfarisi" // ini adalah value.
	fmt.Println(name)

	name = "Gavi Alexa" // contoh
	fmt.Println(name)

	var friendName = "Givi" // contoh
	fmt.Println(friendName)

	var age = 45 // Ini adalah contoh menggunakan integer
	fmt.Println(age)

	country := "Jepang"  // ini adalah contoh penggunaan := sebagai pengganti kata var.
	fmt.Println(country) // menggunakan := agar lebih simpel dan enak dilihat.

	var (
		firstName = "Gaja"
		lastName  = "Torka"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
