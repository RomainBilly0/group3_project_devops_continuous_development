# ST2DCE-2526PSA01 — DevOps & Déploiement Continu | Groupe 3

**Efrei Paris** · INGE-3 · 2025–2026

> **Équipe :** Billy / Bussiere / Godfrin

---

## Architecture globale

<img src="Diagramme sans nom.drawio.png" alt="Schéma d'architecture" width="100%" />

---

## Stack technique

| Couche | Outil | Rôle |
| :--- | :--- | :--- |
| **Application** | Go (stdlib) | API REST légère, sans dépendances externes |
| **Conteneurisation** | Docker (multi-stage) | Packaging en image Alpine ~20 MB |
| **Cluster local** | KinD | Kubernetes in Docker pour l'environnement de dev/test |
| **Orchestration** | Kubernetes | Gestion des pods, réplication, service discovery |
| **Packaging K8s** | Helm | Déploiements reproductibles via `values.yaml` |
| **CI** | Jenkins | Build, push image, déclenchement ArgoCD |
| **CD / GitOps** | ArgoCD | Synchronisation Git → Cluster |
| **Métriques** | Prometheus | Scraping auto via annotations pod |
| **Logs** | Loki + Alloy | Agrégation et filtrage des logs |
| **Visualisation** | Grafana | Dashboards métriques + logs, alertes |
| **Stockage** | MinIO | Backend S3 pour les chunks Loki |

---

## Pipeline CI/CD — Enchaînement

```
┌─────────┐    push     ┌─────────┐   build    ┌────────────┐
│   Git   │ ──────────► │ Jenkins │ ─────────► │   Docker   │
│ (main)  │             │  (CI)   │            │ (image tag │
└─────────┘             └─────────┘            │  = git SHA)│
                                               └─────┬──────┘
                                                     │ load
                                                     ▼
                                               ┌─────────────┐
                                               │     KinD    │
                                               │  (cluster)  │
                                               └─────┬───────┘
                                                     │
                                                     ▼
                                               ┌─────────────┐
                            login + sync       │   ArgoCD    │
              Jenkins ────────────────────────►│   (GitOps)  │
                                               └─────┬───────┘
                                                     │ apply
                                                     ▼
                                          ┌──────────────────────┐
                                          │   Kubernetes         │
                                          │   Namespace: dev     │
                                          │   2 replicas go-api  │
                                          └──────────┬───────────┘
                                                     │ validate
                                                     ▼
                                          ┌──────────────────────┐
                                          │  curl /whoami → 200  │
                                          │  rollout status OK   │
                                          └──────────┬───────────┘
                                                     │ (main only)
                                                     ▼
                                          ┌──────────────────────┐
                                          │  ArgoCD sync PROD    │
                                          │  go-api-prod         │
                                          └──────────────────────┘
```

---

## Qui fait quoi

| Outil | Responsabilités |
| :--- | :--- |
| **Git** | Source of truth — versionne le code **et** les manifestes K8s |
| **Jenkins** | Orchestre la CI : checkout → `docker build` → `kind load` → `argocd sync` → validation `curl` |
| **Docker** | Compile le binaire Go (stage builder) et produit l'image finale Alpine (stage runtime) |
| **KinD** | Reçoit l'image via `kind load docker-image` — évite un registry externe en local |
| **ArgoCD** | Détecte chaque commit, compare l'état désiré (Git) à l'état courant (K8s) et applique le delta |
| **Kubernetes** | Planifie les pods, maintient 2 réplicas, expose le service via DNS interne |
| **Helm** | Paramétrise les déploiements K8s via `values.yaml` (Prometheus, Loki, Alloy) |
| **Prometheus** | Scrape automatiquement les métriques des pods annotés (`prometheus.io/scrape: "true"`) |
| **Alloy** | Collecte des logs distants, filtre les entrées Windows, les pousse vers Loki |
| **Loki** | Agrège et indexe les logs (TSDB v13, stockage MinIO, mode SimpleScalable) |
| **Grafana** | Visualise métriques (Prometheus) et logs (Loki) — alerte si pod crash > 5 redémarrages |
| **MinIO** | Stockage objet S3 local pour les chunks et index Loki |

---

## Environnements

| Environnement | Application ArgoCD | Namespace K8s | Déclenchement |
| :--- | :--- | :--- | :--- |
| **Dev** | `go-api-dev` | `development` | Chaque push |
| **Prod** | `go-api-prod` | *(namespace prod)* | Branche `main` uniquement |

---

## Application — API Go

**Port :** `8080`

| Méthode | Endpoint | Réponse |
| :--- | :--- | :--- |
| GET | `/` | `Welcome to the Web API!` |
| GET | `/aboutme` | `A little bit about me...` |
| GET | `/whoami` | JSON `{ "Title": "Group 3", "Names": "Billy/Bussiere/Godfrin", "State": "FR" }` |

L'image Docker est construite en **multi-stage** :
- Stage `builder` : `golang:1.21-alpine` — compile le binaire statique (CGO_ENABLED=0)
- Stage final : `alpine:latest` — copie uniquement le binaire, image ~20 MB

---

## Observabilité

```
Pods K8s
  ├── Métriques ──────► Prometheus ──────► Grafana (dashboards + alertes)
  └── Logs
        └── Alloy ─────► Loki (MinIO) ───► Grafana (explore / dashboards)
```

**Alerte configurée :** `KubernetesPodCrashLooping` — sévérité `critical` si un pod redémarre plus de 5 fois.

**Alloy pipeline :**
1. Fetch logs HTTP (polling 1 min)
2. Label `job: apache_task`
3. Drop des entrées contenant "Windows" (regex)
4. Push vers `loki-gateway.logging.svc.cluster.local`
