
## what is terraform

Terraform is an infrastruct provisioning tool.

That essentially means... if you want to get/create something.
You can allow terraform to do this for you

Terraform takes config files
```hcl
resource "twitch" "channel_point" {
    name = "dance for 5 seconds"
    cost = 4000
}
```

terraform apply 

creates that channel point and remembers that it created it

terraform destory 

it will remove that channel point


