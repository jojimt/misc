#!/bin/bash
GEN_FILES=$(git status | grep generated | awk '{print $3}')
for file in $GEN_FILES
do
git checkout -- $file
done
hack/update-generated-swagger-docs.sh
hack/update-codegen.sh
hack/update-generated-protobuf.sh
hack/update-codecgen.sh
hack/update-swagger-spec.sh
hack/update-openapi-spec.sh
