#!/usr/bin/env groovy

pipeline {
    agent none
    stages {
        stage('Build Synergy') {
            agent { 
                docker {
                    label 'docker'
                    image 'nexus.d.lowes.com:8800/digital/irs-image-terratest:0.11.7-r0-B1'
                    args '-u root:root -v /home/jenkins/.ssh:/root/.ssh'
                } 
            }
            environment {
                GOOGLE_CREDENTIALS = credentials('SANDBOX_GOOGLE_CREDENTIALS')
            }
            steps {
                    script {
                        sh "go mod vendor"
                        sh "export GOOS=linux && export GOARCH=amd64 && go build -o bin/synergy"
                        sh "ls -lhrt"
                    }
                }
            }    
        stage('Push to GCS') {
            agent any
            environment {
                GOOGLE_CREDENTIALS = credentials('SANDBOX_GOOGLE_CREDENTIALS')
            }
            steps {
                    script {
                        sh "curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-232.0.0-linux-x86_64.tar.gz"
                        sh "tar -xvf google-cloud-sdk-232.0.0-linux-x86_64.tar.gz"
                        sh "./google-cloud-sdk/bin/gcloud auth activate-service-account --key-file=$GOOGLE_CREDENTIALS"
                        sh "./google-cloud-sdk/bin/gcloud config set project gcp-ushi-platform-sandbox-npe"
                        sh "./google-cloud-sdk/bin/gsutil cp bin/synergy gs://synergy-tub/build-${env.BUILD_NUMBER}/synergy"
                    }
                }
            post {
                always {
                    cleanWs()
                    script {
                            sh "echo 'Synergy build complete'"
                    }
                }
                failure {
                    script {
                        mail (to: 'shivaprasad.hs@lowes.com',
                            subject: "Synergy Build Failure. JOB:'${env.JOB_NAME}' BUILD_NO:(${env.BUILD_NUMBER}).",
                            body: "Please visit ${env.BUILD_URL} for further information."
                        );
                    }
                }  
                success {
                    script {
                        mail (to: 'shivaprasad.hs@lowes.com',
                            subject: "Synergy Build Successful. JOB:'${env.JOB_NAME}' BUILD_NO:(${env.BUILD_NUMBER}).",
                            body: "Please visit ${env.BUILD_URL} for further information."
                        );
                    }
                }                           
            }
        }     
    }
}