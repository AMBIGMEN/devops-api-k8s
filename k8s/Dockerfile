FROM centos:latest
COPY start.sh /root
COPY main /root
COPY conf/* /root/conf/
COPY data/* /root/data/
COPY script/* /root/script/
RUN chmod u+x /root/main /root/start.sh
ENTRYPOINT ["/root/start.sh"]
