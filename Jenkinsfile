pipeline {
    agent none
    environment {
        IMAGE_NAME = "las3r12/random_numbers"
        HUB_CRED_ID = "docker_hub"
        PROJECT_DIR = 'ponomarenko-random-numbers-back'
    }
    stages {
        stage("build") {
            agent any
            steps {
                sh "docker build -t ${IMAGE_NAME}:${GIT_COMMIT} -t ${IMAGE_NAME}:latest ."
            }
        }

        stage("push") {
            agent any
            steps {
                withCredentials([usernamePassword(credentialsId: "${HUB_CRED_ID}", usernameVariable: "HUB_USERNAME", passwordVariable: "HUB_PASSWORD")]){
                    sh 'docker login -u ${HUB_USERNAME} -p ${HUB_PASSWORD}'
                    sh 'docker push ${IMAGE_NAME}:${GIT_COMMIT}'
                    sh 'docker push ${IMAGE_NAME}:latest'
                }
            }
        }

        stage ("deploy") {
            agent any
            steps {
                withCredentials([
                    string(credentialsId: "production_ip", variable: "SERVER_IP"),
                    sshUserPrivateKey(credentialsId: "production_key", keyFileVariable: "SERVER_KEY", usernameVariable: "SERVER_USERNAME")
                    ]){
                        sh 'scp -i ${SERVER_KEY} ${SERVER_USERNAME}@${SERVER_IP} docker-compose.yaml ${SERVER_USERNAME}@${SERVER_IP}:${PROJECT_DIR}'
                        sh 'ssh -i ${SERVER_KEY} ${SERVER_USERNAME}@${SERVER_IP} docker-compose -f ${PROJECT_DIR}/docker-compose.yaml pull'
                        sh 'ssh -i ${SERVER_KEY} ${SERVER_USERNAME}@${SERVER_IP} docker-compose -f ${PROJECT_DIR}/docker-compose.yaml up'
                }
            }
        }
    }
}