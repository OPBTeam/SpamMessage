package util

import (
	"regexp"
	"strings"

	"github.com/mgutz/ansi"
)

var MinecraftAsciiMap = make(map[string]string)

func InitColor() {
	MinecraftAsciiMap["§4"] = ansi.Red
	MinecraftAsciiMap["§c"] = ansi.LightRed
	MinecraftAsciiMap["§6"] = ansi.Yellow
	MinecraftAsciiMap["§e"] = ansi.LightYellow
	MinecraftAsciiMap["§2"] = ansi.Green
	MinecraftAsciiMap["§a"] = ansi.LightGreen
	MinecraftAsciiMap["§b"] = ansi.LightBlue
	MinecraftAsciiMap["§3"] = ansi.Blue
	MinecraftAsciiMap["§1"] = ansi.Blue
	MinecraftAsciiMap["§9"] = ansi.LightBlue
	MinecraftAsciiMap["§d"] = ansi.LightMagenta
	MinecraftAsciiMap["§5"] = ansi.Magenta
	MinecraftAsciiMap["§f"] = ansi.Reset
	MinecraftAsciiMap["§7"] = ansi.White
	MinecraftAsciiMap["§8"] = ansi.LightBlack
	MinecraftAsciiMap["§0"] = ansi.Black

	MinecraftAsciiMap["§k"] = ""
	MinecraftAsciiMap["§l"] = ""
	MinecraftAsciiMap["§m"] = ""
	MinecraftAsciiMap["§n"] = ""
	MinecraftAsciiMap["§o"] = ""
	MinecraftAsciiMap["§r"] = ""
}

func MinecraftToAscii(minecraft string) string {
	for key, value := range MinecraftAsciiMap {
		minecraft = strings.ReplaceAll(minecraft, key, value)
	}
	return minecraft
}

var re = regexp.MustCompile(`(?i)§[0-9A-GK-OR]`)

func StripColor(msg string) string {
	return re.ReplaceAllString(msg, "")
}
