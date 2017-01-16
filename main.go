package main

import (
	config "github.com/gostash/config"
	"flag"
	"log"
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
	var config_path string
    flag.StringVar(&config_path, "config", "./conf/gostash.yml", "Load the gostash config from a specific file")
	flag.StringVar(&config_path, "f", "./conf/gostash.yml", "Load the gostash config from a specific file")
    flag.Parse()

	gostash_config := config.ReadConfig(config_path)
	
	log.Println(gostash_config)
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