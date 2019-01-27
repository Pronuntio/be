package word

import (
	"database/sql"

	"go.uber.org/zap"
)

type PostgresWordDao struct {
	conn *sql.DB
	log  *zap.Logger
}

func NewPostgresWordDao(conn *sql.DB, log *zap.Logger) *PostgresWordDao {
	return &PostgresWordDao{
		conn: conn,
		log:  log,
	}
}

func (p *PostgresWordDao) GetWord(ID uint64) (*Word, error) {
	rows, err := p.conn.Query("SELECT text_original, text_english, status, filename FROM pronuntio.words")
	if err != nil {
		p.log.Error("unable to get word", zap.Uint64("id", ID), zap.Error(err))
		return nil, err
	}
	rows.Next()

	w := &Word{
		ID: ID,
	}

	err = rows.Scan(&w.NativeName, &w.EnglishName, &w.Status, &w.Filename)
	if err != nil {
		p.log.Error("unable to scan values", zap.Uint64("id", ID), zap.Error(err))
		return nil, err
	}

	return w, nil
}

func (p *PostgresWordDao) DeleteWord(ID uint64) error {
	_, err := p.conn.Exec("DELETE FROM pronuntio.words WHERE ID = $1", ID)
	return err
}

func (p *PostgresWordDao) CreateWord(word *Word) (uint64, error) {
	res, err := p.conn.Exec("INSERT INTO pronuntio.words (text_original, text_english, status, filename) VALUES ($1, $2, $3, $4)", word.NativeName, word.EnglishName, word.Status, word.Filename)
	if err != nil {
		p.log.Error("unable to create word", zap.Error(err))
		return 0, err
	}
	wId, err := res.LastInsertId()
	return uint64(wId), err
}

func (p *PostgresWordDao) UpdateWord(word *Word) error {
	return nil
}
