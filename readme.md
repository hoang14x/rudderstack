# Setup

## Control Plane Cloud
### Control plane
Sign up and log into the https://app.rudderlabs.com/signup

### Data plane
With Source: https://github.com/rudderlabs/rudder-server

With docker: `docker-compose -f docker-compose-cloud.yml up`

## Control Plane Self-Host

### Control plane
Docs: https://www.rudderstack.com/docs/user-guides/how-to-guides/rudderstack-control-plane-lite/

Launch app:
```
    git clone https://github.com/rudderlabs/config-generator
    npm install
    npm start
```

### Data plane
With Source: https://github.com/rudderlabs/rudder-server

With docker: `docker-compose -f docker-compose.yml up`