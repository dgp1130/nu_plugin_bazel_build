package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Read input JSON.
		var input map[string]interface{}
		line := scanner.Text()
		err := json.Unmarshal([]byte(line), &input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to deserialize request JSON: %s\n", err)
			return
		}

		// Decide which action to take.
		method := input["method"].(string)
		switch method {
		case "config":
			// Respond with config information so nushell knows how to call this plugin.
			// Declare self as a filter with no parameters.
			respond(map[string]interface{}{
				"name":       "from-build",
				"usage":      "Parses BUILD file input into a structured format.",
				"positional": []interface{}{},
				"named":      map[string]interface{}{},
				"is_filter":  true,
			})
			return // Terminate after loading configuration.
		case "begin_filter":
			// Do nothing before filtering.
			respond([]interface{}{})
		case "end_filter":
			// Do nothing after filtering.
			respond([]interface{}{})
			return // Terminate when ending filter.
		case "filter":
			// Filter the input.
			respond([]interface{}{
				map[string]interface{}{
					"Ok": map[string]interface{}{
						"Value": map[string]interface{}{
							// Echo back the provided tag parameter.
							"tag": input["params"].(map[string]interface{})["tag"].(map[string]interface{}),

							// Just return "Hello World!" for now.
							"item": map[string]interface{}{
								"Primitive": map[string]interface{}{
									"String": "Hello World!",
								},
							},
						},
					},
				},
			})
		default:
			fmt.Fprintf(os.Stderr, "Unknown method: %s\n", method)
			return // Terminate on unknown request.
		}
	}
}

func respond(data interface{}) {
	j, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "response",
		"params": map[string]interface{}{
			"Ok": data,
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to serialize response JSON: %s\n", err)
	}

	fmt.Printf("%s\n", string(j))
}
