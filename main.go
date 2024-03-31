package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// Set script start time
	scriptStartTime := time.Now()

	// Define log types
	eventTypes := []string{"Application", "Security", "Setup", "System"}

	// Output file folder
	logOutputFolder := "C:\\data\\EventLogs\\"

	// Start infinite loop
	for {
		for _, eventType := range eventTypes {
			// Build output file path
			outputFile := logOutputFolder + eventType + ".csv"

			// Execute PowerShell command to retrieve event logs
			powershellCommand := fmt.Sprintf(`Get-WinEvent -LogName %s -MaxEvents 10 | Select-Object TimeCreated, Id, ProviderName, Message | Export-Csv -Path '%s' -Append -NoTypeInformation`, eventType, outputFile)
			cmd := exec.Command("powershell", "-Command", powershellCommand)

			// Run the command
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error executing PowerShell command:", err)
			}
		}

		// Sleep for 60 seconds
		time.Sleep(60 * time.Second)
	}

	// Calculate script duration
	scriptEndTime := time.Now()
	scriptDuration := scriptEndTime.Sub(scriptStartTime)

	fmt.Println("Script duration:", scriptDuration)
}
