package configstore

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/consul/api"
)

type ConfigStore struct {
	cli *api.Client
}

func New() (*ConfigStore, error) {
	db := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", db, dbport)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ConfigStore{
		cli: client,
	}, nil
}

func (cs *ConfigStore) FindConf(id string, ver string) (*Config, error) {
	kv := cs.cli.KV()
	key := constructConfigKey(id, ver)
	data, _, err := kv.Get(key, nil)

	if err != nil || data == nil {
		return nil, errors.New("That item does not exist!")
	}

	config := &Config{}
	err = json.Unmarshal(data.Value, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (cs *ConfigStore) DeleteConfig(id, ver string) (map[string]string, error) {
	kv := cs.cli.KV()
	_, err := kv.Delete(constructConfigKey(id, ver), nil)
	if err != nil {
		return nil, err
	}

	return map[string]string{"Deleted": id + ver}, nil
}

func (cs *ConfigStore) FindConfVersions(id string) ([]*Config, error) {
	kv := cs.cli.KV()

	key := constructConfigIdKey(id)
	data, _, err := kv.List(key, nil)
	if err != nil {
		return nil, err
	}

	configs := []*Config{}

	for _, pair := range data {
		config := &Config{}
		err := json.Unmarshal(pair.Value, config)
		if err != nil {
			return nil, err
		}

		configs = append(configs, config)
	}

	return configs, nil
}

func (cs *ConfigStore) CreateConfig(config *Config) (*Config, error) {
	kv := cs.cli.KV()

	sid, rid := generateConfigKey(config.Version)
	config.ID = rid

	data, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	c := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(c, nil)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (cs *ConfigStore) UpdateConfigVersion(config *Config) (*Config, error) {
	kv := cs.cli.KV()

	data, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	_, err = cs.FindConf(config.ID, config.Version)

	//Does exist
	if err == nil {
		return nil, errors.New("Given config version already exists! ")
	}

	c := &api.KVPair{Key: constructConfigKey(config.ID, config.Version), Value: data}
	_, err = kv.Put(c, nil)
	if err != nil {
		return nil, err
	}
	return config, nil

}
