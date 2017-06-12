package conversions

type MyMapType map[string]string

func ValidConversions(){
	// Convert a map
	var map1 map[string]string = map[string]string {"a":"a"}

	var map2 MyMapType

	map2 = MyMapType(map1)

	println(map2)

	// Convert constants
	r := rune(3)
	println(r) // Conversion unchanged
	println(float32(0.49999999)) // Conversion with rounding


	// Numeric types converstions
	var myFloat32 float32 = 1.1
	toInt := int8(myFloat32)
	println(toInt)

}

func InvalidConversions(){
	//f := uint32(1.1)
	// Error: constant 1.1 truncated to integer
	//println(f)
}