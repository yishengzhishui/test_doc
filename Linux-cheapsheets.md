1service 相关

```shell
journalctl -u content_update@vi.service  --since='1 hours ago'
find . -name 'push_queue_no_timed*' | xargs rm -rf
systemctl list-units | grep running | grep content|awk '{print $1}'|xargs systemctl status
systemctl status mangatoon-push-hyperf-distribute.service
```

2.redis 

--批量删除

```shell
redis-cli keys a_* | xargs ./redis-cli del
```

3 du

```
使用du -h -s /* | sort -nr命令查看那个目录占用空间大：
```

4.循环

```shell
for language in cn en id vi es pt th fr ja ar; do
  systemctl restart content_update@$language
done
for i in $(seq 0 125); do
    day=$(date -d "2022-04-25 +$i day" "+%Y-%m-%d")
    hyperf mangatoon-data-export sql:task runFile --sql_file_path='temp/world_travel_game.sql' --date=$day
done
```

