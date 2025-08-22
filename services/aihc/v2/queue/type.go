package queue

import (
	v2 "github.com/baidubce/bce-sdk-go/services/aihc/v2/api"
)

type Interface interface {
	DescribeResourceQueue(queueID string) (*v2.DescribeResourceQueueResponse, error)
	DescribeResourceQueues(describeResourceQueuesRequest *v2.DescribeResourceQueuesRequest) (*v2.DescribeResourceQueuesResponse, error)
}
