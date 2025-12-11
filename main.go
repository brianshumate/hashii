package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const version = "2.0.2"

// Base64-encoded logo art in three sizes
const (
	haLg = "CiAgICAgICAgICAgICAgICAgICAgICAnL3MvICAgICAgICAgICAgK3kvJyAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAtK2hISEgrICAgICAgICAgICAgb0hISGhvLSAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICc6c21ISEhISEgrICAgICAgICAgICAgb0hISEhISG1zLycgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgIC4raEhISEhISEhISEgrICAgICAgICAgICAgb0hISEhISEhISGggICAgICAgICAgICAgICAgICAgCiAgICAgICAgJzpzbUhISEhISEhISEhISEgrICAgICAgICAgICAgb0hISEhISEhISGggICAgLTonICAgICAgICAgICAgCiAgICAgLitoSEhISEhISEhISEhISEhISEgrICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hIaCsuICAgICAgICAgCiAnOnNkSEhISEhISEhISEhISEhISEhoKy0gICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhIZHM6JyAgICAgCidtSEhISEhISEhISEhISEhISGRvOicgICAgICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEhISEhIbXkvLiAgICAgICAnICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEhIaCstICAgICAgIC4veW1vICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgICAgJzpvZEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIC9oSEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhoc3Nzc3Nzc3Nzc3NzaEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhISEhISEhISEhISEhISEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhob29vb29vb29vb29vaEhISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISGggICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhteS8gICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhoby0gICAgICAgK0hISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb2RzLycgICAgICAgOm9kSEhISEhISEhISCcgICAgCidISEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgJyAgICAgICAuK3ltSEhISEhISEhISEhISCcgICAgCidtSEhISEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgICAgICcvc21ISEhISEhISEhISEhISEhIbScgICAgCiAgLW9kSEhISEgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgIC1vaEhISEhISEhISEhISEhISEhIZG86ICAgICAgCiAgICAgLi95bUgrICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISEhISEhISEhteSsuICAgICAgICAgCiAgICAgICAgIDotICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISEhISEhkczogICAgICAgICAgICAgCiAgICAgICAgICAgICAgIGhISEhISEhISEhvICAgICAgICAgICAgb0hISEhISEhISEh5Ky4gICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICc6c2RISEhISEhvICAgICAgICAgICAgb0hISEhISGRzOicgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAuK2hISEhvICAgICAgICAgICAgb0hISGgrLiAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAnOnMrICAgICAgICAgICAgK3M6Jwo="
	haMd = "CiAgICAgICAgICAgICAgICAgICAuICAgICAgICAuICAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAnL3ltcyAgICAgICAgeW15LycgICAgICAgICAgICAgICAKICAgICAgICAgICAgLW9oSEhISHMgICAgICAgIHlISEhIaG8nICAgICAgICAgICAgCiAgICAgICAgJy9zbUhISEhISEhzICAgICAgICB5SEhISEhILiAgLScgICAgICAgIAogICAgIC0raEhISEhISEhISEhtKyAgICAgICAgeUhISEhISC4gIHNIaCstICAgICAKICAgK21ISEhISEhISEhIaCstICAgICAgICAgIHlISEhISEguICBzSEhISG0rICAgCiAgIHlISEhISEhIZHM6JyAgIC4tICAgICAgICB5SEhISEhILiAgc0hISEhIeSAgIAogICB5SEhISEhoLiAgICc6c2RIeSAgICAgICAgeUhISEhISC4gIHNISEhISHkgICAKICAgeUhISEhIcyAgLmhISEhISHkgICAgICAgIHlISEhISEguICBzSEhISEh5ICAgCiAgIHlISEhISHMgIC5ISEhISEh5ICAgICAgICB5SEhISEhILiAgc0hISEhIeSAgIAogICB5SEhISEhzICAuSEhISEhIbXl5eXl5eXl5bUhISEhISC4gIHNISEhISHkgICAKICAgeUhISEhIcyAgLkhISEhISEhISEhISEhISEhISEhISEguICBzSEhISEh5ICAgCiAgIHlISEhISHMgIC5ISEhISEhISEhISEhISEhISEhISEhILiAgc0hISEhIeSAgIAogICB5SEhISEhzICAuSEhISEhIbXl5eXl5eXl5bUhISEhISC4gIHNISEhISHkgICAKICAgeUhISEhIcyAgLkhISEhISHkgICAgICAgIHlISEhISEguICBzSEhISEh5ICAgCiAgIHlISEhISHMgIC5ISEhISEh5ICAgICAgICB5SEhISEhoJyAgc0hISEhIeSAgIAogICB5SEhISEhzICAuSEhISEhIeSAgICAgICAgeUhkbzonICAgLmhISEhISHkgICAKICAgeUhISEhIcyAgLkhISEhISHkgICAgICAgIC0uICAgJy9zbUhISEhISEh5ICAgCiAgICtkSEhISHMgIC5ISEhISEh5ICAgICAgICAgIC1vaEhISEhISEhISEhkKyAgIAogICAgIC4raEhzICAuSEhISEhIeSAgICAgICAgb21ISEhISEhISEhIaCsuICAgICAKICAgICAgICAnLiAgLkhISEhISHkgICAgICAgIHlISEhISEhIbXM6JyAgICAgICAgCiAgICAgICAgICAgICcraEhISEh5ICAgICAgICB5SEhISGgrLSAgICAgICAgICAgIAogICAgICAgICAgICAgICAnYmxzfCAgICAgICAgeUhkLycgICAgICAgICAgICAgICAKICAgICAgICAgICAgICAgICAgIC4gICAgICAgIC4K"
	haSm = "CiAgICAgICAgICAgJy8nICAgICcvJyAgICAgICAgICAgCiAgICAgICAgLStoSEgnICAgIC5ISGgrLiAgICAgICAgCiAgICAnOnNtSEhISEgnICAgIC5ISEhIKyAtOicgICAgCiAgL2hISEhISEhoby0gICAgIC5ISEhIKyArSEhoLyAgCiAgaEhISG1zOicgLSsnICAgIC5ISEhIKyArSEhIaCAgCiAgaEhISC8gLXltSEguICAgIC5ISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEgtLi4uLi1ISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEhISEhISEhISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEhISEhISEhISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEgtLi4uLi1ISEhIKyArSEhIaCAgCiAgaEhISC8gK0hISEguICAgIC5ISG1zLSArSEhIaCAgCiAgaEhISC8gK0hISEguICAgICcrLiAnL3NtSEhIaCAgCiAgL3lISC8gK0hISEguICAgICA6b2RISEhISEhoLyAgCiAgICAnOi0gK0hISEguICAgIC5ISEhISGRzOicgICAgCiAgICAgICAgLitoSEguICAgIC5ISGgrLiAgICAgICAgCiAgICAgICAgICAgJzonICAgICc6Jwo="
)

// ColorCode represents an ANSI color escape sequence
type ColorCode string

// ANSI color codes
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

// allColors contains all available colors for random selection
var allColors = []ColorCode{Cyan, Red, Green, Yellow, Blue, Magenta, White}

// colorMap maps color names to their ANSI codes
var colorMap = map[string]ColorCode{
	"red":     Red,
	"green":   Green,
	"yellow":  Yellow,
	"blue":    Blue,
	"magenta": Magenta,
	"cyan":    Cyan,
	"white":   White,
}

// Config holds the application configuration
type Config struct {
	DazzleMode bool
	Color      string
	Size       string
	PlainMode  bool
}

// App holds the application state
type App struct {
	config Config
	logo   string
	output io.Writer
}

// NewApp creates a new App with the given config
func NewApp(cfg Config, output io.Writer) *App {
	return &App{
		config: cfg,
		output: output,
	}
}

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, output io.Writer) error {
	cfg, err := parseFlags(args)
	if err != nil {
		return err
	}

	app := NewApp(cfg, output)

	if err := app.loadLogo(); err != nil {
		return fmt.Errorf("failed to load logo: %w", err)
	}

	if app.config.DazzleMode {
		return app.runDazzle()
	}

	color := app.selectColor()
	app.printLogo(color)
	return nil
}

func parseFlags(args []string) (Config, error) {
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)

	var cfg Config
	fs.BoolVar(&cfg.DazzleMode, "dazzle", false, "Engage dazzle mode")
	fs.StringVar(&cfg.Color, "color", "blue", "Color (cyan, red, green, yellow, blue, magenta, white, mix, random)")
	fs.StringVar(&cfg.Size, "size", "medium", "Size (small, medium, or large)")
	fs.BoolVar(&cfg.PlainMode, "plain", false, "Display logo without color")

	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage of hashii v%s:\n", version)
		fmt.Fprintln(fs.Output(), "  hashii [options]")
		fmt.Fprintln(fs.Output(), "  hashii version")
		fmt.Fprintln(fs.Output(), "  hashii help")
		fmt.Fprintln(fs.Output(), "\nOptions:")
		fs.PrintDefaults()
	}

	// Handle subcommands before parsing flags
	if len(args) > 1 {
		switch args[1] {
		case "version":
			fmt.Printf("hashii v%s\n", version)
			os.Exit(0)
		case "help":
			fs.Usage()
			os.Exit(0)
		}
	}

	if err := fs.Parse(args[1:]); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (a *App) loadLogo() error {
	var encoded string
	switch a.config.Size {
	case "small":
		encoded = haSm
	case "large":
		encoded = haLg
	case "medium":
		encoded = haMd
	default:
		encoded = haMd
	}

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return fmt.Errorf("decode base64: %w", err)
	}
	a.logo = string(decoded)
	return nil
}

func (a *App) selectColor() ColorCode {
	if a.config.PlainMode {
		return Reset
	}

	switch a.config.Color {
	case "mix", "random":
		return randomColor()
	case "plain":
		a.config.PlainMode = true
		return Reset
	default:
		if c, ok := colorMap[a.config.Color]; ok {
			return c
		}
		return Blue
	}
}

func randomColor() ColorCode {
	return allColors[rand.IntN(len(allColors))]
}

func (a *App) printLogo(color ColorCode) {
	scanner := bufio.NewScanner(strings.NewReader(a.logo))
	for scanner.Scan() {
		if a.config.PlainMode {
			fmt.Fprintln(a.output, scanner.Text())
		} else {
			fmt.Fprintf(a.output, "%s%s%s\n", color, scanner.Text(), Reset)
		}
	}
	// Scanner errors on strings.Reader are not possible, so we skip error checking
}

func (a *App) runDazzle() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		cancel()
	}()

	// Hide cursor
	fmt.Fprint(a.output, "\033[?25l")

	// Ensure cleanup on exit
	defer func() {
		fmt.Fprint(a.output, "\033[?25h") // Show cursor
		fmt.Fprint(a.output, "\033[H\033[2J")
		fmt.Fprintf(a.output, "%sBye!%s\n", Blue, Reset)
	}()

	ticker := time.NewTicker(142 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			fmt.Fprint(a.output, "\033[H\033[2J")
			a.printLogo(randomColor())
		}
	}
}
