terraform {
  required_version = ">= 0.12.0"
}

provider "aws" {
  version = "~> 2.49"
}

provider "kubernetes" {
  version = "~> 1.10"
}

module "k8s" {
  source = "../k8s"

  providers = {
    kubernetes = kubernetes
  }

  namespace = var.namespace
  rack      = var.rack
  release   = var.release
}

resource "kubernetes_service" "resolver-external" {
  metadata {
    namespace = var.namespace
    name      = "resolver-external"

    annotations = {
      "service.beta.kubernetes.io/aws-load-balancer-type" : "nlb"
    }
  }

  spec {
    type = "NodePort"

    port {
      name        = "dns"
      protocol    = "UDP"
      target_port = 5453
    }

    selector = module.k8s.selector
  }

  lifecycle {
    ignore_changes = [metadata[0].annotations]
  }
}
