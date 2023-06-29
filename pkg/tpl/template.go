package tpl

import (
	_ "embed"
)

//go:embed docker-compose.yml
var DockerCompose []byte

//go:embed meta.yml
var Meta []byte

//go:embed config.yml
var Config []byte

//go:embed flag.sh
var Flag []byte

//go:embed start.sh
var Start []byte

//go:embed db.sql
var DB_SQL []byte

//go:embed db.json
var DB_JSON []byte

//go:embed README.md
var Readme []byte
