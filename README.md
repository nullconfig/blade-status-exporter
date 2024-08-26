# Blade Status Exporter
This exporter logs into a chassis that supports the redfish API and extracts out the status of the blades inside of the chassis.

# CSM Redfish Emulator
Testing and development of the exporter relies on the [csm redfish emulator](https://github.com/Cray-HPE/csm-redfish-interface-emulator)

You'll have to pull the repo and build the image locally before launching it. I've provided an example dockerfile for building the container after pulling the repository so that it will work with the exporter.

I'm using podman, and be sure to update the build and compose command for your container manager.

### Building and starting the container
```
cd emulator
git clone https://github.com/Cray-HPE/csm-redfish-interface-emulator.git csm-redfish
mv ./dockerfile ./csm-redfish/Dockerfile
cd csm-redfish
podman build -t redfish-emulator -f dockerfile .
cd ../..
podman-compose -f compose.yaml up
```