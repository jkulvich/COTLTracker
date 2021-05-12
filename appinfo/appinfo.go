// appinfo - Module contains the app info like version, build & etc

package appinfo

// Global vars which injected by -ldflags.
// App should be build with something like: -ldflags="-X cmdline.AppVersion=0.1.0"
var (
	// Version - Build version with 3.digit.dotted format string or "dev"
	// First digit - major app version
	// Second digit - minor app version
	// Third digit - patch app version
	Version = "dev"
	// Build - Build time in format "DD.MM.YY HH:MM"
	Build = "eternal"
	// Name - Actual app name
	Name = "COTLTrack"
	// HomePage - App's home page or site
	HomePage = "it's a dev version, so I don't know who is your distributor"
	// Support - Tech support contact info
	Support = "¯\\_(-_-)_/¯"
)
