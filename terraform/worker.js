export default {
  async fetch(request, env) {
    const url = new URL(request.url);

    if (url.pathname === "/ping") {
      await env.TF_GRAPH_KV.put("last-ping", new Date().toISOString());
      return new Response("pong");
    }

    const lastPing = await env.TF_GRAPH_KV.get("last-ping");
    return new Response(
      `tf-graph worker is alive. Last ping: ${lastPing || "never"}`
    );
  },
};
