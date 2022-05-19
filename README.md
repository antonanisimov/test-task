# Test task

## Requirements for running:
- Installed Vagrant
- Installed VirtualBox



Exchange: my-amq
Queue: my-queue

---
## Runtime project
Create VMs 
> **vagrant up**<br>

!!! Need choose your Internet interface for bridge 
[public_network] 
example: 
1) en0: Wi-Fi(?)
2) en5: USB Ethernet(?)

Run ansible-playbook:

> **ansible-playbook -i ansible/inventories/hosts.yml ansible/playbook.yml**


metrics rabbitMQ:
localhost:15692/metrics 

---
## Access monitoring system
prometheus:
http://192.168.60.10:9090/

grafana:
http://192.168.60.10:3000/
>login: admin<br>
>password: vagrant

---
## Access UI RabbitMQ
rabbitMQ:
http://192.168.60.10:15672/
> login: ansible<br>
> password: ansible

---
## For increase scale worker use:
> **vagrant up worker-2**<br>
> **vagrant up worker-3**<br>
> **vagrant up worker-4**<br>

## How this working
<br>

## How to run stress test
run -f

## Producer (application Go)
> ./producer -count *'count'*
## Consumer (application Go)
![This is an image](scheme.png)
