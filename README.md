# virt-ssh-config

Get the IP addresses of running libvirt virtual machines and create ssh-config file from them

## Example

    $ virt-ssh-config > my-ssh-config
    $ cat my-ssh-config
    Host my-ubuntu-vm-3
      HostName 10.200.0.47

    Host my-ubuntu-vm-1
      HostName 10.200.0.233

    Host my-ubuntu-vm-2
      HostName 10.200.0.56

    $ ssh -F my-ssh-config ubuntu@my-ubuntu-vm-1
