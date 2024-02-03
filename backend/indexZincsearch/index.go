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
					"type": "text",
					"index": false,
					"store": false
				},
				"from": {
					"type": "text",
					"index": false,
					"store": false
				},
				"to": {
					"type": "text",
					"index": false,
					"store": false
				},
				"subject": {
					"type": "text",
					"index": true,
					"store": false
				},
				"content": {
					"type": "text",
					"index": true,
					"store": false,
					"highlightable": true
				},
				"date": {
					"type": "text",
					"index": true,
					"store": false,
					"highlightable": true
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

	httpPOST(url, indexConfig)
}
