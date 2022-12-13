pipeline {
    agent any
    stages {
        stage('Build') {
            agent {
                docker {
                    // Run the container on the node specified at the
                    // top-level of the Pipeline, in the same workspace,
                    // rather than on a new node entirely:
                    image 'golang:latest'
                    reuseNode true
                }
            }
            steps {
                sh 'docker -v'
            }
        }
    }
}