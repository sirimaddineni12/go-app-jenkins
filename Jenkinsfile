pipeline {
    agent any

    tools {
        go 'Go 1.20' // Replace with the name of Go version installed in Jenkins
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/your-username/your-go-repo.git'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -v'
            }
        }

        stage('Test') {
            steps {
                sh 'go test -v ./...'
            }
        }
    }

    post {
        always {
            echo 'Pipeline finished.'
        }
    }
}
