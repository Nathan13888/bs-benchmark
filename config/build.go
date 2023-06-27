package config

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

var (
	BuildVersion = "development"
	BuildTime    = "unknown"
	BuildUser    = "unknown"
	BuildGOOS    = "unknown"
	BuildARCH    = "unknown"
	GOOS         = runtime.GOOS
	GOARCH       = runtime.GOARCH
)

func PrintVersion(cCtx *cli.Context) {
	fmt.Printf("version=%s buildTime=%s buildUser=%s buildGOOS=%s buildARCH=%s\n",
		cCtx.App.Version, BuildTime, BuildUser, BuildGOOS, BuildARCH)
}
