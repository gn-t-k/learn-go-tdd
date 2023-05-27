package poker

type CLI struct {
	playerStore PlayerStore
}

func (cli *CLI) PlayPoler() {
	cli.playerStore.RecordWin("Cleo")
}

// TODO: つぎここ
// https://andmorefine.gitbook.io/learn-go-with-tests/build-an-application/command-line#nitesutowoku-1
