FROM docker.elastic.co/logstash/logstash:8.5.3

RUN /opt/logstash/bin/logstash-plugin install logstash-integration-jdbc
RUN mkdir /opt/logstash/vendor/jdbc
RUN curl -o /opt/logstash/vendor/jdbc/postgresql-42.5.1.jar https://jdbc.postgresql.org/download/postgresql-42.5.1.jar

ENTRYPOINT ["/usr/local/bin/docker-entrypoint"]