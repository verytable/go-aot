package mikasa

import (
	_ "github.com/go-sql-driver/mysql"
	_ "go.opencensus.io/tag"
	_ "golang.org/x/crypto/acme"
)

func Abilities() []string {
	return []string{"Vertical Maneuvering Equipment", "Usage of Thunder Spears", "Strength"}
}
