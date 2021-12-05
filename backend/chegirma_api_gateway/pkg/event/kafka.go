package event

import (
	"context"

	"chegirma_api_gateway/config"
	"chegirma_api_gateway/pkg/logger"

	"github.com/Shopify/sarama"
)

type Kafka struct {
	ctx        context.Context
	log        logger.Logger
	cfg        config.Config
	publishers map[string]*Publisher
	// consumers    map[string]*Consumer
	saramaConfig *sarama.Config
}

func NewKafka(ctx context.Context, cfg config.Config, log logger.Logger) *Kafka {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0
	// saramaConfig.Consumer.Group.Heartbeat.Interval = time.Second * 20
	// saramaConfig.Consumer.Group.Session.Timeout = time.Minute

	kafka := &Kafka{
		ctx:        ctx,
		log:        log,
		cfg:        cfg,
		publishers: make(map[string]*Publisher),
		// consumers:    make(map[string]*Consumer),
		saramaConfig: saramaConfig,
	}

	kafka.RegisterPublishers()

	return kafka
}

func (k *Kafka) RegisterPublishers() {
	k.AddPublisher("v1.ek_analytic_service.action_history.create")
}
