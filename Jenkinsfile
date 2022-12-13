pipeline {
    agent {
        docker { image 'golang:latest' }
    }
    stages {
        stage('Deploy') {
            steps {
                sh 'docker -v'
                sh 'docker  build  -t golang . '
                sh 'docker run -p 7777:7777 golang . '
            }
        }
    }
}