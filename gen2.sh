go get github.com/txix-open/gowsdl/...
mkdir -p wsdl2/stubsV1/findV1
mkdir -p wsdl2/stubsV2/findV2
mkdir -p wsdl2/stubsV1/sudirV1
mkdir -p wsdl2/stubsV2/sudirV2
gowsdl -o=wsdl2/stubsV1/findV1/findService.go -p="" wsdl/find/findV1.wsdl
gowsdl -o=wsdl2/stubsV2/findV2/findService.go -p="" wsdl/find/findV2.wsdl
gowsdl -o=wsdl2/stubsV1/sudirV1/sudirService.go -p="" wsdl/sudir/sudirV1.wsdl
gowsdl -o=wsdl2/stubsV2/sudirV2/sudirService.go -p="" wsdl/sudir/sudirV2.wsdl
