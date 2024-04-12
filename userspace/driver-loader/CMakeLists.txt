cmake_minimum_required(VERSION 3.12)
project(go-scap-driver-loader)

set(GO_MODULE_PATH ${CMAKE_CURRENT_SOURCE_DIR})

string(TIMESTAMP BUILD_DATE "%Y-%m-%dT%H:%M:%SZ")

set(RELEASE ${SYSDIG_VERSION})
set(COMMIT "")

set(GO_LDFLAGS "")
string(APPEND GO_LDFLAGS "-X 'github.com/falcosecurity/falcoctl/cmd/version.semVersion=${RELEASE}' ")
string(APPEND GO_LDFLAGS "-X 'github.com/falcosecurity/falcoctl/cmd/version.gitCommit=${COMMIT}' ")
string(APPEND GO_LDFLAGS "-X 'github.com/falcosecurity/falcoctl/cmd/version.buildDate=${BUILD_DATE}'")

add_custom_target(scap_driver_loader_exe ALL DEPENDS ${CMAKE_CURRENT_BINARY_DIR}/scap-driver-loader)
add_custom_command(
    OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/scap-driver-loader
    COMMAND ${GOENV} go build -ldflags ${GO_LDFLAGS}
        -o ${CMAKE_CURRENT_BINARY_DIR}/scap-driver-loader ${GO_MODULE_PATH}
    WORKING_DIRECTORY ${GO_MODULE_PATH}
    #DEPENDS ${CMAKE_CURRENT_SOURCE_DIR}/${target}/go.mod ${ARGN}
)
