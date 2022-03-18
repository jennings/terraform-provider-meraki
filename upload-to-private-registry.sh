#!/usr/bin/env -S bash -euo pipefail

# Uploads artifacts created by goreleaser to a Terraform private registry.

BASE_URL=${BASE_URL:-https://app.terraform.io/api/v2}
PROVIDER=${PROVIDER:-meraki}
VERSION=${VERSION:-0.0.1}

function createAndUploadPlatform()
{
        echo "creating $1 $2"
        res=$(curl --header "Authorization: Bearer $TERRAFORM_API_KEY" \
             --header "Content-Type: application/vnd.api+json" \
             --request POST \
             --data '{"data":{
                     "type": "registry-provider-platform",
                     "attributes": {
                        "os": "'$1'",
                        "arch": "'$2'",
                        "shasum": "'$(grep <dist/terraform-provider-meraki_${VERSION}_SHA256SUMS "_$1_$2\\." | cut -f 1 -d " ")'",
                        "filename": "'$3'"
                     }}}' \
             $BASE_URL/organizations/$ORGANIZATION/registry-providers/private/$ORGANIZATION/$PROVIDER/versions/$VERSION/platforms)
        echo "$res"
        upload_url=$(echo "$res" | jq --raw-output '
                .data.links."provider-binary-upload"
                ')

        echo "uploading $4"
        curl -T "$4" "$upload_url"
}

ARCHIVES=$(jq --raw-output <dist/artifacts.json '
        sort_by(.goos + .goarch)
        | .[]
        | select(.type == "Archive")
        | .goos + " " + .goarch + " " + .name + " " + .path')

echo "$ARCHIVES"  | while read line; do
        createAndUploadPlatform $line
done
