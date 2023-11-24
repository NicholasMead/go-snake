package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

var (
	exc   uint = 27
	arrow uint = exc | 91<<8
	up    uint = arrow | 65<<16
	down  uint = arrow | 66<<16
	right uint = arrow | 67<<16
	left  uint = arrow | 68<<16
)

func main() {
	stdin := os.Stdin

	prev, err := makeRaw(stdin)
	if err != nil {
		panic(err)
	} else {
		defer restore(stdin, prev)
	}

	w, h := getSize(os.Stdout)
	fmt.Printf("Size: %v, %v\n", w, h)

	for {
		buff := make([]byte, 4)
		i, _ := stdin.Read(buff)

		if i > 0 {
			val := uint(buff[0])
			for i, v := range buff[1:] {
				val |= uint(v) << (8 * (i + 1))
			}

			carrageReturn(os.Stdout)
			fmt.Printf("%032b", val)

			if val == up {
				fmt.Print(" ^")
			} else if val == down {
				fmt.Print(" v")
			} else if val == right {
				fmt.Print(" >")
			} else if val == left {
				fmt.Print(" <")
			} else if val < 127 {
				fmt.Printf(" %s", buff)
			} else {
				fmt.Printf("  ")
			}

			if buff[0] == 3 || buff[0] == 4 { //ctrl+c, ctrl+d
				return
			}
		} else {
			fmt.Println("Empty Buffer")
		}
	}
}

func makeRaw(console *os.File) (*uint32, error) {
	var h = windows.Handle(console.Fd())
	var st uint32
	if err := windows.GetConsoleMode(h, &st); err != nil {
		return nil, err
	}
	raw := makeRawState(st)

	if err := windows.SetConsoleMode(h, raw); err != nil {
		return nil, err
	}

	return &st, nil
}

func makeRawState(mode uint32) uint32 {
	//not these
	mode &^= windows.ENABLE_ECHO_INPUT
	mode &^= windows.ENABLE_PROCESSED_INPUT
	mode &^= windows.ENABLE_LINE_INPUT

	//and these
	mode |= windows.ENABLE_VIRTUAL_TERMINAL_INPUT

	return mode
}

func getSize(console *os.File) (w int, h int) {
	var info windows.ConsoleScreenBufferInfo
	err := windows.GetConsoleScreenBufferInfo(windows.Handle(console.Fd()), &info)
	if err != nil {
		panic(err)
	}

	w = int(info.Window.Right - info.Window.Left + 1)
	h = int(info.Window.Bottom - info.Window.Top + 1)
	return
}

func carrageReturn(console *os.File) {
	var (
		h    = windows.Handle(console.Fd())
		info windows.ConsoleScreenBufferInfo
	)

	if err := windows.GetConsoleScreenBufferInfo(h, &info); err != nil {
		panic(err)
	}

	if err := windows.SetConsoleCursorPosition(h, windows.Coord{
		X: 0,
		Y: info.CursorPosition.Y,
	}); err != nil {
		panic(err)
	}
}

func restore(console *os.File, state *uint32) error {
	var h = windows.Handle(console.Fd())
	if err := windows.SetConsoleMode(h, *state); err != nil {
		return err
	}
	return nil
}
