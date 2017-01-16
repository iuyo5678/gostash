package config

type KafkaConfig struct {
	group_id string
	topic string
	log_type string
    zk_connect []string
}
