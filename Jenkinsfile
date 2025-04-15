pipeline {
    agent any

    environment {
        APP_NAME = "go-app"
    }

    stages {
        stage('Checkout Code') {
            steps {
                echo 'Checking out code from GitHub repository...'
                git url: 'https://github.com/sirimaddineni12/go-app-jenkins.git', branch: 'main'
            }
        }

        stage('Install Go') {
            steps {
                echo 'Checking Go version...'
                sh 'go version'
            }
        }

        stage('Build Go Application') {
            steps {
                echo 'Building Go application...'
                sh 'go build -o ${APP_NAME}'
            }
        }

        stage('Run Tests') {
            steps {
                echo 'Running Go tests...'
                sh 'go test ./...'
            }
        }

        stage('Run Application') {
            steps {
                echo 'Running Go application...'
                sh './${APP_NAME} &'
            }
        }
    }

    post {
        always {
            echo 'Cleaning up...'
        }
    }
}
