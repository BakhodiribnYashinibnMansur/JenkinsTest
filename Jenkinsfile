pipeline {
    agent {
        docker { image 'golang:latest' }
    }
    stages {
        stage('Deploy') {
            steps {
                sh 'docker compose up -d '
            }
        }
    }
}