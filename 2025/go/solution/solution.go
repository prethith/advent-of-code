package solution

type Solution interface {
	Part1(input string) (string, error)
	Part2(input string) (string, error)
}
