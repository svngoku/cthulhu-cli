# Cthulhu CLI

[![Go Version](https://img.shields.io/badge/go-1.24.4-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Cthulhu is a Golang‑based command‑line application that makes Docker‑to‑Kubernetes workflows effortless for DevOps teams. It is built on top of the [Cobra](https://github.com/spf13/cobra) framework with configuration powered by [Viper](https://github.com/spf13/viper).

## 🚀 What's New in Latest Update

- **Enhanced EKS Integration**: Improved AWS EKS cluster management with AWS SDK v2
- **Updated Dependencies**: Updated to Go 1.24.4 with latest Kubernetes client-go v0.33.2
- **Docker Integration**: Enhanced Docker dependency management and cleaner module structure
- **Better Logging**: Integrated Zap logger for structured logging across all operations
- **Task Runner**: Added Taskfile for streamlined build and development workflows

## Core Features

*   **Declarative Configuration & Templating:** Helm‑compatible templating for easy configuration
*   **Automated Deployments:** Automated image build, push, and rollout
*   **One-Command Rollbacks:** Simple, one-command rollbacks to previous deployment versions
*   **Real-Time Monitoring:** Real-time monitoring hooks, ready for Prometheus integration
*   **Multi-Cluster Operations:** First-class support for multi-cluster operations with AWS EKS
*   **AWS EKS Native Support:** Built-in EKS cluster creation and management
*   **Structured Logging:** Advanced logging with Zap for better observability

## 📋 Prerequisites

- Go 1.24.4 or later
- Docker installed and running
- kubectl configured for your Kubernetes clusters
- AWS CLI configured (for EKS operations)
- [Task](https://taskfile.dev/) (optional, for development)

## 🛠️ Installation

### From Source

```bash
git clone https://github.com/your-username/cthulhu-cli.git
cd cthulhu-cli
go build -o bin/cthulhu main.go
```

### Using Task (Recommended for Development)

```bash
# Install Task if you haven't already
go install github.com/go-task/task/v3/cmd/task@latest

# Build the application
task build

# Run directly
task run
```

## 📖 Commands

*   `cthulhu deploy [path]`: Build, push and deploy a Docker image to Kubernetes
*   `cthulhu rollback [deployment]`: Rollback a deployment to the previous revision
*   `cthulhu monitor [deployment]`: Tail pod logs & show Prometheus metrics for a deployment
*   `cthulhu config`: View or edit configuration
*   `cthulhu cluster`: List, create, or switch Kubernetes clusters

## ☸️ Running on AWS EKS

Cthulhu provides native support for AWS EKS clusters with simplified setup and management.

### EKS Setup Prerequisites

1. **AWS CLI Configuration**:
   ```bash
   aws configure
   # Ensure your AWS credentials have EKS permissions
   ```

2. **IAM Role for EKS**:
   Create an EKS service role with the following managed policies:
   - `AmazonEKSClusterPolicy`
   - `AmazonEKSWorkerNodePolicy`
   - `AmazonEKS_CNI_Policy`
   - `AmazonEC2ContainerRegistryReadOnly`

### Configuration for EKS

1. **Update your configuration file** (`~/.cthulhu.yaml` or use `--config`):

```yaml
registry: your-ecr-registry.amazonaws.com
namespace: default
imageTag: latest
aws:
  region: us-east-1
  roleArn: arn:aws:iam::YOUR-ACCOUNT-ID:role/eks-service-role
clusters:
  - name: my-prod-cluster
    provider: eks
    region: us-east-1
  - name: my-dev-cluster
    provider: eks
    region: us-west-2
```

2. **Create an EKS cluster**:

```bash
# Create a new EKS cluster
cthulhu cluster create my-eks-cluster --provider eks --region us-east-1

# List available clusters
cthulhu cluster list

# Switch to EKS cluster context
cthulhu cluster switch my-eks-cluster
```

### EKS Deployment Workflow

1. **Deploy to EKS**:

```bash
# Deploy your application to EKS
cthulhu deploy ./my-app --cluster my-eks-cluster

# Monitor the deployment
cthulhu monitor my-app-deployment --cluster my-eks-cluster

# Check deployment status
kubectl get deployments -n default
```

2. **ECR Integration**:

Cthulhu automatically handles ECR authentication when using EKS:

```bash
# Docker login to ECR is handled automatically
# Just ensure your AWS credentials have ECR permissions
cthulhu deploy ./my-app --registry your-account.dkr.ecr.us-east-1.amazonaws.com
```

3. **Rollback on EKS**:

```bash
# Rollback to previous version
cthulhu rollback my-app-deployment --cluster my-eks-cluster
```

### EKS Best Practices

- **Security**: Always use IAM roles with least privilege principles
- **Networking**: Configure VPC and security groups appropriately
- **Monitoring**: Use CloudWatch integration for EKS cluster monitoring
- **Cost Management**: Use spot instances for non-critical workloads
- **Scaling**: Configure cluster autoscaler for dynamic scaling

### Troubleshooting EKS Issues

```bash
# Check EKS cluster status
aws eks describe-cluster --name my-eks-cluster --region us-east-1

# Verify kubectl context
kubectl config current-context

# Update kubeconfig for EKS
aws eks update-kubeconfig --region us-east-1 --name my-eks-cluster

# Enable verbose logging
cthulhu deploy ./my-app --verbose
```

## ⚙️ Configuration

Cthulhu uses a YAML configuration file. By default, it looks for `~/.cthulhu.yaml`, but you can specify a custom path:

```bash
cthulhu --config /path/to/config.yaml deploy ./my-app
```

### Example Configuration

```yaml
registry: registry.example.com
namespace: default
imageTag: latest
aws:
  region: us-east-1
  roleArn: arn:aws:iam::123456789012:role/eks-service-role
clusters:
  - name: prod
    provider: eks
    region: us-east-1
  - name: staging
    provider: eks
    region: us-west-2
```

## 🏗️ Development

### Project Structure

```
cthulhu-cli/
├── cmd/              # Cobra commands
├── internal/         # Internal packages
│   ├── aws/         # AWS/EKS integration
│   ├── docker/      # Docker operations
│   └── kube/        # Kubernetes operations
├── pkg/             # Public packages
│   └── utils/       # Utilities and logging
├── config/          # Configuration files
├── Taskfile.yml     # Task runner configuration
└── main.go          # Application entry point
```

### Building and Running

```bash
# Using Go directly
go build -o bin/cthulhu main.go
./bin/cthulhu --help

# Using Task
task build
task run

# Clean build artifacts
task clean
```

### Dependencies

Key dependencies include:
- AWS SDK Go v2 for EKS integration
- Kubernetes client-go v0.33.2
- Docker SDK for container operations
- Cobra for CLI framework
- Viper for configuration management
- Zap for structured logging

## 🛣️ Roadmap

Our short-term goals are focused on stability and expanding cloud support:

1. **Add Unit Tests:** Provide test coverage for `internal` and `pkg` packages.
2. **Canary Deployments:** Implement canary release logic for safer rollouts.
3. **GKE/AKS Support:** Abstract provider layer to support GKE and AKS clusters.
4. **Plugin System:** Allow custom commands through a plugin architecture.
5. **v1 Release:** Polish the CLI and publish a stable v1.0.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
