package main

/*
Complete the method/function so that it converts dash/underscore delimited words into camel casing. The first word within the output should be capitalized only if the original word was capitalized (known as Upper Camel Case, also often referred to as Pascal case).

Examples
ToCamelCase("the-stealth-warrior"); // returns "theStealthWarrior"

ToCamelCase("The_Stealth_Warrior"); // returns "TheStealthWarrior"
*/


func main(){
	ToCamelCase("the-Stealth-warrior")
}


func ToCamelCase(s string) string {
	var textResult string
	var nextWordMustBeUpeer bool
	for i, charVariable := range s {
		if i==0 {
			upper := IsNotUpper(int(charVariable))
			if upper == false {
				textResult = string(int(charVariable) - 32)
			}
			textResult = string(charVariable)
		}else{
			if charVariable == 45 || charVariable == 95{
				nextWordMustBeUpeer = true
			}else{
				if nextWordMustBeUpeer == true {
					if IsNotUpper(int(charVariable)) == false {
						textResult += string(int(charVariable) - 32)
					}else{
						textResult += string(charVariable)
					}					
					nextWordMustBeUpeer = false
				}else{
					textResult += string(charVariable)
				}
			}
		}			
	}
	return textResult
}

func IsNotUpper(charVariable int) bool {
	if  charVariable > 'z' || charVariable < 'a' {  
		return true  
	}
	return false
}