pipeline {
    agent { 
        node {
            label 'jenkins-aci-agent'
            }
      }
    triggers {
        pollSCM '* * * * *'
    }
    stages {
        stage('Build') {
            steps {
                echo 'Building'
                sh 'echo "Build Step"'
            }
        }
        
        stage('Test') {
            steps {
                echo 'Testing'
                sh 'echo "Test Step"'
            }
        }

        stage('Deliver') {
            steps {
                echo 'Delivering'
                sh 'echo "Deliver Step"'
            }
        }
    }
}
