package main

func ConvertToRoman(arabic int) string {
	switch arabic {
	case 1:
		return "I"
	case 2:
		return "II"
	}
	return ""
}
