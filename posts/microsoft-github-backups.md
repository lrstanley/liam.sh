Title: Let's talk about GitHub backups!
Show: true
Time: 06 Jun 18 19:19 EDT

### Let's talk about Github backups

Following [Microsoft's announcement](https://blog.github.com/2018-06-04-github-microsoft/) to
[acquire GitHub](https://news.microsoft.com/2018/06/04/microsoft-to-acquire-github-for-7-5-billion/),
this shed some light on how I've been too generous to centralized services offering free services.
After all, you _do get what you pay for_.

Although I have had tendencies to jump on the Microsoft hate bandwagon a few times, due to the
(in my opinion) terrible customer focus -- I don't totally dislike the idea of GitHub being
owned by Microsoft. I will save that spiel for another though. Today, I wanted to focus on getting
a bit more control of the data I house on GitHub. After all, GitHub is almost like a resume for
developers, and if I were to lose that data, I don't know what I'd do.

### python-github-backup

I began browsing for utilities/services that could backup my account info, repositories, gists,
and stars, etc. I came across [python-github-backup](https://github.com/josegonzalez/python-github-backup).
Today, I will briefly go over how I use it.

I wanted to setup weekly backups, to my NAS (with over 40TB worth of free space).

### Installation

Let's begin with the installation. It's silly simple (using root):

```bash
$ sudo pip install github-backup
```

Or (using root):

```bash
$ sudo pip install git+https://github.com/josegonzalez/python-github-backup.git#egg=github-backup
```

Or (as a regular user):

```bash
$ pip install --user github-backup
```

### Usage

Take a look at `github-backup --help`, or the README in the above linked repo. It's pretty extensively
amazing.

### On a cron you say?

I've setup a simple wrapper script, which runs backups in a temporary directory, then archives
them and places them in a desired location. Take a look at it [here](https://gist.github.com/lrstanley/34ee3b024223af142258ce5d1a7c16c9).

I've dropped it in `~/bin/gbackup`, but you can place it anywhere. I also have two cronjobs, which
run every Sunday, at ~3:00AM and ~3:15AM respectively:

```cron
0 3 * * 0 /root/bin/gbackup lrstanley /net/media/archives/github-backups/github-backup-lrstanley.$(date +%Y.%m.%d-%H.%M.%S).tar.gz > /dev/null
15 3 * * 0 find /net/media/archives/github-backups/ -type f -mtime +365 -delete > /dev/null
```

Notice that second command -- it finds any archives older than 365 days old, and removes them. That's
around 52 backups across an entire year, at about ~45MB `.tar.gz`'d per backup. Not bad!

Lastly, the script needs a variable set. `GH_BACKUP_TOKEN`, which you can generate using [this GitHub page](https://github.com/settings/tokens/new).

You can invoke the github-backup script with it using `GH_BACKUP_TOKEN=<token> /root/bin/gbackup` or
place it in e.g. `~/.bashrc`, using `export GH_BACKUP_TOKEN=<token>` where it will be accessible
when you run the script.

Welp, now I have GitHub backups. It's not amazingly robust, but since GitHub doesn't have a full-account
backup system in place, this is sufficient enough for me at this time. Happy commiting.
