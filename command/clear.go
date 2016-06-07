package command

// Imports.
import "github.com/codegangsta/cli"
import "github.com/chenzhe07/snap/action"
import "log"

// Command.
var Clear = cli.Command{
	Name:      "clear",
	ShortName: "cl",
	Usage:     "<database>",
	Description: `Clear specified database's all revision. return empty when use list

EXAMPLE:

    snap clear <database>
`,

	Action: func(ctx *cli.Context) error {
		args := ctx.Args()

		if len(args) > 0 {
			database := args.Get(0)
			action.ClearDatabases(database)
			return nil
		}

		log.Println("Not enough arguments supplied.")
		log.Fatalf("Run '%s help clear' for more information.\n", ctx.App.Name)
		return nil
	},
}
