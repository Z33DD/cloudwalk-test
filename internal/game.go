package internal

import "regexp"

type Game struct {
	MapName  string
	GameName string
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
}
func ParseNewLine(line string) {

}
