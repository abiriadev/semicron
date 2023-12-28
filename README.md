<h1 align="center">semicron</h1>
<p align="center">Minimal crontab for fast and expressive shell scripting</p>

## Why not traditional Cron?

Creating and managing a crontab can often be a cumbersome task. It typically involves several challenges:

-   **Sudo permission**. In many cases, crontab modifications require elevated permissions, which can be inconvenient or even restrictied in certain environments.
-   **Absolute paths**. Crontab demands the use of absolute paths, which makes it more cumbersome when what you want to do is just executing a simple shell script that you just wrote.
-   **Lack of second-precision**. The smallest possible unit of traditional cron is minutes, which can be too coarse for more granular or sophisticated automation needs.

`semicron` is minimalistc and modern alternative to crontab, perfect for quick and dirty shell scripting.

-   **Relative path support**. Unlike crontab, `semicron` allows the use of relative paths. You can run everything in CWD right now!
-   **Easy termination**. You can simply terminate `semicron` with a normal `Ctrl + C` command, which is far intuitive, and you no longer have to worry about cron running in the background that you haven't cleared.
-   Second-precision: `semicron` supports optional second-precision, enabling more precise and repeatitive work. You don't need to wait almost 10 minutes to confirm cron is working any longer!
-   **Cross platform**. Yeah. `semicron` runs perfectly on Windows!

## Usage

```sh
Usage:
  semicron [OPTIONS] cron cmd arguments...

Application Options:
  -s, --shell=     Shell to execute the command

Help Options:
  -h, --help       Show this help message
```

## Examples

### Ping

Ping every 15 seconds.

```sh
$ semicron '*/15 * * * * *' ping google.com -- -c 1
```

### Simple timer

```sh
$ semicron '* * * * * *' -s bash figlet '$(date)' -- -w 140

 _____     _   ____              ____   ___     ___ ____    _  __    _____ _  _     _  ______ _____   ____   ___ ____  _____
|  ___| __(_) |  _ \  ___  ___  |___ \ / _ \   / _ \___ \ _/ |/ /_ _|___ /| || |   | |/ / ___|_   _| |___ \ / _ \___ \|___ /
| |_ | '__| | | | | |/ _ \/ __|   __) | (_) | | | | |__) (_) | '_ (_) |_ \| || |_  | ' /\___ \ | |     __) | | | |__) | |_ \
|  _|| |  | | | |_| |  __/ (__   / __/ \__, | | |_| / __/ _| | (_) | ___) |__   _| | . \ ___) || |    / __/| |_| / __/ ___) |
|_|  |_|  |_| |____/ \___|\___| |_____|  /_/   \___/_____(_)_|\___(_)____/   |_|   |_|\_\____/ |_|   |_____|\___/_____|____/

 _____     _   ____              ____   ___     ___ ____    _  __    _________    _  ______ _____   ____   ___ ____  _____
|  ___| __(_) |  _ \  ___  ___  |___ \ / _ \   / _ \___ \ _/ |/ /_ _|___ / ___|  | |/ / ___|_   _| |___ \ / _ \___ \|___ /
| |_ | '__| | | | | |/ _ \/ __|   __) | (_) | | | | |__) (_) | '_ (_) |_ \___ \  | ' /\___ \ | |     __) | | | |__) | |_ \
|  _|| |  | | | |_| |  __/ (__   / __/ \__, | | |_| / __/ _| | (_) | ___) |__) | | . \ ___) || |    / __/| |_| / __/ ___) |
|_|  |_|  |_| |____/ \___|\___| |_____|  /_/   \___/_____(_)_|\___(_)____/____/  |_|\_\____/ |_|   |_____|\___/_____|____/

 _____     _   ____              ____   ___     ___ ____    _  __    _____  __     _  ______ _____   ____   ___ ____  _____
|  ___| __(_) |  _ \  ___  ___  |___ \ / _ \   / _ \___ \ _/ |/ /_ _|___ / / /_   | |/ / ___|_   _| |___ \ / _ \___ \|___ /
| |_ | '__| | | | | |/ _ \/ __|   __) | (_) | | | | |__) (_) | '_ (_) |_ \| '_ \  | ' /\___ \ | |     __) | | | |__) | |_ \
|  _|| |  | | | |_| |  __/ (__   / __/ \__, | | |_| / __/ _| | (_) | ___) | (_) | | . \ ___) || |    / __/| |_| / __/ ___) |
|_|  |_|  |_| |____/ \___|\___| |_____|  /_/   \___/_____(_)_|\___(_)____/ \___/  |_|\_\____/ |_|   |_____|\___/_____|____/
```
