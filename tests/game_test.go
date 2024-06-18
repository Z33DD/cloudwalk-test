package internal_test

import (
	internal "cloudwalk-test/internal"
	"reflect"
	"testing"
)

func TestParseKill(t *testing.T) {
	g := internal.Game{}
	line := "Kill: 1 2 3: John Doe killed Jane Smith by gun"
	killer, killed, mean := g.ParseKill(line)

	expectedKiller := "John Doe"
	if killer != expectedKiller {
		t.Errorf("Expected killer to be %s, but got %s", expectedKiller, killer)
	}

	expectedKilled := "Jane Smith"
	if killed != expectedKilled {
		t.Errorf("Expected killed to be %s, but got %s", expectedKilled, killed)
	}

	expectedMean := "gun"
	if mean != expectedMean {
		t.Errorf("Expected mean to be %s, but got %s", expectedMean, mean)
	}
}
func TestProcessKill(t *testing.T) {
	g := internal.Game{}
	g.New("")

	line := "Kill: 1 2 3: John Doe killed Jane Smith by gun"
	g.ProcessKill(line)

	expectedKills := map[string]int{
		"John Doe": 1,
	}
	if !reflect.DeepEqual(g.Kills, expectedKills) {
		t.Errorf("Expected Kills to be %v, but got %v", expectedKills, g.Kills)
	}

	expectedKillsByMean := map[string]int{
		"gun": 1,
	}
	if !reflect.DeepEqual(g.KillsByMean, expectedKillsByMean) {
		t.Errorf("Expected KillsByMean to be %v, but got %v", expectedKillsByMean, g.KillsByMean)
	}

	expectedTotalKills := 1
	if g.TotalKills != expectedTotalKills {
		t.Errorf("Expected TotalKills to be %d, but got %d", expectedTotalKills, g.TotalKills)
	}
}
