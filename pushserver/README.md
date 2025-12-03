# Simple SSE Message Push Server

This small example implements a message push system using Server-Sent Events (SSE) in Go.

Endpoints:

- `GET /events?id=<client_id>` — subscribe to messages as an SSE stream. If `id` is omitted, a generated id is returned in the `connected` event.
- `POST /publish` — publish a message. Accepts raw text body (broadcast) or JSON `{ "to": "id", "message": "..." }` to send to a specific client.
- `GET /clients` — list connected client IDs (JSON).

Run:

```powershell
cd pushserver
go run .
```

Subscribe example (keep connection open):

```powershell
curl -N http://localhost:8085/events?id=alice
```

Broadcast a message (raw body):

```powershell
curl -X POST -d "hello all" http://localhost:8085/publish
```

Send to a specific client (JSON):

```powershell
curl -X POST -H "Content-Type: application/json" -d '{"to":"alice","message":"hi alice"}' http://localhost:8085/publish
```

This implementation is intentionally small and dependency-free. It demonstrates a simple, reliable way to push messages to browsers or other HTTP clients without adding a WebSocket dependency.
