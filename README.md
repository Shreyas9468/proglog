# 🚀 proglog — Distributed Commit Log in Go

[![CI](https://github.com/Shreyas9468/proglog/actions/workflows/ci.yml/badge.svg)](https://github.com/Shreyas9468/proglog/actions/workflows/ci.yml)
[![Deploy](https://github.com/Shreyas9468/proglog/actions/workflows/deploy.yml/badge.svg)](https://github.com/Shreyas9468/proglog/actions/workflows/deploy.yml)
[![Docker Image](https://ghcr.io-badge.vercel.app/lv/Shreyas9468/proglog)](https://github.com/Shreyas9468/proglog/pkgs/container/proglog)

**proglog** is a high-performance distributed commit log engine built from scratch in Go, implementing append-only storage, gRPC APIs, mutual TLS (mTLS), Casbin ACL authorization, Serf service discovery, Raft consensus replication, and Kubernetes/Cloud deployment pipelines.

---

## 🏛️ System Architecture

```
                       ┌─────────────────────────┐
                       │   Client / getservers   │
                       └────────────┬────────────┘
                                    │ gRPC (mTLS)
                      ┌─────────────▼─────────────┐
                      │    Custom Load Balancer   │
                      │ (gRPC Resolver + Picker)  │
                      └──────┬──────────┬──────┬──┘
                             │          │      │
                      ┌──────▼──┐  ┌────▼───┐  ┌▼────────┐
                      │proglog-0│  │proglog-1│  │proglog-2│
                      │ Leader  │  │Follower│  │Follower │
                      │  Raft   │◄─┤  Raft  │  │  Raft   │
                      │  Serf   │  │  Serf  │  │  Serf   │
                      └─────────┘  └────────┘  └─────────┘
```

- **Raft Consensus**: Handles leader election, log replication, and strong consistency across nodes.
- **Serf Gossip Protocol**: Manages cluster membership discovery and failure detection.
- **Custom gRPC Load Balancer**: Directs `Produce` requests to the Raft leader and load-balances `Consume` requests across followers.

---

## ✨ Key Features

- 📁 **Append-Only Storage Log** — Segment-based, offset-indexed, memory-mapped (`gommap`) storage engine.
- ⚡ **gRPC API** — Protobuf definitions supporting streaming (`ProduceStream`, `ConsumeStream`) and RPCs.
- 🔒 **Mutual TLS (mTLS)** — Encrypted inter-node and client-to-server communications.
- 🛡️ **ACL Authorization** — Role-based authorization using Casbin model and policy rules.
- 📡 **Service Discovery** — Automatic node membership discovery powered by Serf.
- 🤝 **Consensus Replication** — Distributed consensus managed by HashiCorp Raft.
- 🐳 **Docker & Helm** — Multi-stage Docker container build and Helm Chart for Kubernetes StatefulSets.
- 🔄 **Automated CI/CD** — GitHub Actions workflow for testing, publishing to `ghcr.io`, and auto-deploying.

---

## 💻 Quick Start & Running Locally

### Prerequisites
- [Go 1.25+](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Kind](https://kind.sigs.k8s.io/) or [Minikube](https://minikube.sigs.k8s.io/) *(for local Kubernetes)*
- [Helm 3+](https://helm.sh/)

### 1. Build and Run Tests
```powershell
# Compile all packages
go build ./...

# Execute unit and integration test suite
go test -v ./...
```

### 2. Run Local 3-Node Kubernetes Cluster with Helm
```powershell
# Build Docker image
make build-docker

# Create local Kind cluster & load image
kind create cluster
kind load docker-image ghcr.io/shreyas9468/proglog:latest

# Deploy via Helm
helm upgrade --install proglog deploy/proglog

# Verify running pods
kubectl get pods -w
```

### 3. Inspect Cluster Status
```powershell
# Port forward gRPC port of leader node
kubectl port-forward pod/proglog-0 8400:8400

# Inspect cluster servers using CLI tool
go run ./cmd/getservers/main.go -addr=localhost:8400
```

---

## 🌐 Free Cloud Deployment (Render.com)

**proglog** is set up for **100% Free** cloud deployment with zero credit card required.

See **[`deploy/README.md`](deploy/README.md)** for a complete step-by-step guide on:
- Deploying the Docker image to **Render.com**.
- Setting up the `RENDER_DEPLOY_HOOK_URL` GitHub Secret for 1-click CI/CD auto-deploy.
- Testing live cloud gRPC endpoints using `grpcurl` or Go CLI clients.

---

## 📂 Project Structure

```
proglog/
├── cmd/
│   ├── proglog/          # Main Agent CLI (Cobra + Viper)
│   └── getservers/       # Cluster discovery inspection CLI
├── internal/
│   ├── agent/            # Agent orchestrator (wires log, raft, serf, grpc)
│   ├── api/v1/           # Protobuf definitions & generated gRPC code
│   ├── auth/             # Casbin ACL authorizer
│   ├── config/           # TLS certificate loader helpers
│   ├── discovery/        # Serf membership & gossip handler
│   ├── loadbalance/      # Custom gRPC resolver & leader-aware picker
│   ├── log/              # Segmented commit log & Raft FSM integration
│   └── server/           # gRPC server implementation & health checks
├── deploy/
│   ├── proglog/          # Helm chart (StatefulSet, Service, Config)
│   └── README.md         # Step-by-step free cloud deployment guide
├── render.yaml           # Render infrastructure blueprint
└── .github/workflows/    # CI (test + build + ghcr) & Deploy pipelines
```
