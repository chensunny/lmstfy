package engine

import (
	"fmt"
	"strings"
)

type Queue interface {
	Name() string
	Tag() string
	String() string
}

type QueueImpl struct {
	name string
	tag  string
}

func NewQueue(queue string) Queue {
	if idx := strings.Index(queue, ".{"); idx > 0 {
		return &QueueImpl{
			name: queue[:idx],
			tag:  strings.TrimSuffix(queue[idx+2:], "}"),
		}
	}
	return &QueueImpl{
		name: queue,
		tag:  "",
	}
}

func (q QueueImpl) Name() string {
	return q.name
}

func (q QueueImpl) Tag() string {
	return q.tag
}

func (q QueueImpl) String() string {
	return fmt.Sprintf("%s.{%s}", q.Name, q.Tag)
}
