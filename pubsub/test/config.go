package main

import (
	"errors"
	"math/rand"
	"time"
)

// instance config
type InstanceConfig struct {
	id        int           // the id of instance
	latency   time.Duration // the latency of instance
	publisher bool          // whether nodes of this instance are publishing or not
}

// pubsub message log
type PubsubMessageLog struct {
	sender   string        // sender peer id
	receiver string        // receiver peer id
	sendTime int           // message send time (unix)
	recvTime int           // message receive time (unix)
	elapsed  time.Duration // total time elapsed for packet travel (transport time)
	message  string        // the message itself (currently denotes send time)
}

// Config of each instance mapped with global id
var Instances map[int]InstanceConfig

// min and max latency in ms
var MIN_LATENCY = 100
var MAX_LATENCY = 500

// max number of configs
var MAX_SIZE = 0

// Map for logging pubsub topic messages
var MessageLog map[string]map[string]PubsubMessageLog

func configureInstanceLatency(maxSize int) {

	MAX_SIZE = maxSize

	// initialise a new instances map
	Instances = make(map[int]InstanceConfig, maxSize)

	// initialise a new message log map
	MessageLog = make(map[string]map[string]PubsubMessageLog)

	rand.Seed(time.Now().UnixNano())

	// fill with random latency
	for i := 0; i < maxSize; i++ {
		latency := time.Millisecond * time.Duration((MIN_LATENCY + rand.Intn(MAX_LATENCY-MIN_LATENCY+1)))
		publisher := rand.Float32() < 0.5
		config := InstanceConfig{latency: latency, publisher: publisher, id: i}
		Instances[i] = config
	}
}

func getInstanceConfig(id int) (*InstanceConfig, error) {
	if id >= MAX_SIZE {
		err := errors.New("Instance id out of range")
		return nil, err
	}
	instance := Instances[id]
	return &(instance), nil
}

func postMessageLog(topic, peer, seqNo string, log PubsubMessageLog) {
	MessageLog[topic][peer+seqNo] = log
}

func getMessageLogByTopic(topic string) map[string]PubsubMessageLog {
	return MessageLog[topic]
}
