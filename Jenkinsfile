pipeline{
    agent any

    tools {
        go 'NATS-LOG GO'
    }

    stages{
            stage('Checkout'){
                steps{
                    checkout scm
                }
            }

            stage('Go Vet'){
                steps{
                    dir('docs/intern/NATS-LOG/go-api'){
                        sh 'go vet ./...'
                    }
                }
            }

            stage('Docker Build'){
                steps{
                    sh 'docker compose build'
                }
            }

            stage('Deploy'){
                when {
                    branch 'main'
                }
                steps{
                    sh 'docker compose up -d --build'
                }
            }
        }
    }
