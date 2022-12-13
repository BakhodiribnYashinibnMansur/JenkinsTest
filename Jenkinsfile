pipeline {
    agent
    stages {
        stage("Build") {
            steps {
                cleanWs()
                git branch: 'sprint-42', changelog: false, credentialsId: 'gitlab-credentials-suhrob', poll: false, url: 'https://gitlab.com/thinkland-frontend/sevimli-frontend-app.git'
                    sh "docker compose up -d "
            }
        }
    }
}