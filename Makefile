.PHONY: *

gogo: stop-services build truncate-logs start-services bench kataribe

build:
	cd webapp/go && make build

stop-services:
	sudo systemctl stop nginx
	sudo systemctl stop isubata.golang.service
	sudo systemctl stop mysql

start-services:
	sudo systemctl start mysql
	sleep 5
	sudo systemctl start isubata.golang.service
	sudo systemctl start nginx

truncate-logs:
	sudo truncate --size 0 /var/log/nginx/access.log
	sudo truncate --size 0 /var/log/nginx/error.log
	sudo truncate --size 0 /var/log/mysql/mysql-slow.log
	sudo chmod 777 /var/log/mysql/mysql-slow.log
	sudo journalctl --vacuum-size=1K

bench:
	ssh bench "cd ~/isubata/bench && ./bin/bench -remotes 172.31.1.138"

kataribe:
	sudo cat /var/log/nginx/access.log | ./kataribe -conf kataribe.toml | grep --after-context 20 "Top 20 Sort By Total"
