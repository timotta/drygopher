pipeline {
    agent {
        docker {
            image 'golang:1.10'
            reuseNode true
            customWorkspace 'workspace/drygopher'
            args '-v workspace/drygopher:/go/src/drygopher -w /go/src/drygopher'
        }
    }
    stages {
        stage('Prepare') {
            steps {
                echo 'Preparing build environment...'
                sh 'go get -u github.com/vektra/mockery/.../'
                sh 'go get -u github.com/golang/dep/cmd/dep'
            }
        }
        stage('Build') {
            steps {
                echo 'Building...'
                sh 'echo pwd && pwd'
                sh 'echo drygopher directory && ls go/src/drygopher'
                sh 'make build'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
                sh 'make test'
            }
        }
    }
}