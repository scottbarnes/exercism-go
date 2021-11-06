/*
Take CSV input in the form of:
Allegoric Alaskans;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskans;win

Then parse it and return a 'pretty' table that looks like:
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskans             |  3 |  2 |  0 |  1 |  6
*/
package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name   string
	drawn  int
	lost   int
	played int
	points int
	won    int
}

type teamsMap map[string]*team

// Adjust a winning team's number of matches played, matches won, and total points.
func (t *team) win() {
	t.played++
	t.won++
	t.points += 3
}

// Adjust a losing team's number of matches played and lost.
func (t *team) loss() {
	t.played++
	t.lost++
}

// Adjust a drawn team's number of matches played, drawn, and total points.
func (t *team) draw() {
	t.played++
	t.drawn++
	t.points += 1
}

// Look up a team in a map, create it if it doesn't exist, and return it.
func (t teamsMap) getOrCreateTeam(n string) *team {
	if _, ok := t[n]; !ok {
		t[n] = &team{name: n}
	}
	return t[n]
}

// Parse input in the form of a three-value CSV and update team scores as appropriate.
func (t teamsMap) parseResults(r io.Reader) error {
	// Create a scanner and read through the input line by line.
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt := scanner.Text()

		// Ignore blank lines or comments.
		if txt == "" || txt[0] == '#' {
			continue
		}

		// Split the CSV into three parts and assign the first value to team1, the second to team2 and the third to the result.
		chunk := strings.Split(txt, ";")
		if len(chunk) != 3 {
			return errors.New("wrong number of fields")
		}
		t1name, t2name, result := chunk[0], chunk[1], chunk[2]

		// For the teams and result in this line of text, get or create them.
		t1 := t.getOrCreateTeam(t1name)
		t2 := t.getOrCreateTeam(t2name)

		// Depending on the result, handle the team scores.
		switch result {
		case "win":
			t1.win()
			t2.loss()
		case "loss":
			t1.loss()
			t2.win()
		case "draw":
			t1.draw()
			t2.draw()
		default:
			return errors.New("unknown result type")
		}
	}
	return nil
}

// Sort the teams based on scores, with alphabetical order being the tie-breaker. Returns a slice of teams.
func (t teamsMap) sort() []*team {
	// Put the map's struct values (i.e. the teams) into a slice for sorting.
	sortedKeys := make([]*team, 0, len(t))
	for _, v := range t {
		sortedKeys = append(sortedKeys, v)
	}

	// Sort the structs in descending order of total points, unless tied
	sort.Slice(sortedKeys, func(i, j int) bool {
		// Handle ties scores by sorting alphabetically
		if sortedKeys[i].points == sortedKeys[j].points {
			return sortedKeys[i].name < sortedKeys[j].name
		}
		return sortedKeys[i].points > sortedKeys[j].points
	})

	return sortedKeys
}

// Tally scores from CSV input and output to a Writer.
func Tally(r io.Reader, w io.Writer) error {
	teams := make(teamsMap)
	err := teams.parseResults(r)
	if err != nil {
		return err
	}

	// Get a slice of the winning teams, sorted by points, and pretty-print them to the screen.
	sortedKeys := teams.sort()
	fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n", "Team", "MP", "W", "D", "L", "P")
	for _, v := range sortedKeys {
		fmt.Fprintf(w, "%-31v|%3d |%3d |%3d |%3d |%3d\n", v.name, v.played, v.won, v.drawn, v.lost, v.points)
	}

	return nil
}
