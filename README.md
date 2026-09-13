# NexusNode — Decentralized P2P Storage Network

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

NexusNode is a peer-to-peer (P2P) decentralized file storage network built for **enterprise data sovereignty**. Files are split using **Reed-Solomon erasure coding**, encrypted with **AES-256 GCM**, and distributed across edge nodes — without any central server. Even if a portion of nodes go offline, the original file can be fully recovered.

## Architecture

```
┌───────────────────────────────────────────────────────┐
│                      CLI (cmd/)                       │
├─────────────────┬────────────────┬────────────────────┤
│  network/       │  sharding/     │  crypto/           │
│  libp2p + mDNS  │  Reed-Solomon  │  AES-256 GCM       │
│  P2P Discovery  │  Split/Join    │  Encrypt/Decrypt   │
└─────────────────┴────────────────┴────────────────────┘
```

## Features

| Phase | Feature | Status |
|-------|---------|--------|
| 1 | P2P networking via `go-libp2p` + mDNS peer discovery | ✅ Done |
| 1 | File sharding and reassembly | ✅ Done |
| 2 | AES-256 GCM encryption per shard | ✅ Done |
| 2 | Reed-Solomon erasure coding (fault tolerance) | ✅ Done |
| 2 | Metadata-based exact file reconstruction | ✅ Done |
| 3 | React/Electron Desktop Client & Dashboard | 🔜 Planned |
| 4 | Decentralized Identity (DID) + Token Economics | 🔜 Planned |

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.21 or higher

### Build

```bash
git clone https://github.com/meryemgcl/NexusNode-Decentralized-Storage.git
cd NexusNode-Decentralized-Storage

# Download dependencies and generate go.sum
go mod tidy

# Build the CLI binary
go build -o nexusnode cmd/nexusnode/main.go
```

### Usage

#### Split a file into encrypted shards

```bash
./nexusnode -chunk path/to/document.pdf \
            -key "your-super-secret-key-32chars!!" \
            -data 10 \
            -parity 4
```

Output: `document.pdf.shard.0` … `document.pdf.shard.13` and `document.pdf.meta.json`

#### Assemble shards back into the original file

```bash
# All shards present
./nexusnode -assemble "document.pdf.shard.0,document.pdf.shard.1,...,document.pdf.shard.13" \
            -meta "document.pdf.meta.json" \
            -key "your-super-secret-key-32chars!!" \
            -out "document_restored.pdf"

# Simulating 2 offline nodes (Reed-Solomon recovers them automatically)
./nexusnode -assemble "missing,document.pdf.shard.1,missing,...,document.pdf.shard.13" \
            -meta "document.pdf.meta.json" \
            -key "your-super-secret-key-32chars!!" \
            -out "document_restored.pdf"
```

#### Start a P2P node

```bash
# Terminal 1
./nexusnode -port 4001

# Terminal 2 (will auto-discover Terminal 1 via mDNS)
./nexusnode -port 4002
```

### Run Tests

```bash
go test ./...
```

## Security Model

- **Zero-Trust Sharding:** No single node holds the complete file.
- **AES-256 GCM:** Each shard is independently encrypted. Nodes cannot read the data without the key.
- **Parity Tolerance:** With `-data 10 -parity 4`, any 4 of 14 nodes can be offline and data is still fully recoverable.

## Project Structure

```
.
├── cmd/nexusnode/    # CLI entry point
├── crypto/           # AES-256 GCM encryption/decryption
├── network/          # libp2p host, mDNS peer discovery
├── sharding/         # Reed-Solomon split, assemble, metadata
├── docs/             # Project background and phase planning documents
├── go.mod
└── README.md
```

## License

[MIT](LICENSE)
