# proglog

Proglog — A distributed commit log built with Go, implementing storage, gRPC, TLS/mTLS authentication, authorization, service discovery, consensus, and distributed systems concepts.

[![CI](https://github.com/Shreyas9468/proglog/actions/workflows/ci.yml/badge.svg)](https://github.com/Shreyas9468/proglog/actions/workflows/ci.yml)
[![Deploy](https://github.com/Shreyas9468/proglog/actions/workflows/deploy.yml/badge.svg)](https://github.com/Shreyas9468/proglog/actions/workflows/deploy.yml)
[![Docker Image](https://ghcr.io-badge.vercel.app/lv/Shreyas9468/proglog)](https://github.com/Shreyas9468/proglog/pkgs/container/proglog)

---

## Architecture

```
┌─────────────────────────────────────────────────┐
│  Client  (getservers / custom gRPC client)      │
└──────────────────────┬──────────────────────────┘
                       │ gRPC (TLS)
         ┌─────────────▼─────────────┐
         │      Load Balancer        │
         │  (custom gRPC resolver    │
         │   + round-robin picker)   │
         └──┬──────────┬──────────┬──┘
            │          │          │
     ┌──────▼──┐  ┌────▼────┐  ┌─▼───────┐
     │proglog-0│  │proglog-1│  │proglog-2│
     │ Leader  │  │Follower │  │Follower │
     │  Raft   │◄─┤  Raft   │  │  Raft   │
     │  Serf   │  │  Serf   │  │  Serf   │
     └─────────┘  └─────────┘  └─────────┘
     (Raft consensus + Serf gossip for membership)
```

---

## Features

- **Distributed Commit Log** — Append-only, offset-indexed, segment-based storage
- **gRPC API** — Protobuf-defined Produce / Consume / ProduceStream / ConsumeStream / GetServers
- **mTLS** — Mutual TLS between clients and servers
- **Authorization** — ACL-based (Casbin) with model + policy files
- **Service Discovery** — Serf gossip protocol for membership
- **Consensus** — Raft-based replication (HashiCorp Raft)
- **Load Balancing** — Custom gRPC resolver + leader-aware picker
- **Health Checks** — Standard gRPC Health Checking Protocol (`grpc.health.v1`)
- **Kubernetes** — Helm chart with StatefulSet, headless Service, init-container config bootstrap
- **CI/CD** — GitHub Actions → ghcr.io → Koyeb (auto-deploy on push to main)

---

## Running Locally

### Prerequisites

- Go 1.25+
- Docker
- [Kind](https://kind.sigs.k8s.io/) (for local Kubernetes)
- [Helm](https://helm.sh/)

### Run tests

```powershell
make test
```

### Local Kubernetes cluster (3-node)

```powershell
docker build -t proglog:0.0.1 .
kind create cluster
kind load docker-image proglog:0.0.1
helm install proglog deploy/proglog
kubectl get pods -w
```

Once all 3 pods are ready:

```powershell
kubectl port-forward pod/proglog-0 8400:8400
go run ./cmd/getservers/main.go -addr=localhost:8400
```

---

## Cloud Deployment (Free, No Credit Card)

See [deploy/README.md](deploy/README.md) for full instructions on:
- **GitHub Actions** CI/CD pipeline (auto test + build on every push)
- **GitHub Container Registry** (ghcr.io) for Docker image hosting
- **Koyeb.com** (free tier, always-on, gRPC support) for cloud hosting

---

## Project Structure

```
proglog/
├── cmd/
│   ├── proglog/          # Agent CLI (Cobra + Viper)
│   └── getservers/       # Cluster inspection CLI
├── internal/
│   ├── agent/            # Wires all components together
│   ├── api/v1/           # Protobuf definitions
│   ├── auth/             # Casbin ACL authorization
│   ├── config/           # TLS setup helpers
│   ├── discovery/        # Serf membership
│   ├── loadbalance/      # gRPC resolver + picker
│   ├── log/              # Commit log + Raft integration
│   └── server/           # gRPC server + health check
├── deploy/
│   ├── proglog/          # Helm chart
│   └── README.md         # Cloud deployment guide
└── .github/workflows/    # CI + Deploy pipelines
```
