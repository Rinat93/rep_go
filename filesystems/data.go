package filesystems

import (
	"encoding/json"
	"reg_go/exceptions"
)

type serializers interface {}

func (c *DataBody) Add(body string)  {
	c.body = body
	c.len = len(c.body)
	if c.serialize != nil {
		err := c.Serialize(c.serialize)
		if err != nil {
			exceptions.ErrorFy(err)
		}
	}
}

// Сериализуем данные файла
func (c *DataBody) Serialize(ser serializers) error {
	var data serializers = ser
	err := json.Unmarshal([]byte(c.body), &data)
	if err != nil {
		return err
	}
	c.serialize = data
	return nil
}