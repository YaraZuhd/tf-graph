terraform {
  required_providers {
    cloudflare = {
        source  = "cloudflare/cloudflare"
        version = "~> 4.0"
    }
  }
}


provider "cloudflare" {
  api_token = var.cloudflare_api_token
}


# KV namespace - simple key-value store the Worker will read/write from
resource "cloudflare_workers_kv_namespace" "tf_graph_kv" {
  account_id = var.cloudflare_account_id
  title      = "tf-graph-kv"
}

# Worker script - depends on the KV namespace via its binding
# Worker script - depends on the KV namespace via its binding
resource "cloudflare_workers_script" "tf_graph_worker" {
  account_id = var.cloudflare_account_id
  name       = "tf-graph-worker"
  content    = file("${path.module}/worker.js")
  module     = true

  kv_namespace_binding {
    name         = "TF_GRAPH_KV"
    namespace_id = cloudflare_workers_kv_namespace.tf_graph_kv.id
  }
}
