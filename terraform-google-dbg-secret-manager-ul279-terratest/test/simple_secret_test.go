package test

import (
    "context"
    "testing"
    "os/exec"
    "strings"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
    "google.golang.org/api/secretmanager/v1"
)

func TestSecretManagerSimple(t *testing.T) {
    t.Parallel()

    projectID:= "dbg-tfe-modules-95de4071"

    cmd := exec.Command("gcloud", "projects", "describe", projectID, "--format=value(projectNumber)")
    output, err := cmd.CombinedOutput()
    if err != nil {
	    t.Fatalf("Error fetching project ID: %v, Output: %s", err, output)
    }
    projectNumber := strings.TrimSpace(string(output))
    numericProjectNumber := strings.TrimPrefix(projectNumber, "projects/")
    secretNamewithProjectNumber := "projects/" + numericProjectNumber + "/secrets/simple-secret"

    SecretName := "projects/dbg-tfe-modules-95de4071/secrets/simple-secret"
    Labels := map[string]string{
        "name":           "simple-secret",
        "environment":    "dev",
        "product":        "web",
        "productlineid":  "3001",
        "creator":        "admin",
        "supportgroupid": "admin",
        "applicationid":  "4001",
        "costcenter":     "11001",
    }

    // Define the Terraform options
    terraformOptions := &terraform.Options{
        // Set the path to your Terraform code
        TerraformDir: "../",

        // Variables to pass to our Terraform configuration using -var options
        Vars: map[string]interface{}{
            "project_id": "dbg-tfe-modules-95de4071",
            "secrets": map[string][]string{
                "simple-secret": []string{"europe-west3"},
            },
            "labels": map[string]map[string]string{
                "simple-secret": {
                    "name":           "simple-secret",
                    "environment":    "dev",
                    "product":        "web",
                    "productlineid":  "3001",
                    "creator":        "admin",
                    "supportgroupid": "admin",
                    "applicationid":  "4001",
                    "costcenter":     "11001",
                },
            },
        },
    }

    // Defer the destruction of resources to the end of the test
    defer terraform.Destroy(t, terraformOptions)
	
    // Init and apply the Terraform configuration
    terraform.InitAndApply(t, terraformOptions)

    // Create a context and client for Secret Manager
    ctx := context.Background()
    client, err := secretmanager.NewService(ctx)
    if err != nil {
        t.Fatalf("Error creating Secret Manager client: %v", err)
    }

    // Fetch the secret
    secret, err := client.Projects.Secrets.Get(SecretName).Context(ctx).Do()
    if err != nil {
        t.Fatalf("Error retrieving secret: %v", err)
    }
	
    // Assertions for the name of the fetched secret
    assert.Equal(t, secretNamewithProjectNumber, secret.Name, "Secret name does not match")

    // Assertions for the labels of the fetched secret
    assert.Equal(t, Labels, secret.Labels, "Secret labels do not match")
}
