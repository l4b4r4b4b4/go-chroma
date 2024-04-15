# Go-Chroma
[ChromaDB](https://www.trychroma.com/) port written in Go.

**Short-Term Features**:
- Feature paritiy with ChromaDB (Python)
- Chroma API-compatibility with [protobuf](https://raw.githubusercontent.com/chroma-core/chroma/main/idl/chromadb/proto/chroma.proto)

- Transport layer in RPC and HTTP
- Powered by GraphDB: https://github.com/grails/gorm-neo4j

**Feature Roadmap**:
- RPC-powered Clients in Python, JS and Go
- Optional HNSW or DiskANN retrieval
- Top-k and max. relevance retrieval
- Multi-Modality
- Native Cross-encoder ReRanking
- Knowledge Graph generation, attachment and retrieval
- Optional SIMD-optimization
https://github.com/alivanz/go-simd