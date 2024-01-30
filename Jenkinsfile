pipeline {
    agent none
    environment {
        IMAGE_NAME = "las3r12/random_numbers"
        HUB_CRED_ID = "docker_hub"
    }
    stages {
        stage("build") {
            agent any
            steps {
                sh "docker build -t ${IMAGE_NAME}:${GIT_COMMIT} -t ${IMAGE_NAME}:latest ."
            }
        }
    }
}