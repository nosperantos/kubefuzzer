# KubeFuzzer

KubeFuzzer is a tool designed to fuzz Kubernetes API endpoints and discover hidden, undocumented, or misconfigured resources within a Kubernetes cluster. By leveraging a customizable dictionary of endpoint words and patterns, it systematically probes the Kubernetes API server for common, obscure, or forbidden paths that may indicate misconfigurations, excessive permissions, or attack surface exposure.

## What Are We Trying to Achieve?

KubeFuzzer aims to help security professionals, cloud engineers, and penetration testers identify Kubernetes API endpoints in a cluster that may be:

- Overly permissive or misconfigured
- Exposing sensitive resources
- Unintentionally reachable due to RBAC errors or software flaws
- Useful for Red Team or Blue Team operations, lateral movement, or privilege escalation

By fuzzing the endpoints, the tool assists in mapping the attack surface and locating endpoints that are not immediately visible from standard Kubernetes documentation or tooling.

## Background: Kubernetes & API Security

Kubernetes manages clusters using a rich API-centric architecture. Access to resources (pods, secrets, nodes, etc.) is controlled through the API server, which enforces access policies via [Role-Based Access Control (RBAC)](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) and other mechanisms.

In Kubernetes security, API exposure and misconfiguration are frequent concerns:

- **Excessive Permissions:** Service accounts or users may have more access than needed (e.g., cluster-admin).
- **Hidden Endpoints:** Internal or legacy endpoints may remain undocumented or unprotected.
- **Attack Surface Mapping:** Red Teams attempt to enumerate reachable endpoints for lateral movement or privilege escalation.
- **Defense in Depth:** Blue Teams need visibility into what is exposed and ensure security best practices are followed.

KubeFuzzer supports these security objectives by automating API endpoint discovery—helping teams find what is truly accessible given their credentials and permissions.

## Typical Use Cases

1. **Red Team Engagements:**  
   Enumerate accessible API endpoints using low-privileged, compromised, or default service accounts. Discover resources to further lateral movement (e.g. escalating privileges or extracting secrets).

2. **Blue Team Audits:**  
   Map exposed API endpoints within the cluster. Validate RBAC policies by simulating potential adversary enumeration tactics.

3. **Security Research:**  
   Identify undocumented or legacy endpoints that could indicate outdated versions or misconfigurations.

4. **Cloud Security Automation:**  
   Integrate with CI/CD pipelines to routinely test for new or changed API exposure after infrastructure updates.

5. **Kubernetes Hardening:**  
   Systematically inventory API surface area, assist with tightening RBAC, and develop tailored mitigations.

## Getting Started

### Prerequisites

- Python 3.x
- Access to Kubernetes cluster (direct, kubeconfig, or in-cluster)
- [kubectl](https://kubernetes.io/docs/reference/kubectl/) (optional, for manual verification)

### Install

Clone this repository:

```bash
git clone https://github.com/nosperantos/K8S_LOTL_Red-Team_Arsenal.git
cd K8S_LOTL_Red-Team_Arsenal/KubeFuzzer
```

Install dependencies:

```bash
pip install -r requirements.txt
```

### Usage

1. **Dictionary Configuration**  
   Prepare or customize your endpoint dictionary file (e.g. `endpoints.txt`), containing common and custom API resource paths.

2. **Run KubeFuzzer**  
   Use your kubeconfig file or service account token to authenticate to the Kubernetes API.

   ```bash
   python kube_fuzzer.py --dictionary endpoints.txt --config ~/.kube/config
   ```

   Adjust parameters as needed (see script flags for details).

### Output

KubeFuzzer generates a list of reachable, forbidden, and unrecognized endpoints along with HTTP status codes and response analysis for each attempted API path.

## Contributing

PRs and issue reports are welcome! Please see [CONTRIBUTING.md](../CONTRIBUTING.md).

## References

- [Kubernetes API Reference](https://kubernetes.io/docs/reference/kubernetes-api/)
- [Kubernetes RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)
- [Attack Surface Reduction in Kubernetes](https://kubernetes.io/blog/2020/09/03/kubernetes-attack-surface/)
- [OWASP Kubernetes Top 10](https://owasp.org/www-project-kubernetes-top-10/)

## License

MIT

---

*For more tools for Kubernetes offensive and defensive security, visit the main [K8S_LOTL_Red-Team_Arsenal](https://github.com/nosperantos/K8S_LOTL_Red-Team_Arsenal) repository.*
