package word

type Dao interface {
	GetWord(ID uint64) (*Word, error)
	DeleteWord(ID uint64) error
	CreateWord(word *Word) (uint64, error)
	UpdateWord(word *Word) error
	ListWords() ([]*Word, error)
}
