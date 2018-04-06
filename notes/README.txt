
Deployment Notes
----------------
1. Build for linux destination architecture amd64:

cd <PROJECT>/
env GOOS=linux GOARCH=amd64 go build -o webapp -v

2. Copy your project excluding the Go source *.go to your server in /var/www/html/<PROJECT>.

rsync -azv --exclude '*.go' ../<PROJECT> user@example.com:~/var/www/html/

3. Configure supervisor to run the app.

sudo vim /etc/supervisor/conf.d/<PROJECT>.conf
# add the following

[program:<PROJECT>webapp]
environment =
    DBUSER=<DB_USER_HERE>,
    DBPASS=<DB_PASS_HERE>,
    DBNAME=<DB_NAME_HERE>,
    DBSOCKET=<DB_SOCKET_HERE>,
    DBCONNTYPE=<DB_CONNECTION_TYPE_HERE 1 (ip) or 2 (socket)>
    PORT=<APP_PORT_HERE>,
    APPSOCKET=<APP_SOCKET_HERE>,
    LISTENIP=<LISTEN_ON_IP_HERE>,
    LISTENTYPE=<LISTEN_TYPE_HERE 1 (ip) or 2 (socket)>,
    PRODUCTION=<ENVIRONMENT_HERE 1 or 0>,
    SMTPUSER=<SMTP_USER_HERE>,
    SMTPPASS=<SMTP_PASS_HERE>,
    SMTPSERVER=<SMTP_SERVER_ADDRESS:PORT_HERE>,
    UPLOADDIR=</PATH/TO/STATIC/UPLOAD/DIR>,
    SECRET=<APPLICATION_SECRET>,
command=/var/www/html/<PROJECT>/webapp
autostart=true
autorestart=true
startretries=10
user=www-data
directory=/var/www/html/<PROJECT>/
redirect_stderr=true
stdout_logfile=/var/log/supervisor/<PROJECT>webapp.log
stdout_logfile_maxbytes=50MB
stdout_logfile_backups=10



sudo supervisorctl reload

4. Configure nginx as your application frontend.

sudo vim /etc/nginx/sites-available/<PROJECT>.conf
# add the following

upstream <PROJECT> {
    server unix:/tmp/<PROJECT>.sock;
}

server {
    listen 80;
    server_name <PROJECT>;
    access_log /var/log/nginx/<PROJECT>-access.log;
    error_log /var/log/nginx/<PROJECT>-error.log error;
    location /static/ { alias /var/www/html/<PROJECT>/res/static/; }
    location / {
        proxy_pass http://<PROJECT>;
    }
}



sudo ln -s /etc/nginx/sites-available/<PROJECT>.conf /etc/nginx/sites-enabled
sudo service nginx reload


PROD DEPLOY
-----------
cd <PROJECT>/
env GOOS=linux GOARCH=amd64 go build -o webapp -v
rsync -avz --exclude '*.go' --exclude 'res/static/uploads' --exclude 'notes' --exclude '.DS_Store' <PROJECT>/ user@example.com:/var/www/html/<PROJECT>/
ssh -i <key.pem> user@example.com
sudo supervisorctl reload
