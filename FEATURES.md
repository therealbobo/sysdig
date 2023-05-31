# Features

`sysdig` is a universal system visibility tool with native support for containers:

![sysdig](/img/sysdig_screenshot.png?raw=true "sysdig")

- **System Call Capturing:** sysdig captures system calls at the OS level, providing deep visibility into the behavior of the system. It captures system events such as *process execution*, *file system activity*, *network connections*, *interprocess communication*, and more.

- **Container Support:** sysdig natively supports containers, allowing users to gain insights into the behavior of *containerized applications*. It automatically correlates system-level activities with specific containers, providing a comprehensive view of containerized environments.

- **Scap File Generation:** sysdig can create *trace files* that store system activity for analysis at a later time. These trace files capture rich system state, enabling comprehensive contextual analysis of system behavior and troubleshooting.

- **Command-line Interface (CLI):** sysdig offers a powerful and user-friendly command-line interface, allowing users to interactively explore system events, apply filters and custom scripts, and perform real-time analysis. The CLI provides flexible control and access to system-level visibility and debugging capabilities.

`csysdig` is a simple, intuitive, and fully customizable curses UI for sysdig:

![csysdig](/img/csysdig_screenshot.png?raw=true "csysdig")

- **Curses-based UI:** csysdig provides a curses-based UI (text-based graphical user interface) for sysdig, making it accessible in terminal environments. It leverages the ncurses library to create an interactive and visually appealing interface.

- **Real-time Exploration:** csysdig allows real-time exploration of system activity by presenting captured events in a dynamic and continuously updated interface. Users can navigate through the captured data and observe system behavior as it happens.

- **Customization:** csysdig offers extensive customization options, allowing users to configure the display and focus on specific system events or metrics. It supports the selection of different views, filters, and output formats, enabling users to tailor the interface to their specific needs.

- **Interactive Navigation:** csysdig provides interactive navigation capabilities, allowing users to drill down into system events and explore related details. It enables efficient troubleshooting and debugging by providing a high-level overview as well as the ability to zoom in on specific areas of interest.

- **Efficient Analysis:** csysdig enhances the analysis process by presenting captured system events in a visual and structured manner. It helps identify patterns, anomalies, and potential issues by visualizing data in real-time, making it easier to understand system behavior and make informed decisions.

sysdig and csysdig provide a powerful feature called `chisels` that allows users to extend its functionality and perform specialized analysis and data processing. Chisels are scripts written in Lua that can be executed within sysdig or csysdig to extract, manipulate, and visualize system data. Here are some key points about chisels:

- **Scripting Capabilities:** Chisels leverage the scripting capabilities of Lua to perform advanced analysis and data processing tasks. Users can write custom chisels to extract specific information from system events, apply complex filters, aggregate data, and create custom visualizations.

- **Rich Chisel Library:** sysdig has a rich ecosystem of pre-built chisels that cover a wide range of use cases. They provide ready-to-use functionality for tasks such as container monitoring, security analysis, performance profiling, network forensics, and more.

- **Flexible Integration:** Chisels seamlessly integrate with the sysdig and csysdig interfaces, allowing users to execute chisels directly within the tools. Users can apply chisels to live system data or trace files, enabling real-time analysis or offline investigation.

- **Customization and Extensibility:** Chisels offer a high degree of customization and extensibility, allowing users to adapt sysdig to their specific needs. Users can modify existing chisels or create new ones to extract the desired information, generate custom reports, create visualizations, or integrate with other tools and workflows.

By leveraging chisels, users can unlock the full potential of Sysdig and Csysdig by tailoring the tools to their specific use cases and analysis requirements. The availability of a wide range of chisels and the ability to create custom ones make Sysdig a highly flexible and powerful system visibility and analysis platform.

These technical features empower users to effectively monitor, analyze, and troubleshoot system activities, providing deep visibility into both the host system and containerized environments.

To get a quick overview of Sysdig, you can try it out for yourself, as it is designed to be user-friendly and easy to use. Additionally, you can watch a quick introductory video on csysdig, the customizable curses-based UI for Sysdig, which demonstrates its simplicity and intuitiveness. Check out the video [here](https://www.youtube.com/watch?v=UJ4wVrbP-Q8).

