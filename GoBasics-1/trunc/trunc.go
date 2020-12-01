package main 

import "fmt"

func main() {
	var getFloat = enterFloat()
	var updatedFloat = convertFloat(getFloat)
	
	printConvertedValue(updatedFloat)
}

func enterFloat() float32{
	
	fmt.Println("Enter Float Data Type:")
	var floatOn float32
	
	fmt.Scanln(&floatOn)

	return floatOn
}

func convertFloat(enteredFloat float32) int{
	var convertToInt int = int(enteredFloat)

	return convertToInt
}

func printConvertedValue(convertedInt int){
	fmt.Println(convertedInt)
}
