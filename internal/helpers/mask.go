package helpers


func Mask(s string) string {
	if len(s) == 0 {
		return ""
	}
	return "••••••"
}