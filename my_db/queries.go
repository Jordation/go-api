package my_db

const (
	PlayersList         = "list"
	PlayersListDistinct = "listDistinct"
	PlayersListMAX      = "listMax"
	PlayersListAVG      = "listAVG"
)

func GetPlayerQueries() map[string]string {
	return map[string]string{
		PlayersList: `
			SELECT
				*
			FROM players
		`,
		PlayersListDistinct: `
			SELECT DISTINCT
			?
			FROM players
		`,
		PlayersListMAX: `
			SELECT MAX
			?
			FROM players
		`,
		PlayersListAVG: `
		SELECT AVG
		?
		FROM players
	`,
	}
}
