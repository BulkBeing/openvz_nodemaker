# openvz_nodemaker
Configures a CentOS 6 installation with all settings and packages to use as an OpenVZ node

Sets up OpenVZ repo.

Installs vzkernel, vzctl, vzquota.

Modifies /etc/sysctl.conf to contain all necessary settings including ip forwarding.

Disables selinux.

Creating binary:
GOOS=linux go build .
