pipeline {
    agent any

    triggers {
        githubPush()
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Create version tag with short git commit
                    commit = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim()

                    // Build the image: go-api:<commit>
                    sh "docker build -t go-api:${commit} ."
                }
            }
        }

        stage('Run Container (test)') {
            steps {
                script {
                    // 1. Run the container in the background
                    // We give it a name 'test-api' so we can stop it easily later
                    sh "docker run -d --name test-api -p 9090:9090 go-api:${commit}"

                    try {
                        // 2. Wait for it to start
                        sh "sleep 5"

                        // 3. Test it
                        sh "curl --fail http://localhost:9090/"
                        echo "Test Passed!"
                    } finally {
                        // 4. Clean up: Stop and remove the container whether test passed or failed
                        sh "docker stop test-api || true"
                        sh "docker rm test-api || true"
                    }
                }
            }
        }

        stage('Tag and Save Latest') {
            steps {
                script {
                    sh "docker tag go-api:${commit} go-api:latest"
                }
            }
        }
    }
}
