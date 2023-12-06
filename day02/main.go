package day02

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ndfsa/advent-of-code-2023/day02/parser"
	"github.com/ndfsa/advent-of-code-2023/util"
)

type Game struct {
	Id    int
	Turns []Turn
}

type Turn struct {
	Colors map[string]int
}

type GameListener struct {
	*parser.BaseGameGrammarListener
	Games []Game
	Game  Game
	Turn  Turn
}

func (l *GameListener) EnterInitial(ctx *parser.InitialContext) {
	l.Games = make([]Game, 0)
}

func (l *GameListener) EnterGame(ctx *parser.GameContext) {
	l.Game = Game{Turns: make([]Turn, 0)}
}

func (l *GameListener) ExitGame(ctx *parser.GameContext) {
	count, _ := strconv.Atoi(ctx.INT().GetText())
	l.Game.Id = count
	l.Games = append(l.Games, l.Game)
}

func (l *GameListener) EnterTurn(ctx *parser.TurnContext) {
	l.Turn = Turn{Colors: make(map[string]int)}
}

func (l *GameListener) ExitTurn(ctx *parser.TurnContext) {
	l.Game.Turns = append(l.Game.Turns, l.Turn)
}

func (l *GameListener) ExitColor(ctx *parser.ColorContext) {
	count, _ := strconv.Atoi(ctx.INT().GetText())
	color := ctx.ID().GetText()
	l.Turn.Colors[color] = count
}

func parse(input string) []Game {
	// Create the lexer and parser
	lexer := parser.NewGameGrammarLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewGameGrammarParser(stream)

	listener := &GameListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Initial())

	return listener.Games
}

func possible(r, g, b, rMax, gMax, bMax int) bool {
	return r <= rMax &&
		g <= gMax &&
		b <= bMax
}

func gameMax(g Game) (int, int, int) {
	rMax, gMax, bMax := 0, 0, 0
	for _, turn := range g.Turns {
		r, g, b := turn.Colors["red"], turn.Colors["green"], turn.Colors["blue"]
		if r > rMax {
			rMax = r
		}
		if g > gMax {
			gMax = g
		}
		if b > bMax {
			bMax = b
		}
	}

	return rMax, gMax, bMax
}

func SolvePart1(filePath string) (int, error) {
	input, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	games := parse(input)
	res := 0
	for idx, game := range games {
		r, g, b := gameMax(game)
		if possible(r, g, b, 12, 13, 14) {
			res += idx + 1
		}
	}

	return res, nil
}

func calculatePower(game Game) int {
	r, g, b := gameMax(game)
	return r * g * b
}

func SolvePart2(filePath string) (int, error) {
	input, err := util.ReadFile(filePath)

	if err != nil {
		return 0, err
	}

	games := parse(input)
	res := 0
	for _, game := range games {
		res += calculatePower(game)
	}

	return res, nil
}
