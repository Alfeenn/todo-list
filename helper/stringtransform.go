package helper

import "strings"

func DashString(name string) string {

	slugToLower := strings.ToLower(name)
	dash := strings.Replace(slugToLower, " ", "-", -1)
	return dash
}

func LowerAndDash(name string) []string {

	dash := strings.Split(name, ",")
	return dash
}
