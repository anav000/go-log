package cli

import (
	"fmt"

	"go-log/cmd"
)

func printVersion() {
	fmt.Printf("%s\n\n", cmd.Banner)
	fmt.Printf("repository: %s\n", cmd.Repo)
	fmt.Printf("buildDate: %s\n", cmd.BuildDate)
	fmt.Printf("version: %s\n", cmd.Version)
	fmt.Printf("go-version: %s\n", cmd.GoVersion)
}
