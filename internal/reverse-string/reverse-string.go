package reverse_string

// Add your solution here

// Function to reverse string
func reverseString(s string) string {

	// Convert string to rune array
	r := []rune(s)

	// Take length of rune array with len
	length := len(r)

	// Create reversed string
	rev := ""

	i := length
	for i >= 1 {

		rev = rev + string(r[i-1])
		i = i - 1
	}

	//Return reversed string
	return rev

}





