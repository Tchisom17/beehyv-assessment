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
                bat 'echo Building...'
                // For Go project, you might use: bat 'go build -o yourproject.exe'
            }
        }

        stage('Test') {
            steps {
                // Example test step, replace with your test commands
                bat 'echo Running tests...'
                // For Go project, you might use: bat 'go test ./...'
            }
        }

        stage('Deploy') {
            steps {
                // Example deploy step, replace with your deploy commands
                bat 'echo Deploying...'
                // Add actual deployment steps here
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
