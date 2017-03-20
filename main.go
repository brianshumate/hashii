package main

import (
    "bufio"
    base64 "encoding/base64"
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

var color, h, logo, mix, scanner, size string
var dazzlePtr *bool
var version = "github-master"
var cyn = ansi.ColorCode("cyan")
var grn = ansi.ColorCode("green")
var blu = ansi.ColorCode("blue")
var mag = ansi.ColorCode("magenta")
var red = ansi.ColorCode("red")
var rst = ansi.ColorCode("reset")
var wht = ansi.ColorCode("white")
var ylw = ansi.ColorCode("yellow")
var haLg = "CiAgICAgICAgICAgICAgICAgICAgICAnL3MvICAgICAgICAgICAgK3kvJyAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAtK2hISEgrICAgICAgICAgICAgb0hISGhvLSAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICc6c21ISEhISEgrICAgICAgICAgICAgb0hISEhISG1zLycgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgIC4raEhISEhISEhISEgrICAgICAgICAgICAgb0hISEhISEhISGggICAgICAgICAgICAgICAgICAgCiAgICAgICAgJzpzbUhISEhISEhISEhISEgrICAgICAgICAgICAgb0hISEhISEhISGggICAgLTonICAgICAgICAgICAgCiAgICAgLitoSEhISEhISEhISEhISEhISEgrICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hIaCsuICAgICAgICAgCiAnOnNkSEhISEhISEhISEhISEhISEhoKy0gICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhIZHM6JyAgICAgCidtSEhISEhISEhISEhISEhISGRvOicgICAgICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEhISEhIbXkvLiAgICAgICAnICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEhIaCstICAgICAgIC4veW1vICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgICAgJzpvZEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIC9oSEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhoc3Nzc3Nzc3Nzc3NzaEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhob29vb29vb29vb29vaEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhteS8gICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhoby0gICAgICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb2RzLycgICAgICAgOm9kSEhISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgJyAgICAgICAuK3ltSEhISEhISEhISEhISCcgICAgCidtSEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgICAgICcvc21ISEhISEhISEhISEhISEhIbScgICAgCiAgLW9kSEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgIC1vaEhISEhISEhISEhISEhISEhIZG86ICAgICAgCiAgICAgLi95bUgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISEhISEhISEhteSsuICAgICAgICAgCiAgICAgICAgIDotICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISEhISEhkczogICAgICAgICAgICAgCiAgICAgICAgICAgICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISEh5Ky4gICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICc6c2RISEhISEhvICAgICAgICAgICAgb0hISEhISGRzOicgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAuK2hISEhvICAgICAgICAgICAgb0hISGgrLiAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAnOnMrICAgICAgICAgICAgK3M6Jwo="
var haMd = "CiAgICAgICAgICAgICAgICAgICAuICAgICAgICAuICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAnL3ltcyAgICAgICAgeW15LycgICAgICAgICAgICAgICAKICAgICAgICAgICAgLW9oSEhISHMgICAgICAgIHlISEhIaG8nICAgICAgICAgICAgCiAgICAgICAgJy9zbUhISEhISEhzICAgICAgICB5SEhISEhILiAgLScgICAgICAgIAogICAgIC0raEhISEhISEhISEhtKyAgICAgICAgeUhISEhISC4gIHNIaCstICAgICAKICAgK21ISEhISEhISEhIaCstICAgICAgICAgIHlISEhISEguICBzSEhISG0rICAgCiAgIHlISEhISEhIZHM6JyAgIC4tICAgICAgICB5SEhISEhILiAgc0hISEhIeSAgIAogICB5SEhISEhoLiAgICc6c2RIeSAgICAgICAgeUhISEhISC4gIHNISEhISHkgICAKICAgeUhISEhIcyAgLmhISEhISHkgICAgICAgIHlISEhISEguICBzSEhISEh5ICAgCiAgIHlISEhISHMgIC5ISEhISEh5ICAgICAgICB5SEhISEhILiAgc0hISEhIeSAgIAogICB5SEhISEhzICAuSEhISEhIbXl5eXl5eXl5bUhISEhISC4gIHNISEhISHkgICAKICAgeUhISEhIcyAgLkhISEhISEhISEhISEhISEhISEhISEguICBzSEhISEh5ICAgCiAgIHlISEhISHMgIC5ISEhISEhISEhISEhISEhISEhISEhILiAgc0hISEhIeSAgIAogICB5SEhISEhzICAuSEhISEhIbXl5eXl5eXl5bUhISEhISC4gIHNISEhISHkgICAKICAgeUhISEhIcyAgLkhISEhISHkgICAgICAgIHlISEhISEguICBzSEhISEh5ICAgCiAgIHlISEhISHMgIC5ISEhISEh5ICAgICAgICB5SEhISEhoJyAgc0hISEhIeSAgIAogICB5SEhISEhzICAuSEhISEhIeSAgICAgICAgeUhkbzonICAgLmhISEhISHkgICAKICAgeUhISEhIcyAgLkhISEhISHkgICAgICAgIC0uICAgJy9zbUhISEhISEh5ICAgCiAgICtkSEhISHMgIC5ISEhISEh5ICAgICAgICAgIC1vaEhISEhISEhISEhkKyAgIAogICAgIC4raEhzICAuSEhISEhIeSAgICAgICAgb21ISEhISEhISEhIaCsuICAgICAKICAgICAgICAnLiAgLkhISEhISHkgICAgICAgIHlISEhISEhIbXM6JyAgICAgICAgCiAgICAgICAgICAgICcraEhISEh5ICAgICAgICB5SEhISGgrLSAgICAgICAgICAgIAogICAgICAgICAgICAgICAnYmxzfCAgICAgICAgeUhkLycgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgIC4gICAgICAgIC4K"
var haSm = "CiAgICAgICAgICAgJy8nICAgICcvJyAgICAgICAgICAgCiAgICAgICAgLStoSEgnICAgIC5ISGgrLiAgICAgICAgCiAgICAnOnNtSEhISEgnICAgIC5ISEhIKyAtOicgICAgCiAgL2hISEhISEhoby0gICAgIC5ISEhIKyArSEhoLyAgCiAgaEhISG1zOicgLSsnICAgIC5ISEhIKyArSEhIaCAgCiAgaEhISC8gLXltSEguICAgIC5ISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEgtLi4uLi1ISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEhISEhISEhISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEhISEhISEhISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEgtLi4uLi1ISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEguICAgIC5ISG1zLSArSEhIaCAgCiAgaEhISC8gK0hISEguICAgICcrLiAnL3NtSEhIaCAgCiAgL3lISC8gK0hISEguICAgICA6b2RISEhISEhoLyAgCiAgICAnOi0gK0hISEguICAgIC5ISEhISGRzOicgICAgCiAgICAgICAgLitoSEguICAgIC5ISGgrLiAgICAgICAgCiAgICAgICAgICAgJzonICAgICc6Jwo="

func init() {
    // Quick n dirty version info
    if len(os.Args) > 1 {
        if os.Args[1] == "version" {
            if version == "github-master" {
                fmt.Printf("hashii %s\n", version)
            } else {
                fmt.Printf("hashii v%s\n", version)
            }
            os.Exit(0)
        }
    }

    // CLI flags
    dazzlePtr = flag.Bool("dazzle", false, "Engage dazzle mode")
    flag.StringVar(&color, "color", "blue", "color= cyan, red, green, yellow, blue, magenta, white, mix, or random")
    flag.StringVar(&size, "size", "medium", "size= small, medium, or large")
}

func cleanup() {
    // Restore cursor
    fmt.Print("\033[?25h")
    fmt.Printf("\033[H\033[2J")
    bye := "\nGoodbye!\n"
    fmt.Print(blu, bye, rst)
}

func main() {
    flag.Parse()

    // Random color mode setup
    rand.Seed(time.Now().Unix())
    colors := make([]string, 0)
    colors = append(colors, cyn, red, grn, ylw, blu, mag, wht)
    rando := colors[rand.Intn(len(colors))]

    // Setup CTRL-C handling for cleanup
    c := make(chan os.Signal, 2)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        cleanup()
        os.Exit(0)
    }()

    switch size {
    case "small":
        decLogo, _ := base64.StdEncoding.DecodeString(haSm)
        logo = string(decLogo)
    case "medium":
        decLogo, _ := base64.StdEncoding.DecodeString(haMd)
        logo = string(decLogo)
    case "large":
        decLogo, _ := base64.StdEncoding.DecodeString(haLg)
        logo = string(decLogo)
    default:
        decLogo, _ := base64.StdEncoding.DecodeString(haMd)
        logo = string(decLogo)
    }

    // Time to make the 🍩🍩🍩🍩🍩🍩🍩

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
                fmt.Println(mix, h, rst)
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

        // Choose a color with the switch statement of doom™!
        switch color {
        case "plain":
            fmt.Println(h)
        case "red":
            fmt.Println(red, h, rst)
        case "green":
            fmt.Println(grn, h, rst)
        case "yellow":
            fmt.Println(ylw, h, rst)
        case "blue":
            fmt.Println(blu, h, rst)
        case "magenta":
            fmt.Println(mag, h, rst)
        case "cyan":
            fmt.Println(cyn, h, rst)
        case "white":
            fmt.Println(wht, h, rst)
        case "mix":
            mix := colors[rand.Intn(len(colors))]
            fmt.Println(mix, h, rst)
        case "random":
            fmt.Println(rando, h, rst)
        default:
            fmt.Println(blu, h, rst)
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
