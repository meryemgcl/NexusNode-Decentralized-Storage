# NexusNode Decentralized Storage

NexusNode is a peer-to-peer (P2P) decentralized file storage and sharing network designed for enterprise data sovereignty. It ensures zero-trust architecture, scalable redundant storage, and self-healing properties without relying on centralized cloud providers.

## Features (Phase 1)
- **P2P Networking:** Decentralized discovery and communication using `go-libp2p` and mDNS.
- **File Sharding:** Splits large files into mathematical chunks (sharding) for distributed storage.
- **Reassembly:** Reconstructs the original file from chunks completely losslessly.

## Future Phases
- **Phase 2:** AES-256 Encryption & Reed-Solomon Erasure Coding for extreme fault tolerance.
- **Phase 3:** React/Electron Desktop Client & Dashboard.
- **Phase 4:** Decentralized Identity (DID) & Token Economics.

## Getting Started

### Prerequisites
- [Go (Golang)](https://go.dev/dl/) 1.21 or higher.

### Build
```bash
git clone https://github.com/meryemgcl/NexusNode-Decentralized-Storage.git
cd NexusNode-Decentralized-Storage
go mod tidy
go build -o nexusnode cmd/nexusnode/main.go
```

### Usage
Start a P2P node:
```bash
./nexusnode -port 4001
```

Split a file into chunks:
```bash
./nexusnode -chunk path/to/file.pdf
```

Assemble chunks back into a file:
```bash
./nexusnode -assemble path/to/file.pdf.chunk.0,path/to/file.pdf.chunk.1 -out restored_file.pdf
```

## License
MIT License
