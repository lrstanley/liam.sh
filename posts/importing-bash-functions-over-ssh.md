Title: Importing bash functions over SSH
Show: true
Time: 20 Aug 15 23:23 EST

Ever wanted to import or carry over custom made functions and aliases over SSH, when logging into other desktop computers, and servers? Well, here's how I do it.

**Note: This can pose a large security concern. I would highly ensure that the location you are logging in from is completely secure before implementing this. It has the full fledged functionality of running anything on the remote server. Regardless of how safe it is.**

## Installing the sshrc script:

Simply start off by taking the below script, and dropping it into `/usr/local/bin/sshrc` (or your equivalent **$PATH** location), and running `chmod +x /usr/local/bin/sshrc`.

```
#!/usr/bin/env bash

function sshrc {
    local SSHHOME=${SSHHOME:=~}
    if [ ! -f $SSHHOME/.sshrc ];then echo "No such file: $SSHHOME/.sshrc";exit 1;fi

    local sshrcfile=.sshrc
    SIZE=$(tar cz -h -C $SSHHOME .sshrc | wc -c)
    if [ $SIZE -gt 65536 ];then echo >&2 $'.sshrc must be less than 64kb\ncurrent size: '$SIZE' bytes';exit 1;fi
    ssh -t "$@" "
        command -v xxd > /dev/null 2>&1 || { echo >&2 \"sshrc requires xxd to be installed on the server, but it's not. Aborting.\"; exit 1; }
        export SSHHOME=/tmp;echo $'"$(cat "$0" | xxd -p)"' | xxd -p -r > \$SSHHOME/sshrc;chmod +x \$SSHHOME/sshrc

        echo $'"$( cat << 'EOF' | xxd -p
if [ -e /etc/etc/profile ];then source /etc/profile;fi
if [ -e /etc/bashrc ];then source /etc/bashrc;fi
if [ -e /etc/bash.bashrc ];then source /etc/bash.bashrc;fi
if [ -e ~/.bashrc ];then source ~/.bashrc;fi
export PATH=$PATH:$SSHHOME;source $SSHHOME/.sshrc;
cat $SSHHOME/.sshrc | sed -rn "s:function ([a-zA-Z]+) .*:\1:p" | while read fnc;do export -f $fnc;done > /dev/null 2>&1
rm -f $SSHHOME/.post-sshrc $SSHHOME/.sshrc $SSHHOME/sshrc > /dev/null 2>&1 || echo "Unable to cleanup sshrc"
EOF
        )"' | xxd -p -r > \$SSHHOME/.post-sshrc
        echo $'"$(tar cz -h -C $SSHHOME .sshrc | xxd -p)"' | xxd -p -r | tar mxz -C \$SSHHOME
        export SSHHOME=\$SSHHOME;bash --rcfile \$SSHHOME/.post-sshrc
    "
}
if [ $1 ];then
    command -v xxd >/dev/null 2>&1 || { echo >&2 "sshrc requires xxd to be installed locally, but it's not. Aborting.";exit 1; }
    sshrc $@
else ssh;fi
```

## Aliasing ssh

Since the functionality is useful for every server I login to, I alias **ssh** to the above command. As such, I simply open up my `~/.bashrc` file, and add the following (be sure to modify this if you stored the **sshrc** script in another location):

```
function ssh() {
        /usr/local/bin/sshrc -oStrictHostKeyChecking=no $@
        clear
}
```

The above will ensure running `ssh <args>` passes over to `sshrc <args>`. I've added `-oStrictHostKeyChecking=no` to the above function, as this prevents me from having to accept public keys being added to my workstation every time I connect to a server, or if the public key changes. **If you still want this, simply remove the section outlined beforehand.**

## Setting up the scripts to import

The simplest way to get started, is simply creating a file stored at `~/.sshrc`.

```
$ touch ~/.sshrc
```

## Where from here?

Well, there are many things you could add to be extremely helpful in real world situations.

---------------------------------------

#### For example, a simple password generating function:

Function:

```
function genpw { < /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-15};echo; }
```

Usage:

```
$ genpw
cxD3i4_1gtGaEkn
$ genpw
Hy9VopX5nnV8i2r
$ genpw
oqdh7HplyInBLVz
```

---------------------------------------

#### A geoip one liner, that adds GeoIP information to anything passed to it:

Function:

```
function geoip { while IFS= read x;do ip=$(echo $x | egrep -o "[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}");if [ -n "$ip" ];then location=$(curl -s geoip.cf/api/$ip/country);echo "$x" | sed -r "s#$ip#$ip ($location)#g"; else echo "$x";fi;done; }
```

Usage:

```
$ cat /home/domlogs/liamstanley.io | awk '{print $1}' | tail -5 | geoip
54.149.154.72 (United States)
198.190.131.2 (United States)
207.46.13.40 (United States)
201.250.35.137 (Argentina)
187.87.199.209 (Brazil)
```

---------------------------------------

#### Disable SSH within the server you are SSH'd into:

Function:

```
function ssh { echo -e "\n\n\e[0;36mSSH has been disabled, you are daisychaining!\e[m\n\n"; }
```

Usage:

```
$ ssh root@server-2.liam.sh


SSH has been disabled, you are daisychaining!


```

---------------------------------------

#### Pull human-readable file modes

Function:

```
function mode { while read filename;do echo "($(if [ -f "$filename" ];then stat -c '%a' "$filename";else echo "---";fi)) $filename";done; }
```

Usage:

```
$ find /home/liam/public_html/test -maxdepth 1 -type f | mode
(644) /home/liam/public_html/test/config.default.php
(644) /home/liam/public_html/test/index.php
(644) /home/liam/public_html/test/config.php
(644) /home/liam/public_html/test/readme.md
```

There are simply too many possibilities here, that I cannot grasp as of writing this. Some day I may just simply integrate the function into my site to show all the bash functions I use, to share with others. For now, here you go.
