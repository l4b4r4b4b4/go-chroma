# Go-Chroma
[ChromaDB](https://www.trychroma.com/) port written in Go.

**Short-Term Features**:
- Feature paritiy with ChromaDB (Python)
- Chroma API-compatibility with [protobuf](https://raw.githubusercontent.com/chroma-core/chroma/main/idl/chromadb/proto/chroma.proto)

- Transport layer in RPC and HTTP
- Powered by libSQL (SQLite)

**Feature Roadmap**:
- RPC-powered Clients in Python, JS and Go
- Optional HNSW or DiskANN retrieval
- Top-k and max. relevance retrieval
- Multi-Modality
- Native Cross-encoder ReRanking
- Synchronisation with central vector storage in the cloud
