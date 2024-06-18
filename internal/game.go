package internal

import (
	"log"
	"regexp"
	"strings"
	"sync"
)

type Game struct {
	mu          sync.Mutex
	Id          int            `json:"id"`
	MapName     string         `json:"map_name"`
	GameName    string         `json:"game_name"`
	TotalKills  int            `json:"total_kills"`
	Kills       map[string]int `json:"kills"`
	KillsByMean map[string]int `json:"kills_by_mean"`
}

func (g *Game) New(line string) {
	re := regexp.MustCompile(`\\(\w+)\\([^\\]+)`)

	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		switch match[1] {
		case "mapname":
			g.MapName = match[2]
		case "gamename":
			g.GameName = match[2]
		}
	}
	g.Kills = make(map[string]int)
	g.KillsByMean = make(map[string]int)
}

func (g *Game) ParseNewLogLine(line string) {

	if strings.HasPrefix(line, "Kill:") {
		g.processKill(line)
	}
}

func (g *Game) processKill(line string) {
	killer, killed, mean := g.parseKill(line)
	log.Println(line)
	log.Printf("killer: %s, killed: %s, mean: %s\n", killer, killed, mean)
	log.Println()

	g.mu.Lock()
	defer g.mu.Unlock()

	if killer == "<world>" {
		g.Kills[killed]--
		if g.Kills[killed] < 0 {
			g.Kills[killed] = 0
		}
	} else {
		g.Kills[killer]++
	}

	g.KillsByMean[mean]++
	g.TotalKills++
}
func (g *Game) parseKill(line string) (string, string, string) {
	re := regexp.MustCompile(`Kill: \d+ \d+ \d+: ([^\s]+(?: [^\s]+)*) killed ([^\s]+(?: [^\s]+)*) by ([^\s]+(?: [^\s]+)*)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 4 {
		panic("Invalid kill line")
	}
	killer := matches[1]
	killed := matches[2]
	mean := matches[3]
	return killer, killed, mean
}
