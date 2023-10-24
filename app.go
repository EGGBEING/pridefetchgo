//TODO: Better name. More flags. More versatile color system (allow hex code. should be easy to parse). Switch existing palettes to getAsciiColorCode functions. Add more complex flags (intersex, demi, progress, etc.)
//TODO: Rewrite the argument parser to actually make some sense. Add package count, uptime, and GPU to system output. Add some formatting to the output strings to make them look better.
//NOTE: For flags without enough lines to display everything, either solely use the large variant, or cut off package count and uptime.

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// UNIX terminal color escape codes
// they go like this "\u001B[48;2;R;G;Bm"

// Trans
const transbluergb string = "\u001B[48;2;95;215;255m"
const transbluergbfg string = "\u001B[38;2;95;215;255m"
const transpinkrgb string = "\u001b[48;2;255;135;175m"
const transwhitrgb string = "\u001b[48;2;255;255;255m"

// Non-binary
const nbyellow string = "\u001b[48;2;255;255;0m"
const nbpurple string = "\u001b[48;2;135;95;215m"
const nbblack string = "\u001b[48;2;58;58;58m"

// Bi
const bipink string = "\u001b[48;2;255;0;135m"
const bipurple string = "\u001b[48;2;135;95;175m"
const biblue string = "\u001b[48;2;0;95;175m"

// Deutschland (inside joke)
var (
	schwarz string = getAsciiColorCode(1, 1, 1)
	rot     string = getAsciiColorCode(222, 0, 0)
	gelb    string = getAsciiColorCode(253, 205, 1)
)

// Gay
var (
	gayred    string = getAsciiColorCode(228, 3, 3)
	gayorange string = getAsciiColorCode(255, 140, 0)
	gayyellow string = getAsciiColorCode(255, 237, 0)
	gaygreen  string = getAsciiColorCode(0, 128, 38)
	gayblue   string = getAsciiColorCode(0, 77, 255)
	gaypurple string = getAsciiColorCode(117, 7, 135)
)

// Ace
var (
	aceblack  string = getAsciiColorCode(0, 0, 0)
	acegray   string = getAsciiColorCode(115, 121, 123)
	acewhite  string = getAsciiColorCode(255, 255, 255)
	acepurple string = getAsciiColorCode(95, 0, 255)
)

// Pan
var (
	panpink   string = getAsciiColorCode(255, 34, 140)
	panyellow string = getAsciiColorCode(255, 216, 0)
	panblue   string = getAsciiColorCode(33, 177, 255)
)

// Genderfluid
var (
	gfpink   string = getAsciiColorCode(255, 115, 161)
	gfpurple string = getAsciiColorCode(189, 23, 213)
	gfblue   string = getAsciiColorCode(50, 59, 187)
	gfwhite  string = getAsciiColorCode(255, 255, 255)
	gfblack  string = getAsciiColorCode(0, 0, 0)
)

// Genderqueer
var (
	gqlavender string = getAsciiColorCode(184, 152, 223)
	gqwhite    string = getAsciiColorCode(254, 254, 254)
	gqgreen    string = getAsciiColorCode(105, 140, 56)
)

// Aro
var (
	arogreen1 string = getAsciiColorCode(60, 164, 65)
	arogreen2 string = getAsciiColorCode(169, 210, 120)
	arowhite  string = getAsciiColorCode(255, 255, 255)
	arogray   string = getAsciiColorCode(168, 168, 168)
	aroblack  string = getAsciiColorCode(0, 0, 0)
)

// Agender
var (
	agenderblack string = getAsciiColorCode(0, 0, 0)
	agendergray  string = getAsciiColorCode(187, 197, 198)
	agendergreen string = getAsciiColorCode(182, 245, 130)
	agenderwhite string = getAsciiColorCode(255, 255, 255)
)

// Color Palettes
var (
	deutschland = []string{schwarz, schwarz, rot, rot, gelb, gelb}
	trans       = []string{transbluergb, transpinkrgb, transwhitrgb, transpinkrgb, transbluergb}
	enby        = []string{nbyellow, nbyellow, transwhitrgb, transwhitrgb, nbpurple, nbpurple, nbblack, nbblack}
	bi          = []string{bipink, bipink, bipurple, biblue, biblue}
	gay         = []string{gayred, gayorange, gayyellow, gaygreen, gayblue, gaypurple}
	aro         = []string{arogreen1, arogreen2, arowhite, arogray, aroblack}
	ace         = []string{aceblack, acegray, acewhite, acepurple}
	pan         = []string{panpink, panpink, panyellow, panyellow, panblue, panblue}
	genderfluid = []string{gfpink, gfwhite, gfpurple, gfblack, gfblue}
	genderqueer = []string{gqlavender, gqlavender, gqwhite, gqwhite, gqgreen, gqgreen}
	agender     = []string{agenderblack, agendergray, agenderwhite, agendergreen, agenderwhite, agendergray, agenderblack}
)

const reset string = "\u001B[0m"

func main() {

	var args []string = nil

	if os.Args[1:] != nil {
		args = os.Args[1:]
	}

	flag := ""
	noCpu := ""
	printFlavortext := 0
	flavortext := ""

	if len(args) > 0 {
		if args[0] != "" {
			flag = args[0]
			if len(args) > 1 {
				noCpu = args[1]
				if len(args) > 2 {
					flavortext = args[2]
					printFlavortext = 1
				}
			}
		}
	}

	cpuStatus := 1

	if noCpu == "1" {
		cpuStatus = 0
	}

	// TODO: Turn this into a switch statement (done. leaving this because it's kinda funny to me)
	switch flag {
	case "--help":
		showHelp()
	case "--flags":
		listFlags()
	case "deutschland":
		drawFlag(deutschland, reset, cpuStatus, flavortext, printFlavortext)
	case "trans":
		drawFlag(trans, reset, cpuStatus, flavortext, printFlavortext)
	case "bi":
		drawFlag(bi, reset, cpuStatus, flavortext, printFlavortext)
	case "enby":
		drawFlag(enby, reset, cpuStatus, flavortext, printFlavortext)
	case "gay":
		drawFlag(gay, reset, cpuStatus, flavortext, printFlavortext)
	case "ace":
		drawFlag(ace, reset, cpuStatus, flavortext, printFlavortext)
	case "aro":
		drawFlag(aro, reset, cpuStatus, flavortext, printFlavortext)
	case "pan":
		drawFlag(pan, reset, cpuStatus, flavortext, printFlavortext)
	case "genderfluid":
		drawFlag(genderfluid, reset, cpuStatus, flavortext, printFlavortext)
	case "genderqueer":
		drawFlag(genderqueer, reset, cpuStatus, flavortext, printFlavortext)
	case "agender":
		drawFlag(agender, reset, cpuStatus, flavortext, printFlavortext)
	default:
		fmt.Println("Unrecognized argument, displaying help text:")
		showHelp()
	}
}

func listFlags() {
	fmt.Println("FLAGS")
	fmt.Print("\n")
	fmt.Println("gay")
	drawOnlyFlag(gay)
	fmt.Println("bi")
	drawOnlyFlag(bi)
	fmt.Println("trans")
	drawOnlyFlag(trans)
	fmt.Println("enby")
	drawOnlyFlag(enby)
	fmt.Println("aro")
	drawOnlyFlag(aro)
	fmt.Println("ace")
	drawOnlyFlag(ace)
	fmt.Println("pan")
	drawOnlyFlag(pan)
	fmt.Println("genderfluid")
	drawOnlyFlag(genderfluid)
	fmt.Println("genderqueer")
	drawOnlyFlag(genderqueer)
	fmt.Println("agender")
	drawOnlyFlag(agender)
}

func showHelp() {
	fmt.Println("argument 1: the flag to display (use the --flags argument for a list)")
	fmt.Println("argument 2: whether or not to display system information (1 for don't print, any other value for print)")
	fmt.Println("argument 3: flavortext that gets printed next to the hostname. will not appear if set to not show system info")
}

func drawFlag(colors []string, reset string, cpu int, extratext string, printft int) {
	flagLine := ""
	for i := 0; i < len(colors)*4; i++ {
		flagLine += " "
	}
	for i := 0; i < len(colors); i++ {
		fmt.Print(" " + colors[i] + flagLine + reset)
		if i == 0 && cpu == 1 {
			hostname, err := os.Hostname()
			if err != nil {
				fmt.Println(err)
			}
			mainString := getLinuxUser() + "@" + hostname
			if printft == 1 {
				mainString += " " + extratext
			}
			fmt.Print(" " + transbluergbfg + mainString + reset)
		}
		if i == 1 && cpu == 1 {
			fmt.Print(" - " + getLinuxDistro())
		}
		if i == 2 && cpu == 1 {
			fmt.Print(" - " + getCpuModel())
		}
		if i == 3 && cpu == 1 {
			used, total, avail := getRamInfo()
			fmt.Print(" - " + used + " / " + total + " (" + avail + " available)")
		}
		fmt.Print("\n")
	}
}

func drawOnlyFlag(colors []string) {
	flagLine := ""
	for i := 0; i < len(colors)*4; i++ {
		flagLine += " "
	}
	for i := 0; i < len(colors); i++ {
		fmt.Print(" " + colors[i] + flagLine + reset)
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func getCpuModel() string {
	cpuFile, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println(err)
	}
	defer cpuFile.Close()

	scanner := bufio.NewScanner(cpuFile)
	hasPrinted := false
	stringCpu := ""

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "model name") && !hasPrinted {
			// fmt.Println(scanner.Text())
			stringCpu = strings.ReplaceAll(scanner.Text(), "	", "")
			stringCpu = strings.ReplaceAll(stringCpu, "model name: ", "")
			// fmt.Println(stringCpu)
			hasPrinted = true
		}
	}
	return stringCpu
}

func getLinuxDistro() string {
	distroFile, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println(err)
	}
	defer distroFile.Close()

	scanner := bufio.NewScanner(distroFile)
	hasPrinted := false
	stringDistro := ""

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "PRETTY_NAME=") && !hasPrinted {
			stringDistro = strings.ReplaceAll(scanner.Text(), "PRETTY_NAME=", "")
			stringDistro = strings.ReplaceAll(stringDistro, "\"", "")
			hasPrinted = true
		}
	}

	return stringDistro
}

func getLinuxUser() string {
	app := "id"
	arg0 := "-un"
	cmd := exec.Command(app, arg0)
	ausgabe, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	ausgabe2 := string(ausgabe)
	if strings.Contains(ausgabe2, "\n") {
		ausgabe2 = strings.ReplaceAll(ausgabe2, "\n", "")
	}
	return ausgabe2

}

func getAsciiColorCode(r int, g int, b int) string {
	// they go like this "\u001B[48;2;R;G;Bm"
	colorCode := "\u001B[48;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
	return colorCode
}

func getRamInfo() (string, string, string) {
	ramFile, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Println(err)
	}
	defer ramFile.Close()

	scanner := bufio.NewScanner(ramFile)
	haveTotal := false
	haveAvailable := false

	totalRamString := ""
	availableRamString := ""

	for scanner.Scan() {
		if haveAvailable && haveTotal {
			break
		}
		if strings.Contains(scanner.Text(), "MemTotal:") && !haveTotal {
			totalRamString = strings.ReplaceAll(scanner.Text(), "MemTotal:", "")
			totalRamString = strings.ReplaceAll(totalRamString, "kB", "")
			totalRamString = strings.Trim(totalRamString, " ")
			totalRamString = strings.Trim(totalRamString, "\n")
			haveTotal = true
		}
		if strings.Contains(scanner.Text(), "MemAvailable:") && !haveAvailable {
			availableRamString = strings.ReplaceAll(scanner.Text(), "MemAvailable:", "")
			availableRamString = strings.ReplaceAll(availableRamString, "kB", "")
			availableRamString = strings.Trim(availableRamString, " ")
			availableRamString = strings.Trim(availableRamString, "\n")
			haveAvailable = true
		}
	}
	availRamInt, err := strconv.Atoi(availableRamString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	totalRamInt, err := strconv.Atoi(totalRamString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gbUsed := strconv.Itoa((totalRamInt-availRamInt)/1000000) + "GB"
	gbTotal := strconv.Itoa(totalRamInt/1000000) + "GB"
	gbAvail := strconv.Itoa(availRamInt/1000000) + "GB"

	return gbUsed, gbTotal, gbAvail
}
