pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Clone the repository
                git 'https://github.com/Tchisom17/beehyv-assessment.git'
            }
        }

        stage('Build') {
            steps {
                // Example build step, replace with your build commands
                sh 'echo "Building..."'
                // For a Maven project, you might use: sh 'mvn clean install'
            }
        }

        stage('Test') {
            steps {
                // Example test step, replace with your test commands
                sh 'echo "Running tests..."'
                // For a Maven project, you might use: sh 'mvn test'
            }
        }

        stage('Deploy') {
            steps {
                // Example deploy step, replace with your deploy commands
                sh 'echo "Deploying..."'
                // For deployment, you might copy files, upload artifacts, etc.
            }
        }
    }

    post {
        always {
            // Clean up, notify, or other steps to always run
            echo 'This will always run'
        }
        success {
            // Actions to perform if the pipeline succeeds
            echo 'Pipeline succeeded!'
        }
        failure {
            // Actions to perform if the pipeline fails
            echo 'Pipeline failed!'
        }
    }
}
