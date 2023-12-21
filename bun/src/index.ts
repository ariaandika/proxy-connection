
Bun.serve({
  async fetch(request, server) {
    const u = new URL(request.url)

    const file = serveFile(u)

    if (file) {
      return new Response(file)
    }

    if (u.pathname === "/") {
      return new Response(Bun.file("../static/index.html"))
    }

    if (u.pathname.startsWith("/echo")) {
      let text = [
        `${request.method} ${request.url}`,
        `keepalive: ${request.keepalive}`,
        `body: ${await request.text()}`,
      ].join("\n")

      return new Response(text)
    }

    if (u.pathname.startsWith("/ws")) {
      return new Response(Bun.file("../static/ws.html"))
    }

    if (u.pathname.startsWith("/api/ws")) {
      if (server.upgrade(request)) {
        return;
      }
    }

    return new Response()
  },

  websocket: {
    open(ws) {
      ws.subscribe("proto")
      ws.send(`id: ${Math.floor(Math.random() * 90)}`)
      console.log("WS open")
    },
    message(ws, message) {
      ws.publish("proto", message)
      console.log("WS messsage:", message)
    },
    close(ws, code, reason) {
      ws.unsubscribe("proto")
      console.log("WS close", reason, code)
    },
  },

  port: 8000
})

const files = [
  "index.html",
  "script.js",
  "ws.html",
  "ws.js"
]

function serveFile(u: URL) {
  let file = files.find(e=>u.pathname.endsWith(e))

  return file ? Bun.file(`../static/${file}`) : null
}


