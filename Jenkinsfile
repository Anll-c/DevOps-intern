pipeline {
    agent any

    options {
        disableConcurrentBuilds()
        timestamps()
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Go Vet') {
            steps {
                dir('docs/intern/NATS-LOG/go-api') {
                    sh 'go vet ./...'
                }
            }
        }

        stage('Docker Build') {
            steps {
                dir('docs/intern/NATS-LOG/Jenkins') {
                    sh 'docker compose build'
                }
            }
        }

        stage('Deploy') {
            when {
                branch 'main'
            }
            steps {
                dir('docs/intern/NATS-LOG/Jenkins') {
                    sh 'docker compose up -d --remove-orphans'
                }
            }
        }
    }
}