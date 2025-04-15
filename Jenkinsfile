pipeline {
    agent any

    environment {
        GO111MODULE = 'on'
        GOPATH = "${env.WORKSPACE}/go"
        PATH = "${env.PATH}:${env.GOPATH}/bin"
    }

    stages {
        stage('Checkout Code') {
            steps {
                git url: 'https://github.com/sirimaddineni12/go-app-jenkins.git', branch: 'main'
            }
        }

        stage('Install Go') {
            steps {
                sh '''
                    echo "Updating packages..."
                    sudo yum update -y
                    
                    echo "Installing Go..."
                    sudo yum install -y golang
                    
                    echo "Go version:"
                    go version
                '''
            }
        }

        stage('Build Go Application') {
            steps {
                sh '''
                    echo "Building Go application..."
                    go build -o myapp main.go
                '''
            }
        }

        stage('Run Tests') {
            steps {
                sh '''
                    echo "Running Go tests..."
                    go test ./...
                '''
            }
        }

        stage('Run Application') {
            steps {
                sh '''
                    echo "Running the Go application..."
                    ./myapp &
                '''
            }
        }
    }
}

