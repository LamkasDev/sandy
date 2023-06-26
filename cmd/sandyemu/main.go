package main

import (
	"flag"
	"os"
	"os/user"
	"path"
	"runtime/pprof"

	"github.com/LamkasDev/sandy/cmd/common/flags"
	"github.com/LamkasDev/sandy/cmd/sandyemu/ps"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	u, _ := user.Current()
	defaultBios := path.Join(u.HomeDir, "Documents", "sandy", "resources", "bios", "SCPH1001.bin")
	defaultLog := path.Join(u.HomeDir, "Documents", "sandy", "log.txt")
	defaultProfile := path.Join(u.HomeDir, "Documents", "sandy", "cpu.prof")

	var biosPath string
	var logPath string
	var profilePath string
	flag.StringVar(&biosPath, "bios", defaultBios, "path to bios .bin file")
	flag.StringVar(&logPath, "log", defaultLog, "path to log file")
	flag.StringVar(&profilePath, "profile", defaultProfile, "path to profile file")
	flag.Parse()
	flags.ResolveColor()

	logFile, err := os.Create(logPath)
	if err != nil {
		log.Error().Msgf("failed to open log file: %v", err)
		panic(err)
	}
	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, logFile)
	// multi := zerolog.MultiLevelWriter(logFile)
	log.Logger = zerolog.New(multi).Level(zerolog.DebugLevel)
	// log.Logger = zerolog.New(multi).Level(zerolog.InfoLevel)

	profileFile, err := os.Create(profilePath)
	if err != nil {
		log.Error().Msgf("failed to open profile file: %v", err)
		panic(err)
	}
	pprof.StartCPUProfile(profileFile)

	emulator := ps.NewPlaystation()
	ps.LoadPlaystation(&emulator, biosPath)
	ps.RunPlaystation(&emulator)
}
