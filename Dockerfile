FROM envoyproxy/envoy:latest

COPY envoy_front_proxy_v2.yaml /etc/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy.yaml