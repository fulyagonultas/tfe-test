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

func TestSecretManagerCMEK(t *testing.T) {
    t.Parallel()

    projectID := "dbg-tfe-modules-95de4071"
    cmd := exec.Command("gcloud", "projects", "describe", projectID, "--format=value(projectNumber)")
    output, err := cmd.CombinedOutput()
    if err != nil {
        t.Fatalf("Error fetching project ID: %v, Output: %s", err, output)
    }

    projectNumber := strings.TrimSpace(string(output))
    numericProjectNumber := strings.TrimPrefix(projectNumber, "projects/")
    secretNamewithProjectNumber := "projects/" + numericProjectNumber + "/secrets/cmek-encrypted-secret"

    SecretName := "projects/dbg-tfe-modules-95de4071/secrets/cmek-encrypted-secret"

    // Define the Terraform options
    terraformOptions := &terraform.Options{
        TerraformDir: "../",
        Vars: map[string]interface{}{
            "project_id": "dbg-tfe-modules-95de4071",
            "secrets": map[string][]string{
                "cmek-encrypted-secret": []string{"europe-west3"},
            },
            "encryption_keyenr": map[string]string{
                "europe-west3": "projects/dbg-tfe-modules-95de4071/locations/europe-west3/keyRings/test-secret-encryption/cryptoKeys/test-encryption",
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

    // Fetch the encryption key name from the secret details
    encryptionKeyName := secret.KmsKeyName

    // Expected encryption key for your specific region (europe-west3)
    expectedEncryptionKey := "projects/dbg-tfe-modules-95de4071/locations/europe-west3/keyRings/test-secret-encryption/cryptoKeys/test-encryption"

    // Assertions for the encryption key
    assert.Equal(t, expectedEncryptionKey, encryptionKeyName, "Encryption key name does not match")
}
