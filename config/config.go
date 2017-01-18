package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
	"encoding/json"
)

func parseKafkaConfig(config interface{}, t *KafkaConfig) {
	d, _ := json.Marshal(config)
	json.Unmarshal(d, t)
	log.Println(t)

    /*
    for key, value := range kafka_config {
		key_str, ok := key.(string)
		if !ok {
			log.Printf("Invalid configuration k->: %s\n", key)
        } else {
			log.Println(key_str)
			t.key_str = value
    	}
		
        //log.Printf("%s = %s", key, value)
		log.Println(value)
    }
    */
}


func ReadConfig(config_path string) (map[string]interface{}) {
    var config map[string]interface{}
    
    reader, err := os.Open(config_path)
    if err != nil {
        log.Fatal("Failed to open %s", config_path)
		panic("err to load config fiel")
    }
    defer reader.Close()
	
	content, err := ioutil.ReadAll(reader)

	err1 := yaml.Unmarshal(content, &config)
    if err1 != nil {
        log.Fatalf("error: %v", err1)
    }

    input_array, ok := config["input"].([]interface{})
    if !ok {
        log.Println("type:", reflect.TypeOf(config["input"]))
    } 

    for _, value := range input_array {
        input_item, ok := value.(map[interface{}]interface{})
        if !ok {
            log.Fatalf("parse input errer %s", value)
        }
        for k, v := range input_item {
            switch k{
            case "kafka":
				var t KafkaConfig
				parseKafkaConfig(v, &t)
            }
        }

    }

	return config
}

type KafkaConfig struct {
    GroupID int `json:"group_id"`
    Topic []string `json:"topic"`
    Tag []string `json:"tag"`
    ZkHosts []string `json:"zk_hosts"`
    Type string `json:"type"`
}

func SetField(obj interface{}, name string, value interface{}) error {
    structValue := reflect.ValueOf(obj).Elem()
    structFieldValue := structValue.FieldByName(name)

    if !structFieldValue.IsValid() {
        return fmt.Errorf("No such field: %s in obj", name)
    }

    if !structFieldValue.CanSet() {
        return fmt.Errorf("Cannot set %s field value", name)
    }

    structFieldType := structFieldValue.Type()
    val := reflect.ValueOf(value)
    if structFieldType != val.Type() {
        return errors.New("Provided value type didn't match obj field type")
    }

    structFieldValue.Set(val)
    return nil
}


func (s *KafkaConfig) FillStruct(m map[interface{}]interface{}) error {
    for k, v := range m {
		key, _ := k.(string)
        err := SetField(s, key, v)
        if err != nil {
            return err
        }
    }
    return nil
}

