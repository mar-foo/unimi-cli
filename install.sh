#!/bin/sh
INSTALL_PATH="${HOME}/.local/bin"
case "${OSTYPE}" in
				"linux-gnu"*, *"bsd"*)	test -d ${INSTALL_PATH} || mkdir -p ${INSTALL_PATH} ; cp bin/linux/university-cli university-cli ${INSTALL_PATH} ;;
				"darwin"*) test -d ${INSTALL_PATH} || mkdir -p ${INSTALL_PATH} ; cp bin/macos/university-cli university-cli ${INSTALL_PATH} ;;
esac
