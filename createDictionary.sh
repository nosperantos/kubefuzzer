#!/usr/bin/env	bash
# Fetches Kubernetes API paths and replaces placeholders (e.g., {namespace}) with example values.
curl -s https://raw.githubusercontent.com/kubernetes/kubernetes/master/api/openapi-spec/swagger.json | jq -r '.paths | keys[]' | sed \
    -e 's/{namespace}/default/g' \
    -e 's/{name}/example-resource/g' \
    -e 's/{path}/example-path/g' \
    -e 's/{logpath}/example-logpath/g' > kubernetes_endpoints.txt
