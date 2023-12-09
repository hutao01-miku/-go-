package tools

import (
	"fmt"
	"github.com/google/uuid"
)

func GetUUID() string {
	id := uuid.New() // 默认V4版本,基于一个随机数的。
	fmt.Printf("uuid: %s, version: %s\n", id.String(), id.Version().String())
	return id.String()
}
