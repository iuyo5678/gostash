package config

type KafkaConfig struct {
	group_id string
	topic string
	tag []string
    zk_connect []string
	infoType string
}
