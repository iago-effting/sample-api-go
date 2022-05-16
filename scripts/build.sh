
#!/bin/sh

export PATH_BIN=${PATH_BIN}

go build                                                                                                    \
    -ldflags="-X '$(go list -m)/pkg/version.Version=$VERSION' -X '$(go list -m)/pkg/version.Time=$(date)'"  \
    -o ${PATH_BIN}/app                                                                                      \
    "$@"
