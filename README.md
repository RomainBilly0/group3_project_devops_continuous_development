# ST2DCE-2526PSA01 - DevOps and Continuous Deployment (INGE-3-SEM-A, INGE-3, INGE-3-PRO) Group 3


## Part 1
### Question 1
<img src="Diagramme sans nom.drawio.png" alt="Schema" />

### 1. Architecture Cible : Cloud-Native

| Couche | Outil(s) Principal(aux) | Rôle dans l'Architecture |
| :--- | :--- | :--- |
| **Conteneurisation (Base)** | **Docker** | Packaging des microservices en images isolées. |
| **Orchestration** | **Kubernetes** | Gère le cycle de vie, la mise à l'échelle et la résilience des conteneurs.  |
| **Gestion Déploiement** | **Helm** | Outil de packagement (Charts) pour des déploiements complexes et reproductibles sur K8s. |

### 2. Tool Chain pour le Déploiement Continu

La chaîne est divisée en trois phases clés : Intégration, Déploiement et Surveillance.

##### A. Intégration Continue (CI)

| Outil(s) | Action | Résultat |
| :--- | :--- | :--- |
| **Jenkins** | 1. Exécute les tests. 2. Construit l'image **Docker**. 3. Pousse l'image vers le Registre. | Image conteneur vérifiée et prête au déploiement. |

##### B. Déploiement Continu (CD) : Approche GitOps

| Outil(s) | Action | Résultat |
| :--- | :--- | :--- |
| **Jenkins / Git** | Met à jour le tag d'image dans le fichier **Helm** (`values.yaml`) et *commit* cette modification dans un dépôt **GitOps**. | Le dépôt GitOps reflète le nouvel état désiré (nouvelle version de l'application). |
| **ArgoCD / Flux** (Outil GitOps) | Détecte le *commit* Git, compare avec l'état actuel de **Kubernetes**, et applique le changement. | Déploiement automatique et continu de l'application en Production.  |

##### C. Observabilité et Surveillance

| Outil(s) | Fonction | Objectif dans le CD |
| :--- | :--- | :--- |
| **Prometheus** | Collecte des **Métriques** de performance et de santé. | Détection rapide des anomalies et des goulots d'étranglement après le déploiement. |
| **Loki** | Agrège les **Logs** des conteneurs. | Diagnostic et analyse des erreurs en temps réel. |
| **Grafana** | Visualisation unifiée des données de Prometheus et Loki. | Tableaux de bord pour le monitoring post-déploiement et les alertes. |
## 2
## 3

### Part 2
## 1
## 2
## 3

### Part 3
## 1
## 2 
## 3

### Bonus
 ## Part 1: Build the docker image using the buildpack utility and describe what you observe in comparison with the Dockerfile option.
 ## Part 2: Configure another alert and send it by e-mail to abdoul-aziz.zakarimadougou@intervenants.efrei.net.
 
### Tools
- Graphana
- Prometheus
- Jenkins
- Graphan/Loki
- Kubernetes
- Docker
- Kubernetes/helm
