package pulsar

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/sirupsen/logrus"
	"performance_producer_go/conf"
	"performance_producer_go/util"
)

func Start() error {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: fmt.Sprintf("pulsar://%s:%d", conf.PulsarHost, conf.PulsarPort),
	})
	if err != nil {
		return err
	}
	for i := 0; i < conf.RoutineNum; i++ {
		go startProducer(client)
	}
	return nil
}

func startProducer(client pulsar.Client) {
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: conf.PulsarTopic,
	})
	if err != nil {
		logrus.Errorf("create producer %s error: %v", conf.PulsarTopic, err)
	}
	for {
		messageID, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte(util.RandStr(conf.PulsarMessageSize)),
		})
		if err != nil {
			logrus.Errorf("send message %s error: %v", conf.PulsarTopic, err)
		} else {
			logrus.Infof("send message %s success, messageID: %s", conf.PulsarTopic, messageID)
		}
	}
}
