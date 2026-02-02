# Chat Server

A lightweight, real-time TCP-based chat server implementation written in Go. This chat server allows multiple clients to connect simultaneously and communicate with each other through a central hub.

## Features

- **Real-time messaging**: Instant message delivery to all connected clients
- **Multi-client support**: Handle multiple concurrent client connections
- **User identification**: Each client can set a username upon joining
- **Server notifications**: Automatic notifications when users join or leave
- **Simple TCP protocol**: Easy to connect using standard TCP clients (telnet, netcat, etc.)
- **Goroutine-based concurrency**: Efficient handling of concurrent connections using Go's goroutines and channels
- **Channel-based message hub**: Central hub for broadcasting messages to all connected clients

## Prerequisites

Before you begin, ensure you have the following installed:
- **Go**: Version 1.25.1 or higher ([Download Go](https://golang.org/dl/))

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Kian-Abdalkhani/chat-server.git
cd chat-server
```

2. Initialize Go modules (if needed):
```bash
go mod tidy
```

3. Build the server:
```bash
go build -o chat-server
```

## Setup and Running

### Starting the Server

Run the compiled binary:
```bash
./chat-server
```

Or run directly with Go:
```bash
go run main.go
```

The server will start listening on **port 9000** by default. You should see:
```
Listening on port 9000
```

### Connecting to the Server

You can connect to the server using any TCP client. Here are some examples:

**Using telnet:**
```bash
telnet localhost 9000
```

**Using netcat:**
```bash
nc localhost 9000
```

**Using a custom TCP client:**
```go
// Example Go client code would go here
```

### Using the Chat

1. Once connected, you'll be prompted to enter your name
2. After entering your name, you'll receive a welcome message
3. Type messages and press Enter to send them to all connected users
4. All connected users will see your messages prefixed with your username
5. Server notifications appear when users join or leave

## Project Structure

```
chat-server/
├── main.go           # Entry point, server setup, and connection handling
├── client/
│   └── client.go     # Client struct and methods for reading/writing messages
├── message/
│   └── message.go    # Message types and message kind definitions
├── server/
│   └── server.go     # Server utility functions
└── go.mod            # Go module definition
```

### Architecture Overview

- **Hub Pattern**: The server uses a hub-based architecture where all messages flow through a central hub that broadcasts to all connected clients
- **Channels**: Go channels are used for communication between goroutines (client registration, unregistration, and message broadcasting)
- **Concurrent Client Handling**: Each client connection is handled in its own goroutine for reading and writing

## Usage Examples

<!-- TODO: Add specific usage examples here -->
<!-- Example: Demonstrate a typical chat session between multiple users -->
<!-- Example: Show how to integrate with a custom client application -->

## Future Improvements

This project has potential for several enhancements:

- **Message Encryption**: Implement TLS/SSL for secure communication between clients and server
- **Message Persistence**: Store chat history in a database for message retrieval and logging
- **Authentication**: Add user authentication and authorization mechanisms
- **Private Messaging**: Support for direct messages between specific users
- **Chat Rooms**: Create multiple chat rooms or channels
- **Message History**: Allow new users to see recent message history upon joining
- **User Commands**: Implement commands like `/list` (show online users), `/whisper` (private message), etc.
- **WebSocket Support**: Add WebSocket protocol support for web-based clients
- **Rate Limiting**: Prevent spam by implementing message rate limiting
- **User Profiles**: Extended user information and profiles
- **File Sharing**: Support for sharing files between users
- **Emojis and Rich Text**: Support for formatting and emojis
- **Logging and Monitoring**: Add comprehensive logging and monitoring capabilities
- **Configuration File**: External configuration for port, max connections, etc.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

<!-- TODO: Add license information here -->
