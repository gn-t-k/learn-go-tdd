package poker

import "io"

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayPoler() {
	cli.playerStore.RecordWin("Chris")
}

// TODO: https://andmorefine.gitbook.io/learn-go-with-tests/build-an-application/command-line#nitesutowoku-2
