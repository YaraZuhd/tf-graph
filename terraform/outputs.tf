output "worker_name" {
  description = "The name of the deployed Worker script"
  value       = cloudflare_workers_script.tf_graph_worker.name
}

output "kv_namespace_id" {
  description = "The ID of the created KV namespace"
  value       = cloudflare_workers_kv_namespace.tf_graph_kv.id
}
