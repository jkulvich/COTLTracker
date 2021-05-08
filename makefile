# Use next command to build a release version
# > make build v=0.1.0

# You need next linux tools to use this script:
#
# go - GoLang compiler // sudo snap install go --classic
# rsrc - Tool to make windows resources files (.syso) // go get github.com/akavel/rsrc
# imagemagick - Tool to manipulate images (convert etc) // sudo apt install imagemagick
# librsvg2-bin - Tool to convert SVG images // sudo apt install librsvg2-bin
# UPX - Tool to compress binaries // sudo apt install upx
#
# If you're under linux system execute next shell to install it all:
#
# sudo snap install go --classic
# go get github.com/akavel/rsrc
# sudo apt install imagemagick
# sudo apt install librsvg2-bin
# sudo apt install upx

# Builded app name
# The name will be extended with .exe when building for Windows OS
NAME = cotl

# App homepage or distribution page
HOME_PAGE = https://github.com/jkulvich/COTLTracker

# App tech support contacts
TECH_SUPPORT = Kulagin Yuri: TG: @jkulvich, EMail: jkulvichi@gmail.com

# The app build version
# This flag is required for release versions because a specific version-depended folder will be created inside ./dist/
# In other way, the VERSION value will be equal "dev"
VERSION = $(if $(strip $(v)),$(v),dev)

# Build time in format "DAY.MONTH.YEAR HOURS.MINUTES"
BUILD = $(shell date "+%d.%m.%y %H:%M")

#####################################
##########[ INTERNAL VARS ]##########
#####################################

_LDFLAGS = -ldflags="-s -w -X 'player/appinfo.Version=${VERSION}' -X 'player/appinfo.Build=${BUILD}' -X 'player/appinfo.HomePage=${HOME_PAGE}' -X 'player/appinfo.Support=${TECH_SUPPORT}'"

_DIST_FOLDER = dist
_VERSION_FOLDER = ${_DIST_FOLDER}/${VERSION}

_LINUX64_FOLDER = ${_VERSION_FOLDER}/linux_64
_LINUX32_FOLDER = ${_VERSION_FOLDER}/linux_32
_LINUXARM64_FOLDER = ${_VERSION_FOLDER}/linux_arm_64
_LINUXARM32_FOLDER = ${_VERSION_FOLDER}/linux_arm_32
_MAC64_FOLDER = ${_VERSION_FOLDER}/mac_64
_MACARM64_FOLDER = ${_VERSION_FOLDER}/mac_arm_64
_WIN64_FOLDER = ${_VERSION_FOLDER}/win_64
_WIN32_FOLDER = ${_VERSION_FOLDER}/win_32

_ASSETS_FOLDER = assets

_MANIFEST = ${_ASSETS_FOLDER}/manifest.xml

_LOGO_BIG_SVG = ${_ASSETS_FOLDER}/logo_big.svg
_LOGO_MIDDLE_SVG = ${_ASSETS_FOLDER}/logo_middle.svg
_LOGO_SMALL_SVG = ${_ASSETS_FOLDER}/logo_small.svg

_LOGO_BIG = ${_ASSETS_FOLDER}/logo_big.png
_LOGO_MIDDLE = ${_ASSETS_FOLDER}/logo_middle.png
_LOGO_SMALL = ${_ASSETS_FOLDER}/logo_small.png

_ICONS_FOLDER = ${_ASSETS_FOLDER}/icons
_ICON256 = ${_ICONS_FOLDER}/icon256.ico
_ICON72 = ${_ICONS_FOLDER}/icon72.ico
_ICON48 = ${_ICONS_FOLDER}/icon48.ico
_ICON32 = ${_ICONS_FOLDER}/icon32.ico
_ICON16 = ${_ICONS_FOLDER}/icon16.ico

####################################
##########[ PUBLIC CALLS ]##########
####################################

# Shows help info for this building script
help:
	@echo "help               - Show this help"
	@echo "build              - Make build for all platforms inside ./${_DIST_FOLDER} folder"
	@echo "clear              - Remove ./${_DIST_FOLDER}"
	@echo "build-linux-64     - Make build for linux x64"
	@echo "build-linux-32     - Make build for linux x32"
	@echo "build-linux-arm-64 - Make build for linux ARM x64"
	@echo "build-linux-arm-32 - Make build for linux ARM x32"
	@echo "build-mac-64       - Make build for MacOS x64"
	@echo "build-mac-arm-64   - Make build for MacOS ARM x64"
	@echo "build-win-64       - Make build for Windows x64"
	@echo "build-win-32       - Make build for Windows x32"
.PHONY: help

# Build for all platforms
build: build-linux-64 build-linux-32 build-linux-arm-64 build-linux-arm-32 build-mac-64 build-mac-arm-64 build-win-64 build-win-32
	@make _clean-up
	@echo BUILD WAS SUCCESSFULLY, CURRENT VERSION: ${VERSION}
.PHONY: build

# Clear ./${_DIST_FOLDER}
clear:
	@echo Removing ./${_DIST_FOLDER}
	rm -rf ./${_DIST_FOLDER}
.PHONY: clear

# Build the app bundle for linux x64
build-linux-64: _prepare
	@echo Start building for Linux x64: ${NAME} ${VERSION} ${BUILD}
	@GOOS=linux GOARCH=amd64 go build ${_LDFLAGS} -o ./${_LINUX64_FOLDER}/${NAME}
	@upx ./${_LINUX64_FOLDER}/${NAME} > /dev/null
.PHONY: build-linux-64

# Build the app bundle for linux x32
build-linux-32: _prepare
	@echo Start building for Linux x32: ${NAME} ${VERSION} ${BUILD}
	@GOOS=linux GOARCH=386 go build ${_LDFLAGS} -o ./${_LINUX32_FOLDER}/${NAME}
	@upx ./${_LINUX32_FOLDER}/${NAME} > /dev/null
.PHONY: build-linux-32

# Build the app bundle for linux ARM x64
build-linux-arm-64: _prepare
	@echo Start building for Linux ARM x64: ${NAME} ${VERSION} ${BUILD}
	@GOOS=linux GOARCH=arm64 go build ${_LDFLAGS} -o ./${_LINUXARM64_FOLDER}/${NAME}
	@upx ./${_LINUXARM64_FOLDER}/${NAME} > /dev/null
.PHONY: build-linux-arm-64

# Build the app bundle for linux ARM x32
build-linux-arm-32: _prepare
	@echo Start building for Linux ARM x32: ${NAME} ${VERSION} ${BUILD}
	@GOOS=linux GOARCH=arm go build ${_LDFLAGS} -o ./${_LINUXARM32_FOLDER}/${NAME}
	@upx ./${_LINUXARM32_FOLDER}/${NAME} > /dev/null
.PHONY: build-linux-arm-32

# Build the app bundle for MacOS x64
build-mac-64: _prepare
	@echo Start building for MacOS x64: ${NAME} ${VERSION} ${BUILD}
	@GOOS=darwin GOARCH=amd64 go build ${_LDFLAGS} -o ./${_MAC64_FOLDER}/${NAME}
	@upx ./${_MAC64_FOLDER}/${NAME} > /dev/null
.PHONY: build-mac-64

# Build the app bundle for MacOS ARM x64
# UPX unsopports this pair OS+ARCH yet: upx ./${_MACARM64_FOLDER}/${NAME}
build-mac-arm-64: _prepare
	@echo Start building for MacOS ARM x64: ${NAME} ${VERSION} ${BUILD}
	@GOOS=darwin GOARCH=arm64 go build ${_LDFLAGS} -o ./${_MACARM64_FOLDER}/${NAME}
.PHONY: build-mac-arm-64

# Build the app bundle for Windows x64
build-win-64: _prepare
	@echo Start building for Windows x64: ${NAME} ${VERSION} ${BUILD}
	@GOOS=windows GOARCH=amd64 go build ${_LDFLAGS} -o ./${_WIN64_FOLDER}/${NAME}.exe
	@upx ./${_WIN64_FOLDER}/${NAME}.exe > /dev/null
.PHONY: build-win-64

# Build the app bundle for Windows x32
build-win-32: _prepare
	@echo Start building for Windows x32: ${NAME} ${VERSION} ${BUILD}
	@GOOS=windows GOARCH=386 go build ${_LDFLAGS} -o ./${_WIN32_FOLDER}/${NAME}.exe
	@upx ./${_WIN32_FOLDER}/${NAME}.exe > /dev/null
.PHONY: build-win-32

###############################
##########[ HELPERS ]##########
###############################

# Prepares for build
_prepare: _generate-syso _make-dist-folder _clear-dev-folder _make-version-folder _make-platforms-folders _copy-tracks-to-platforms
.PHONY: _prepare

# Makes build folder if it doesn't exists
_make-dist-folder:
	@echo Making dist folder: ${_DIST_FOLDER}
	@mkdir -p ./${_DIST_FOLDER}
.PHONY: _make-dist-folder

# Makes version folder
# It falls if same version are exists
_make-version-folder:
	@echo Making version folder: ${_VERSION_FOLDER}
	@mkdir ./${_VERSION_FOLDER}
.PHONY: _make-version-folder

# Makes platforms folders
_make-platforms-folders:
	@echo Making platforms folders
	@mkdir ./${_LINUX64_FOLDER}
	@mkdir ./${_LINUX32_FOLDER}
	@mkdir ./${_LINUXARM32_FOLDER}
	@mkdir ./${_LINUXARM64_FOLDER}
	@mkdir ./${_MAC64_FOLDER}
	@mkdir ./${_MACARM64_FOLDER}
	@mkdir ./${_WIN64_FOLDER}
	@mkdir ./${_WIN32_FOLDER}
.PHONY: _make-platforms-folders

# Copies track folder into platforms folders
_copy-tracks-to-platforms:
	@echo Copying tracks to platforms
	@cp -r ./tracks ./${_LINUX64_FOLDER}
	@cp -r ./tracks ./${_LINUX32_FOLDER}
	@cp -r ./tracks ./${_LINUXARM32_FOLDER}
	@cp -r ./tracks ./${_LINUXARM64_FOLDER}
	@cp -r ./tracks ./${_MAC64_FOLDER}
	@cp -r ./tracks ./${_MACARM64_FOLDER}
	@cp -r ./tracks ./${_WIN64_FOLDER}
	@cp -r ./tracks ./${_WIN32_FOLDER}
.PHONY: _copy-tracks-to-platforms

# Clears dev folder inside ${_DIST_FOLDER}
_clear-dev-folder:
	@echo Clearing dev folder
	@rm -rf ./${_DIST_FOLDER}/dev
.PHONY: _clear-dev-folder

# Generate icons
_generate-icons:
	@echo Generating icons
	@rm -rf ./${_ICONS_FOLDER}
	@mkdir ./${_ICONS_FOLDER}

	@convert ./${_LOGO_BIG_SVG} ./${_LOGO_BIG}
	@convert ./${_LOGO_MIDDLE_SVG} ./${_LOGO_MIDDLE}
	@convert ./${_LOGO_SMALL_SVG} ./${_LOGO_SMALL}

	@convert -background none ./${_LOGO_BIG} -resize 256x256 ./${_ICON256}
	@convert -background none ./${_LOGO_BIG} -resize 72x72 ./${_ICON72}
	@convert -background none ./${_LOGO_BIG} -resize 48x48 ./${_ICON48}
	@convert -background none ./${_LOGO_MIDDLE} -resize 32x32 ./${_ICON32}
	@convert -background none ./${_LOGO_SMALL} -resize 16x16 ./${_ICON16}
.PHONY: _generate-icons

# Generate syso Windows rsrc file
_generate-syso: _generate-icons
	@echo Generating syso
	@rsrc -manifest ./${_MANIFEST} -ico ./${_ICON256},./${_ICON72},./${_ICON48},./${_ICON32},./${_ICON16} -arch 386 -o ./app_386.syso
	@rsrc -manifest ./${_MANIFEST} -ico ./${_ICON256},./${_ICON72},./${_ICON48},./${_ICON32},./${_ICON16} -arch amd64 -o ./app_amd64.syso
.PHONY: _generate-syso

# Clean up extra files created during compilation process
_clean-up:
	@echo Cleaning manifests
	@rm -rf ./*.syso
	@echo Cleaning generated icons
	@rm -rf ./${_ICONS_FOLDER}
	@rm -rf ./${_LOGO_BIG} ./${_LOGO_MIDDLE} ./${_LOGO_SMALL}
.PHONY: _clean-up