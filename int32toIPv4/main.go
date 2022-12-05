package int32toipv4

func Int32ToIp(n uint32) string {
	// step0 handle exceptions or known cases
	if n == 0 {
		return "0.0.0.0"
	}
	step1 := toBinary(n)
	// step2 toDecimal(octets) // used in func toIPv4
	step3 := toIPv4(step1)
	return step3
}

func toBinary(x uint32) string {
	result := ""
	number := int(x)
	resultInt := 0

	for number > 0 { // get that binary
		resultInt = number % 2                       // tells you if you append a 1 or 0
		number /= 2                                  // binary has a base 2
		result = string(rune(resultInt+48)) + result // prepend the new one or zero
	}
	return result
}

func toDecimal(bin string) string {
	output := ""
	outNum := 0
	for i := 0; i < len(bin); i++ {
		bit := int(bin[i] - 48)
		num := bit
		for exp := 1; exp < len(bin)-i; exp++ { // num^(len(bin)-1)
			num *= 2
		}
		outNum += num // add each one together to get the decimal
	}
	if outNum == 0 {
		return "0"
	} // the for loop below breaks if outNum is 0
	for outNum > 0 { // put decimal back as a string
		output = string(rune((outNum%10)+48)) + output // each ones place gets prepended to keep the digits in proper order.
		outNum /= 10
	}
	return output
}

func toIPv4(bin string) (address string) { // returns the 4 parts of an IPv4 address
	i := 8 - (32 - len(bin)) // end index of first octet
	octet1 := toDecimal(bin[:i])
	octet2 := toDecimal(bin[i : i+8])
	octet3 := toDecimal(bin[i+8 : i+16])
	octet4 := toDecimal(bin[i+16:])
	address = octet1 + "." + octet2 + "." + octet3 + "." + octet4
	return address
}
