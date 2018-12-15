def project = 'britzc-devops'
def appName = 'icedoapp'
def svcName = "${appName}"
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

    // triggers {
    //     pollSCM "* * * * *"
    // }

    stages {
/*
        stage("Execute Unit Tests") {
            steps{
                container('golang'){

                    script{
                        try{
                            sh "go test -tags=unit_test"
                        } catch (error){
                            throw error
                        }
                    }

                }
            }
        }
        */

        stage("Compile Binary") {
            steps{
                container('golang'){

                        script{
                            try{
                                sh "GOOS=linux go build -o icedoapp"
                                //sh "mkdir -p /opt/app/shared/${env.BUILD_NUMBER}"
                                //sh "mv icedoapp /opt/app/shared/${env.BUILD_NUMBER}/."
                                //sh "mv Dockerfile /opt/app/shared/${env.BUILD_NUMBER}/."
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

/*
        stage('Deploy Dev') {
            steps {
                container('kubectl') {
                    dir("build"){
                        // Create namespace if it doesn't exist
                        sh("kubectl get ns sandbox || kubectl create ns sandbox")
                        // Don't use public load balancing for development branches
                        sh("sed -i.bak 's#LoadBalancer#ClusterIP#' ./k8s/services/frontend.yaml")
                        sh("sed -i.bak 's#gcr.io/cloud-solutions-images/icedo:1.0.0#${imageTag}#' ./k8s/dev/*.yaml")
                        sh("kubectl --namespace=sandbox apply -f k8s/services/")
                        sh("kubectl --namespace=sandbox apply -f k8s/dev/")
                        echo 'To access your environment run `kubectl proxy`'
                        echo "Then access your service via http://localhost:8001/api/v1/proxy/namespaces/sandbox/services/${feSvcName}:80/"
                    }
                }
          }     
    }
    */

    }

    //     steps {
    //         lock('storageupload'){
    //             dir("code") {
    //                 googleStorageUpload bucket: 'gs://icedo', credentialsId: 'britz-devops', pattern: 'icedoapp'
    //             }
    //         }
    //         step([$class: 'InfluxDbPublisher', customData: null, customDataMap: null, customPrefix: null, target: 'grafana'])
    //     }
    // }

    // stage("Execute Environment UAT"){
    //     steps{
    //         build job: "Deploy Environment", parameters: [
    //             string(name: 'ENVIRONMENT', value: "uat")
    //         ]
    //         step([$class: 'InfluxDbPublisher', customData: null, customDataMap: null, customPrefix: null, target: 'grafana'])
    //     }
    // }

post {
    always {
        echo "One way or another, I have finished"
            deleteDir()
            // step([$class: 'InfluxDbPublisher', customData: null, customDataMap: null, customPrefix: null, target: 'grafana'])
    }
    success {
        echo "I succeeeded!"
    }
    unstable {
        echo "I am unstable :/"
    }
    failure {
        echo "I failed :("
    }
    changed {
        echo "Things were different before..."
    }
}
}
