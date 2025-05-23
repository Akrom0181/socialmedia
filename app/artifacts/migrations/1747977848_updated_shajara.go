package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2876785954")
		if err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"hidden": false,
			"id": "select1561212630",
			"maxSelect": 1,
			"name": "shajara_role",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "select",
			"values": [
				"ota",
				"ona"
			]
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2876785954")
		if err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"hidden": false,
			"id": "select1561212630",
			"maxSelect": 1,
			"name": "shajara_role",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"ota",
				"ona"
			]
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
