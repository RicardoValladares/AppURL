# AppURL
Abrir Exe o Bin local desde el Navegador Web, como lo hace la aplicacion zoom...


<html>
	<head>
		<title>URL que abre un Ejecutable Local</title>
		<script type='text/javascript'>
			function appurl(){
				window.location.assign("appurl://argumentos");
				self.focus();
			}
		</script>
	</head>
	<body>
		<input type='button' onclick='appurl()' value='Ejecutar'>
	</body>
</html>
