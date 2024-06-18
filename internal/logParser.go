package internal

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type LogParser struct {
	FilePath string
}

func (lp *LogParser) Parse() {
	file, err := os.Open(lp.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var games []*Game

	var game *Game = nil

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, remaining := splitTimeFromLog(line)

		if strings.Contains(remaining, "InitGame") {
			game = &Game{}
			game.New(remaining)
		} else if strings.Contains(remaining, "ShutdownGame") || strings.Contains(remaining, "------------------------------------------------------------") {
			if game != nil {
				games = append(games, game)
				game = nil
			}
		}

	}
	log.Println(len(games))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// splitTimeFromLog splits the time from a log line and returns the time and the remaining part of the line.
// If the log line is invalid, it panics with an error message.
func splitTimeFromLog(line string) (string, string) {
	line = strings.TrimSpace(line)
	index := strings.Index(line, " ")
	if index == -1 {
		panic("Invalid log line: " + line)
	}
	time := line[:index]
	remaining := line[index+1:]
	return time, remaining
}
