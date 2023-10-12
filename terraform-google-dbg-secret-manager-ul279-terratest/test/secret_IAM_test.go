package test

import (
	"context"
	"testing"

	"google.golang.org/api/option"
	"google.golang.org/api/secretmanager/v1"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestSecretManagerIAM(t *testing.T) {
	t.Parallel()
  
  SecretName := "projects/dbg-tfe-modules-95de4071/secrets/simple-secret"

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
			"iam": map[string]interface{}{
				"secret-iam-binding": map[string][]string{
					"roles/secretmanager.secretAccessor": []string{"serviceAccount:dbg-tfe-modules-95de4071@appspot.gserviceaccount.com"},
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
	client, err := secretmanager.NewService(ctx, option.WithoutAuthentication())
	if err != nil {
		t.Fatalf("Error creating Secret Manager client: %v", err)
	}

	// Fetch the IAM policy of the secret
	iamPolicy, err := client.Projects.Secrets.GetIamPolicy(SecretName).Context(ctx).Do()
	if err != nil {
		t.Fatalf("Error retrieving IAM policy: %v", err)
	}

	// Define the expected IAM binding
	IamBinding := &secretmanager.Policy{
		Bindings: []*secretmanager.Binding{
			{
				Role:    "roles/secretmanager.secretAccessor",
				Members: []string{"serviceAccount:dbg-tfe-modules-95de4071@appspot.gserviceaccount.com"},
			},
		},
	}
  
	// Assertions for the IAM binding
	assert.Equal(t, IamBinding.Bindings, iamPolicy.Bindings, "IAM binding does not match")
}
