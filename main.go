package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	version = "2.0.1"
	haLg    = "CiAgICAgICAgICAgICAgICAgICAgICAnL3MvICAgICAgICAgICAgK3kvJyAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAtK2hISEgrICAgICAgICAgICAgb0hISGhvLSAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICc6c21ISEhISEgrICAgICAgICAgICAgb0hISEhISG1zLycgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgIC4raEhISEhISEhISEgrICAgICAgICAgICAgb0hISEhISEhISGggICAgICAgICAgICAgICAgICAgCiAgICAgICAgJzpzbUhISEhISEhISEhISEgrICAgICAgICAgICAgb0hISEhISEhISGggICAgLTonICAgICAgICAgICAgCiAgICAgLitoSEhISEhISEhISEhISEhISEgrICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hIaCsuICAgICAgICAgCiAnOnNkSEhISEhISEhISEhISEhISEhoKy0gICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhIZHM6JyAgICAgCidtSEhISEhISEhISEhISEhISGRvOicgICAgICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEhISEhIbXkvLiAgICAgICAnICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEhIaCstICAgICAgIC4veW1vICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgICAgJzpvZEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIC9oSEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhoc3Nzc3Nzc3Nzc3NzaEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhob29vb29vb29vb29vaEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhteS8gICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhoby0gICAgICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb2RzLycgICAgICAgOm9kSEhISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgJyAgICAgICAuK3ltSEhISEhISEhISEhISCcgICAgCidtSEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgICAgICcvc21ISEhISEhISEhISEhISEhIbScgICAgCiAgLW9kSEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgIC1vaEhISEhISEhISEhISEhISEhIZG86ICAgICAgCiAgICAgLi95bUgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISEhISEhISEhteSsuICAgICAgICAgCiAgICAgICAgIDotICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISEhISEhkczogICAgICAgICAgICAgCiAgICAgICAgICAgICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISEh5Ky4gICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICc6c2RISEhISEhvICAgICAgICAgICAgb0hISEhISGRzOicgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAuK2hISEhvICAgICAgICAgICAgb0hISGgrLiAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAnOnMrICAgICAgICAgICAgK3M6Jwo="
	haMd    = "CiAgICAgICAgICAgICAgICAgICAuICAgICAgICAuICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAnL3ltcyAgICAgICAgeW15LycgICAgICAgICAgICAgICAKICAgICAgICAgICAgLW9oSEhISHMgICAgICAgIHlISEhIaG8nICAgICAgICAgICAgCiAgICAgICAgJy9zbUhISEhISEhzICAgICAgICB5SEhISEhILiAgLScgICAgICAgIAogICAgIC0raEhISEhISEhISEhtKyAgICAgICAgeUhISEhISC4gIHNIaCstICAgICAKICAgK21ISEhISEhISEhIaCstICAgICAgICAgIHlISEhISEguICBzSEhISG0rICAgCiAgIHlISEhISEhIZHM6JyAgIC4tICAgICAgICB5SEhISEhILiAgc0hISEhIeSAgIAogICB5SEhISEhoLiAgICc6c2RIeSAgICAgICAgeUhISEhISC4gIHNISEhISHkgICAKICAgeUhISEhIcyAgLmhISEhISHkgICAgICAgIHlISEhISEguICBzSEhISEh5ICAgCiAgIHlISEhISHMgIC5ISEhISEh5ICAgICAgICB5SEhISEhILiAgc0hISEhIeSAgIAogICB5SEhISEhzICAuSEhISEhIbXl5eXl5eXl5bUhISEhISC4gIHNISEhISHkgICAKICAgeUhISEhIcyAgLkhISEhISEhISEhISEhISEhISEhISEguICBzSEhISEh5ICAgCiAgIHlISEhISHMgIC5ISEhISEhISEhISEhISEhISEhISEhILiAgc0hISEhIeSAgIAogICB5SEhISEhzICAuSEhISEhIbXl5eXl5eXl5bUhISEhISC4gIHNISEhISHkgICAKICAgeUhISEhIcyAgLkhISEhISHkgICAgICAgIHlISEhISEguICBzSEhISEh5ICAgCiAgIHlISEhISHMgIC5ISEhISEh5ICAgICAgICB5SEhISEhoJyAgc0hISEhIeSAgIAogICB5SEhISEhzICAuSEhISEhIeSAgICAgICAgeUhkbzonICAgLmhISEhISHkgICAKICAgeUhISEhIcyAgLkhISEhISHkgICAgICAgIC0uICAgJy9zbUhISEhISEh5ICAgCiAgICtkSEhISHMgIC5ISEhISEh5ICAgICAgICAgIC1vaEhISEhISEhISEhkKyAgIAogICAgIC4raEhzICAuSEhISEhIeSAgICAgICAgb21ISEhISEhISEhIaCsuICAgICAKICAgICAgICAnLiAgLkhISEhISHkgICAgICAgIHlISEhISEhIbXM6JyAgICAgICAgCiAgICAgICAgICAgICcraEhISEh5ICAgICAgICB5SEhISGgrLSAgICAgICAgICAgIAogICAgICAgICAgICAgICAnYmxzfCAgICAgICAgeUhkLycgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgIC4gICAgICAgIC4K"
	haSm    = "CiAgICAgICAgICAgJy8nICAgICcvJyAgICAgICAgICAgCiAgICAgICAgLStoSEgnICAgIC5ISGgrLiAgICAgICAgCiAgICAnOnNtSEhISEgnICAgIC5ISEhIKyAtOicgICAgCiAgL2hISEhISEhoby0gICAgIC5ISEhIKyArSEhoLyAgCiAgaEhISG1zOicgLSsnICAgIC5ISEhIKyArSEhIaCAgCiAgaEhISC8gLXltSEguICAgIC5ISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEgtLi4uLi1ISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEhISEhISEhISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEhISEhISEhISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEgtLi4uLi1ISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEguICAgIC5ISG1zLSArSEhIaCAgCiAgaEhISC8gK0hISEguICAgICcrLiAnL3NtSEhIaCAgCiAgL3lISC8gK0hISEguICAgICA6b2RISEhISEhoLyAgCiAgICAnOi0gK0hISEguICAgIC5ISEhISGRzOicgICAgCiAgICAgICAgLitoSEguICAgIC5ISGgrLiAgICAgICAgCiAgICAgICAgICAgJzonICAgICc6Jwo="
)

var (
	dazzleMode bool
	color      string
	size       string
	logo       string
	plainMode  bool // New variable for plain mode
)

type ColorCode string

const (
	Reset   ColorCode = "\033[0m"
	Red     ColorCode = "\033[31m"
	Green   ColorCode = "\033[32m"
	Yellow  ColorCode = "\033[33m"
	Blue    ColorCode = "\033[34m"
	Magenta ColorCode = "\033[35m"
	Cyan    ColorCode = "\033[36m"
	White   ColorCode = "\033[37m"
)

var colorMap = map[string]ColorCode{
	"red":     Red,
	"green":   Green,
	"yellow":  Yellow,
	"blue":    Blue,
	"magenta": Magenta,
	"cyan":    Cyan,
	"white":   White,
}

func init() {
	flag.BoolVar(&dazzleMode, "dazzle", false, "Engage dazzle mode")
	flag.StringVar(&color, "color", "blue", "Color (cyan, red, green, yellow, blue, magenta, white, mix, random, or plain)")
	flag.StringVar(&size, "size", "medium", "Size (small, medium, or large)")
	flag.BoolVar(&plainMode, "plain", false, "Display logo without color") // New flag for plain mode
	flag.Parse()

	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("hashii v%s\n", version)
		os.Exit(0)
	}
}

func cleanup() {
	fmt.Print("\033[?25h") // Show cursor
	fmt.Print("\033[H\033[2J")
	fmt.Printf("%sBye!%s\n", Blue, Reset)
}

func decodeBase64(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalf("Failed to decode base64 string: %v", err)
	}
	return string(decoded)
}

func getRandomColor() ColorCode {
	colors := []ColorCode{Cyan, Red, Green, Yellow, Blue, Magenta, White}
	return colors[rand.Intn(len(colors))]
}

func printLogo(color ColorCode) {
	scanner := bufio.NewScanner(strings.NewReader(logo))
	for scanner.Scan() {
		if plainMode {
			fmt.Println(scanner.Text())
		} else {
			fmt.Printf("%s%s%s\n", color, scanner.Text(), Reset)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading logo: %v", err)
	}
}

func dazzle() {
	fmt.Print("\033[?25l") // Hide cursor
	for {
		fmt.Print("\033[H\033[2J")
		printLogo(getRandomColor())
		time.Sleep(142 * time.Millisecond)
	}
}

func main() {
	// defer cleanup()

	rand.Seed(time.Now().UnixNano())

	switch size {
	case "small":
		logo = decodeBase64(haSm)
	case "large":
		logo = decodeBase64(haLg)
	default:
		logo = decodeBase64(haMd)
	}

	if dazzleMode {
		go func() {
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			<-c
			cleanup()
			os.Exit(0)
		}()
		dazzle()
	}

	var selectedColor ColorCode
	switch color {
	case "mix":
		selectedColor = getRandomColor()
	case "random":
		selectedColor = getRandomColor()
	case "plain":
		plainMode = true
		selectedColor = Reset // Set to Reset, though it won't be used
	default:
		if c, ok := colorMap[color]; ok {
			selectedColor = c
		} else {
			selectedColor = Blue
		}
	}

	printLogo(selectedColor)
}
