package executor

import (
	"fmt"
	"log"
	"net/http"

	"github.com/s8sg/faas-flow/sdk/executor"
)

func pauseFlowHandler(w http.ResponseWriter, req *http.Request, id string, ex executor.Executor) ([]byte, error) {
	log.Printf("Pausing flow %s\n", id)

	flowExecutor := executor.CreateFlowExecutor(ex, nil)
	err := flowExecutor.Pause(id)
	if err != nil {
		return nil, fmt.Errorf("failed to pause request %s, check if request is active", id)
	}

	return []byte("Successfully paused request " + id), nil
}
