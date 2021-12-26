package problems

type (
	// Handler contains reference to whatever I need across problems
	Handler struct {
		Env map[string]string
		// IGDB     *igdb.Client
		// Postgres *sql.DB
		// Trie     *patricia.Trie
	}
)
