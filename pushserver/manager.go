package main

import "sync"

// Manager manages connected clients and message distribution.
type Manager struct {
    mu      sync.RWMutex
    clients map[string]chan string
}

// NewManager creates a new Manager.
func NewManager() *Manager {
    return &Manager{clients: make(map[string]chan string)}
}

// Register adds a client with the given id and returns a receive-only channel for messages.
func (m *Manager) Register(id string) <-chan string {
    m.mu.Lock()
    defer m.mu.Unlock()
    ch := make(chan string, 16)
    m.clients[id] = ch
    return ch
}

// Unregister removes a client and closes its channel.
func (m *Manager) Unregister(id string) {
    m.mu.Lock()
    defer m.mu.Unlock()
    if ch, ok := m.clients[id]; ok {
        delete(m.clients, id)
        close(ch)
    }
}

// Broadcast sends a message to all connected clients (best-effort, non-blocking).
func (m *Manager) Broadcast(msg string) {
    m.mu.RLock()
    defer m.mu.RUnlock()
    for _, ch := range m.clients {
        select {
        case ch <- msg:
        default:
        }
    }
}

// SendTo sends a message to a specific client. Returns true if the client exists and the message was queued.
func (m *Manager) SendTo(id, msg string) bool {
    m.mu.RLock()
    ch, ok := m.clients[id]
    m.mu.RUnlock()
    if !ok {
        return false
    }
    select {
    case ch <- msg:
        return true
    default:
        return false
    }
}

// ListClients returns the list of currently connected client IDs.
func (m *Manager) ListClients() []string {
    m.mu.RLock()
    defer m.mu.RUnlock()
    ids := make([]string, 0, len(m.clients))
    for id := range m.clients {
        ids = append(ids, id)
    }
    return ids
}
