# **S**plunk **C**loud **KN**owledge **OB**ject **B**ackup

'`scknobb`' is a tool to quickly snapshot your Splunk cloud search history, dashboards and reports to your local filesystem. This is useful when you don't have access to the backend filesystems or want to develop a workflow with version control.

Just run '`scknobb backup all`' and enjoy the benefits of your local copies, timestamped from their last updated values - ie. sort newest to oldest.

- Ready to run compiled binaries for Windows and Linux:
  - The '`release/`' folder contains builds for the latest release; download and run for your platform.
- To run from source:
  - '`git clone https://github.com/whereiskurt/splunkcloudknobb.git`'
  - '`go run cmd/scknobb.go help`'
- Build your own binary with Mage or Docker:

  - To build the whole project and binaries (see '`magefile.go`' for more details):
    -'`mage -v release`'
  - For Docker build/run:
    - '`docker build --tag scknobb.v0.1 .`'
    - '`docker run --tty --interactive --rm scknobb.v0.1`'

This is very proof-of-concept, but it's built using the best Go libraries money can buy.

The very first time you run the '`scknobb backup all`' the URL, username/password details that are prompted for are stored in the users '`$HOMEDIR`' in the '`.scknobb.v1.yaml`' configuration file encrypted using the crypto key that was prompted. It is not recommended to store that crypto key in the same '`.scknobb.v1.yaml`' file (unless you're aware of the implications) but instead pass the '`--key=XYZ`' parameter.

**WARNING:** Running '`scknobb backup all`' will also backup your search history (aka. '`scknobb backup history`'). If you've done any sensitive searches during (e.g. during an incident investigation / HR review) those searches will be downloaded locallly to the folder you are in.

**NOTE: This tool is written and maintained by @whereiskurt and is not officially supported or endorsed by Splunk in any way.**

<img src="https://github.com/whereiskurt/splunkcloudknobb/blob/master/doc/images/kphgopher.png" width="250">

## Release Notes (v0.0.1):

This release has '`backup`' command with support for '`reports`', '`dashboards`' and '`history`' (aka '`all`') subcommands.

Use '`scknobb help backup`' to see more details or '`scknobb backup all`' to back to take a time snapshot.

**Dedicated to: PTaddy, DeeMaCK, KMaze, sc0ttys1n, trizz13.**

### Backup All - Search History, Reports, and Dashboards to Local Filesystem

```shell

## Run the scknobb binary to backup all
$ scknobb backup all --key=secret1234

#              _                _     _
#     ___  ___| | ___ __   ___ | |__ | |__
#    / __|/ __| |/ / '_ \ / _ \| '_ \| '_ \  2020-04-18
#    \__ \ (__|   <| | | | (_) | |_) | |_) | v0.0.1-development
#    |___/\___|_|\_\_| |_|\___/|_.__/|_.__/  0x0123abcd
#
#    > Splunk Cloud KNowledge OBject Backup
#
#    An CLI interface into the Splunk Cloud using Go!
#
#    Find more information at:
#        https://github.com/whereiskurt/scknobb/
#
#    Starting backup for dashboards...
#
#    √ Successful login with 'some_user@domain.tld' to instance:
#        https://instanceX.splunkcloud.com/en-US/
#
#    √ Fetched dashboard listing for ALL dashboards ...
#    √ Preparing to backup '314' dashboards
#    √ Writing files to folder: '/home/suser/splunkcloudknobb/20200421T010602.scknobb/'
#    ..........................................................................
#    ..........................................................................
#    ..........................................................................
#    .......... done! :-)
#
#    √ Success! Wrote '3141kb' across '314' dashboard backup files.
#
#    Congratulations you now have a local backup! :-)
...

```

## Why?

Once upon a time someone unintentionally deleted a dashboard from their Splunk Cloud instance and a bad feeling was felt. The feeling worsened without a backup. **"Don't be knobb, have a backup!"** - and so the Splunk Cloud Knowledge Object Backup - was created born.

Primary use case: 1) You want a local copy of all the dashboards/reports/search history :-)

## Some details about the code:

I've [curated a YouTube playlist](https://www.youtube.com/playlist?list=PLa1qVAzg1FHthbIaRRbLyA4sNE4PmLmn6) of videos which help explain how I ended up with this structure and 'why things are the way they are.' I've leveraged 'best practices' I've seen and that have been explicted called out by others. Of course **THERE ARE SOME WRINKLES** and few **PURELY DEMONSTRATION** portions of code. I hope to be able to keep improving on this.

- [x] Fundamental Go features like modules, channels, generate, templates, tags, "ldflags"
- [x] Build file using [`mage`](https://github.com/magefile/mage) for build/release etc.
- [x] Structured using folders `cmd`, `internal`, `pkg` with `magefile.go` + `Dockerfile`
      [x] Flags and configuration with [`cobra`](https://github.com/spf13/cobra) and [`viper`](https://github.com/spf13/viper) (without func inits!!!)
  - Cleanly separated CLI/configuration invocation from client library calls - by calling `viper.Unmarshal` to transfer our `pkg.Config`
  - **NOTE**: A lot of sample Cobra/Viper code rely on `func init()` making it more difficult to reuse.
- [x] Embeded templates for 'one-big-binary' using [`vfsgen`](https://github.com/shurcooL/vfsgen)
- [x] Logging with [`logrus`](https://github.com/sirupsen/logrus) library and written to `log/`
- [x] [Retry](https://github.com/matryer/try) using @matryer's idiomatic `try.Do(..)`
- [x] An example Dockerfile and build recipe `(docs/recipe/)` for a docker workflow
  - Use `docker build --tag scknobb.v0.1 .` to create a full golang image
  - Use `docker run -it --rm scknobb.v0.1` to work from with the container
