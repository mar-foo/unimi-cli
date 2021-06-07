#!/bin/bash
INSTALL_PATH="${HOME}/.local/bin"
case "${OSTYPE}" in
				"linux-"*)	test -d ${INSTALL_PATH} || mkdir -p ${INSTALL_PATH} ; cp bin/linux/university-cli university-cli.sh ${INSTALL_PATH} ;;
				"darwin"*) test -d ${INSTALL_PATH} || mkdir -p ${INSTALL_PATH} ; cp bin/macos/university-cli university-cli.sh ${INSTALL_PATH} ;;
				*) echo "Operating system ${OSTYPE} not supported" ;;
esac
