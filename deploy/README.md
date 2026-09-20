# Deploying proglog to the Cloud (Free, No Credit Card)

This guide deploys **proglog** to [Koyeb](https://koyeb.com) — a free, always-on cloud platform that supports Docker and gRPC — with full CI/CD via GitHub Actions.

---

## Stack (100% Free, No Credit Card Required)

| Component | Service |
|---|---|
| CI/CD Pipeline | GitHub Actions |
| Container Registry | GitHub Container Registry (`ghcr.io`) |
| Cloud Host | Koyeb.com (free tier, always on, gRPC support) |

---

## One-Time Setup (~5 minutes)

### Step 1: Make your GitHub repository public

Go to your repo **Settings → General → Danger Zone → Change visibility → Public**.

> This allows `ghcr.io` to serve your Docker image publicly so Koyeb can pull it.

### Step 2: Create a Koyeb account (no credit card)

1. Go to [https://app.koyeb.com](https://app.koyeb.com)
2. Click **Continue with GitHub** — no credit card ever required
3. Complete sign-up

### Step 3: Get your Koyeb API Token

1. In Koyeb dashboard → click your avatar (top right) → **Account**
2. Go to **API** tab → **Create API Token**
3. Name it `github-actions`, copy the token

### Step 4: Add the secret to GitHub Actions

1. Go to your GitHub repo → **Settings → Secrets and variables → Actions**
2. Click **New repository secret**
3. Name: `KOYEB_API_KEY`
4. Value: paste the token you copied
5. Click **Add secret**

### Step 5: Create the Koyeb app (first-time only)

1. In Koyeb dashboard → click **Create App**
2. Select **Docker** as the deployment method
3. Docker image: `ghcr.io/shreyas9468/proglog:latest`
4. App name: `proglog`
5. Service name: `proglog`
6. Port: `8400`
7. Environment variables:
   ```
   PROGLOG_BOOTSTRAP=true
   PROGLOG_RPC_PORT=8400
   PROGLOG_DATA_DIR=/tmp/proglog
   ```
8. Click **Deploy**

After this one-time setup, **every push to `main`** auto-deploys a new version via GitHub Actions!

---

## CI/CD Pipeline Flow

```
git push → GitHub Actions CI runs tests
            ↓ (if tests pass)
           docker build
            ↓
           push to ghcr.io/shreyas9468/proglog:latest
            ↓
           deploy to Koyeb (auto, via GitHub Actions)
```

---

## Testing the Live Deployment

After deployment, Koyeb gives you a URL like:
`https://proglog-<hash>.koyeb.app`

Test it with the `getservers` tool:
```powershell
# TLS is handled by Koyeb's edge, so connect with TLS
go run ./cmd/getservers/main.go -addr=proglog-<hash>.koyeb.app:443
```

---

## Deploy to Your Own Kubernetes Cluster

If you have your own Kubernetes cluster, use the production Helm values:

```powershell
# Pull the latest image from ghcr.io
helm upgrade --install proglog deploy/proglog -f deploy/proglog/values-prod.yaml
```

---

## Architecture Note

The Koyeb free deployment runs proglog in **single-node mode** (`PROGLOG_BOOTSTRAP=true`). This means:
- One Raft leader (itself)
- No peer replication
- Ephemeral storage (data resets on restart)

This is **perfect for demos**. For production multi-node clustering, use the Helm chart with a Kubernetes StatefulSet (from Chapter 10).
