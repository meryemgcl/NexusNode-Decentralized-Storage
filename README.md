# NexusNode — Decentralized P2P Storage Network

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

NexusNode is a peer-to-peer (P2P) decentralized file storage network built for **enterprise data sovereignty**. Files are split using **Reed-Solomon erasure coding**, encrypted with **AES-256 GCM**, and distributed across edge nodes — without any central server. Even if a portion of nodes go offline, the original file can be fully recovered using parity shards.

## Architecture

```
┌───────────────────────────────────────────────────────┐
│                      CLI (cmd/)                       │
├─────────────────┬────────────────┬────────────────────┤
│  network/       │  sharding/     │  crypto/           │
│  libp2p + mDNS  │  Reed-Solomon  │  AES-256 GCM       │
│  + Kademlia DHT │  Split/Join    │  Encrypt/Decrypt   │
└─────────────────┴────────────────┴────────────────────┘
           ↑ spawned by
┌───────────────────────────────────────────────────────┐
│           desktop/ (Electron + React + Vite)          │
│  Dashboard │ Upload │ Files │ Network                 │
└───────────────────────────────────────────────────────┘
```

## Features

| Phase | Feature | Status |
|-------|---------|--------|
| 1 | P2P networking via `go-libp2p` + mDNS local discovery | ✅ Done |
| 1 | File sharding and reassembly | ✅ Done |
| 2 | AES-256 GCM encryption per shard (key via stdin, not args) | ✅ Done |
| 2 | Reed-Solomon erasure coding (fault tolerance) | ✅ Done |
| 2 | Metadata-based exact file reconstruction | ✅ Done |
| 2 | Kademlia DHT global peer discovery | ✅ Done |
| 3 | Electron/React Desktop Client & Dashboard | ✅ Done |
| 4 | Decentralized Identity (DID) + Token Economics | 🔜 Planned |

## Getting Started

### Prerequisites

| Tool | Version | Install |
|------|---------|---------|
| [Go](https://go.dev/dl/) | 1.21+ | Required to build the CLI backend |
| [Node.js](https://nodejs.org/) | 18+ | Required for the desktop UI |

### 1 — Build the Go CLI Backend

> **This step is required before running the desktop app.**

```bash
git clone https://github.com/meryemgcl/NexusNode-Decentralized-Storage.git
cd NexusNode-Decentralized-Storage

# Download Go dependencies and generate go.sum
go mod tidy

# Build the CLI binary (produces nexusnode.exe on Windows, nexusnode on Linux/macOS)
go build -o nexusnode cmd/nexusnode/main.go
```

### 2 — Run the Desktop App

```bash
cd desktop
npm install
npm run dev
```

The Electron app will automatically find the `nexusnode` binary in the repository root.

### CLI Usage (without the desktop app)

**Security note:** The AES-256 key is always read from **stdin** (not command-line arguments) to prevent it from appearing in process lists.

#### Split a file into encrypted shards

```bash
echo "your-super-secret-key-32chars!!" | ./nexusnode -chunk path/to/document.pdf -data 10 -parity 4
```

Output: `document.pdf.shard.0` … `document.pdf.shard.13` and `document.pdf.meta.json`

#### Assemble shards back into the original file

```bash
# All shards present
echo "your-super-secret-key-32chars!!" | ./nexusnode \
  -assemble "document.pdf.shard.0,...,document.pdf.shard.13" \
  -meta "document.pdf.meta.json" \
  -out "document_restored.pdf"

# Simulating 2 offline nodes — Reed-Solomon recovers them automatically
echo "your-super-secret-key-32chars!!" | ./nexusnode \
  -assemble "missing,document.pdf.shard.1,missing,...,document.pdf.shard.13" \
  -meta "document.pdf.meta.json" \
  -out "document_restored.pdf"
```

#### Start a P2P node

```bash
# Terminal 1 — node starts, bootstraps DHT, advertises on mDNS
./nexusnode -port 4001

# Terminal 2 — discovers Terminal 1 via mDNS (LAN) or Kademlia DHT (internet)
./nexusnode -port 4002
```

### Run Tests

```bash
go test ./...
```

## Security Model

- **Zero-Trust Sharding:** No single node holds the complete file.
- **AES-256 GCM:** Each shard is independently encrypted. The key is passed via stdin, never via process arguments.
- **Parity Tolerance:** With `-data 10 -parity 4`, any 4 of 14 nodes can be offline and data is fully recoverable.
- **Electron Sandboxing:** `sandbox: true` + `contextIsolation: true` enforced. `nodeIntegration` is disabled.

## Project Structure

```
.
├── cmd/nexusnode/      # CLI entry point (reads key from stdin)
├── crypto/             # AES-256 GCM encryption/decryption + tests
├── network/            # libp2p host, mDNS + Kademlia DHT discovery
├── sharding/           # Reed-Solomon split, assemble, metadata + tests
├── desktop/            # Electron + React + Vite desktop client
│   └── src/
│       ├── main/       # Electron main process (IPC, CLI bridge)
│       ├── preload/    # contextBridge API surface
│       └── renderer/   # React pages & components
├── docs/               # Project background and phase planning documents
├── go.mod
└── README.md
```

## License

[MIT](LICENSE)
