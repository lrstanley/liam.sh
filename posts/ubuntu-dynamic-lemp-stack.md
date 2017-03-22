Title: Ubuntu Dynamic LEMP Stack
Show: true
Time: 07 Aug 15 23:20 EST

LEMP is a variation of the LAMP stack used for developing and deploying web applications. Traditionally, LAMP consists of **Linux**, **Apache**, **MySQL**, and **PHP**. Due to its modular nature, the components can easily be swapped out. With LEMP, Apache is replaced with the lightweight yet powerful [Nginx](http://nginx.org/). If you are wondering why it's called LEMP, not _LNMP_, this is due to the [pronunciation](http://wiki.nginx.org/Pronunciation) of Nginx (_en-juhn-ecks_).

With a LEMP stack, you can expand, and scale in a very powerful fashion. I have been using LEMP stacks for quite some time now, however when setting up my servers, the configuration is made to be managed easily by one user. If this is the case for you, then this post will be an interesting read.

With the following configuration given in this post, you will be able to point your domains to the server in mention, make a directory named `/var/www/<your-domain>`, and content will be directly served from that directory, _without any fruther configuration_. Simple as that.

> Before we begin, ensure that you are running the following the following commands as a regular, non-root user account on your server with `sudo` privileges. Digital Ocean has a [good guide on how to get that setup](https://www.digitalocean.com/community/tutorials/initial-server-setup-with-ubuntu-14-04) if you are new to Linux.

---
## Step one -- Installing the Web Services

I am running these commands on an **Ubuntu 14.04 x64** standard server installation. Due to this, we will be using the standard Ubuntu repositories accessible via `apt`. We will need to begin by updating the local package index, and downloading the needed packages:

```
$ sudo apt-get update
$ sudo apt-get install nginx mysql-server php5-fpm php5-mysql
```

In Ubuntu, the Nginx service is started up upon installing the package. Just to ensure it's up and running, you can head to `http://<your-server-hostname>` and you should be able to see something like this:

![Default Nginx Installation](https://i.imgur.com/KoR1DSm.png)

---
## Step Two -- Securely Setting up MySQL

In the previous steps, you should have installed MySQL, and during the installation process it should prompt you to setup a _root MySQL password._ However, given that MySQL should be containing all important site data, much of which should be quite secure. By default, MySQL isn't going to be very secure. Run the following command to properly setup MySQL, and ensure the configuration isn't set to default.

This command will basically properly structure the MySQL **data** directory:

```
$ sudo mysql_install_db
```

Next, the following command will walk you through properly securing MySQL:

**Note, this will ask you if you would like to re-setup the root MySQL password. If the password you setup during the installation is secure, you shouldn't need to change it during this step.**

```
$ sudo mysql_secure_installation
```


---
## Step Three -- Configuring the PHP handler for Nginx

Due to Nginx being very lightweight, it doesn't function with PHP by default, unlike other web servers. To get PHP (or _php5-fpm_ in this case) to function with Nginx, we will need to configure Nginx to pass PHP based requests to the `php5-fpm` service.


The following will make a small change to the `php5-fpm` configuration file, allowing a more secure installation:

```
$ sudo sed -ri "s/;?cgi\.fix_pathinfo=[0-1]$/cgi.fix_pathinfo=0/g" /etc/php5/fpm/php.ini
```


If you would like to modify the file yourself, feel free to open up the configuration file with a text editor:

```
$ sudo nano /etc/php5/fpm/php.ini
```


The parameter in the configuration file you are going to be looking for is `cgi.fix_pathinfo`. Changing it to the following, will ensure remote users can't execute PHP files out of the accessible directories provided to them (**Note:** ensure there is no `;` in front of the line, this is an inline comment character, and it won't be applied if it exists):

```
cgi.fix_pathinfo=0
```


Once done, simply restart the php5-fpm service with the following:

```
$ sudo service php5-fpm restart
```


---
## Step Four -- Configure Nginx

Due to the way that `php5-fpm` works, it requires passing requests to a sock file to the php5-fpm service. An Nginx configuration section needs to be added to any server block to tell Nginx that `.php` requests are special. Add the below to `/etc/nginx/fastcgi_php`, which we can later include in an Nginx configuration file for easy setup:

```
location ~ \.php$ {
    include /etc/nginx/fastcgi_params;
    fastcgi_index index.php;
    fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    if (-f $request_filename) {
        fastcgi_pass unix:/var/run/php5-fpm.sock;
    }
}
```


Now is the core of this post. Basically, we can tell Nginx that it should look in a specific directory for the web content of each site pointed towards the server. As mentioned at the start of the post, website roots will be in `/var/www/<domain>`. Put the following into `/etc/nginx/sites-enabled/core`:

```
server {
    server_name ~^(www\.)?(?<domain>.+)$;
    if ($host ~* ^www\.(.*)) {
        set $host_without_www $1;
        rewrite ^(.*) http://$host_without_www$1 permanent;
    }
    root /var/www/$host;
    include fastcgi_php;
    index index.php index.html index.htm index.xhtml index.cgi;
    location / {
        autoindex on;
        autoindex_exact_size off;
    }
}
```


Now, make the needed directory for Nginx to be able to pull content for the websites:

```
$ sudo mkdir -p /var/www
```

While you are at it, assuming you are setting up a domain whilst setting up the LEMP stack, you might as well create that directory too:

```
$ sudo mkdir -p /var/www/<domain>
```


Now, restart Nginx. Assuming nothing was done improperly, this should show as successful:

```
$ sudo service nginx restart
```


---
## Step Five -- Verify the Installation

At this point, everything should be setup properly. To test, we can create a PHP test file, and if all goes well, you should see the information from your PHP installation.

Create a file named `index.php` in `/var/www/<domain>/index.php` with the following contents:

```
<?php phpinfo(); ?>
```


As mentioned, if all goes well, going to `http://<your-domain>` should yield a page that looks like this:

![PHP info page](https://i.imgur.com/eVx2T2E.png)

Once verified, make sure to remove this file, as it exposes configuration locations and information about the installation which can pose as an advantage to malicious users.

You should now have a LEMP stack installed on your Ubuntu 14.04 server. Happy hosting!
