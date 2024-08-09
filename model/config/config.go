package config

type Config struct {
	Datasource Datasource
	Discordbot Discordbot
}

type Datasource struct {
	Postgres Postgres
	Redis    Redis
}

type Postgres struct {
	Uri string
}

type Redis struct {
	Uri string
}

type Discordbot struct {
	Token string
}
