# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/bionic64"
  config.vm.provider "virtualbox" do |vb|
    vb.memory = 1024
    vb.cpus = 1
    vb.gui = false
  end

  # CENTRAL host
  config.vm.define "central" do |central|
    central.vm.hostname = "central"
    central.vm.network "forwarded_port", guest: 15672, host: 15672, host_ip: "0.0.0.0" #rabbitmq
    central.vm.network "forwarded_port", guest: 9090, host: 9090, host_ip: "0.0.0.0" #prometheus
    central.vm.network "forwarded_port", guest: 3000, host: 3000, host_ip: "0.0.0.0" #prometheus
    central.vm.network "private_network", ip: "192.168.60.10" 
  end 

  # Workers
  COUNT = 1
  (1..COUNT).each do |worker_id|
    config.vm.define "worker-#{worker_id}" do |worker|
    worker.vm.hostname = "worker#{worker_id}"
    worker.vm.network "private_network", ip: "192.168.60.#{10+worker_id}"
      if worker_id == COUNT
        worker.vm.provision :ansible do |ansible|
            ansible.limit = "all"
            ansible.playbook = "./ansible/playbook.yml"
            ansible.groups = {
              "central_group" => ["central"],
              "worker_group" => ["worker-[1:#{COUNT}]"],
              "all_groups:children" => ["central_group", "worker_group"],
              "central_group:vars" => {
                "name" => "central",
                "rabbitmq_default_vhost" => "/",
                "rabbitmq_default_user" => "ansible",
                "rabbitmq_default_pass" => "ansible",
                "rabbitmq_default_user_tags" => "administrator",
                "rabbitmq_configure_priv" => ".*",
                "rabbitmq_read_priv" => ".*",
                "rabbitmq_write_priv" => ".*",
                "prometheus_node_exporter_group" => "all",
                "prometheus_dir_configuration" => "/etc/prometheus",
                "grafana_admin_password" => "vagrant",
                "grafana_url" => "http://192.168.60.10:3000",
                "grafana_dashboards_dir" => "dashboards",
                "grafana_data_dir" => "/var/lib/grafana",
                "grafana_dashboards" => ["node_exporter", "rabbitmq"],
                "datasource" => "Prometheus",
                "prometheus_url" => "http://localhost:9090",
                "count_messages" => "999999999999999"
              },
              "worker_group:vars" => {"name" => "worker",
                  "ansible_ssh_private_key_file" => "./.vagrant/machines/worker/virtualbox/private_key"          
              },
              "all_groups:vars" => {
                "ubuntu_distribution" => "bionic",
                "home_directory" => "/home/vagrant",
                "arch" => "amd64",
                "node_exporter_version" => "1.3.1", 
                "node_exporter_bin_path" => "/usr/local/bin",
                "node_exporter_download_url" => "https://github.com/prometheus/node_exporter/releases/download/v1.3.1/node_exporter-1.3.1.linux-amd64.tar.gz",
                "node_exporter_options" => "",
                "_node_exporter_system_group" => "node-exporter",
                "_node_exporter_system_user" => "node-exporter"
              }
            }
        end
      end
    end
  end
end
