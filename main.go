package main

import (
	viper "github.com/spf13/viper"
	"flag"
	"log"
	"os"
	"reflect"
	/*
	"github.com/Shopify/sarama"
	"gopkg.in/olivere/elastic.v3"
	"fmt"
	"encoding/json"
	"os"
	"os/signal"
    */
)


func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	
	var config_path string
    flag.StringVar(&config_path, "config", "./conf/gostash.yml", "Load the gostash config from a specific file")
	flag.StringVar(&config_path, "f", "./conf/gostash.yml", "Load the gostash config from a specific file")
    flag.Parse()

	reader, err := os.Open(config_path)
    if err != nil {
        log.Fatal("Failed to open %s", config_path)
        panic("err to load config fiel")
    }
    defer reader.Close()
	//设置配置文件ymal
	viper.SetConfigType("yaml")
	
	viper.ReadConfig(reader)

	inputSlice := viper.Get("input")
	v := reflect.ValueOf(inputSlice)

    if v.Kind() == reflect.Slice {
		num := v.Len()
        for i := 0; i < num; i++ {
            inputItem := v.Index(i).Interface()
			inputItemType := reflect.ValueOf(inputItem)
			log.Println(inputItemType.Kind())
        }
    }
    
	
	log.Println(v)
	
	log.Println(viper.GetString("input.kafka.topic"))

	
	
	
	/*
	config := sarama.NewConfig()
	config.ClientID = "go-client"
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		panic(err)
	}
	
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Ping the Elasticsearch server to get e.g. the version number
	
    client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:8200"))
    if err != nil {
        panic(err)
	}

	info, code, err := client.Ping("http://127.0.0.1:8200").Do()
    if err != nil {
        // Handle error
        panic(err)
    }
    fmt.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)


    
    partitionConsumer, err := consumer.ConsumePartition("test_biwebtrends", 0, 0)
	if err != nil {
		panic(err)
	}
	
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	exists, err := client.IndexExists("test").Do()
    if err != nil {
        panic(err)
    }
    if !exists {
		createIndex, err := client.CreateIndex("test").Do()
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged  {
			log.Printf("create %s index faild!" , "test")
			panic("error, create index faild!")
		}
	}
	
	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	
	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var message interface{}
			err := json.Unmarshal(msg.Value, &message)
			if err != nil {
				panic(err)
			}
			_, err = client.Index().
                Index("test").
                Type("test").
                BodyJson(message).
                Do()
            if err != nil {
                panic(err)
            }
            
			_, err = client.Index().
				Index("test").
				Type("test").
				BodyString(string(msg.Value)).
				Do()
			if err != nil {
				panic(err)
			}
            //log.Printf(string(msg.Value))
			log.Printf("Consumed message offset %d\n", msg.Offset)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}
	
	log.Printf("Consumed: %d\n", consumed)
*/
}
