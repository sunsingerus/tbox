#!/bin/bash

# Error on unset variables
set -o nounset

# Source configuration
CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source "${CUR_DIR}/go_build_config.sh"

echo "Generating code with the following options:"
echo "      SRC_ROOT=${SRC_ROOT}"

# Check whether required binaries available
PROTOC="${PROTOC:-protoc}"
PROTOC_GEN_DOC="${PROTOC_GEN_DOC:-protoc-gen-doc}"

# Check protoc is available
"${PROTOC}" --help > /dev/null
if [[ $? ]]; then
    :
else
    echo "${PROTOC} is not available. Abort"
    exit 1
fi

# Check protoc-gen-doc is available
"${PROTOC_GEN_DOC}" --help > /dev/null
if [[ $? ]]; then
    :
else
    echo "${PROTOC_GEN_DOC} is not available. Abort"
    exit 1
fi

# Setup folders
# Where to look for .proto files
PROTO_ROOT="${SRC_ROOT}/proto"
# Where to write generated *.pb.go files
PB_GO_ROOT="${SRC_ROOT}/pkg"
# Where to write generated *_pb.js files
PB_JS_ROOT="${SRC_ROOT}/js"

DOC_ROOT="${SRC_ROOT}/docs/"
LOGGING="no"
#
#
#
function log() {
  if [[ "${LOGGING}" == "yes" ]]; then
      echo "${1}"
  fi
}

#
# Build full set of "-I ..." statements for protoc.
# Should be inserted "as is" as protoc parameter.
#
function imports() {
    local PROTO_ROOT_FOLDER="${1}"

    # List folders in which to search for imports.
    # It will contain 2 folders:
    # 1. Proto root in vendor (if any)
    # 2. Provided PROTO_ROOT_FOLDER

    # Add proto folder with all its sub-folders to "imports"
    local IMPORTS="-I ${PROTO_ROOT_FOLDER}"

    local PROTO_ROOT_FOLDER=$(realpath "${SRC_ROOT}/vendor/github.com/sunsingerus/tbbox/proto/" 2>/dev/null)
    if [[ -d "${PROTO_ROOT_FOLDER}" ]]; then
        # Specified folder exists
        IMPORTS="${IMPORTS} -I ${PROTO_ROOT_FOLDER}"
    fi

    echo -n "${IMPORTS}"
}

#
#
#
function clean_grpc_code_go() {
    local CODE_FOLDER="${1}"

    log "Go code generator. Clean previously generated .pb.go files in ${CODE_FOLDER}"
    rm -f "${CODE_FOLDER}"/*.pb.go
    for ABS_CODE_SUB_FOLDER in $(find "${CODE_FOLDER}" -type d ! -path "${CODE_FOLDER}"); do
        log "Go code generator. Clean previously generated .pb.go files in ${ABS_CODE_SUB_FOLDER}"
        rm -f "${ABS_CODE_SUB_FOLDER}"/*.pb.go
    done
}

#
# Compile .proto files in specified folder
#
function generate_grpc_code_go_folder() {
    local PROTO_FOLDER="${1}"
    local RESULT_FOLDER="${2}"
    local IMPORTS="${3}"
    local CLEAN="${4}"

    # Prepare target folder - clean target folder in case it exsists
    if [[ "${CLEAN}" == "clean" ]]; then
        if [[ -d "${RESULT_FOLDER}" ]]; then
            clean_grpc_code_go "${RESULT_FOLDER}"
        fi
    fi

    # What is the number of .proto files in the specified folder?
    local N=$(ls -1q "${PROTO_FOLDER}"/*.proto 2>/dev/null | wc -l)
    if [[ "${N}" == "0" ]]; then
        log "No .proto files in ${PROTO_FOLDER} skip it"
    else
        # Prepare target folder
        mkdir -p "${RESULT_FOLDER}"

        log "Compile .proto files in ${PROTO_FOLDER} into ${RESULT_FOLDER}"
        "${PROTOC}" \
            ${IMPORTS} \
            --go_opt=paths=source_relative \
            --go_out="${RESULT_FOLDER}" \
            --go-grpc_opt=paths=source_relative \
            --go-grpc_out="${RESULT_FOLDER}" \
            "${PROTO_FOLDER}"/*.proto
    fi
}

#
# Generate Go code from .proto files
#
function generate_grpc_code_go() {
    local PROTO_FOLDER="${1}"
    local RESULT_FOLDER="${2}"

    log "Go code generator. Generate code from .proto files in ${PROTO_FOLDER} into ${RESULT_FOLDER}"

    if [[ -z "${PROTO_FOLDER}" ]]; then
        echo "Go code generator. Need to specify folder where to look for .proto files to generate code from "
        exit 1
    fi

    log "Go code generator. Compile .proto files in ${PROTO_FOLDER}"
    log "${PROTOC} details:"
    log "Executable:" $(which "${PROTOC}")
    log "Version:" $("${PROTOC}" --version)
    log "protoc-gen-go details:"
    log "Executable:" $(which protoc-gen-go)

    # Specify the directory in which to search for imports.
    local IMPORTS=$(imports "${PROTO_FOLDER}")
    log "Imports: ${IMPORTS}"

    # Compile specified folder and all it's sub-folders
    generate_grpc_code_go_folder "${PROTO_FOLDER}" "${RESULT_FOLDER}" "${IMPORTS}" clean
    for ABS_PROTO_SUB_FOLDER in $(find "${PROTO_FOLDER}" -type d ! -path "${PROTO_FOLDER}"); do
        generate_grpc_code_go_folder "${ABS_PROTO_SUB_FOLDER}" "${RESULT_FOLDER}" "${IMPORTS}" noclean
    done
    log ""
}

#
#
#
function clean_grpc_code_js() {
    local CODE_FOLDER="${1}"

    log "JS code generator. Clean previously generated *_pb.js files in ${CODE_FOLDER}"
    rm -f "${CODE_FOLDER}"/*_pb.js
    for ABS_CODE_SUB_FOLDER in $(find "${CODE_FOLDER}" -type d ! -path "${CODE_FOLDER}"); do
        log "JS code generator. Clean previously generated *_pb.js files in ${ABS_CODE_SUB_FOLDER}"
        rm -f "${ABS_CODE_SUB_FOLDER}"/*_pb.js
    done
}

#
# Compile .proto files in specified folder
#
function generate_grpc_code_js_folder() {
    local PROTO_FOLDER="${1}"
    local RESULT_FOLDER="${2}"
    local IMPORTS="${3}"
    local CLEAN="${4}"

    # What is the number of .proto files in the specified folder?
    local N=$(ls -1q "${PROTO_FOLDER}"/*.proto 2>/dev/null | wc -l)
    if [[ "${N}" == "0" ]]; then
        log "No .proto files in ${PROTO_FOLDER} skip it"
    else
        # Prepare target folder
        mkdir -p "${RESULT_FOLDER}"
        if [[ "${CLEAN}" == "clean" ]]; then
            clean_grpc_code_js "${RESULT_FOLDER}"
        fi

        log "Compile .proto files in ${PROTO_FOLDER} into ${RESULT_FOLDER}"
        "${PROTOC}" \
            ${IMPORTS} \
            --js_out=import_style=commonjs:"${RESULT_FOLDER}" \
            "${PROTO_FOLDER}"/*.proto

        log "Generate the service client stub from .proto files in ${PROTO_FOLDER} into ${RESULT_FOLDER}"
        # In the --grpc-web_out param:
        # import_style can be closure (default) or commonjs
        # mode can be grpcwebtext (default) or grpcweb
        "${PROTOC}" \
            ${IMPORTS} \
            --grpc-web_out=import_style=commonjs,mode=grpcwebtext:"${RESULT_FOLDER}" \
            "${PROTO_FOLDER}"/*.proto
    fi
}

#
# Generate JS code from .proto files
#
function generate_grpc_code_js() {
    local PROTO_FOLDER="${1}"
    local RESULT_FOLDER="${2}"

    log "JS code generator. Generate code from .proto files in ${PROTO_FOLDER} into ${RESULT_FOLDER}"

    if [[ -z "${PROTO_FOLDER}" ]]; then
        echo "JS code generator. Need to specify folder where to look for .proto files to generate code from."
        exit 1
    fi

    log "JS code generator. Compile .proto files in ${PROTO_FOLDER} into ${RESULT_FOLDER}"

    # Specify the directory in which to search for imports
    local IMPORTS=$(imports "${PROTO_FOLDER}")
    log "Imports: ${IMPORTS}"

    # Compile specified folder and all its sub-folders
    generate_grpc_code_js_folder "${PROTO_FOLDER}" "${RESULT_FOLDER}" "${IMPORTS}" clean
    for ABS_PROTO_SUB_FOLDER in $(find "${PROTO_FOLDER}" -type d ! -path "${PROTO_FOLDER}"); do
        generate_grpc_code_js_folder "${ABS_PROTO_SUB_FOLDER}" "${RESULT_FOLDER}" "${IMPORTS}" noclean
    done
    log ""

    # Copy *.js files into RESULT_FOLDER
    local JS_ROOT_FOLDER=$(realpath "${SRC_ROOT}/vendor/github.com/sunsingerus/tbox/js/lib/api" 2>/dev/null)
    if [[ -d "${JS_ROOT_FOLDER}" ]]; then
        # Add specified folder itself
        cp -r "${JS_ROOT_FOLDER}" "${RESULT_FOLDER}"
    fi
}

#
# Delete String() function from generated *.pb.go files
# This function is not that human-friendly and it is better to introduce own function for each type
function delete_string_function_folder() {
    local PB_GO_FILES_FOLDER="${1}"

    if [[ -z "${PB_GO_FILES_FOLDER}" ]]; then
        echo "need to specify folder where to look for .pb.go files to process"
        exit 1
    fi

    # What is the number of .proto files in the specified folder?
    local N=$(ls -1q "${PB_GO_FILES_FOLDER}"/*.pb.go 2>/dev/null | wc -l)
    if [[ "${N}" == "0" ]]; then
        log "No .pb.go files in ${PB_GO_FILES_FOLDER} skip it"
    else
        log "Delete String() in ${PB_GO_FILES_FOLDER}"
        RERUN="yes"
        while [[ "${RERUN}" == "yes" ]]; do
            RERUN="no"
            # /path/to/file:LINE_NUMBER:line
            # /path/to/file:31:func (x *Address) String() string {
            FILE_AND_LINEs=$(grep -nH "String() string {" "${PB_GO_FILES_FOLDER}"/*.pb.go | cut -f1,2 -d:)

            PREV_FILE="-"
            for FILE_AND_LINE in $FILE_AND_LINEs; do
                # Cut filename from the grep-output line
                FILE=$(echo "${FILE_AND_LINE}" | cut -f1 -d:)
                if [[ "${FILE}" == "${PREV_FILE}" ]]; then
                    log "See dup for file ${FILE}, need to re-run"
                    RERUN="yes"
                    break
                fi
                PREV_FILE="${FILE}"
                # Find lines where String() func starts and ends
                LINE=$(echo "${FILE_AND_LINE}" | cut -f2 -d:)
                LINE_END=$((LINE + 2))
                #echo "${FILE}:${LINE}:${LINE_END}"
                # Cut specified line(s) from the file and rewrite the file
                FILE_NEW="${FILE}".new
                sed "${LINE},${LINE_END}d" "${FILE}" > "${FILE_NEW}"
                LEN=$(wc -l "${FILE}" | awk '{print $1}')
                LEN_NEW=$(wc -l "${FILE_NEW}" | awk '{print $1}')
                mv "${FILE_NEW}" "${FILE}"
                if [[ "${LEN}" == "${LEN_NEW}" ]]; then
                    :
                else
                    log "Cut String() from $FILE"
                fi
            done
        done
    fi
    log "--- Delete String() Done ---"
}

#
#
#
function delete_string_function() {
    local PB_GO_FOLDER="${1}"

    log "String() function deleter. Delete in ${PB_GO_FOLDER}"

    #delete_string_function_folder "${PB_GO_FOLDER}"
    #echo $(find "${PB_GO_FOLDER}" -type d ! -path "${PB_GO_FOLDER}")
    for ABS_SUB_FOLDER in $(find "${PB_GO_FOLDER}" -type d ! -path "${PB_GO_FOLDER}"); do
        delete_string_function_folder "${ABS_SUB_FOLDER}"
    done
}

#
#
#
BUILD_DOCS_HTML="yes"
BUILD_DOCS_MD="yes"

#
#
#
function clean_doc() {
    local CODE_FOLDER="${1}"

    log "Docs generator. Clean previously generated doc files in ${CODE_FOLDER}"
    rm -f "${CODE_FOLDER}"/*.html
    rm -f "${CODE_FOLDER}"/*.md
    for ABS_CODE_SUB_FOLDER in $(find "${CODE_FOLDER}" -type d ! -path "${CODE_FOLDER}"); do
        log "Docs generator. Clean previously generated doc files in ${ABS_CODE_SUB_FOLDER}"
        rm -f "${ABS_CODE_SUB_FOLDER}"/*.html
        rm -f "${ABS_CODE_SUB_FOLDER}"/*.md
    done
}

#
# Compile .proto files in specified folder
#
function generate_doc_folder() {
    local PROTO_FOLDER="${1}"
    local RESULT_FOLDER="${2}"
    local IMPORTS="${3}"
    local CLEAN="${4}"

    # Prepare target folder - clean target folder in case it exsists
    if [[ "${CLEAN}" == "clean" ]]; then
        if [[ -d "${RESULT_FOLDER}" ]]; then
            clean_doc "${RESULT_FOLDER}"
        fi
    fi

    # What is the number of .proto files in the specified folder?
    local N=$(ls -1q "${PROTO_FOLDER}"/*.proto 2>/dev/null | wc -l)
    if [[ "${N}" == "0" ]]; then
        log "No .proto files in ${PROTO_FOLDER} skip it"
    else
        # Prepare target folder
        mkdir -p "${RESULT_FOLDER}"

        if [[ "${BUILD_DOCS_HTML}" == "yes" ]]; then
            local DOC_FILE_NAME_HTML="$(basename ${PROTO_FOLDER}).html"
            log "Compile .proto files in ${PROTO_FOLDER} into ${RESULT_FOLDER}/${DOC_FILE_NAME_HTML}"
            "${PROTOC}" \
                ${IMPORTS} \
                --doc_out="${RESULT_FOLDER}" \
                --doc_opt=html,"${DOC_FILE_NAME_HTML}" \
                "${PROTO_FOLDER}"/*.proto
        fi

        if [[ "${BUILD_DOCS_MD}" == "yes" ]]; then
            local DOC_FILE_NAME_HTML="$(basename ${PROTO_FOLDER}).md"
            log "Compile .proto files in ${PROTO_FOLDER} into ${RESULT_FOLDER}/${DOC_FILE_NAME_HTML}"
            "${PROTOC}" \
                ${IMPORTS} \
                --doc_out="${RESULT_FOLDER}" \
                --doc_opt=markdown,"${DOC_FILE_NAME_HTML}" \
                "${PROTO_FOLDER}"/*.proto
        fi
    fi
}

#
# Generate Go code from .proto files
#
function generate_doc() {
    local PROTO_FOLDER="${1}"
    local RESULT_FOLDER="${2}"

    log "Go code generator. Generate code from .proto files in ${PROTO_FOLDER} into ${RESULT_FOLDER}"

    if [[ -z "${PROTO_FOLDER}" ]]; then
        echo "Go code generator. Need to specify folder where to look for .proto files to generate code from "
        exit 1
    fi

    log "Go code generator. Compile .proto files in ${PROTO_FOLDER}"
    log "${PROTOC} details:"
    log "Executable:" $(which "${PROTOC}")
    log "Version:" $("${PROTOC}" --version)
    log "protoc-gen-go details:"
    log "Executable:" $(which protoc-gen-go)

    # Specify the directory in which to search for imports.
    local IMPORTS=$(imports "${PROTO_FOLDER}")
    log "Imports: ${IMPORTS}"

    # Compile specified folder and all it's sub-folders
    generate_doc_folder "${PROTO_FOLDER}" "${RESULT_FOLDER}" "${IMPORTS}" clean
    for ABS_PROTO_SUB_FOLDER in $(find "${PROTO_FOLDER}" -type d ! -path "${PROTO_FOLDER}"); do
        generate_doc_folder "${ABS_PROTO_SUB_FOLDER}" "${RESULT_FOLDER}" "${IMPORTS}" noclean
    done
    log ""
}

function generate_docs() {
    local AREA="${1}"
    local PACKAGE_NAME="${2}"
    local PROTO_FILES_FOLDER="${3}"

    log "AutoDoc generator. Generate docs start."

    local DOC_FILES_FOLDER="${DOCS_ROOT}/${AREA}/${PACKAGE_NAME}"
    log "AutoDoc generator. Prepare folder for docs ${DOC_FILES_FOLDER}"
    mkdir -p "${DOC_FILES_FOLDER}"

    if [[ "${BUILD_DOCS_HTML}" == "yes" ]]; then
        local DOC_FILE_NAME_HTML="${PACKAGE_NAME}.html"

        rm -f "${DOC_FILES_FOLDER}/${DOC_FILE_NAME_HTML}"
        "${PROTOC}" \
            -I "${PROTO_FILES_FOLDER}" \
            --doc_out="${DOC_FILES_FOLDER}" \
            --doc_opt=html,"${DOC_FILE_NAME_HTML}" \
            "${PROTO_FILES_FOLDER}"/*.proto
    fi

    if [[ "${BUILD_DOCS_MD}" == "yes" ]]; then
        local DOC_FILE_NAME_MD="${PACKAGE_NAME}.md"
        rm -f "${DOC_FILES_FOLDER}/${DOC_FILE_NAME_MD}"
        "${PROTOC}" \
            -I "${PROTO_FILES_FOLDER}" \
            --doc_out="${DOC_FILES_FOLDER}" \
            --doc_opt=markdown,"${DOC_FILE_NAME_MD}" \
            "${PROTO_FILES_FOLDER}"/*.proto
    fi
}

generate_grpc_code_go "${PROTO_ROOT}"  "${PB_GO_ROOT}"
generate_grpc_code_js "${PROTO_ROOT}"  "${PB_JS_ROOT}"/lib

delete_string_function "${PB_GO_ROOT}"/api

#generate_docs api common  "${PROTO_ROOT}"/api/common
#generate_docs api service "${PROTO_ROOT}"/api/service
#generate_docs api health "${PROTO_ROOT}"/api/health

generate_doc "${PROTO_ROOT}"  "${DOC_ROOT}/api"

echo "--- Done ---"
