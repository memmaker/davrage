cd ..
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o container/davrage -trimpath -ldflags '-s -w'
cd container || exit
podman build -t ghcr.io/memmaker/davrage .
rm davrage

while true; do
    read -rp "Do you wish to push this image? [yN]" yn
    case $yn in
        [Yy]* ) podman push ghcr.io/memmaker/davrage; break;;
        [Nn]* ) exit;;
        * ) exit;;
    esac
done