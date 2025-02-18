package reset

import (
	"github.com/deviceinsight/kafkactl/internal/consumergroupoffsets"
	"github.com/deviceinsight/kafkactl/internal/consumergroups"
	"github.com/deviceinsight/kafkactl/internal/k8s"
	"github.com/deviceinsight/kafkactl/output"
	"github.com/spf13/cobra"
)

func newResetOffsetCmd() *cobra.Command {

	var offsetFlags consumergroupoffsets.ResetConsumerGroupOffsetFlags

	var cmdResetOffset = &cobra.Command{
		Use:     "consumer-group-offset GROUP",
		Aliases: []string{"cgo", "offset"},
		Short:   "reset a consumer group offset",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if !(&k8s.Operation{}).TryRun(cmd, args) {
				if err := (&consumergroupoffsets.ConsumerGroupOffsetOperation{}).ResetConsumerGroupOffset(offsetFlags, args[0]); err != nil {
					output.Fail(err)
				}
			}
		},
		ValidArgsFunction: consumergroups.CompleteConsumerGroups,
	}

	cmdResetOffset.Flags().BoolVarP(&offsetFlags.OldestOffset, "oldest", "", false, "set the offset to oldest offset (for all partitions or the specified partition)")
	cmdResetOffset.Flags().BoolVarP(&offsetFlags.NewestOffset, "newest", "", false, "set the offset to newest offset (for all partitions or the specified partition)")
	cmdResetOffset.Flags().BoolVarP(&offsetFlags.AllTopics, "all-topics", "", false, "do the operation for all topics in the consumer group")
	cmdResetOffset.Flags().Int64VarP(&offsetFlags.Offset, "offset", "", -1, "set offset to this value. offset with value -1 is ignored")
	cmdResetOffset.Flags().Int32VarP(&offsetFlags.Partition, "partition", "p", -1, "partition to apply the offset. -1 stands for all partitions")
	cmdResetOffset.Flags().StringArrayVarP(&offsetFlags.Topic, "topic", "t", offsetFlags.Topic, "one ore more topics to change offset for")
	cmdResetOffset.Flags().BoolVarP(&offsetFlags.Execute, "execute", "e", false, "execute the reset (as default only the results are displayed for validation)")
	cmdResetOffset.Flags().StringVarP(&offsetFlags.OutputFormat, "output", "o", offsetFlags.OutputFormat, "output format. One of: json|yaml")

	return cmdResetOffset
}
