package neo4j

import "encoding/json"

type Cypher struct {
	Query   map[string]string
	Payload interface{}
}

type CypherResponse struct {
	Columns map[string]interface{} `json:"columns"`
	Data    map[string]interface{} `json:"data"`
}

func (c *Cypher) mapBatchResponse(neo4j *Neo4j, data interface{}) (bool, error) {
	encodedData, err := jsonEncode(data)
	err = c.decodeResponse(encodedData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Cypher) getBatchQuery(operation string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"method": "POST",
		"to":     "/cypher",
		"body":   c.Query,
	}, nil
}

func (c *Cypher) decodeResponse(data string) error {
	resp := map[string]interface{}{}

	err := json.Unmarshal([]byte(data), &resp)
	if err != nil {
		return err
	}

	columnData := resp["data"].([]interface{})
	jsonizedData, err := json.Marshal(columnData)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonizedData, &c.Payload)
	if err != nil {
		return err
	}

	return nil
}