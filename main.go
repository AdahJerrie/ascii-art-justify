package main

import (
	"flag"
	"fmt"
	"syscall"
	"unsafe"
)

func GetTerminalWidth() int {
	var dimensions [4]uint16
	syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&dimensions)),
	)
	width := int(dimensions[1])
	if width == 0 {
		width = 80
	}
	return width
}

func main() {
	aligns := flag.String("align", "left", "align text")
	flag.Parse()

	args := flag.Args()

	align := *aligns

	terminalWidth := GetTerminalWidth()

	bannerfile := "standard"
	if len(args) == 2 {
		bannerfile = args[1]
	}
	banner := fmt.Sprintf("%s.txt", bannerfile)

	loaded, err := LoadBanner(banner)
	if err != nil {
		fmt.Println("error loading banner")
		return
	}

	input := args[0]

	built := BUildArt(input, align, terminalWidth, loaded)
	fmt.Println(built)
}
