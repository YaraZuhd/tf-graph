variable "cloudflare_api_token" {
  description = "Cloudflare API token with Workers Scripts, Workers KV Storage, and DNS edit permissions"
  type        = string
  sensitive   = true
}

variable "cloudflare_account_id" {
  description = "Your Cloudflare account ID (found in the dashboard under 'My Profile' > 'API Tokens')"
  type        = string
}
