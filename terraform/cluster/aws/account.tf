resource "kubernetes_service_account" "system" {
  provider = kubernetes.direct

  metadata {
    namespace = "kube-system"
    name      = "${var.name}-system"
  }

  lifecycle {
    ignore_changes = all
  }
}

resource "kubernetes_cluster_role" "system" {
  provider = kubernetes.direct

  metadata {
    name = "${var.name}-system"
  }

  rule {
    api_groups = ["*"]
    resources  = ["*"]
    verbs      = ["*"]
  }

  lifecycle {
    ignore_changes = all
  }
}

resource "kubernetes_cluster_role_binding" "system" {
  provider = kubernetes.direct

  metadata {
    name = "${var.name}-system"
  }

  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "${var.name}-system"
  }

  subject {
    kind      = "ServiceAccount"
    name      = "${var.name}-system"
    namespace = "kube-system"
  }

  lifecycle {
    ignore_changes = all
  }
}

data "kubernetes_secret" "system" {
  provider = kubernetes.direct

  metadata {
    namespace = "kube-system"
    name      = kubernetes_service_account.system.default_secret_name
  }
}
