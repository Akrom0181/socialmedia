package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2576396437")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_nA4o0ropd8` + "`" + ` ON ` + "`" + `shajara_person` + "`" + ` (` + "`" + `person` + "`" + `)"
			]
		}`), &collection); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": true,
			"collectionId": "_pb_users_auth_",
			"hidden": false,
			"id": "relation886886774",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "person",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"hidden": false,
			"id": "select2585007050",
			"maxSelect": 1,
			"name": "ralative",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "select",
			"values": [
				"ota",
				"ona",
				"bobo",
				"buvi",
				"aka",
				"uka",
				"singil",
				"opa",
				"amaki",
				"toga"
			]
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2576396437")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": []
		}`), &collection); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(1, []byte(`{
			"cascadeDelete": false,
			"collectionId": "_pb_users_auth_",
			"hidden": false,
			"id": "relation886886774",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "person",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"hidden": false,
			"id": "select2585007050",
			"maxSelect": 1,
			"name": "ralative",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"ota",
				"ona",
				"bobo",
				"buvi",
				"aka",
				"uka",
				"singil",
				"opa",
				"amaki",
				"toga"
			]
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
