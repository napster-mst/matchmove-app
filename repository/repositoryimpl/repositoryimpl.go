package repositoryimpl

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"matchmove.com/demo/common/alias"
	"matchmove.com/demo/dtos"
	"matchmove.com/demo/ent"
	"matchmove.com/demo/repository"
)

type repositoryImpl struct {
	db orm.DB
}

func NewRepository(db orm.DB) repository.Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) InsertToken(token *ent.Token) error {
	_, err := r.db.Model(token).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (r *repositoryImpl) GetAdminCredentials(username string) (user ent.User, err error) {
	err = r.db.Model(&user).Where("username = ?", username).Select()
	if err != nil && err != pg.ErrNoRows {
		return
	}

	return user, nil
}

func (r *repositoryImpl) GetStatus(status alias.TokenStatus) int {
	s := ent.Status{}
	r.db.QueryOne(&s, `Select * FROM status WHERE name = ?`, string(status))

	return s.ID
}

func (r *repositoryImpl) CheckUserExist(id string) (exists bool, err error) {
	_, err = r.db.QueryOne(pg.Scan(&exists), `SELECT exists (SELECT 1 FROM users WHERE id = ?)`, id)

	return
}

func (r *repositoryImpl) GetToken(token string) (t ent.Token, err error) {
	err = r.db.Model(&t).Where("token = ?", token).Limit(1).Select()

	return
}

func (r *repositoryImpl) CountTokenStatus() (tsr dtos.TokenStatusResponse, err error) {
	query := "SELECT status_id,COUNT(*) AS \"status_count\" FROM tokens GROUP BY status_id"

	type result struct {
		StatusID    int
		StatusCount int
	}

	var resultRows []result
	_, err = r.db.Query(&resultRows, query)
	if err != nil {
		return
	}

	for _, r2 := range resultRows {
		if r2.StatusID == 1 {
			tsr.ActiveTokens = r2.StatusCount
		} else {
			tsr.InActiveTokens = r2.StatusCount
		}
	}

	return
}

func (r *repositoryImpl) DisableAllUserTokens() (err error) {
	query := `update tokens set status_id = ? 
				where user_id NOT IN (SELECT id FROM users WHERE role_id = (SELECT id FROM roles WHERE name = 'user'))`

	_, err = r.db.Exec(query, r.GetStatus(alias.TokenStatusActive))
	if err != nil && err != pg.ErrNoRows {
		return
	}

	return
}
