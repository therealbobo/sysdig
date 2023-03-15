sysdig
======

<p align="center"><img src="https://raw.githubusercontent.com/draios/sysdig/dev/img/logo_large.png" width="360"></p>
<p align="center"><b>Dig Deeper</b></p>

<hr>

`~$ sysdig` is a system-level exploration and troubleshooting tool. It provides deep visibility into the behavior of Linux systems and applications in real-time. Sysdig captures system calls and other OS events, and then uses a filtering language to extract useful information from the captured data. This allows users to diagnose and troubleshoot complex problems, such as system hangs, application slowdowns, and network issues.

`~$ csysdig` is a command-line tool built on top of Sysdig that provides an interactive interface for exploring system behavior. It allows users to drill down into captured data, view system calls and events in real-time, and apply filters to focus on specific processes or activities. `~$ csysdig` also includes a set of pre-built views for common use cases, such as network analysis and container monitoring.

Together, Sysdig and csysdig offer a powerful set of tools for understanding and troubleshooting complex Linux systems. They are widely used in DevOps and system administration contexts to improve system performance and reliability.


Getting Started
---

Run Sysdig in a container:

```
sudo docker run --rm -i -t --privileged --net=host \
    -v /var/run/docker.sock:/host/var/run/docker.sock \
    -v /dev:/host/dev \
    -v /proc:/host/proc:ro \
    -v /boot:/host/boot:ro \
    -v /src:/src \
    -v /lib/modules:/host/lib/modules:ro \
    -v /usr:/host/usr:ro \
    -v /etc:/host/etc:ro \
    docker.io/sysdig/sysdig
```

And then run the `sysdig` or `csysdig` tool from the container shell!

Or install the [latest release](https://github.com/draios/sysdig/releases/latest) with a `deb` or `rpm` package for your distribution.

What does sysdig do and why should I use it?
---
**Sysdig is a simple tool for deep system visibility, with native support for containers.**

The best way to understand sysdig is to [try it](https://github.com/draios/sysdig/wiki/How-to-Install-Sysdig-for-Linux) - its super easy! Or here's a quick video introduction to csysdig, the simple, intuitive, and fully customizable curses-based UI for sysdig: https://www.youtube.com/watch?v=UJ4wVrbP-Q8

Far too often, system-level monitoring and troubleshooting still involves logging into a machine with SSH and using a plethora of dated tools with very inconsistent interfaces. And many of these classic Linux tools breakdown completely in containerized environments. Sysdig unites your Linux toolkit into a single, consistent, easy-to-use interface. And sysdig's unique architecture allows deep inspection into containers, right out of the box, without having to instrument the containers themselves in any way.

Sysdig instruments your physical and virtual machines at the OS level by installing into the Linux kernel and capturing system calls and other OS events. Sysdig also makes it possible to create trace files for system activity, similarly to what you can do for networks with tools like tcpdump and Wireshark. This way, problems can be analyzed at a later time, without losing important information. Rich system state is stored in the trace files, so that the captured activity can be put into full context.

Think about sysdig as strace + tcpdump + htop + iftop + lsof + ...awesome sauce.

Features
---
**sysdig**:

- Captures system calls and other OS events in real-time.
- Provides deep visibility into the behavior of Linux systems and applications.
- Uses a filtering language to extract useful information from captured data.
- Supports a wide range of filter options, including time-based filters and user-defined filters.
- Includes pre-built chisels (scripts that extract specific data from captured events) for common use cases, such as process monitoring and network analysis.
- Supports custom chisels for user-specific use cases.
- Can capture and analyze data from containers and Kubernetes clusters.
- Provides a command-line interface as well as a GUI for visualization and exploration of captured data.
- Integrates with popular monitoring and logging tools like Prometheus, Grafana, and Elasticsearch.

**csysdig**:

- Provides an interactive terminal-based interface for exploring system behavior.
- Displays system calls and events in real-time, with color coding for easy identification.
- Includes pre-built views for common use cases, such as process monitoring and network analysis.
- Supports custom views and filters for user-specific use cases.
- Allows users to navigate captured data using keyboard shortcuts.
- Provides detailed documentation on all available views and commands.
- Supports output in JSON format for integration with other tools and automation.

Extending sysdig
---
 In Sysdig, a chisel is a script written in Lua that extracts specific data from captured system events. Chisels can be used to automate the process of analyzing system behavior and identifying issues.

There are a number of pre-built chisels included with Sysdig that cover common use cases, such as process monitoring, network analysis, and container monitoring. For example, the "topfiles" chisel can be used to identify the files that are being accessed most frequently by processes on the system, while the "httplog" chisel can be used to log HTTP requests made by applications.

Chisels can also be customized or created from scratch to suit specific use cases. This makes Sysdig and its chisels a powerful tool for system administrators and developers who need to diagnose and troubleshoot complex system issues.

To use a chisel in Sysdig, you simply specify it on the command line when running the tool, along with any relevant parameters. For example, to use the "topfiles\_time" chisel, you would run:

```
sudo sysdig -c topfiles_time
```

This would launch Sysdig with the topfiles\_time chisel enabled, and it would begin capturing data on the most frequently accessed files on the system.

Overall, chisels are a key feature of Sysdig that enable users to quickly and easily extract valuable insights from captured system data.

Examples
---
1. Monitoring system performance: Sysdig can be used to monitor system performance in real-time, providing deep visibility into CPU, memory, and disk usage. For example, to monitor CPU usage, you can run the following command:

```
# sysdig -c topprocs_cpu
```
This will show you the processes that are consuming the most CPU cycles, along with other useful information like process ID and command.

2. Debugging application issues: Sysdig can be used to debug issues with specific applications, providing detailed information on system calls, network activity, and more. For example, to monitor HTTP requests made by a particular application, you can run the following command:

```
# sudo sysdig -A -pc -c echo_fds fd.port=80 and proc.name=myapp
```

This will capture all HTTP requests made by the "myapp" process and display them in real-time.

3. Analyzing container behavior: Sysdig can be used to analyze the behavior of containers and Kubernetes clusters, providing visibility into container metrics, network activity, and more. For example, to monitor network activity within a container, you can run the following command:

```
# sysdig -pc -c topconns 'container.id != host'
```
This will display a live view of network activity within all running containers on the system.

4. Investigating security issues: Sysdig can be used to investigate security issues, providing deep visibility into system events and process activity. For example, to monitor file access events on the system, you can run the following command:

```
# sysdig -p '%proc.cmdline %evt.arg.name' "evt.type=openat"
```
This will display a live stream of file open events on the system, showing the path of the file being accessed and the process that is accessing it.


Documentation / Support
---
[Visit the wiki](https://github.com/draios/sysdig/wiki) for full documentation on sysdig and its APIs.  

For support using sysdig, please contact [the official mailing list](https://groups.google.com/forum/#!forum/sysdig).  

Join the Community
---
* Contact the [official mailing list](https://groups.google.com/forum/#!forum/sysdig) for support and to talk with other users
* Follow us on [Twitter](https://twitter.com/sysdig)
* This is our [blog](https://sysdig.com/blog/). There are many like it, but this one is ours.
* Join our [Public Slack](https://slack.sysdig.com) channel for sysdig announcements and discussions.

Our [code of conduct](CODE_OF_CONDUCT.md) applies across all our projects and community places.

License Terms
---
The sysdig userspace programs and supporting code are licensed to you under the [Apache 2.0](./COPYING) open source license.

Developer Certification of Origin (DCO)
---
The Apache 2.0 license tells you what rights you have that are provided by the copyright holder. It is important that the contributor fully understands what rights they are licensing and agrees to them. Sometimes the copyright holder isn't the contributor, such as when the contributor is doing work on behalf of a company.

To make a good faith effort to ensure these criteria are met, we require the Developer Certificate of Origin (DCO) process to be followed.

The DCO is an attestation attached to every contribution made by every developer. In the commit message of the contribution, the developer simply adds a Signed-off-by statement and thereby agrees to the DCO, which you can find at http://developercertificate.org.

### DCO Sign-Off Methods
The DCO requires a sign-off message in the following format appear on each commit in the pull request:

```
Signed-off-by: John Doe <john.doe@sysdig.com>
```

You have to use your real name (sorry, no pseudonyms or anonymous contributions).

The DCO text can either be manually added to your commit body, or you can add either `-s` or `--signoff` to your usual `git commit` commands. If you are using the GitHub UI to make a change, you can add the sign-off message directly to the commit message when creating the pull request. If you forget to add the sign-off you can also amend a previous commit with the sign-off by running `git commit --amend -s`. If you've pushed your changes to GitHub already you'll need to force push your branch after this with `git push -f`.

Commercial Support
---
Interested in a fully supported, fully distributed version of sysdig? Check out [Sysdig Monitor](https://sysdig.com/products/monitor/)!

Open source sysdig is proudly supported by [Sysdig Inc](https://sysdig.com/company/).  

Interested in what we're doing? [Sysdig is hiring](https://sysdig.com/jobs/).

Reporting a vulnerability
---
Please refer to [SECURITY.md](SECURITY.md).
