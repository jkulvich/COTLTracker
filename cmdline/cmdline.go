// cmdline - Module to handle CmdLine interactions.
// Also, it contains CmdLine & app documentation

package cmdline

import (
	"fmt"
	"github.com/alecthomas/kong"
	"gopkg.in/gookit/color.v1"
	"os"
	"player/appinfo"
	"runtime"
)

// CmdLine - Parsed command line args & flags
type CmdLine struct {
	Doc struct{} `cmd:"" aliases:"d,about,info,inf,help,h,version,v" help:"Show additional documentation and references" default:"1"`

	Play struct {
		Track   string `arg:"" help:"Path to track file or URL with or without file extension" required:""`
		Start   int    `short:"s" help:"First block position to start playing" default:"0"`
		Delay   int    `short:"d" help:"Delay between 'taps'" default:"40"`
		Tick    int    `short:"i" help:"Time of 'tick' in ms. -1 will be overridden by track comment or 200 if not specified" default:"-1"`
		Loader  string `short:"l" help:"Mod specific tracks loader for specific types of trackers' files" enum:"cotl" default:"cotl"`
		Mod     string `short:"m" help:"Mod specific tracker module to play tracks different ways" enum:"stdout,report,virtual" default:"virtual"`
	} `cmd:"" help:"Start track file playing"`
}

// Parse - Parse command line args into struct
func Parse() (string, CmdLine) {
	cli := CmdLine{}
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
