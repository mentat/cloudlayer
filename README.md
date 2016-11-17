# The Cloud Layer

A provisioning interface for various cloud systems.


# Setup

    go get -u github.com/gophercloud/gophercloud
    go get -u github.com/aws/aws-sdk-go


# Usage

    layer, err := NewCloudLayer("openstack")
    if err != nil {
        return err
    }
    err = layer.DetailedAuthorize(map[string]string{
        "identityEndpoint":"https://...",
        "password":"...",
        "tenantId":"...",
    })
