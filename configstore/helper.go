package configstore

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	allConfigs = "config"
	configId   = "config/%s"
	config     = "config/%s/%s"

	allGroups = "group"
	groupId   = "group/%s"
	groupVer  = "group/%s/%s"
	group     = "group/%s/%s/%s"
)

func generateConfigKey(ver string) (string, string) {
	id := uuid.New().String()
	return fmt.Sprintf(config, id, ver), id
}

func constructConfigKey(id string, ver string) string {
	return fmt.Sprintf(config, id, ver)
}

func constructConfigIdKey(id string) string {
	return fmt.Sprintf(configId, id)
}

func generateGroupKey(ver string) (string, string) {
	id := uuid.New().String()
	return fmt.Sprintf(groupVer, id, ver), id
}

func constructGroupKey(id string, ver string) string {
	return fmt.Sprintf(groupVer, id, ver)
}

func constructGroupIdKey(id string) string {
	return fmt.Sprintf(groupId, id)
}
