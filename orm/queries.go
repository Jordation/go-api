package orm

const (
	PlayersList         = "list"
	PlayersListDistinct = "listDistinct"
	PlayersListMAX      = "listMax"
	PlayersListAVG      = "listAVG"
)

func GetStatQueries() map[string]string {
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
			(?)
			FROM players
		`,
		PlayersListAVG: `
		SELECT AVG
		(?)
		FROM players
	`,
	}
}
