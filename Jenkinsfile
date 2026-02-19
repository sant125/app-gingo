pipeline {
    agent {
        kubernetes {
            defaultContainer 'go'
            yaml '''
apiVersion: v1
kind: Pod
spec:
  serviceAccountName: jenkins-agent
  containers:
    - name: go
      image: golang:1.22
      command: [sleep, infinity]
      resources:
        requests: { cpu: "500m", memory: "512Mi" }
        limits:   { cpu: "1",   memory: "1Gi"   }
    - name: dind
      image: docker:24-dind
      securityContext:
        privileged: true
      env:
        - name: DOCKER_TLS_CERTDIR
          value: ""
      resources:
        requests: { cpu: "250m", memory: "256Mi" }
        limits:   { cpu: "500m", memory: "512Mi" }
'''
        }
    }

    environment {
        ECR_REPO_URI       = '123456789012.dkr.ecr.us-east-1.amazonaws.com/gin-tattoo'
        AWS_DEFAULT_REGION = 'us-east-1'
        INFRA_REPO         = 'https://github.com/sant125/aws-devops.git'
    }

    stages {
        stage('Setup') {
            steps {
                script {
                    env.COMMIT    = sh(script: 'git rev-parse --short=8 HEAD', returnStdout: true).trim()
                    env.NAMESPACE = env.BRANCH_NAME == 'main' ? 'prod' : 'homolog'
                    env.IMAGE_TAG = env.BRANCH_NAME == 'main' ? env.COMMIT : "dev-${env.COMMIT}"
                    echo "branch=${env.BRANCH_NAME} | namespace=${env.NAMESPACE} | tag=${env.IMAGE_TAG}"
                }
            }
        }

        stage('CI') {
            steps {
                sh 'go vet ./...'
                sh 'go test ./... -count=1 -timeout 60s -coverprofile=coverage.out -covermode=atomic'
                sh 'go install golang.org/x/vuln/cmd/govulncheck@latest && govulncheck ./...'

                withSonarQubeEnv('sonarqube') {
                    sh '''
                        wget -q https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-6.0.0.4432-linux-x64.zip \
                          -O /tmp/sonar.zip && unzip -q /tmp/sonar.zip -d /tmp/
                        /tmp/sonar-scanner-6.0.0.4432-linux-x64/bin/sonar-scanner \
                          -Dsonar.go.coverage.reportPaths=coverage.out
                    '''
                }
                timeout(time: 5, unit: 'MINUTES') {
                    waitForQualityGate abortPipeline: true
                }
            }
        }

        stage('Build & Push') {
            when { anyOf { branch 'developer'; branch 'main' } }
            steps {
                container('dind') {
                    sh """
                        apk add --no-cache aws-cli
                        aws ecr get-login-password --region ${AWS_DEFAULT_REGION} | \
                          docker login --username AWS --password-stdin ${ECR_REPO_URI}
                        docker build -t ${ECR_REPO_URI}:${IMAGE_TAG} .
                        docker push ${ECR_REPO_URI}:${IMAGE_TAG}
                    """
                }
            }
        }

        stage('Deploy') {
            when { anyOf { branch 'developer'; branch 'main' } }
            steps {
                withCredentials([string(credentialsId: 'github-token', variable: 'GH_TOKEN')]) {
                    sh """
                        git clone https://x-access-token:\${GH_TOKEN}@\${INFRA_REPO#https://} /tmp/infra
                        sed -i "s|image: .*gin-tattoo.*|image: ${ECR_REPO_URI}:${IMAGE_TAG}|g" \
                          /tmp/infra/manifests/gin-tattoo-${NAMESPACE}/deployment.yaml
                        cd /tmp/infra
                        git config user.email "jenkins@noreply" && git config user.name "Jenkins"
                        git add . && git diff --cached --quiet || \
                          git commit -m "ci(${NAMESPACE}): gin-tattoo → ${IMAGE_TAG} [skip ci]"
                        git push
                    """
                }
            }
        }
    }

    post {
        failure {
            echo "Pipeline falhou — branch: ${env.BRANCH_NAME}"
        }
        success {
            echo "Pipeline OK — branch: ${env.BRANCH_NAME}"
        }
    }
}
