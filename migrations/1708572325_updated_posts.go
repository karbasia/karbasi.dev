package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("zwjghnxbt3d0awt")
		if err != nil {
			return err
		}

		// add
		new_attachments := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "0wqmgchj",
			"name": "attachments",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [],
				"thumbs": [],
				"maxSelect": 99,
				"maxSize": 52428800,
				"protected": false
			}
		}`), new_attachments)
		collection.Schema.AddField(new_attachments)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("zwjghnxbt3d0awt")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("0wqmgchj")

		return dao.SaveCollection(collection)
	})
}
