# Cthulhu CLI

Cthulhu is a Golang‑based command‑line application that makes Docker‑to‑Kubernetes workflows effortless for DevOps teams. It is built on top of the [Cobra](https://github.com/spf13/cobra) framework with configuration powered by [Viper](https://github.com/spf13/viper).

## Core Features

*   **Declarative Configuration & Templating:** Helm‑compatible templating for easy configuration.
*   **Automated Deployments:** Automated image build, push, and rollout.
*   **One-Command Rollbacks:** Simple, one-command rollbacks to previous deployment versions.
*   **Real-Time Monitoring:** Real-time monitoring hooks, ready for Prometheus integration.
*   **Multi-Cluster Operations:** First-class support for multi-cluster operations with AWS EKS.

## Commands

*   `cthulhu deploy [path]`: Build, push and deploy a Docker image to Kubernetes.
*   `cthulhu rollback [deployment]`: Rollback a deployment to the previous revision.
*   `cthulhu monitor [deployment]`: Tail pod logs & show Prometheus metrics for a deployment.
*   `cthulhu config`: View or edit configuration.
*   `cthulhu cluster`: List, create, or switch Kubernetes clusters.