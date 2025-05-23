package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_719629581")
		if err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"hidden": false,
			"id": "select1542800728",
			"maxSelect": 1,
			"name": "field",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"Neque porro quisquam est qui dolorem ipsum quia"
			]
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_719629581")
		if err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"hidden": false,
			"id": "select1542800728",
			"maxSelect": 1,
			"name": "field",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"hi"
			]
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
