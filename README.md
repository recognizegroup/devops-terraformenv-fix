# devops-terraformenv-fix

DevOps Pipeline has the strange behavior of converting environment variables to uppercase. This makes it inconvenient to use Terraform variables from the environment, since the convention for Terraform is to use something like TF_VAR_env_name . This tool converts the environment back to what it should be.

## Usage

eval $(./devops-terraformenv-fix)