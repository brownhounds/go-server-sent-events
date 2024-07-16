## Server Sent Events With GO

### Pros ✅

1. **Simplicity:** SSE is easy to implement and use for real-time updates in web applications. It is built into the HTTP protocol, meaning you don’t need additional libraries or complex setup.
2. **Built-in Reconnection:** SSE has automatic reconnection built-in, so if the connection drops, it will attempt to reconnect without additional effort from the developer.
3. **Efficient for Unidirectional Data:** SSE is efficient for scenarios where the server needs to push updates to the client, such as live feeds or notifications. The connection remains open, reducing the need for frequent polling.

### Cons 💀

1. **Text-Based Only:** SSE only supports the transmission of text data, making it unsuitable for scenarios that require the transfer of binary data, such as images or files, without additional encoding and decoding processes.
2. **One-Directional Transfer:** SSE supports only server-to-client communication.

### Notes 🗒️

When **not used over HTTP/2**, SSE suffers from a limitation to the maximum number of open connections, which can be especially painful when opening multiple tabs as the limit is per browser and is set to a very low number (6). **When using HTTP/2**, the maximum number of simultaneous HTTP streams is negotiated between the server and the client (defaults to 100).

## Example 1

**Dispatching Events Using Redis Pub/Sub**

In this approach, events are dispatched via the Redis Pub/Sub mechanism, which allows different parts of an application or multiple applications to communicate by publishing and subscribing to messages over Redis.

## Example 2

**Dispatching Events Using an In-Application Event Bus with Channels**

This example involves dispatching events within a single application using an internal event bus system that utilizes channels for communication. This method is well-suited for handling events within a single application instance but does not facilitate communication across different applications or services.
