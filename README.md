# Developer Experience On Kubernertes

## Part I: The Inner Development Loop
- **Chapter 1** - Crafting Experiences for Cloud Native Development
  - Tutorials:
    - Clone source code, kind create , kubectl apply -f k8s/ (prove that it sucks) 
- **Chapter 2** - Inner Loop: Development Environments
  - Tutorials:
    - DevContainers development environment: allow a developer to start working on a simple application (UI + REST Endpoints) with DevPod
    - Changing the application lifecycle Tools (Tilt, Skaffold, ko, Spring Boot / Quarkus plugins)
    - Application Infrastructure / Dependencies Tools (docker-compose, operators for Kubernetes) 
- **Chapter 3** - Inner Loop: Containers and Kubernetes
  - Integration Tests as part of developers workflows (Scripts, TestContainers, Dagger)
    - Running with Kubernetes (Microcks)
  - Developer's and Kubernetes
    - Visibility with Headlamp
    - Load images local images in Kubernetes with KinD
    - Debug? Is it possible? Telepresence approach vs `mirrord`
   
## Part II: Cloud Native Runtimes
- **Chapter 4** - Crafting Experiences for Complex Environments
  - Frameworks, abstractions and Cloud Provider SDKs
  - Integrating with AI APIs, use case and highlight with local model
- **Chapter 5** - Developer Experience: Working with Cloud Native Runtimes
  - Crosscutting concerns, architectural patterns and complex integrations with Dapr
  - Feature flagging and local development with Open Feature
- **Chapter 6** - Developer Experience: Working with AI

## Part III: The Path to Production
