package word

type Word struct {
	ID          uint64
	EnglishName string
	NativeName  string
	Filename    string
	Status      string
	payload     []byte
}
