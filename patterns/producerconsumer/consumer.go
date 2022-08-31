package producerconsumer

import (
	"github.com/luaxlou/gomessagequeue/mqengines"
)

type Consumer struct {
	engine mqengines.MqEngine
}

func NewConsumer(engine mqengines.MqEngine) *Consumer {

	return &Consumer{
		engine: engine,
	}
}

func (c *Consumer) ConsumeBlock(key string, count int64, onRead func(contents []string) error) {

	c.engine.ReadBlock(key, "single", count, onRead)
}

func (c *Consumer) Consume(key string, count int64, onRead func(contents []string) error) {

	c.engine.Read(key, "single", count, onRead)
}

func (c *Consumer) ConsumeBlockWithGroup(key string, group string, count int64, onRead func(contents []string) error) {

	c.engine.ReadBlock(key, group, count, onRead)
}

func (c *Consumer) ConsumeWithGroup(key string, group string, count int64, onRead func(contents []string) error) {

	c.engine.Read(key, group, count, onRead)
}
