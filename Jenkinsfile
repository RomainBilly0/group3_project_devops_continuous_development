pipeline {
    agent any
    environment {
        ARGOCD_SERVER = "localhost:8080" // requires forwarding 443 -> 8080
        APP_NAME = "go-api-dev"
        CLUSTER_NAME = "efrei-devops-project"
    }
    stages {
        stage('Checkout') {
            steps { checkout scm }
        }
        stage('Build & Load Image') {
            steps {
                script {
                    commit = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim()
                    sh "docker build -t go-api:${commit} ."
                    sh "kind load docker-image go-api:${commit} --name ${CLUSTER_NAME}"
                    sh "docker tag go-api:${commit} go-api:latest"
                    sh "kind load docker-image go-api:latest --name ${CLUSTER_NAME}"
                }
            }
        }
        
        stage('Deploy via ArgoCD') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'argo-admin-pwd', passwordVariable: 'ARGO_PWD', usernameVariable: 'ARGO_USER')]) {
                    script {
                        sh 'argocd login ${ARGOCD_SERVER} --username ${ARGO_USER} --password ${ARGO_PWD} --insecure --grpc-web'
                        
                        sh "argocd app sync ${APP_NAME}"
                        sh "argocd app wait ${APP_NAME}"
                    }
                }
            }
        }

        stage('Validate Dev') {
            steps {
                sh "kubectl run test-curl --rm -i --restart=Never --image=curlimages/curl -- curl --fail http://go-api.development.svc.cluster.local:8080/whoami"
            }
        }
        stage('Promote to Prod') {
            when {
                branch 'main'
            }
            steps {
                sh "argocd app sync go-api-prod"
            }
        }
    }
}