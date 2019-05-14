package publisher

import (
	"context"

	"github.com/chef/automate/components/compliance-service/api/common"
	"github.com/chef/automate/components/ingest-service/pipeline/message"
	"github.com/chef/automate/components/nodemanager-service/api/manager"
	"github.com/chef/automate/components/nodemanager-service/api/nodes"
	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
)

func BuildNodeManagerPublisher(nodeManagerClient manager.NodeManagerServiceClient) message.ChefRunPipe {
	return func(in <-chan message.ChefRun) <-chan message.ChefRun {
		return nodeManagerPublisher(in, nodeManagerClient)
	}
}

func nodeManagerPublisher(in <-chan message.ChefRun, nodeManagerClient manager.NodeManagerServiceClient) <-chan message.ChefRun {
	ctx := context.Background()
	maxNumberOfBundledMsgs := 100
	out := make(chan message.ChefRun, maxNumberOfBundledMsgs)
	go func() {
		for msg := range in {
			// send to node manager from here.
			log.Infof("send info about node %s to node manager", msg.Node.NodeName)

			// convert node check in time to proto timestamp
			timestamp, err := ptypes.TimestampProto(msg.Node.Checkin)
			if err != nil {
				log.Errorf("unable to translate time %v to timestamp. aborting call to NodeManager.ProcessNode for node %s", msg.Node.Checkin, msg.Node.NodeName)
				out <- msg
				continue
			}

			// translate status
			status := nodes.LastContactData_UNKNOWN
			switch msg.Node.Status {
			case "success":
				status = nodes.LastContactData_PASSED
			case "failure":
				status = nodes.LastContactData_FAILED
			}

			// translate tags
			tags := make([]*common.Kv, len(msg.Node.ChefTags))
			for i, tag := range msg.Node.ChefTags {
				tags[i] = &common.Kv{
					Key:   "chef-tag",
					Value: tag,
				}
			}

			_, err = nodeManagerClient.ProcessNode(ctx, &manager.NodeMetadata{
				Uuid:            msg.Node.EntityUuid,
				Name:            msg.Node.NodeName,
				PlatformName:    msg.Node.Platform,
				PlatformRelease: msg.Node.PlatformVersion,
				LastContact:     timestamp,
				SourceId:        msg.Node.Ec2.InstanceId,
				SourceRegion:    msg.Node.Ec2.PlacementAvailabilityZone,
				Tags:            tags,
				RunData: &nodes.LastContactData{
					Id:      msg.Node.LatestRunID,
					EndTime: timestamp,
					Status:  status,
				},
			})
			if err != nil {
				log.Errorf("unable to send info about node %s to node manager", msg.Node.NodeName)
			}

			out <- msg
		}
		close(out)
	}()

	return out
}
