# 🌐 Free Cloud Deployment Guide for proglog (Render.com)

Deploy **proglog** — a distributed gRPC commit log service — to **Render.com** on their **100% Free Tier** with zero credit card required. Includes automated CI/CD via GitHub Actions.

---

## 🏗️ Deployment Architecture

| Layer | Service | Cost |
|---|---|---|
| **Code Repository** | GitHub (`Shreyas9468/proglog`) | Free |
| **CI / Automated Tests** | GitHub Actions | Free |
| **Container Registry** | GitHub Container Registry (`ghcr.io`) | Free |
| **Cloud Hosting** | Render.com Web Service (Docker) | **100% Free (No Credit Card)** |

---

## ⚡ 3-Minute Quick Setup

### Step 1: Create a Free Render Account
1. Visit [https://render.com](https://render.com).
2. Click **Sign Up** and log in with your **GitHub account**.
3. *No credit card is required.*

---

### Step 2: Deploy the Service on Render
1. Click **New +** -> **Web Service** in the Render Dashboard.
2. Select **Build and deploy from a Git repository**.
3. Connect your **`Shreyas9468/proglog`** repository.
4. Configure the service parameters:
   - **Name**: `proglog`
   - **Environment**: `Docker`
   - **Region**: Any (e.g., Oregon, US)
   - **Instance Type**: `Free`
   - **Dockerfile Path**: `./Dockerfile`
5. Under **Environment Variables**, add:
   - `PROGLOG_BOOTSTRAP` = `true`
   - `PROGLOG_RPC_PORT` = `8400`
   - `PROGLOG_DATA_DIR` = `/tmp/proglog`
6. Click **Create Web Service**. Render will start building the Docker container and deploy it automatically!

---

### Step 3: Setup Automated CI/CD (Auto-Deploy on Push)
1. In your Render Dashboard, select your `proglog` Web Service.
2. Go to **Settings** -> scroll down to **Deploy Hook**.
3. Copy the unique **Deploy Hook URL** (looks like `https://api.render.com/deploy/srv-xxxxx?key=yyyyy`).
4. In your GitHub repository (`Shreyas9468/proglog`), go to:
   **Settings -> Secrets and variables -> Actions -> New repository secret**.
5. Name: `RENDER_DEPLOY_HOOK_URL`  
   Value: *Paste the copied URL*.
6. Click **Add secret**.

🎉 **That's it!** Whenever you merge or push code to `main`, GitHub Actions will run tests, build the image, and trigger Render to deploy the live service automatically!

---

## 🧪 Testing Your Live Endpoint (Shareable on Resume!)

Once deployed, Render gives you a public HTTPS endpoint, for example:
`https://proglog-service.onrender.com`

### 1. Test using `grpcurl` (Command Line)
```bash
# List available gRPC services on your live cloud endpoint
grpcurl proglog-service.onrender.com:443 list
```

### 2. Test using the Go CLI tool (`getservers`)
```powershell
go run ./cmd/getservers/main.go -addr=proglog-service.onrender.com:443
```

---

## 📜 Resume Showcase Example

Add this section to your resume or portfolio:

> **Distributed Commit Log (Go, gRPC, Raft, Docker, CI/CD)**  
> - Designed and built a high-performance distributed append-only commit log in Go implementing gRPC, Serf discovery, and Raft consensus.
> - Containerized with Docker and implemented automated CI/CD workflows using GitHub Actions.
> - Live cloud deployment hosted at `https://proglog-service.onrender.com:443`.

---

## ☸️ Local Kubernetes / Helm Option

If you prefer to run a full 3-node Raft cluster locally using Helm and Kubernetes:

```powershell
# 1. Build local Docker image
make build-docker

# 2. Deploy Helm chart to local Kubernetes (Minikube / Docker Desktop)
helm upgrade --install proglog deploy/proglog
```
