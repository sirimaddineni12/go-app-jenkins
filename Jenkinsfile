pipeline {
    agent any

    stages {
        stage('Checkout Code') {
            steps {
                git branch: 'main', url: 'https://github.com/sirimaddineni12/go-app-jenkins.git'
            }
        }

        stage('Install Go') {
            steps {
                sh 'sudo apt-get update && sudo apt-get install -y golang'
            }
        }

        stage('Build Go Application') {
            steps {
                sh 'go build -o app main.go'
            }
        }

        stage('Run Tests') {
            steps {
                sh 'go test'
            }
        }

        stage('Run Application') {
            steps {
                sh './app'
            }
        }
    }
}
