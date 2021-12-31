package mapping

const (
	Mapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text"
				},
				"retweets":{
					"type":"integer"
				},
				"created":{
					"type":"date"
				},
				"attributes":{
					"type":"object"
				}
			}
	}
}`
)
