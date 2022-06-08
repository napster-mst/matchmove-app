package seeder

import (
	"log"
	"strings"
	"time"

	"github.com/go-pg/pg/orm"
	"golang.org/x/crypto/bcrypt"
	"matchmove.com/demo/common/alias"
	"matchmove.com/demo/ent"
)

func Seed(db orm.DB) {
	seedRole(db)
	seedStatus(db)
	seedAdmin(db)
}

func seedRole(db orm.DB) {
	roles := []*ent.Roles{
		{
			Name:      string(alias.UserRoleAdmin),
			CreatedAt: time.Now(),
		},
		{
			Name:      string(alias.UserRoleUser),
			CreatedAt: time.Now(),
		},
	}
	_, err := db.Model(&roles).Insert()
	if err != nil && !strings.Contains(err.Error(), "duplicate key value") {
		log.Fatal(err)
	}
}

func seedStatus(db orm.DB) {
	_, err := db.Exec(`INSERT INTO public.status(name)
		VALUES (?), (?)`, alias.TokenStatusActive, alias.TokenStatusInactive)
	if err != nil && !strings.Contains(err.Error(), "duplicate key value") {
		log.Fatal(err)
	}
}

func seedAdmin(db orm.DB) {
	password := "adminpassword"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	var adminRole ent.Roles

	db.Model(&adminRole).Where("name = ?", alias.UserRoleAdmin).Select()

	admin := ent.User{
		Username: "admin",
		Password: string(hash),
		RoleID:   adminRole.ID,
	}

	_, err = db.Model(&admin).Insert()
	if err != nil && !strings.Contains(err.Error(), "duplicate key value") {
		log.Fatal(err)
	}
}
