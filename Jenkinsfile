def project = 'britzc-devops'
def appName = 'icedoapp-frontend'
def feSvcName = "${appName}"
def imageTag = "gcr.io/${project}/${appName}:${env.BRANCH_NAME}.${env.BUILD_NUMBER}"

pipeline {
  agent {
    kubernetes {
      label 'icedo-app'
      defaultContainer 'jnlp'
      yaml """
apiVersion: v1
kind: Pod
metadata:
labels:
  component: ci
spec:
  # Use service account that can deploy to all namespaces
  serviceAccountName: cd-jenkins 
  containers:
  - name: golang
    image: golang:1.10
    command:
    - cat
    tty: true
    volumeMounts:
    - mountPath: '/opt/app/shared'
      name: sharedvolume
  - name: gcloud
    image: gcr.io/cloud-builders/gcloud
    command:
    - cat
    tty: true
    volumeMounts:
    - mountPath: '/opt/app/shared'
      name: sharedvolume
  - name: kubectl
    image: gcr.io/cloud-builders/kubectl
    command:
    - cat
    tty: true
  volumes:
  - name: sharedvolume
    emptyDir: {}
"""
}
  }

    stages {

        stage("Compile Binary") {
            steps{
                slackSend message:"${feSvcName} Build started ${env.BUILD_NUMBER}"

                container('golang'){

                        script{
                            try{
                                sh "go get github.com/nats-io/go-nats"
                                sh "GOOS=linux go build -o icedoapp"
                            } catch (error){
                                throw error
                            }
                        }

                }
            }
        }
        
        stage('Build and push image with Container Builder') {
            steps {
                container('gcloud') {

                    //sh "mv /opt/app/shared/${env.BUILD_NUMBER}/icedoapp ."
                    //sh "mv /opt/app/shared/${env.BUILD_NUMBER}/Dockerfile ."
                    sh "PYTHONUNBUFFERED=1 gcloud builds submit -t ${imageTag} ."

                }
            }
        }

    stage('Deploy Canary') {
      // Canary branch
      when { branch 'canary' }
      steps {
        slackSend message:"${feSvcName} canary deployment started"

        container('kubectl') {
          // Change deployed image in canary to the one we just built
          sh("sed -i.bak 's#gcr.io/cloud-solutions-images/icedoapp:1.0.0#${imageTag}#' ./k8s/canary/*.yaml")
          sh("kubectl --namespace=production apply -f k8s/services/")
          sh("kubectl --namespace=production apply -f k8s/canary/")
          sh("echo http://`kubectl --namespace=production get service/${feSvcName} -o jsonpath='{.status.loadBalancer.ingress[0].ip}'` > ${feSvcName}")
        } 

        slackSend message:"${feSvcName} canary deployment completed"
      }
    }
    stage('Deploy Production') {
      // Production branch
      when { branch 'master' }
      steps{
        slackSend message:"${feSvcName} production deployment started"

        container('kubectl') {
        // Change deployed image in canary to the one we just built
          sh("sed -i.bak 's#gcr.io/cloud-solutions-images/icedoapp:1.0.0#${imageTag}#' ./k8s/production/*.yaml")
          sh("kubectl --namespace=production apply -f k8s/services/")
          sh("kubectl --namespace=production apply -f k8s/production/")
          sh("echo http://`kubectl --namespace=production get service/${feSvcName} -o jsonpath='{.status.loadBalancer.ingress[0].ip}'` > ${feSvcName}")
        }

        slackSend message:"${feSvcName} production deployment completed"
      }
    }
    stage('Deploy Dev') {
      // Developer Branches
      when { 
        not { branch 'master' } 
        not { branch 'canary' }
      } 
      steps {
        slackSend message:"${feSvcName} ${env.BRANCH_NAME} deployment started"

        container('kubectl') {
          // Create namespace if it doesn't exist
          sh("kubectl get ns ${env.BRANCH_NAME} || kubectl create ns ${env.BRANCH_NAME}")
          // Don't use public load balancing for development branches
          sh("sed -i.bak 's#LoadBalancer#ClusterIP#' ./k8s/services/frontend.yaml")
          sh("sed -i.bak 's#gcr.io/cloud-solutions-images/icedoapp:1.0.0#${imageTag}#' ./k8s/dev/*.yaml")
          sh("kubectl --namespace=${env.BRANCH_NAME} apply -f k8s/services/")
          sh("kubectl --namespace=${env.BRANCH_NAME} apply -f k8s/dev/")
          echo 'To access your environment run `kubectl proxy`'
          echo "Then access your service via http://localhost:8001/api/v1/proxy/namespaces/${env.BRANCH_NAME}/services/${feSvcName}:80/"
        }

        slackSend message:"${feSvcName} ${env.BRANCH_NAME} deployment completed"
      }     
    }

    }

post {
    always {
        echo "One way or another, I have finished"
            deleteDir()
    }
    success {
        slackSend message:"${feSvcName} Build successful ${env.BUILD_NUMBER}"
        echo "I succeeeded!"
    }
    unstable {
        echo "I am unstable :/"
    }
    failure {
        slackSend message:"${feSvcName} Build failed ${env.BUILD_NUMBER}"
        echo "I failed :("
    }
    changed {
        echo "Things were different before..."
    }
}
}
