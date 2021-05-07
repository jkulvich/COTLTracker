// cmdline - Module to handle CLI interactions.
// Also, it contains CLI & app documentation

package cmdline

import (
	"fmt"
	"github.com/alecthomas/kong"
	"gopkg.in/gookit/color.v1"
	"os"
	"player/appinfo"
	"runtime"
)

// CLI - Parsed command line args & flags
type CLI struct {
	Doc  struct{} `cmd:"" aliases:"d,about,info,inf,help,h,version,v" help:"Show additional documentation and references" default:"1"`
	Play struct {
		Track string `short:"t" help:"Path to track file" required:"" type:"existingfile" placeholder:"path/to/track/file" xor:"track"`
		Test  bool   `short:"x" help:"Sound test to check compatibility with your device" xor:"track"`
		Delay int    `short:"d" help:"Minimal delay between screen taps" default:"80"`
		Start int    `short:"s" help:"First block position to start playing" default:"0"`
	} `cmd:"" help:"Start manual player module"`
}

// Parse - Parse command line args into struct
func Parse() (string, CLI) {
	cli := CLI{}
	ctx := kong.Parse(&cli)

	// Handle special commands
	switch ctx.Command() {
	case "doc":
		color.LightCyan.Printf("COTLTrack %s at %s\n", appinfo.Version, appinfo.Build)

		color.LightWhite.Print("Build extra: ")
		color.Printf("%s %s %s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)

		color.LightWhite.Print("Home Page: ")
		color.Printf("%s\n", appinfo.HomePage)

		color.LightWhite.Print("Tech Support: ")
		color.Printf("%s\n", appinfo.Support)

		// COTLTrack actual filename
		filename, _ := os.Stat(os.Args[0])

		// Main doc
		color.Yellow.Printf("\n%s is a command-line console tool ", filename.Name())
		color.Printf("now (I'll change it soon).\n")
		color.Printf("You ")
		color.Yellow.Printf("can't use it as a usual ")
		color.Printf("graphical app and you can't run it\n")
		color.Printf("by double clicking on the app icon.\n\n")

		color.Printf("You ")
		color.Yellow.Printf("need to use your terminal/console ")
		color.Printf("to interact with it.\n")
		color.Printf("I strongly recommend to read our doc before using.\n")
		color.Printf("Just ")
		color.Yellow.Printf("visit our home page ")
		color.Printf("which URL specified above.\n\n")

		color.Printf("Right now %s supports ", filename.Name())
		color.Yellow.Printf("only \"Play\" mode.\n")
		color.Printf("It means that you can play any song with same command line flags\n")
		color.Yellow.Printf("as in older versions ")
		color.Printf("but you need to use \"play\" command before.\n\n")

		color.LightWhite.Printf("Try for more detail doc:\n")
		color.Printf("\t%s --help play\n\n", filename.Name())

		color.LightWhite.Printf("Or use this example:\n")
		color.Printf("\t%s play --track tracks/way_back_home.txt\n\n", filename.Name())

		color.Printf("Check that you have ./tracks folder\n")
		color.Printf("at the same directory where %s located\n\n", filename.Name())

		color.LightCyan.Printf("Now, press [ENTER] to close this window")

		_, _ = fmt.Scanf("%s")
		os.Exit(0)
	}

	return ctx.Command(), cli
}
