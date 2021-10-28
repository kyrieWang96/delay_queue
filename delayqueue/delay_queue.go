package delayqueue

import (
	"github.com/yasin-wu/delay_queue/cronjob"
	"github.com/yasin-wu/delay_queue/logger"
	"github.com/yasin-wu/utils/redis"
)

type DelayQueue struct {
	logger             logger.Logger
	scheduler          *cronjob.Scheduler
	jobExecutorFactory map[string]*jobExecutor
	redisCli           *redisClient
	redisConf          *redis.Config
}

func New(conf *Config) *DelayQueue {
	initDelayQueue(conf)
	return delayQueue
}

func (dq *DelayQueue) StartBackground() {
	dq.scheduler.Start()
}

func (dq *DelayQueue) Register(action JobBaseAction) error {
	jobID := action.ID()
	if _, ok := dq.jobExecutorFactory[jobID]; ok {
		return ErrorsDelayQueueRegisterIDDuplicate
	} else {
		dq.jobExecutorFactory[jobID] = &jobExecutor{
			ID:     jobID,
			action: action,
		}
	}
	return nil
}

func (dq *DelayQueue) AddJob(job DelayJob) error {
	return dq.redisCli.ZAdd(job)
}

func (dq *DelayQueue) SetLogger(logger logger.Logger) {
	dq.logger = logger
}

func (dq *DelayQueue) availableJobIDs() []string {
	var IDs []string
	for k := range dq.jobExecutorFactory {
		IDs = append(IDs, k)
	}
	return IDs
}
