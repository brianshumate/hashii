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

	"github.com/mgutz/ansi"
)

const (
	defaultColor = "blue"
	logoLarge    = "CiAgICAgICAgICAgICAgICAgICAgICAnL3MvICAgICAgICAgICAgK3kvJyAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAtK2hISEgrICAgICAgICAgICAgb0hISGhvLSAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICc6c21ISEhISEgrICAgICAgICAgICAgb0hISEhISG1zLycgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgIC4raEhISEhISEhISEgrICAgICAgICAgICAgb0hISEhISEhISGggICAgICAgICAgICAgICAgICAgCiAgICAgICAgJzpzbUhISEhISEhISEhISEgrICAgICAgICAgICAgb0hISEhISEhISGggICAgLTonICAgICAgICAgICAgCiAgICAgLitoSEhISEhISEhISEhISEhISEgrICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hIaCsuICAgICAgICAgCiAnOnNkSEhISEhISEhISEhISEhISEhoKy0gICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhIZHM6JyAgICAgCidtSEhISEhISEhISEhISEhISGRvOicgICAgICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEhISEhIbXkvLiAgICAgICAnICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEhIaCstICAgICAgIC4veW1vICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgICAgJzpvZEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIC9oSEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhoc3Nzc3Nzc3Nzc3NzaEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhob29vb29vb29vb29vaEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhteS8gICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhoby0gICAgICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb2RzLycgICAgICAgOm9kSEhISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgJyAgICAgICAuK3ltSEhISEhISEhISEhISCcgICAgCidtSEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgICAgICcvc21ISEhISEhISEhISEhISEhIbScgICAgCiAgLW9kSEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgIC1vaEhISEhISEhISEhISEhISEhIZG86ICAgICAgCiAgICAgLi95bUgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISEhISEhISEhteSsuICAgICAgICAgCiAgICAgICAgIDotICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISEhISEhkczogICAgICAgICAgICAgCiAgICAgICAgICAgICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISEh5Ky4gICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICc6c2RISEhISEhvICAgICAgICAgICAgb0hISEhISGRzOicgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAuK2hISEhvICAgICAgICAgICAgb0hISGgrLiAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAnOnMrICAgICAgICAgICAgK3M6Jwo="
	logoMedium   = "CiAgICAgICAgICAgICAgICAgICAuICAgICAgICAuICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAnL3ltcyAgICAgICAgeW15LycgICAgICAgICAgICAgICAKICAgICAgICAgICAgLW9oSEhISHMgICAgICAgIHlISEhIaG8nICAgICAgICAgICAgCiAgICAgICAgJy9zbUhISEhISEhzICAgICAgICB5SEhISEhILiAgLScgICAgICAgIAogICAgIC0raEhISEhISEhISEhtKyAgICAgICAgeUhISEhISC4gIHNIaCstICAgICAKICAgK21ISEhISEhISEhIaCstICAgICAgICAgIHlISEhISEguICBzSEhISG0rICAgCiAgIHlISEhISEhIZHM6JyAgIC4tICAgICAgICB5SEhISEhILiAgc0hISEhIeSAgIAogICB5SEhISEhoLiAgICc6c2RIeSAgICAgICAgeUhISEhISC4gIHNISEhISHkgICAKICAgeUhISEhIcyAgLmhISEhISHkgICAgICAgIHlISEhISEguICBzSEhISEh5ICAgCiAgIHlISEhISHMgIC5ISEhISEh5ICAgICAgICB5SEhISEhILiAgc0hISEhIeSAgIAogICB5SEhISEhzICAuSEhISEhIbXl5eXl5eXl5bUhISEhISC4gIHNISEhISHkgICAKICAgeUhISEhIcyAgLkhISEhISEhISEhISEhISEhISEhISEguICBzSEhISEh5ICAgCiAgIHlISEhISHMgIC5ISEhISEhISEhISEhISEhISEhISEhILiAgc0hISEhIeSAgIAogICB5SEhISEhzICAuSEhISEhIbXl5eXl5eXl5bUhISEhISC4gIHNISEhISHkgICAKICAgeUhISEhIcyAgLkhISEhISHkgICAgICAgIHlISEhISEguICBzSEhISEh5ICAgCiAgIHlISEhISHMgIC5ISEhISEh5ICAgICAgICB5SEhISEhoJyAgc0hISEhIeSAgIAogICB5SEhISEhzICAuSEhISEhIeSAgICAgICAgeUhkbzonICAgLmhISEhISHkgICAKICAgeUhISEhIcyAgLkhISEhISHkgICAgICAgIC0uICAgJy9zbUhISEhISEh5ICAgCiAgICtkSEhISHMgIC5ISEhISEh5ICAgICAgICAgIC1vaEhISEhISEhISEhkKyAgIAogICAgIC4raEhzICAuSEhISEhIeSAgICAgICAgb21ISEhISEhISEhIaCsuICAgICAKICAgICAgICAnLiAgLkhISEhISHkgICAgICAgIHlISEhISEhIbXM6JyAgICAgICAgCiAgICAgICAgICAgICcraEhISEh5ICAgICAgICB5SEhISGgrLSAgICAgICAgICAgIAogICAgICAgICAgICAgICAnYmxzfCAgICAgICAgeUhkLycgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgIC4gICAgICAgIC4K"
	logoSmall    = "CiAgICAgICAgICAgJy8nICAgICcvJyAgICAgICAgICAgCiAgICAgICAgLStoSEgnICAgIC5ISGgrLiAgICAgICAgCiAgICAnOnNtSEhISEgnICAgIC5ISEhIKyAtOicgICAgCiAgL2hISEhISEhoby0gICAgIC5ISEhIKyArSEhoLyAgCiAgaEhISG1zOicgLSsnICAgIC5ISEhIKyArSEhIaCAgCiAgaEhISC8gLXltSEguICAgIC5ISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEgtLi4uLi1ISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEhISEhISEhISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEhISEhISEhISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEgtLi4uLi1ISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEguICAgIC5ISG1zLSArSEhIaCAgCiAgaEhISC8gK0hISEguICAgICcrLiAnL3NtSEhIaCAgCiAgL3lISC8gK0hISEguICAgICA6b2RISEhISEhoLyAgCiAgICAnOi0gK0hISEguICAgIC5ISEhISGRzOicgICAgCiAgICAgICAgLitoSEguICAgIC5ISGgrLiAgICAgICAgCiAgICAgICAgICAgJzonICAgICc6Jwo="
)

// Config some things
type Config struct {
	LogoLarge  string
	LogoMedium string
	LogoSmall  string
	Version    string
}

var color, h, logo, mix, size string
var versionPtr *bool
var dazzlePtr *bool
var reset = ansi.ColorCode("reset")

func init() {

	// CLI flags
	dazzlePtr = flag.Bool("dazzle", false, "Engage dazzle mode")
	flag.StringVar(&color, "color", "blue", "color= cyan, red, green, yellow, blue, magenta, white, mix, or random")
	flag.StringVar(&size, "size", "medium", "size= small, medium, or large")
	versionPtr = flag.Bool("version", false, "Show hashii version")
}

func cleanup() {
	// Restore cursor
	fmt.Print("\033[?25h")
	fmt.Printf("\033[H\033[2J")

	bye := "\nGoodbye!\n"
	fmt.Print(ansi.ColorCode("blue"), bye, reset)
}

// Time to make the 🍩🍩🍩🍩🍩🍩🍩
func main() {
	flag.Parse()

	c := &Config{}

	// Quick n dirty version info
	c.Version = "1.0.3"
	if *versionPtr {
		if c.Version == "github-master" {
			fmt.Printf("hashii %s\n", c.Version)
		} else {
			fmt.Printf("hashii v%s\n", c.Version)
		}
		os.Exit(0)
	}

	// Random color mode setup
	rand.Seed(time.Now().Unix())
	colors := []string{"cyan", "red", "green", "yellow", "blue", "magenta", "white"}
	rando := colors[rand.Intn(len(colors))]

	// CTRL-C handling for cleanup
	chn := make(chan os.Signal, 2)
	signal.Notify(chn, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-chn
		cleanup()
		os.Exit(0)
	}()

	// Panic on these because there really should not be a resonable
	// way that the string cannot be decoded...
	switch size {
	case "small":
		decodedLogo, err := base64.StdEncoding.DecodeString(logoSmall)
		if err != nil {
			panic(err)
		}
		logo = string(decodedLogo)
	case "medium":
		decodedLogo, err := base64.StdEncoding.DecodeString(logoMedium)
		if err != nil {
			panic(err)
		}
		logo = string(decodedLogo)
	case "large":
		decodedLogo, err := base64.StdEncoding.DecodeString(logoLarge)
		if err != nil {
			panic(err)
		}
		logo = string(decodedLogo)
	default:
		decodedLogo, err := base64.StdEncoding.DecodeString(logoMedium)
		if err != nil {
			panic(err)
		}
		logo = string(decodedLogo)
	}

	// ✨ Dazzle mode!
	if *dazzlePtr {
		// Hide the cursor
		fmt.Print("\033[?25l")

		// Yes, it is an infinite loop — we're just havin' a 'lil fun here
		// and we have that goroutine above to handle SIGINT a tiny bit ☠️
		for {
			fmt.Printf("\033[H\033[2J")
			dazzler := bufio.NewScanner(strings.NewReader(logo))
			for dazzler.Scan() {
				h = dazzler.Text()
				mix = colors[rand.Intn(len(colors))]
				fmt.Println(ansi.ColorCode(mix), h, reset)
			}
			time.Sleep(142 * time.Millisecond)
		}
		/*
		   / Due to the infinity above whatever ends up here is unreachable
		   / code... What if there could somehow be a bit of code that was
		   / reachable through magic as explained in this delightful story:
		   /
		   / https://www.cs.utah.edu/~elb/folklore/magic.html
		   /
		*/
	}

	// Basic color mode
	scanner := bufio.NewScanner(strings.NewReader(logo))
	for scanner.Scan() {
		h = scanner.Text()
		switch color {
		case "plain":
			fmt.Println(h)
		case color:
			if color == "mix" {
				mix := colors[rand.Intn(len(colors))]
				fmt.Println(ansi.ColorCode(mix), h, reset)
			} else if color == "random" {
				fmt.Println(ansi.ColorCode(rando), h, reset)
			} else {
				fmt.Println(ansi.ColorCode(color), h, reset)
			}
		default:
			fmt.Println(defaultColor, h, reset)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
