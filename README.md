# Frens


Frens is an open source social networking server similar to [Mastodon](https://github.com/mastodon/mastodon), [Plemora](https://github.com/Hostdon/pleroma), or [GoToSocial](https://github.com/superseriousbusiness/gotosocial).

Anyone interested in the project is more than welcome to contribute.

### Built With
- https://go.dev/
- https://github.com/gofiber/fiber
- https://github.com/spf13/cobra
- https://github.com/spf13/viper
- https://github.com/dgraph-io/badger
- https://github.com/postgres/postgres

# Design
Frens is designed to be modular, supporting multiple forms of submission and retrieval of data. This allows for easy customization of frontends. When designing a frontend, its not necessary to use all features of the server, in fact, it is encouraged not to do so. This allows for a more diverse ecosystem of frontends.

## API

The API is designed to avoid as much nesting as possible, instead perferring to use multiple endpoints. For example, statuses do not incude reactions and media, but instead have their own endpoints. This allows for more flexibility in the frontend, and allows for more efficient caching.

## Reverse Proxy

The frens server is not meant to recieve requests from end users directly. Instead, it uses a reverse proxy to forward requests to the server. Many frontends have the reverse proxy setup included as part of their setup instructions, but in order to support clients such as Pinafore, we need to set up a reverse proxy ourselves.

We start by installing the reverse proxy. In this example, we are using NGINX running on Ubuntu 22.04.

```bash
sudo apt update
sudo apt install nginx
```

Once the reverse proxy is installed, we need to create a configuration file.
    
```bash
sudo nano /etc/nginx/sites-available/frens
```

The following is a known working configuration for frens.moe. Change the website and port as appropriate.

```
server {
        server_name frens.moe www.frens.moe;

        location / {
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header Host $host;
                proxy_set_header X-NginX-Proxy true;
                proxy_pass http://localhost:4000/;
                proxy_redirect http://localhost:4000/ https://$server_name/;
        }
}
```

Once the configuration is created, we need to create a linked file in the sites-enabled directory.

```bash
sudo ln -s /etc/nginx/sites-available/frens /etc/nginx/sites-enabled/frens
```

Now we can use Certbot to create the certificates and automatically finish the configuration. Again, replace the website and port as appropriate.

```bash
sudo certbot --nginx -d frens.moe -d www.frens.moe
```

Now requests can come in from applications to our server on port 80 and 443, which will then be forwarded to the server on port 4000.