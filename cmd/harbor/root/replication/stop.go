// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package replication

import (
	"fmt"
	"strconv"

	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/goharbor/harbor-cli/pkg/prompt"
	"github.com/goharbor/harbor-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func StopCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "stop replication",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Debug("Stopping replication")

			var rpolicyID int64
			var executionID int64
			if len(args) > 0 {
				var err error
				// convert string to int64
				rpolicyID, err = strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					return fmt.Errorf("invalid replication policy ID: %s, %v", args[0], err)
				}
				executionID = prompt.GetReplicationExecutionIDFromUser(rpolicyID)
			} else {
				rpolicyID = prompt.GetReplicationPolicyFromUser()
				executionID = prompt.GetReplicationExecutionIDFromUser(rpolicyID)
			}

			execution, err := api.GetReplicationExecution(executionID)
			if err != nil {
				return fmt.Errorf("failed to get replication execution: %v", utils.ParseHarborErrorMsg(err))
			}
			if execution.Payload.Status != "InProgress" {
				return fmt.Errorf("replication execution with ID: %d is already stopped, succeed or failed", executionID)
			}

			_, err = api.StopReplication(executionID)
			if err != nil {
				return fmt.Errorf("failed to stop replication: %v", utils.ParseHarborErrorMsg(err))
			}
			fmt.Printf("Replication execution with ID: %d stopped successfully\n", executionID)
			return nil
		},
	}

	return cmd
}
