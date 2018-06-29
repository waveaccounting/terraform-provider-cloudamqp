# Cloudamqp Terraform Module

#### Configuring the provider
```HCL
provider "cloudamqp" {
    #required
    apikey = 'YOUR_API_KEY'

    #optional
    base_url = 'YOUR_BASE_URL'
}
```

#### Configuring the RabbitMQ instance
```HCL
resource "cloudamqp_instance" "instance-name" {
  name   = "terraform-provider-test-instance-1"
  plan   = "lemur"
  region = "amazon-web-services::us-east-1"
  nodes = 2
  vpc_subnet = "10.1.0.0/16"
  rmq_version = "3.7.5"
}
```
