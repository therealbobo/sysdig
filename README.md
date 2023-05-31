<p align="center">
<img src="https://raw.githubusercontent.com/draios/sysdig/dev/img/logo_large.png" width="360">
<p align="center"><b>Dig Deeper</b></p>
<p align="center">
<a href="https://github.com/draios/sysdig/actions/workflows/ci.yaml/badge.svg"><img alt="Release" src="https://github.com/draios/sysdig/actions/workflows/ci.yaml/badge.svg"></a>
<img alt="Release" src="https://img.shields.io/github/v/release/draios/sysdig"></a>
<a href="https://slack.sysdig.com/"><img alt="License" src="https://img.shields.io/badge/chat-on%20slack-green"></a>
<a href="https://twitter.com/sysdig"><img alt="Twitter" src="https://img.shields.io/twitter/url/https/twitter.com/sysdig.svg?style=social&label=Follow%20%40sysdig"></a>
<a href="https://sysdig.com/blog"><img alt="Blog" src="https://img.shields.io/badge/Blog-Sysdig-blue"></a>
</p>
</p>

`sysdig` is an open-source system exploration and troubleshooting tool. It provides deep visibility into the behavior of your Linux containers and supports system-level troubleshooting and analysis.

`csysdig` is the interactive counterpart of `sysdig` and offers a terminal-based user interface (TUI) with ncurses. It provides an interactive and intuitive way to explore and troubleshoot Linux systems, allowing users to gain deep visibility into system-level activities, analyze system behavior, and identify performance issues efficiently.

## Features

**sysdig** is a powerful system visibility tool with native support for containers. It provides deep insights into system behavior and offers extensive troubleshooting and analysis capabilities. Sysdig allows you to capture system calls, monitor file system activity, analyze network traffic, and much more.

- **Universal System Visibility**: sysdig offers a comprehensive view of the system by capturing system calls, file system activity, network connections, and other events at the OS level. Think of Sysdig as a powerful **combination of tools like strace, tcpdump, htop, iftop, and lsof, all integrated into a single cohesive solution**.

- **Native Container Support**: sysdig seamlessly integrates with containerized environments, providing deep visibility into the behavior of containerized applications and correlating system-level activities with specific containers.

- **Trace File Generation**: sysdig can create trace files that store system activity, enabling offline analysis and troubleshooting at a later time. These trace files capture rich system state, providing a complete context for captured activity.

- **Command-Line Interface (CLI) and Text-based User Interface (TUI)**: sysdig provides both powerful and user-friendly command-line interface and text-based user interface, allowing interactive exploration, real-time analysis, and the application of custom scripts and filters.

For a detailed list of features and more technical information, please refer to the [FEATURES.md](./FEATURES.md) file.

![sysdig](/img/sysdig_screenshot.png?raw=true "sysdig")

![csysdig](/img/csysdig_screenshot.png?raw=true "csysdig")

## Getting Started

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

## Documentation / Support

[Visit the wiki](https://github.com/draios/sysdig/wiki) for full documentation on sysdig and its APIs.  

For support using sysdig, please contact [the official mailing list](https://groups.google.com/forum/#!forum/sysdig).  

## How to Contribute

Please refer to the [contributing guide](CONTRIBUTING.md) and the [code of conduct](CODE_OF_CONDUCT.md) for more information on how to contribute.

## Join the Community

* Contact the [official mailing list](https://groups.google.com/forum/#!forum/sysdig) for support and to talk with other users
* Follow us on [Twitter](https://twitter.com/sysdig)
* This is our [blog](https://sysdig.com/blog/). There are many like it, but this one is ours.
* Join our [Public Slack](https://slack.sysdig.com) channel for sysdig announcements and discussions.

## License Terms

The sysdig userspace programs and supporting code are licensed to you under the [Apache 2.0](COPYING.md) open source license.

## Commercial Support

Interested in a fully supported, fully distributed version of sysdig? Check out [Sysdig Monitor](https://sysdig.com/products/monitor/)!

Open source sysdig is proudly supported by [Sysdig Inc](https://sysdig.com/company/).  

Interested in what we're doing? [Sysdig is hiring](https://sysdig.com/jobs/).

## Reporting a vulnerability

Please refer to [SECURITY.md](SECURITY.md).

## Packaging status

<p align="center">
    <a href="https://repology.org/project/sysdig/versions">
        <img src="https://repology.org/badge/vertical-allrepos/sysdig.svg?minversion=0.30.0&exclude_unsupported=1" alt="Packaging Status">
    </a>
</p>

