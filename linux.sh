#abre leafpad.dekstop con editor de texto buscandolo en /usr/share/applications/ sino esta instala leafpad
#para ver el argumento que envia la URL en leafpad.desktop en vez de usar %f usa %U
xdg-mime default appurl.desktop x-scheme-handler/appurl
