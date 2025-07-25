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
package execution

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/goharbor/harbor-cli/pkg/views/base/selection"
)

func ReplicationExecutionList(executions []*models.ReplicationExecution, choice chan<- int64) {
	itemsList := make([]list.Item, len(executions))
	for i, p := range executions {
		displayName := fmt.Sprintf("ID: %d, Status: %s, Trigger: %s, Start Time: %s, Succeed: %d, Total: %d",
			p.ID, p.Status, p.Trigger, p.StartTime.String(), p.Succeed, p.Total)
		itemsList[i] = selection.Item(displayName)
	}

	m := selection.NewModel(itemsList, "Select a Replication Execution")

	p, err := tea.NewProgram(m, tea.WithAltScreen()).Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	if p, ok := p.(selection.Model); ok {
		// Extract the ID from p.Choice
		var execID int64
		_, err = fmt.Sscanf(p.Choice, "ID: %d", &execID)
		if err != nil {
			fmt.Println("error parsing execution ID:", err)
			os.Exit(1)
		}
		choice <- execID
	}
}
