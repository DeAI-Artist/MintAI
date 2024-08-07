terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

resource "digitalocean_tag" "cluster" {
  name = "${var.name}"
}

resource "digitalocean_ssh_key" "cluster" {
  name       = "${var.name}"
  public_key = "${file(var.ssh_key)}"
}

resource "digitalocean_droplet" "cluster" {
  name = "${var.name}-node${count.index}"
  image = "ubuntu-20-04-x64"  # Specify the Ubuntu image here
  size = "${var.instance_size}"
  region = "${element(var.regions, count.index)}"
  ssh_keys = ["${digitalocean_ssh_key.cluster.id}"]
  count = "${var.servers}"
  tags = ["${digitalocean_tag.cluster.id}"]


  lifecycle {
    prevent_destroy = false
  }


  connection {
    timeout = "30s"
  }

}

