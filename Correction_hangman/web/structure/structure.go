package structure

//structure of the game
type Stock struct {
	Title       string
	Right       []string
	Wrong       []string
	Attempts    int
	TargetWord  string
	CurrentWord string
}
