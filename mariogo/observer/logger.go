package observer

import (
	"fmt"
	"jkli/mariogo/mariogo/analyzer"
	"time"
)

type Logger struct {
}

func (l *Logger) StateChange(from int, to int) {

	fmt.Print("State change: ")

	if to == analyzer.Idle {
		fmt.Println("Game is idle")
	} else if to == analyzer.Loading {
		fmt.Println("Game loading")
	} else if to == analyzer.Racing {
		if from == analyzer.Pause {
			fmt.Println("Game resumed")
		} else {
			fmt.Println("Racing")
		}
	} else if to == analyzer.Pause {
		fmt.Println("Game paused")
	} else if to == analyzer.RoundResults {
		fmt.Println("Round results")
	} else if to == analyzer.InterimResults {
		fmt.Println("Interim results")
	} else if to == analyzer.EndResults {
		fmt.Println("Cup End")
	}
}

func (l *Logger) PlayerCount(count int) {
	fmt.Println("Player count: ", count)
}

func (l *Logger) PlayerName(player int, name string) {
	fmt.Println("Player", player, "is", name)
}

func (l *Logger) NewRound(name string) {
	fmt.Println("New round: ", name)
}

func (l *Logger) RoundResults(placements [4]int) {
	fmt.Println("Round results: ", placements)
}

func (l *Logger) InterimResults(placements [4]int) {
	fmt.Println("Interim results: ", placements)
}

func (l *Logger) Abort(message string) {
	fmt.Println("Game aborted: ", message)
}

func (l *Logger) RoundFinished(player int, round int, time time.Duration, finished bool) {
	fmt.Println("Player", player, "finished round", round, "in", time)
	if finished {
		fmt.Println("Player", player, "finished")
	}
}

func (l *Logger) PlacementsChanged(old [4]int, new [4]int) {
	for i := 0; i < 4; i++ {
		if old[i] != new[i] && old[i] != 0 && new[i] != 0 {
			fmt.Println("Player", i+1, "moved from", old[i], "to", new[i])
		}
	}
}
