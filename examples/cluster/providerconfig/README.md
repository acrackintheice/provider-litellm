# ProviderConfig and credentials (cluster)

1. Copy the secret template and set your LiteLLM credentials:

   ```bash
   cp secret.yaml.tmpl secret.yaml
   ```

   Edit `secret.yaml` and set `api_key` in the `credentials` JSON to your LiteLLM API key.
   For the sit environment, `api_base` is already set to `https://litellm.sit.q3.questech.io`.

   Or create from the template using an environment variable (keeps the key out of the repo):

   ```bash
   LITELLM_API_KEY=your-key-here bash -c 'sed "s/YOUR_LITELLM_API_KEY/$LITELLM_API_KEY/" secret.yaml.tmpl > secret.yaml'
   ```

2. Apply the ProviderConfig and secret (after creating the namespace if needed):

   ```bash
   kubectl create namespace crossplane-system --dry-run=client -o yaml | kubectl apply -f -
   kubectl apply -f examples/cluster/providerconfig/
   ```

Note: `secret.yaml` is gitignored; do not commit it with real credentials.
