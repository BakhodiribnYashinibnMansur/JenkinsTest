pipeline {
    agent {
        docker { image 'golang:latest' }
    }
    stages {
        stage('Initialize'){
        def dockerHome = tool 'docker'
        env.PATH = "${dockerHome}/bin:${env.PATH}"
    }
        stage('Deploy') {
            steps {
                sh 'docker  build  -t golang . '
                sh 'docker run -p 7777:7777 golang . '
            }
        }
    }
}