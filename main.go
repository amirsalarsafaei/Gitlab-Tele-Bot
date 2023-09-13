package main

import "github.com/amirsalarsafaei/Gitlab-Tele-Bot/cmd"

func main() {
	err := cmd.Execute()
	panic(err.Error())
}
