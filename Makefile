INSTALL_PATH="${HOME}/.local/bin"
install: university-cli
	test -d ${INSTALL_PATH} || mkdir -p ${INSTALL_PATH} ; cp university-cli ${INSTALL_PATH}
