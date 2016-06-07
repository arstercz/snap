package action

// Imports.
import "github.com/nomad-software/snap/database"
import "log"

func exitOnerror(err error, databaseName string) {
	database.Rollback()
	log.Println(err)
	log.Fatalf("Error occurred while retrieving database '%s'", databaseName)
}

// List all managed databases.
func ClearDatabases(databaseName string) {

	database.AssertConfigDatabaseExists()
	database.AssertDatabaseExists(databaseName)
	//use snap config database
	database.AssertUseConfigDatabase()

	query := `SELECT id.id
		FROM initialisedDatabases AS id
		WHERE id.name = ?
		LIMIT 1;`
	databaseIdInfo, err := database.QueryRow(query, databaseName)
	if err != nil {
		exitOnerror(err, databaseName)
	}

	//get database id
	if len(databaseIdInfo) == 0 {
		log.Fatalf("Can not find the '%s' database id", databaseName)
	}
	databaseId := databaseIdInfo.Str(0)

	//As the foreign key, delete id is eque to delete all revisions on this database.
	deleteQuery := `DELETE FROM
		initialisedDatabases WHERE id = ?
		LIMIT 1;
	`
	err = database.Exec(deleteQuery, databaseId)
	if err != nil {
		exitOnerror(err, databaseName)
	} else {
		log.Printf("clear revisions on '%s' ok\n", databaseName)
	}

}
