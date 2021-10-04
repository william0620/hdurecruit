package main

import (
	_ "recruiting_system/internal/app/community"
	"recruiting_system/internal/dep/httpEngine"
	_ "recruiting_system/internal/pkg/db"
)

func main() {
	httpEngine.Run()
}
