package indexzincsearch

func CreateIndex() {
	url := "http://localhost:4080/api/index"

	indexConfig := []byte(`
	{
		"name": "email",
		"storage_type": "disk",
		"shard_num": 1,
		"mappings": {
			"properties": {
				"id": {
					"type": "keyword",
					"index": true,
					"store": true,
					"sortable": true
				},
				"from": {
					"type": "text",
					"index": true,
					"store": true
				},
				"to": {
					"type": "text",
					"index": true,
					"store": true
				},
				"subject": {
					"type": "text",
					"index": true,
					"store": true
				},
				"content": {
					"type": "text",
					"index": true,
					"store": true,
					"highlightable": true
				},
				"date": {
					"type": "date",
					"index": true,
					"store": true,
					"highlightable": true,
					"sortable": true
				}
			}
		},
		"settings": {
			"analysis": {
				"analyzer": {
					"default": {
						"type": "standard"
					}
				}
			}
		}
	}
	`)

	HttpPOST(url, string(indexConfig))
}
