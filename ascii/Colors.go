package ascii

import (
	"fmt"
	"strings"
)

var (
	// Ansii Colors
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Black   = "\033[30m"

	// RGB Colors
	LightRed  = "\033[38;2;255;102;102m"
	DarkRed   = "\033[38;2;139;0;0m"
	Crimson   = "\033[38;2;220;20;60m"
	Firebrick = "\033[38;2;178;34;34m"
	// Shades of Green
	LightGreen  = "\033[38;2;144;238;144m"
	DarkGreen   = "\033[38;2;0;100;0m"
	LimeGreen   = "\033[38;2;50;205;50m"
	ForestGreen = "\033[38;2;34;139;34m"
	// Shades of Blue
	LightBlue = "\033[38;2;173;216;230m"
	DarkBlue  = "\033[38;2;0;0;139m"
	SkyBlue   = "\033[38;2;135;206;235m"
	RoyalBlue = "\033[38;2;65;105;225m"
	// Shades of Yellow
	LightYellow   = "\033[38;2;255;255;224m"
	DarkGoldenrod = "\033[38;2;184;134;11m"
	Gold          = "\033[38;2;255;215;0m"
	Khaki         = "\033[38;2;240;230;140m"
	// Shades of Purple
	Lavender    = "\033[38;2;230;230;250m"
	DarkMagenta = "\033[38;2;139;0;139m"
	Orchid      = "\033[38;2;218;112;214m"
	Plum        = "\033[38;2;221;160;221m"
	//	Shades of Gray
	LightGray = "\033[38;2;211;211;211m"
	DarkGray  = "\033[38;2;169;169;169m"
	DimGray   = "\033[38;2;105;105;105m"
	SlateGray = "\033[38;2;112;128;144m"
	// Shades of Orange
	LightOrange = "\033[38;2;255;200;100m"
	DarkOrange  = "\033[38;2;255;140;0m"
	Coral       = "\033[38;2;255;127;80m"
	OrangeRed   = "\033[38;2;255;69;0m"
	DarkSalmon  = "\033[38;2;233;150;122m"
	Tomato      = "\033[38;2;255;99;71m"
	Orange      = "\033[38;2;255;165;0m"
	// Shades of Pink
	LightPink     = "\033[38;2;255;182;193m"
	HotPink       = "\033[38;2;255;105;180m"
	DeepPink      = "\033[38;2;255;20;147m"
	PaleVioletRed = "\033[38;2;219;112;147m"
	LightCoral    = "\033[38;2;240;128;128m"
	Pink          = "\033[38;2;255;192;203m"
	LightSalmon   = "\033[38;2;255;160;122m"
	MistyRose     = "\033[38;2;255;228;225m"
	// Shades of Brown
	SandyBrown  = "\033[38;2;244;164;96m"
	Chocolate   = "\033[38;2;210;105;30m"
	SaddleBrown = "\033[38;2;139;69;19m"
	Sienna      = "\033[38;2;160;82;45m"
	// Shades of Cyan
	LightCyan = "\033[38;2;224;255;255m"
	DarkCyan  = "\033[38;2;0;139;139m"
	Aqua      = "\033[38;2;0;255;255m"
	// Shades of Magenta
	Fuchsia      = "\033[38;2;255;0;255m"
	Violet       = "\033[38;2;238;130;238m"
	MediumOrchid = "\033[38;2;186;85;211m"
	// Additional Colors
	Teal      = "\033[38;2;0;128;128m"
	Turquoise = "\033[38;2;64;224;208m"
	PeachPuff = "\033[38;2;255;218;185m"
	Salmon    = "\033[38;2;250;128;114m"
)

func Color() string {
	_, _, _, color, _, _, _, _, _ = ArgsManagement()
	if strings.EqualFold(color, "Reset") {
		color = Reset
	} else if strings.EqualFold(color, "Red") {
		color = Red
	} else if strings.EqualFold(color, "Green") {
		color = Green
	} else if strings.EqualFold(color, "Yellow") {
		color = Yellow

	} else if strings.EqualFold(color, "Blue") {
		color = Blue

	} else if strings.EqualFold(color, "Magenta") {
		color = Magenta
	} else if strings.EqualFold(color, "Cyan") {
		color = Cyan
	} else if strings.EqualFold(color, "White") {
		color = White
	} else if strings.EqualFold(color, "Black") {
		color = Black
	} else if strings.EqualFold(color, "LightRed") {
		color = LightRed
	}

	// else if strings.EqualFold(color, "DarkRed") {
	// } else if strings.EqualFold(color, "Crimson") {
	// } else if strings.EqualFold(color, "Firebrick") {
	// } else if strings.EqualFold(color, "LightGreen") {
	// } else if strings.EqualFold(color, "DarkGreen") {
	// } else if strings.EqualFold(color, "LimeGreen") {
	// } else if strings.EqualFold(color, "ForestGreen") {
	// } else if strings.EqualFold(color, "LightBlue") {
	// } else if strings.EqualFold(color, "DarkBlue") {
	// } else if strings.EqualFold(color, "SkyBlue") {
	// } else if strings.EqualFold(color, "RoyalBlue") {
	// } else if strings.EqualFold(color, "LightYellow") {
	// }
	fmt.Println("color in colorfunc is :", color)
	return color
}
